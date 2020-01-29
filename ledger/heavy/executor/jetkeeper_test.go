//
// Copyright 2020 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// +build slowtest

package executor

import (
	"context"
	"os"
	"sort"
	"sync"
	"testing"

	"github.com/insolar/insolar/insolar/gen"

	"github.com/insolar/insolar/ledger/heavy/migration"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/tests/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	fuzz "github.com/google/gofuzz"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/instrumentation/inslogger"
)

var (
	poolLock     sync.Mutex
	globalPgPool *pgxpool.Pool
)

func setPool(pool *pgxpool.Pool) {
	poolLock.Lock()
	defer poolLock.Unlock()
	globalPgPool = pool
}

func getPool() *pgxpool.Pool {
	poolLock.Lock()
	defer poolLock.Unlock()
	return globalPgPool
}

// TestMain does the before and after setup
func TestMain(m *testing.M) {
	ctx := context.Background()
	log.Info("[TestMain] About to start PostgreSQL...")
	pgURL, stopPostgreSQL := common.StartPostgreSQL()
	log.Info("[TestMain] PostgreSQL started!")

	pool, err := pgxpool.Connect(ctx, pgURL)
	if err != nil {
		stopPostgreSQL()
		log.Panicf("[TestMain] pgxpool.Connect() failed: %v", err)
	}

	migrationPath := "../../../migration"
	cwd, err := os.Getwd()
	if err != nil {
		stopPostgreSQL()
		panic(errors.Wrap(err, "[TestMain] os.Getwd failed"))
	}
	log.Infof("[TestMain] About to run PostgreSQL migration, cwd = %s, migration migrationPath = %s", cwd, migrationPath)
	ver, err := migration.MigrateDatabase(ctx, pool, migrationPath)
	if err != nil {
		stopPostgreSQL()
		panic(errors.Wrap(err, "Unable to migrate database"))
	}
	log.Infof("[TestMain] PostgreSQL database migration done, current schema version: %d", ver)

	setPool(pool)
	// Run all tests
	code := m.Run()

	log.Info("[TestMain] Cleaning up...")
	stopPostgreSQL()
	os.Exit(code)
}

func cleanJetsInfoTables() {
	ctx := context.Background()
	conn, err := getPool().Acquire(ctx)
	if err != nil {
		panic("Unable to acquire a database connection")
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "DELETE FROM jets_info")
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, "DELETE FROM key_value")
	if err != nil {
		panic(err)
	}
}

func initDB(t *testing.T, testPulse insolar.PulseNumber) (JetKeeper, *jet.DBStore, *pulse.DB) {
	cleanJetsInfoTables()

	ctx := inslogger.TestContext(t)
	jets := jet.NewDBStore(getPool())
	pulses := pulse.NewDB(getPool())
	err := pulses.Append(ctx, insolar.Pulse{PulseNumber: insolar.GenesisPulse.PulseNumber})
	require.NoError(t, err)

	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: testPulse})
	require.NoError(t, err)

	jetKeeper := NewJetKeeper(jets, getPool(), pulses)

	return jetKeeper, jets, pulses
}

func Test_TruncateHead_TryToTruncateTopSync(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	ji, _, _ := initDB(t, testPulse)
	err := ji.(*DBJetKeeper).TruncateHead(ctx, 1)
	require.EqualError(t, err, "try to truncate top sync pulse")
}

func TestJetInfoIsConfirmed_OneDropOneHot(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	ji, jets, _ := initDB(t, testPulse)
	testJet := insolar.ZeroJetID

	err := jets.Update(ctx, testPulse, true, testJet)
	require.NoError(t, err)

	err = ji.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())

	err = ji.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)

	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())

	err = ji.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)
	require.Equal(t, testPulse, ji.TopSyncPulse())
}

func Test_DifferentSplitFlagsInDropsAndHots(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	ji, _, _ := initDB(t, testPulse)

	testJet := insolar.ZeroJetID

	// AddHotConfirmation: 'true' come first
	err := ji.AddDropConfirmation(ctx, testPulse, testJet, true)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.Contains(t, err.Error(), "try to change split from true to false")

	// AddHotConfirmation: 'false' comes first
	left, _ := jet.Siblings(testJet)
	leftLeft, rightLeft := jet.Siblings(left)
	err = ji.AddHotConfirmation(ctx, testPulse, left, false)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, testPulse, leftLeft, true)
	require.Contains(t, err.Error(), "try to change split from false to true")

	// AddDropConfirmation
	err = ji.AddHotConfirmation(ctx, testPulse, rightLeft, false)
	require.NoError(t, err)
	err = ji.AddDropConfirmation(ctx, testPulse, rightLeft, true)
	require.Contains(t, err.Error(), "try to change split from false to true")
}

