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

package handle

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/ledger/light/proc"
)

type GetJet struct {
	dep    *proc.Dependencies
	meta   payload.Meta
	passed bool
}

func NewGetJet(dep *proc.Dependencies, meta payload.Meta, passed bool) *GetJet {
	return &GetJet{
		dep:    dep,
		meta:   meta,
		passed: passed,
	}
}

func (h *GetJet) Present(ctx context.Context, f flow.Flow) error {
	pl, err := payload.Unmarshal(h.meta.Payload)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal GetJet message")
	}
	msg, ok := pl.(*payload.GetJet)
	if !ok {
		return fmt.Errorf("wrong request type: %T", pl)
	}

	getJet := proc.NewGetJet(h.meta, msg.ObjectID, msg.PulseNumber)
	h.dep.GetJet(getJet)
	return f.Procedure(ctx, getJet, false)
}
