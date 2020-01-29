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

package pulse

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/instrumentation/inslogger"
)

// DB is a pulse.DB storage implementation. It saves pulses to PostgreSQL and does not allow removal.
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates new DB storage instance.
func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

func (s *DB) selectPulse(ctx context.Context, tx pgx.Tx, pn insolar.PulseNumber) (retPulse insolar.Pulse, retErr error) {
	pulseRow := tx.QueryRow(ctx,
		"SELECT pulse_number, prev_pn, next_pn, tstamp, epoch, origin_id, entropy FROM pulses WHERE pulse_number = $1",
		pn)

	retPulse.Signs = make(map[string]insolar.PulseSenderConfirmation)
	var originSlice []byte
	var entropySlice []byte
	err := pulseRow.Scan(
		&retPulse.PulseNumber,
		&retPulse.PrevPulseNumber,
		&retPulse.NextPulseNumber,
		&retPulse.PulseTimestamp,
		&retPulse.EpochPulseNumber,
		&originSlice,
		&entropySlice)

	if err == pgx.ErrNoRows {
		retErr = ErrNotFound
		_ = tx.Rollback(ctx)
		return
	}

	if err != nil {
		retErr = errors.Wrap(err, "Unable to SELECT ... FROM pulses")
		_ = tx.Rollback(ctx)
		return
	}

	copy(retPulse.OriginID[:], originSlice)
	copy(retPulse.Entropy[:], entropySlice)

	signRows, err := tx.Query(ctx,
		"SELECT pulse_number, chosen_public_key, entropy, signature FROM pulse_signs WHERE pulse_number = $1",
		pn)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to SELECT ... FROM pulse_signs")
		_ = tx.Rollback(ctx)
		return
	}
	defer signRows.Close()

	for signRows.Next() {
		var conf insolar.PulseSenderConfirmation
		err = signRows.Scan(&conf.PulseNumber, &conf.ChosenPublicKey, &entropySlice, &conf.Signature)
		if err != nil {
			retErr = errors.Wrap(err, "Unable to scan another pulse_signs row")
			_ = tx.Rollback(ctx)
			return
		}
		copy(conf.Entropy[:], entropySlice)
		retPulse.Signs[conf.ChosenPublicKey] = conf
	}

	return
}

func (s *DB) selectByCondition(ctx context.Context, query string, args ...interface{}) (retPulse insolar.Pulse, retErr error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to acquire a database connection")
		return
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, insolar.PGReadTxOptions)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to start a read transaction")
		return
	}

	var pn insolar.PulseNumber
	row := tx.QueryRow(ctx, query, args...)
	err = row.Scan(&pn)
	if err == pgx.ErrNoRows {
		_ = tx.Rollback(ctx)
		retErr = ErrNotFound
		return
	}
	if err != nil {
		retErr = errors.Wrapf(err, "selectByCondition - request failed query = `%v`, args = %v", query, args)
		_ = tx.Rollback(ctx)
		return
	}

	retPulse, retErr = s.selectPulse(ctx, tx, pn)
	if retErr != nil {
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to commit read transaction. If you see this consider adding a retry or lower the isolation level!")
		return
	}

	return
}

// ForPulseNumber returns pulse for provided a pulse number. If not found, ErrNotFound will be returned.
func (s *DB) ForPulseNumber(ctx context.Context, pn insolar.PulseNumber) (retPulse insolar.Pulse, retErr error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to acquire a database connection")
		return
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, insolar.PGReadTxOptions)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to start a read transaction")
		return
	}

	retPulse, retErr = s.selectPulse(ctx, tx, pn)
	if retErr != nil {
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		retErr = errors.Wrap(err, "Unable to commit read transaction. If you see this consider adding a retry or lower the isolation level!")
		return
	}

	return
}

// Latest returns a latest pulse saved in DB. If not found, ErrNotFound will be returned.
func (s *DB) Latest(ctx context.Context) (retPulse insolar.Pulse, retErr error) {
	retPulse, retErr = s.selectByCondition(ctx, "SELECT max(pulse_number) as latest FROM pulses")
	return
}

