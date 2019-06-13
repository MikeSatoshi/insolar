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

package object

import (
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

//go:generate minimock -i github.com/insolar/insolar/ledger/object.IndexCleaner -o ./ -s _mock.go

// IndexCleaner provides an interface for removing backets from a storage.
type IndexCleaner interface {
	// DeleteForPN method removes indexes from a storage for a provided
	DeleteForPN(ctx context.Context, pn insolar.PulseNumber)
}

//go:generate minimock -i github.com/insolar/insolar/ledger/object.IndexBucketModifier -o ./ -s _mock.go

// IndexBucketModifier provides methods for modifying buckets of index.
// Index contains buckets with pn->objID->Bucket hierarchy.
// With using of IndexBucketModifier there is a possibility to set buckets from outside of an index.
type IndexBucketModifier interface {
	// SetBucket adds a bucket with provided pulseNumber and ID
	SetBucket(ctx context.Context, pn insolar.PulseNumber, bucket FilamentIndex) error
}

//go:generate minimock -i github.com/insolar/insolar/ledger/object.IndexBucketAccessor -o ./ -s _mock.go

// IndexBucketAccessor provides an interface for fetching buckets from an index.
type IndexBucketAccessor interface {
	// ForPNAndJet returns a collection of buckets for a provided pn and jetID
	ForPNAndJet(ctx context.Context, pn insolar.PulseNumber, jetID insolar.JetID) []FilamentIndex
}

//go:generate minimock -i github.com/insolar/insolar/ledger/object.PendingModifier -o ./ -s _mock.go

// PendingModifier provides methods for modifying pending requests
type PendingModifier interface {
	// SetRequest sets a request for a specific object
	SetRequest(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, jetID insolar.JetID, reqID insolar.ID) error
	// SetResult sets a result for a specific object. Also, if there is a not closed request for a provided result,
	// the request will be closed
	SetResult(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, jetID insolar.JetID, resID insolar.ID, res record.Result) error
	// SetFilament adds a slice of records to an object with provided id and pulse. It's assumed, that the method is
	// called for setting records from another light, during the process of filling full chaing of pendings
	SetFilament(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, filPN insolar.PulseNumber, recs []record.CompositeFilamentRecord) error
}

//go:generate minimock -i github.com/insolar/insolar/ledger/object.PendingFilamentStateAccessor -o ./ -s _mock.go

type PendingFilamentStateAccessor interface {
	WaitForRefresh(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (<-chan struct{}, error)
}

type PendingFilamentStateModifier interface {
	// RefreshState recalculates state of the chain, marks requests as closed and opened.
	RefreshState(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) error
	ExpireRequests(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID, jetID insolar.JetID, reqs []insolar.ID) error
}

//go:generate minimock -i github.com/insolar/insolar/ledger/object.PendingAccessor -o ./ -s _mock.go

// PendingAccessor provides methods for fetching pending requests.
type PendingAccessor interface {
	// OpenRequestsForObjID returns a specific number of open requests for a specific object
	OpenRequestsForObjID(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID, count int) ([]record.Request, error)
	// Records returns all the records for a provided object
	Records(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID) ([]record.CompositeFilamentRecord, error)
	FirstPending(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID) (*record.PendingFilament, error)
}

type HeavyPendingAccessor interface {
	// AllOpenRequestsForObjID returns all open requests for the provided object and pulse
	AllOpenRequestsForObjID(ctx context.Context, currentPN insolar.PulseNumber, objID insolar.ID) ([]insolar.ID, error)
	Records(ctx context.Context, readFrom insolar.PulseNumber, readUntil insolar.PulseNumber, objID insolar.ID) ([]record.CompositeFilamentRecord, error)
}