func TestJetInfoIsConfirmed_Split(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	nextPulse := gen.PulseNumber()
	if nextPulse < testPulse {
		nextPulse, testPulse = testPulse, nextPulse
	}

	ji, jets, pulses := initDB(t, testPulse)
	testJet := insolar.ZeroJetID

	err := jets.Update(ctx, testPulse, true, testJet)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())
	err = ji.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())
	err = ji.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)
	require.Equal(t, testPulse, ji.TopSyncPulse())

	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: nextPulse})
	require.NoError(t, err)

	left, right := jet.Siblings(testJet)
	err = jets.Update(ctx, nextPulse, true, testJet)
	require.NoError(t, err)
	err = ji.AddDropConfirmation(ctx, nextPulse, testJet, true)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, nextPulse, left, true)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, nextPulse, right, true)
	require.NoError(t, err)
	err = ji.AddBackupConfirmation(ctx, nextPulse)
	require.NoError(t, err)
	require.Equal(t, nextPulse, ji.TopSyncPulse())
}

func TestJetInfo_BackupConfirmComesFirst(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	jetKeeper, _, _ := initDB(t, testPulse)
	err := jetKeeper.AddBackupConfirmation(ctx, testPulse)
	require.Contains(t, err.Error(), "Received backup confirmation before replication data")
}

func TestJetInfo_ExistingDrop(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	jetKeeper, _, _ := initDB(t, testPulse)
	testJet := gen.JetID()
	err := jetKeeper.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)

	err = jetKeeper.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.Contains(t, err.Error(), "try to rewrite drop confirmation")
	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())
}

func TestJetInfo_ExistingHot(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	jetKeeper, _, _ := initDB(t, testPulse)

	testJet := gen.JetID()
	err := jetKeeper.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.Contains(t, err.Error(), "try add already existing hot confirmation")
}

func TestJetInfo_ExceedNumHotConfirmations(t *testing.T) {
	ctx := inslogger.TestContext(t)

	testPulse := gen.PulseNumber()
	jetKeeper, _, _ := initDB(t, testPulse)

	testJet := gen.JetID()
	left, right := jet.Siblings(testJet)

	err := jetKeeper.AddHotConfirmation(ctx, testPulse, left, true)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, right, true)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, left, true)
	require.Contains(t, err.Error(), "num hot confirmations exceeds")
	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())
}

func TestNewJetKeeper(t *testing.T) {
	jets := jet.NewDBStore(getPool())
	pulses := pulse.NewCalculatorMock(t)
	jetKeeper := NewJetKeeper(jets, getPool(), pulses)
	require.NotNil(t, jetKeeper)
}

func TestDbJetKeeper_DifferentActualAndExpectedJets(t *testing.T) {
	ctx := inslogger.TestContext(t)

	testPulse := gen.PulseNumber()
	jetKeeper, jets, _ := initDB(t, testPulse)
	testJet := gen.JetID()
	left, _ := jet.Siblings(testJet)

	err := jets.Update(ctx, testPulse, true, left)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)

	err = jetKeeper.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)

	require.False(t, jetKeeper.HasAllJetConfirms(ctx, testPulse))

	err = jetKeeper.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())
	require.False(t, jetKeeper.HasAllJetConfirms(ctx, testPulse))
}

func TestDbJetKeeper_DifferentNumberOfActualAndExpectedJets(t *testing.T) {
	ctx := inslogger.TestContext(t)

	testPulse := gen.PulseNumber()
	jetKeeper, jets, _ := initDB(t, testPulse)

	testJet := gen.JetID()
	left, right := jet.Siblings(testJet)

	err := jets.Update(ctx, testPulse, true, testJet)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, left, false)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, testPulse, right, false)
	require.NoError(t, err)

	err = jetKeeper.AddDropConfirmation(ctx, testPulse, right, false)
	require.NoError(t, err)

	err = jetKeeper.AddDropConfirmation(ctx, testPulse, left, false)
	require.NoError(t, err)

	err = jetKeeper.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)

	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())
}

