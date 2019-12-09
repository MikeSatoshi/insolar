///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package jetid

import (
	"fmt"

	"github.com/insolar/insolar/pulse"
)

type JetPrefix = Prefix

type ShortJetId uint32 // JetPrefix + 5bit length

const bitsShortJetIdLen = 5

func (v ShortJetId) Prefix() JetPrefix {
	return JetPrefix(v & ((^ShortJetId(0)) >> bitsShortJetIdLen))
}

func (v ShortJetId) PrefixLength() (uint8, bool) {
	if n := uint8(v >> (32 - bitsShortJetIdLen)); n > 0 {
		return n - 1, true
	}
	return 0, false
}

func (v ShortJetId) HasLength() bool {
	_, ok := v.PrefixLength()
	return ok
}

func (v ShortJetId) String() string {
	if n, ok := v.PrefixLength(); ok {
		return fmt.Sprintf("0x%02X[%d]", v.Prefix(), n)
	} else {
		return fmt.Sprintf("0x%02X[]", v.Prefix())
	}
}

type FullJetId uint64 // ShortJetId + LastSplitPulse

func (v FullJetId) IsValid() bool {
	_, ok := ShortJetId(v).PrefixLength()
	return ok && pulse.IsValidAsPulseNumber(int(v>>32))
}

func (v FullJetId) ShortId() ShortJetId {
	return ShortJetId(v)
}

func (v FullJetId) CreatedAt() pulse.Number {
	return pulse.OfUint32(uint32(v >> 32))
}

func (v FullJetId) String() string {
	return fmt.Sprintf("%v@%d", v.ShortId(), v.CreatedAt())
}