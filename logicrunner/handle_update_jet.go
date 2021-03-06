// Copyright 2020 Insolar Network Ltd.
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

package logicrunner

import (
	"context"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/payload"
)

type HandleUpdateJet struct {
	dep *Dependencies

	meta payload.Meta
}

func (h *HandleUpdateJet) Present(ctx context.Context, _ flow.Flow) error {
	pl := payload.UpdateJet{}
	err := pl.Unmarshal(h.meta.Payload)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal payload")
	}

	err = h.dep.JetStorage.Update(ctx, pl.Pulse, true, pl.JetID)
	if err != nil {
		return errors.Wrap(err, "failed to update jets")
	}

	return nil
}