func TestDbJetKeeper_AddDropConfirmation(t *testing.T) {
	ctx := inslogger.TestContext(t)
	jets := jet.NewDBStore(getPool())
	pulses := pulse.NewCalculatorMock(t)
	pulses.BackwardsMock.Set(func(p context.Context, p1 insolar.PulseNumber, p2 int) (r insolar.Pulse, r1 error) {
		return insolar.Pulse{PulseNumber: p1 - insolar.PulseNumber(p2)}, nil
	})
	jetKeeper := NewJetKeeper(jets, getPool(), pulses)

	var (
		pulse insolar.PulseNumber
		jet   insolar.JetID
	)
	f := fuzz.New()
	f.Fuzz(&pulse)
	f.Fuzz(&jet)
	err := jetKeeper.AddDropConfirmation(ctx, pulse, jet, false)
	require.NoError(t, err)
}

func TestDbJetKeeper_CheckJetTreeFail(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	ji, _, _ := initDB(t, testPulse)

	testJet := insolar.ZeroJetID

	err := ji.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())
	err = ji.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())
	err = ji.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, ji.TopSyncPulse())
	require.False(t, false, ji.HasAllJetConfirms(ctx, testPulse))
}

func TestDbJetKeeper_TopSyncPulse(t *testing.T) {
	ctx := inslogger.TestContext(t)
	jets := jet.NewDBStore(getPool())
	pulses := pulse.NewDB(getPool())
	err := pulses.Append(ctx, insolar.Pulse{PulseNumber: insolar.GenesisPulse.PulseNumber})
	require.NoError(t, err)

	jetKeeper := NewJetKeeper(jets, getPool(), pulses)

	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())

	var (
		currentPulse insolar.PulseNumber
		nextPulse    insolar.PulseNumber
		testJet      insolar.JetID
	)
	currentPulse = gen.PulseNumber()
	nextPulse = gen.PulseNumber()
	if nextPulse < currentPulse {
		currentPulse, nextPulse = nextPulse, currentPulse
	}
	testJet = insolar.ZeroJetID

	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: currentPulse})
	require.NoError(t, err)
	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: nextPulse})
	require.NoError(t, err)

	err = jets.Update(ctx, currentPulse, true, testJet)
	require.NoError(t, err)
	err = jetKeeper.AddDropConfirmation(ctx, currentPulse, testJet, false)
	require.NoError(t, err)
	// it's still top confirmed
	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())

	err = jetKeeper.AddHotConfirmation(ctx, currentPulse, testJet, false)
	require.NoError(t, err)
	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())

	err = jetKeeper.AddBackupConfirmation(ctx, currentPulse)
	require.NoError(t, err)
	require.Equal(t, currentPulse, jetKeeper.TopSyncPulse())

	err = jets.Clone(ctx, currentPulse, nextPulse, true)
	require.NoError(t, err)
	left, right := jet.Siblings(testJet)

	err = jetKeeper.AddDropConfirmation(ctx, nextPulse, testJet, true)
	require.NoError(t, err)

	err = jetKeeper.AddHotConfirmation(ctx, nextPulse, right, true)
	require.NoError(t, err)
	require.Equal(t, currentPulse, jetKeeper.TopSyncPulse())
	err = jetKeeper.AddHotConfirmation(ctx, nextPulse, left, true)
	require.NoError(t, err)
	require.Equal(t, currentPulse, jetKeeper.TopSyncPulse())

	err = jetKeeper.AddBackupConfirmation(ctx, nextPulse)
	require.NoError(t, err)
	require.Equal(t, nextPulse, jetKeeper.TopSyncPulse())
}

