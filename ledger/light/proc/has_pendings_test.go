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

package proc_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/light/proc"
	"github.com/insolar/insolar/ledger/object"
	"github.com/insolar/insolar/pulse"
)

func TestHasPendings_Proceed(t *testing.T) {
	ctx := flow.TestContextWithPulse(inslogger.TestContext(t), pulse.MinTimePulse+10)
	mc := minimock.NewController(t)

	var (
		index  *object.IndexAccessorMock
		sender *bus.SenderMock
	)

	setup := func() {
		index = object.NewIndexAccessorMock(mc)
		sender = bus.NewSenderMock(mc)
	}

	t.Run("ok, has pendings", func(t *testing.T) {
		setup()
		defer mc.Finish()

		pulseNumber := insolar.NewID(pulse.MinTimePulse, []byte{1}).Pulse()

		index.ForIDMock.Return(
			record.Index{
				Lifeline: record.Lifeline{
					EarliestOpenRequest: &pulseNumber,
				},
			},
			nil,
		)

		expectedMsg, _ := payload.NewMessage(&payload.PendingsInfo{
			HasPendings: true,
		})

		sender.ReplyMock.Inspect(func(ctx context.Context, origin payload.Meta, reply *message.Message) {
			assert.Equal(t, expectedMsg.Payload, reply.Payload)
		}).Return()

		p := proc.NewHasPendings(payload.Meta{}, gen.ID())
		p.Dep(index, sender)

		err := p.Proceed(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok, no pendings", func(t *testing.T) {
		setup()
		defer mc.Finish()

		pulseNumber := insolar.NewID(pulse.MinTimePulse+100, []byte{1}).Pulse()

		index.ForIDMock.Return(
			record.Index{
				Lifeline: record.Lifeline{
					EarliestOpenRequest: &pulseNumber,
				},
			},
			nil,
		)

		expectedMsg, _ := payload.NewMessage(&payload.PendingsInfo{
			HasPendings: false,
		})

		sender.ReplyMock.Inspect(func(ctx context.Context, origin payload.Meta, reply *message.Message) {
			assert.Equal(t, expectedMsg.Payload, reply.Payload)
		}).Return()

		p := proc.NewHasPendings(payload.Meta{}, gen.ID())
		p.Dep(index, sender)

		err := p.Proceed(ctx)
		assert.NoError(t, err)
	})
}
