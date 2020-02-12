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

// +build ignore

package logbuf

import "io"

var _ io.Writer = &PlainBuffer{}
var _ io.WriterTo = &PlainBuffer{}

type MissedFunc func(missed uint32)

func NewPlainBuffer(buffer PagedBuffer, idleFn func(), missedFn MissedFunc) PlainBuffer {
	return PlainBuffer{
		buffer:   buffer,
		idleFn:   idleFn,
		missedFn: missedFn,
	}
}

type PlainBuffer struct {
	buffer   PagedBuffer
	idleFn   func()
	missedFn MissedFunc
}

func (p *PlainBuffer) WriteTo(w io.Writer) (int64, error) {
	pg := p.buffer.FlushPages()
	pg.StartAccess(p.idleFn)
	defer pg.StopAccess()

	c := pg.Count()
	totalN := int64(0)
	for i := 0; i < c; i++ {
		writes, buf := pg.Page(i)
		if buf == nil {
			if writes > 0 && p.missedFn != nil {
				p.missedFn(writes)
			}
			continue
		}
		nc, err := w.Write(buf)
		totalN += int64(nc)
		if err != nil {
			return totalN, err
		}
	}
	return totalN, nil
}

func (p *PlainBuffer) Write(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	pg, buf := p.buffer.allocateBuffer(uint32(len(b)))
	copy(buf, b)
	pg.stopAccess()
	return len(b), nil
}