func TestDbJetKeeper_LostDataOnNextPulseAfterSplit(t *testing.T) {
	ctx := inslogger.TestContext(t)

	jets := jet.NewDBStore(getPool())
	pulses := pulse.NewDB(getPool())
	err := pulses.Append(ctx, insolar.Pulse{PulseNumber: insolar.GenesisPulse.PulseNumber})
	require.NoError(t, err)

	jetKeeper := NewJetKeeper(jets, getPool(), pulses)

	require.Equal(t, insolar.GenesisPulse.PulseNumber, jetKeeper.TopSyncPulse())

	var (
		currentPulse insolar.PulseNumber
		nextPulse    insolar.PulseNumber
		futurePulse  insolar.PulseNumber
		testJet      insolar.JetID
	)
	pulsesSlice := make([]insolar.PulseNumber, 3)
	for i := 0; i < len(pulsesSlice); i++ {
		pulsesSlice[i] = gen.PulseNumber()
	}
	sort.Slice(pulsesSlice, func(i, j int) bool {
		return pulsesSlice[i] < pulsesSlice[j]
	})

	currentPulse = pulsesSlice[0]
	nextPulse = pulsesSlice[1]
	futurePulse = pulsesSlice[2]
	testJet = insolar.ZeroJetID

	err = jets.Update(ctx, currentPulse, true, testJet)
	require.NoError(t, err)
	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: currentPulse})
	require.NoError(t, err)
	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: nextPulse})
	require.NoError(t, err)
	err = pulses.Append(ctx, insolar.Pulse{PulseNumber: futurePulse})
	require.NoError(t, err)

	// finalize currentPulse
	{
		err = jetKeeper.AddHotConfirmation(ctx, currentPulse, testJet, false)
		require.NoError(t, err)
		err = jetKeeper.AddDropConfirmation(ctx, currentPulse, testJet, false)
		require.NoError(t, err)
		require.True(t, jetKeeper.HasAllJetConfirms(ctx, currentPulse))
		err = jetKeeper.AddBackupConfirmation(ctx, currentPulse)
		require.Equal(t, currentPulse, jetKeeper.TopSyncPulse())
	}

	left, right := jet.Siblings(testJet)
	// finalize nextPulse
	{
		err = jets.Update(ctx, nextPulse, true, testJet)
		require.NoError(t, err)
		err = jetKeeper.AddDropConfirmation(ctx, nextPulse, testJet, true)
		require.NoError(t, err)
		err = jetKeeper.AddHotConfirmation(ctx, nextPulse, left, true)
		require.NoError(t, err)
		require.False(t, jetKeeper.HasAllJetConfirms(ctx, nextPulse))
		err = jetKeeper.AddHotConfirmation(ctx, nextPulse, right, true)
		require.NoError(t, err)

		require.True(t, jetKeeper.HasAllJetConfirms(ctx, currentPulse))
		require.True(t, jetKeeper.HasAllJetConfirms(ctx, nextPulse))
		err = jetKeeper.AddBackupConfirmation(ctx, nextPulse)
		require.NoError(t, err)
		require.Equal(t, nextPulse, jetKeeper.TopSyncPulse())
	}

	err = jets.Update(ctx, futurePulse, true, left)
	require.NoError(t, err)
	err = jetKeeper.AddDropConfirmation(ctx, futurePulse, left, false)
	require.NoError(t, err)
	err = jetKeeper.AddHotConfirmation(ctx, futurePulse, left, false)
	require.NoError(t, err)
	require.True(t, jetKeeper.HasAllJetConfirms(ctx, currentPulse))
	require.False(t, jetKeeper.HasAllJetConfirms(ctx, futurePulse))

	err = jets.Update(ctx, futurePulse, true, right)
	err = jetKeeper.AddDropConfirmation(ctx, futurePulse, right, false)
	require.NoError(t, err)
	err = jetKeeper.AddHotConfirmation(ctx, futurePulse, right, false)
	require.NoError(t, err)

	require.True(t, jetKeeper.HasAllJetConfirms(ctx, currentPulse))
	require.True(t, jetKeeper.HasAllJetConfirms(ctx, nextPulse))
	require.True(t, jetKeeper.HasAllJetConfirms(ctx, futurePulse))

	err = jetKeeper.AddBackupConfirmation(ctx, futurePulse)
	require.NoError(t, err)
	require.Equal(t, futurePulse, jetKeeper.TopSyncPulse())
}

func Test_TruncateHead(t *testing.T) {
	ctx := inslogger.TestContext(t)
	testPulse := gen.PulseNumber()
	nextPulse := gen.PulseNumber()
	if nextPulse < testPulse {
		testPulse, nextPulse = nextPulse, testPulse
	}
	ji_interface, jets, _ := initDB(t, testPulse)
	ji := ji_interface.(*DBJetKeeper)

	testJet := insolar.ZeroJetID

	err := jets.Update(ctx, testPulse, true, testJet)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	err = ji.AddDropConfirmation(ctx, testPulse, testJet, false)
	require.NoError(t, err)
	err = ji.AddBackupConfirmation(ctx, testPulse)
	require.NoError(t, err)

	require.Equal(t, testPulse, ji.TopSyncPulse())

	_, err = ji.get(testPulse)
	require.NoError(t, err)

	err = ji.AddDropConfirmation(ctx, nextPulse, gen.JetID(), false)
	require.NoError(t, err)
	err = ji.AddHotConfirmation(ctx, nextPulse, gen.JetID(), false)
	require.NoError(t, err)

	_, err = ji.get(nextPulse)
	require.NoError(t, err)

	err = ji.TruncateHead(ctx, nextPulse)
	require.NoError(t, err)

	_, err = ji.get(testPulse)
	require.NoError(t, err)
	_, err = ji.get(nextPulse)
	require.EqualError(t, err, "value not found")
}
