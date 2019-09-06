//
// Copyright 2019 Insolar Technologies GmbH
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

// +build functest

package functest

import (
	"testing"

	"github.com/insolar/insolar/testutils/launchnet"
)

var functestCount int

// TestRotateLogs rotates launchnet logs (removes and reopen it).
// Should be always 'first' test in package.
func TestRotateLogs(t *testing.T) {
	functestCount++
	t.Log("functest iteration:", functestCount)
	if !launchnet.LogRotateEnabled() {
		t.Skip("log rotate disabled")
	}
	// BEWARE: it removes files by pattern!
	launchnet.RotateLogs("../.artifacts/launchnet/logs/*/*/*.log", true)
}