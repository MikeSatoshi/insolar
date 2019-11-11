package executor

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

// DetachedNotifierMock implements DetachedNotifier
type DetachedNotifierMock struct {
	t minimock.Tester

	funcNotify          func(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID)
	inspectFuncNotify   func(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID)
	afterNotifyCounter  uint64
	beforeNotifyCounter uint64
	NotifyMock          mDetachedNotifierMockNotify
}

// NewDetachedNotifierMock returns a mock for DetachedNotifier
func NewDetachedNotifierMock(t minimock.Tester) *DetachedNotifierMock {
	m := &DetachedNotifierMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NotifyMock = mDetachedNotifierMockNotify{mock: m}
	m.NotifyMock.callArgs = []*DetachedNotifierMockNotifyParams{}

	return m
}

type mDetachedNotifierMockNotify struct {
	mock               *DetachedNotifierMock
	defaultExpectation *DetachedNotifierMockNotifyExpectation
	expectations       []*DetachedNotifierMockNotifyExpectation

	callArgs []*DetachedNotifierMockNotifyParams
	mutex    sync.RWMutex
}

// DetachedNotifierMockNotifyExpectation specifies expectation struct of the DetachedNotifier.Notify
type DetachedNotifierMockNotifyExpectation struct {
	mock   *DetachedNotifierMock
	params *DetachedNotifierMockNotifyParams

	Counter uint64
}

// DetachedNotifierMockNotifyParams contains parameters of the DetachedNotifier.Notify
type DetachedNotifierMockNotifyParams struct {
	ctx             context.Context
	openedRequests  []record.CompositeFilamentRecord
	objectID        insolar.ID
	closedRequestID insolar.ID
}

// Expect sets up expected params for DetachedNotifier.Notify
func (mmNotify *mDetachedNotifierMockNotify) Expect(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID) *mDetachedNotifierMockNotify {
	if mmNotify.mock.funcNotify != nil {
		mmNotify.mock.t.Fatalf("DetachedNotifierMock.Notify mock is already set by Set")
	}

	if mmNotify.defaultExpectation == nil {
		mmNotify.defaultExpectation = &DetachedNotifierMockNotifyExpectation{}
	}

	mmNotify.defaultExpectation.params = &DetachedNotifierMockNotifyParams{ctx, openedRequests, objectID, closedRequestID}
	for _, e := range mmNotify.expectations {
		if minimock.Equal(e.params, mmNotify.defaultExpectation.params) {
			mmNotify.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmNotify.defaultExpectation.params)
		}
	}

	return mmNotify
}

// Inspect accepts an inspector function that has same arguments as the DetachedNotifier.Notify
func (mmNotify *mDetachedNotifierMockNotify) Inspect(f func(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID)) *mDetachedNotifierMockNotify {
	if mmNotify.mock.inspectFuncNotify != nil {
		mmNotify.mock.t.Fatalf("Inspect function is already set for DetachedNotifierMock.Notify")
	}

	mmNotify.mock.inspectFuncNotify = f

	return mmNotify
}

// Return sets up results that will be returned by DetachedNotifier.Notify
func (mmNotify *mDetachedNotifierMockNotify) Return() *DetachedNotifierMock {
	if mmNotify.mock.funcNotify != nil {
		mmNotify.mock.t.Fatalf("DetachedNotifierMock.Notify mock is already set by Set")
	}

	if mmNotify.defaultExpectation == nil {
		mmNotify.defaultExpectation = &DetachedNotifierMockNotifyExpectation{mock: mmNotify.mock}
	}

	return mmNotify.mock
}