// TruncateHead remove all records with pulse_number >= `from`
func (s *DB) TruncateHead(ctx context.Context, from insolar.PulseNumber) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "Unable to acquire a database connection")
	}
	defer conn.Release()

	log := inslogger.FromContext(ctx)

	for { // retry loop
		tx, err := conn.BeginTx(ctx, insolar.PGWriteTxOptions)
		if err != nil {
			return errors.Wrap(err, "Unable to start a write transaction")
		}

		_, err = tx.Exec(ctx, "DELETE FROM pulse_signs WHERE pulse_number >= $1", from)
		if err != nil {
			_ = tx.Rollback(ctx)
			return errors.Wrap(err, "Unable to DELETE FROM pulse_signs")
		}

		_, err = tx.Exec(ctx, "DELETE FROM pulses WHERE pulse_number >= $1", from)
		if err != nil {
			_ = tx.Rollback(ctx)
			return errors.Wrap(err, "Unable to DELETE FROM pulses")
		}

		err = tx.Commit(ctx)
		if err == nil { // success
			break
		}

		log.Infof("TruncateHead - commit failed: %v - retrying transaction", err)
	}

	return nil
}

// Append appends provided pulse to current storage. Pulse number should be greater than currently saved for preserving
// pulse consistency. If a provided pulse does not meet the requirements, ErrBadPulse will be returned.
func (s *DB) Append(ctx context.Context, pulse insolar.Pulse) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "Unable to acquire a database connection")
	}
	defer conn.Release()

	log := inslogger.FromContext(ctx)

	for { // retry loop
		tx, err := conn.BeginTx(ctx, insolar.PGWriteTxOptions)
		if err != nil {
			return errors.Wrap(err, "Unable to start a write transaction")
		}

		_, err = tx.Exec(ctx, `
			INSERT INTO pulses(pulse_number, prev_pn, next_pn, tstamp, epoch, origin_id, entropy)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, pulse.PulseNumber, pulse.PrevPulseNumber, pulse.NextPulseNumber, pulse.PulseTimestamp,
			pulse.EpochPulseNumber, pulse.OriginID[:], pulse.Entropy[:])
		if err != nil {
			_ = tx.Rollback(ctx)
			return errors.Wrap(err, "Unable to INSERT pulse")
		}

		for k, s := range pulse.Signs {
			if (s.PulseNumber != pulse.PulseNumber) || (k != s.ChosenPublicKey) {
				_ = tx.Rollback(ctx)
				return ErrBadPulse
			}

			_, err = tx.Exec(ctx, `
				INSERT INTO pulse_signs (pulse_number, chosen_public_key, entropy, signature)
				VALUES ($1, $2, $3, $4)
			`, s.PulseNumber, s.ChosenPublicKey, s.Entropy[:], s.Signature)
			if err != nil {
				_ = tx.Rollback(ctx)
				return errors.Wrap(err, "Unable to INSERT pulse_sign")
			}
		}

		err = tx.Commit(ctx)
		if err == nil { // success
			break
		}

		log.Infof("Append - commit failed: %v - retrying transaction", err)
	}

	return nil
}

// Forwards calculates steps pulses forwards from provided pulse. If calculated pulse does not exist, ErrNotFound will
// be returned.
func (s *DB) Forwards(ctx context.Context, pn insolar.PulseNumber, steps int) (retPulse insolar.Pulse, retErr error) {
	// There can be "holes" in pulses double-linked list, e.g.
	// 1) Between fake genesis pulse and first real pulse
	// 2) If pulsar is separated from the rest of the network for N pulses
	// 3) The platform was down for N pulses
	// Thus we can't use recursive queries here. In the future we are
	// going to refactor the entire pulses logic.
	retPulse, retErr = s.selectByCondition(ctx, `
SELECT pulse_number FROM pulses WHERE pulse_number >= $1 ORDER BY pulse_number asc OFFSET $2 LIMIT 1;
	`, pn, steps)
	return
}

// Backwards calculates steps pulses backwards from provided pulse. If calculated pulse does not exist, ErrNotFound will
// be returned.
func (s *DB) Backwards(ctx context.Context, pn insolar.PulseNumber, steps int) (retPulse insolar.Pulse, retErr error) {
	retPulse, retErr = s.selectByCondition(ctx, `
SELECT pulse_number FROM pulses WHERE pulse_number <= $1 ORDER BY pulse_number desc offset $2 limit 1;
	`, pn, steps)
	return
}