//Set uses given function f to mock the DetachedNotifier.Notify method
func (mmNotify *mDetachedNotifierMockNotify) Set(f func(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID)) *DetachedNotifierMock {
	if mmNotify.defaultExpectation != nil {
		mmNotify.mock.t.Fatalf("Default expectation is already set for the DetachedNotifier.Notify method")
	}

	if len(mmNotify.expectations) > 0 {
		mmNotify.mock.t.Fatalf("Some expectations are already set for the DetachedNotifier.Notify method")
	}

	mmNotify.mock.funcNotify = f
	return mmNotify.mock
}

// Notify implements DetachedNotifier
func (mmNotify *DetachedNotifierMock) Notify(ctx context.Context, openedRequests []record.CompositeFilamentRecord, objectID insolar.ID, closedRequestID insolar.ID) {
	mm_atomic.AddUint64(&mmNotify.beforeNotifyCounter, 1)
	defer mm_atomic.AddUint64(&mmNotify.afterNotifyCounter, 1)

	if mmNotify.inspectFuncNotify != nil {
		mmNotify.inspectFuncNotify(ctx, openedRequests, objectID, closedRequestID)
	}

	mm_params := &DetachedNotifierMockNotifyParams{ctx, openedRequests, objectID, closedRequestID}

	// Record call args
	mmNotify.NotifyMock.mutex.Lock()
	mmNotify.NotifyMock.callArgs = append(mmNotify.NotifyMock.callArgs, mm_params)
	mmNotify.NotifyMock.mutex.Unlock()

	for _, e := range mmNotify.NotifyMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmNotify.NotifyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmNotify.NotifyMock.defaultExpectation.Counter, 1)
		mm_want := mmNotify.NotifyMock.defaultExpectation.params
		mm_got := DetachedNotifierMockNotifyParams{ctx, openedRequests, objectID, closedRequestID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmNotify.t.Errorf("DetachedNotifierMock.Notify got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmNotify.funcNotify != nil {
		mmNotify.funcNotify(ctx, openedRequests, objectID, closedRequestID)
		return
	}
	mmNotify.t.Fatalf("Unexpected call to DetachedNotifierMock.Notify. %v %v %v %v", ctx, openedRequests, objectID, closedRequestID)

}

// NotifyAfterCounter returns a count of finished DetachedNotifierMock.Notify invocations
func (mmNotify *DetachedNotifierMock) NotifyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNotify.afterNotifyCounter)
}

// NotifyBeforeCounter returns a count of DetachedNotifierMock.Notify invocations
func (mmNotify *DetachedNotifierMock) NotifyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNotify.beforeNotifyCounter)
}

// Calls returns a list of arguments used in each call to DetachedNotifierMock.Notify.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmNotify *mDetachedNotifierMockNotify) Calls() []*DetachedNotifierMockNotifyParams {
	mmNotify.mutex.RLock()

	argCopy := make([]*DetachedNotifierMockNotifyParams, len(mmNotify.callArgs))
	copy(argCopy, mmNotify.callArgs)

	mmNotify.mutex.RUnlock()

	return argCopy
}

// MinimockNotifyDone returns true if the count of the Notify invocations corresponds
// the number of defined expectations
func (m *DetachedNotifierMock) MinimockNotifyDone() bool {
	for _, e := range m.NotifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NotifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNotifyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNotify != nil && mm_atomic.LoadUint64(&m.afterNotifyCounter) < 1 {
		return false
	}
	return true
}

// MinimockNotifyInspect logs each unmet expectation
func (m *DetachedNotifierMock) MinimockNotifyInspect() {
	for _, e := range m.NotifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DetachedNotifierMock.Notify with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NotifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNotifyCounter) < 1 {
		if m.NotifyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DetachedNotifierMock.Notify")
		} else {
			m.t.Errorf("Expected call to DetachedNotifierMock.Notify with params: %#v", *m.NotifyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNotify != nil && mm_atomic.LoadUint64(&m.afterNotifyCounter) < 1 {
		m.t.Error("Expected call to DetachedNotifierMock.Notify")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DetachedNotifierMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockNotifyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DetachedNotifierMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *DetachedNotifierMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNotifyDone()
}
