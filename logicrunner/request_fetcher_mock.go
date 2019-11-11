package logicrunner

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/logicrunner/common"
)

// RequestFetcherMock implements RequestFetcher
type RequestFetcherMock struct {
	t minimock.Tester

	funcAbort          func(ctx context.Context)
	inspectFuncAbort   func(ctx context.Context)
	afterAbortCounter  uint64
	beforeAbortCounter uint64
	AbortMock          mRequestFetcherMockAbort

	funcFetchPendings          func(ctx context.Context) (ch1 <-chan *common.Transcript)
	inspectFuncFetchPendings   func(ctx context.Context)
	afterFetchPendingsCounter  uint64
	beforeFetchPendingsCounter uint64
	FetchPendingsMock          mRequestFetcherMockFetchPendings
}

// NewRequestFetcherMock returns a mock for RequestFetcher
func NewRequestFetcherMock(t minimock.Tester) *RequestFetcherMock {
	m := &RequestFetcherMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AbortMock = mRequestFetcherMockAbort{mock: m}
	m.AbortMock.callArgs = []*RequestFetcherMockAbortParams{}

	m.FetchPendingsMock = mRequestFetcherMockFetchPendings{mock: m}
	m.FetchPendingsMock.callArgs = []*RequestFetcherMockFetchPendingsParams{}

	return m
}

type mRequestFetcherMockAbort struct {
	mock               *RequestFetcherMock
	defaultExpectation *RequestFetcherMockAbortExpectation
	expectations       []*RequestFetcherMockAbortExpectation

	callArgs []*RequestFetcherMockAbortParams
	mutex    sync.RWMutex
}

// RequestFetcherMockAbortExpectation specifies expectation struct of the RequestFetcher.Abort
type RequestFetcherMockAbortExpectation struct {
	mock   *RequestFetcherMock
	params *RequestFetcherMockAbortParams

	Counter uint64
}

// RequestFetcherMockAbortParams contains parameters of the RequestFetcher.Abort
type RequestFetcherMockAbortParams struct {
	ctx context.Context
}

// Expect sets up expected params for RequestFetcher.Abort
func (mmAbort *mRequestFetcherMockAbort) Expect(ctx context.Context) *mRequestFetcherMockAbort {
	if mmAbort.mock.funcAbort != nil {
		mmAbort.mock.t.Fatalf("RequestFetcherMock.Abort mock is already set by Set")
	}

	if mmAbort.defaultExpectation == nil {
		mmAbort.defaultExpectation = &RequestFetcherMockAbortExpectation{}
	}

	mmAbort.defaultExpectation.params = &RequestFetcherMockAbortParams{ctx}
	for _, e := range mmAbort.expectations {
		if minimock.Equal(e.params, mmAbort.defaultExpectation.params) {
			mmAbort.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAbort.defaultExpectation.params)
		}
	}

	return mmAbort
}

// Inspect accepts an inspector function that has same arguments as the RequestFetcher.Abort
func (mmAbort *mRequestFetcherMockAbort) Inspect(f func(ctx context.Context)) *mRequestFetcherMockAbort {
	if mmAbort.mock.inspectFuncAbort != nil {
		mmAbort.mock.t.Fatalf("Inspect function is already set for RequestFetcherMock.Abort")
	}

	mmAbort.mock.inspectFuncAbort = f

	return mmAbort
}

// Return sets up results that will be returned by RequestFetcher.Abort
func (mmAbort *mRequestFetcherMockAbort) Return() *RequestFetcherMock {
	if mmAbort.mock.funcAbort != nil {
		mmAbort.mock.t.Fatalf("RequestFetcherMock.Abort mock is already set by Set")
	}

	if mmAbort.defaultExpectation == nil {
		mmAbort.defaultExpectation = &RequestFetcherMockAbortExpectation{mock: mmAbort.mock}
	}

	return mmAbort.mock
}

//Set uses given function f to mock the RequestFetcher.Abort method
func (mmAbort *mRequestFetcherMockAbort) Set(f func(ctx context.Context)) *RequestFetcherMock {
	if mmAbort.defaultExpectation != nil {
		mmAbort.mock.t.Fatalf("Default expectation is already set for the RequestFetcher.Abort method")
	}

	if len(mmAbort.expectations) > 0 {
		mmAbort.mock.t.Fatalf("Some expectations are already set for the RequestFetcher.Abort method")
	}

	mmAbort.mock.funcAbort = f
	return mmAbort.mock
}

// Abort implements RequestFetcher
func (mmAbort *RequestFetcherMock) Abort(ctx context.Context) {
	mm_atomic.AddUint64(&mmAbort.beforeAbortCounter, 1)
	defer mm_atomic.AddUint64(&mmAbort.afterAbortCounter, 1)

	if mmAbort.inspectFuncAbort != nil {
		mmAbort.inspectFuncAbort(ctx)
	}

	mm_params := &RequestFetcherMockAbortParams{ctx}

	// Record call args
	mmAbort.AbortMock.mutex.Lock()
	mmAbort.AbortMock.callArgs = append(mmAbort.AbortMock.callArgs, mm_params)
	mmAbort.AbortMock.mutex.Unlock()

	for _, e := range mmAbort.AbortMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmAbort.AbortMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAbort.AbortMock.defaultExpectation.Counter, 1)
		mm_want := mmAbort.AbortMock.defaultExpectation.params
		mm_got := RequestFetcherMockAbortParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAbort.t.Errorf("RequestFetcherMock.Abort got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmAbort.funcAbort != nil {
		mmAbort.funcAbort(ctx)
		return
	}
	mmAbort.t.Fatalf("Unexpected call to RequestFetcherMock.Abort. %v", ctx)

}

// AbortAfterCounter returns a count of finished RequestFetcherMock.Abort invocations
func (mmAbort *RequestFetcherMock) AbortAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAbort.afterAbortCounter)
}

// AbortBeforeCounter returns a count of RequestFetcherMock.Abort invocations
func (mmAbort *RequestFetcherMock) AbortBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAbort.beforeAbortCounter)
}

// Calls returns a list of arguments used in each call to RequestFetcherMock.Abort.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAbort *mRequestFetcherMockAbort) Calls() []*RequestFetcherMockAbortParams {
	mmAbort.mutex.RLock()

	argCopy := make([]*RequestFetcherMockAbortParams, len(mmAbort.callArgs))
	copy(argCopy, mmAbort.callArgs)

	mmAbort.mutex.RUnlock()

	return argCopy
}

// MinimockAbortDone returns true if the count of the Abort invocations corresponds
// the number of defined expectations
func (m *RequestFetcherMock) MinimockAbortDone() bool {
	for _, e := range m.AbortMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AbortMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAbortCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAbort != nil && mm_atomic.LoadUint64(&m.afterAbortCounter) < 1 {
		return false
	}
	return true
}

// MinimockAbortInspect logs each unmet expectation
func (m *RequestFetcherMock) MinimockAbortInspect() {
	for _, e := range m.AbortMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestFetcherMock.Abort with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AbortMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAbortCounter) < 1 {
		if m.AbortMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestFetcherMock.Abort")
		} else {
			m.t.Errorf("Expected call to RequestFetcherMock.Abort with params: %#v", *m.AbortMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAbort != nil && mm_atomic.LoadUint64(&m.afterAbortCounter) < 1 {
		m.t.Error("Expected call to RequestFetcherMock.Abort")
	}
}

type mRequestFetcherMockFetchPendings struct {
	mock               *RequestFetcherMock
	defaultExpectation *RequestFetcherMockFetchPendingsExpectation
	expectations       []*RequestFetcherMockFetchPendingsExpectation

	callArgs []*RequestFetcherMockFetchPendingsParams
	mutex    sync.RWMutex
}

// RequestFetcherMockFetchPendingsExpectation specifies expectation struct of the RequestFetcher.FetchPendings
type RequestFetcherMockFetchPendingsExpectation struct {
	mock    *RequestFetcherMock
	params  *RequestFetcherMockFetchPendingsParams
	results *RequestFetcherMockFetchPendingsResults
	Counter uint64
}

// RequestFetcherMockFetchPendingsParams contains parameters of the RequestFetcher.FetchPendings
type RequestFetcherMockFetchPendingsParams struct {
	ctx context.Context
}

// RequestFetcherMockFetchPendingsResults contains results of the RequestFetcher.FetchPendings
type RequestFetcherMockFetchPendingsResults struct {
	ch1 <-chan *common.Transcript
}

// Expect sets up expected params for RequestFetcher.FetchPendings
func (mmFetchPendings *mRequestFetcherMockFetchPendings) Expect(ctx context.Context) *mRequestFetcherMockFetchPendings {
	if mmFetchPendings.mock.funcFetchPendings != nil {
		mmFetchPendings.mock.t.Fatalf("RequestFetcherMock.FetchPendings mock is already set by Set")
	}

	if mmFetchPendings.defaultExpectation == nil {
		mmFetchPendings.defaultExpectation = &RequestFetcherMockFetchPendingsExpectation{}
	}

	mmFetchPendings.defaultExpectation.params = &RequestFetcherMockFetchPendingsParams{ctx}
	for _, e := range mmFetchPendings.expectations {
		if minimock.Equal(e.params, mmFetchPendings.defaultExpectation.params) {
			mmFetchPendings.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFetchPendings.defaultExpectation.params)
		}
	}

	return mmFetchPendings
}

// Inspect accepts an inspector function that has same arguments as the RequestFetcher.FetchPendings
func (mmFetchPendings *mRequestFetcherMockFetchPendings) Inspect(f func(ctx context.Context)) *mRequestFetcherMockFetchPendings {
	if mmFetchPendings.mock.inspectFuncFetchPendings != nil {
		mmFetchPendings.mock.t.Fatalf("Inspect function is already set for RequestFetcherMock.FetchPendings")
	}

	mmFetchPendings.mock.inspectFuncFetchPendings = f

	return mmFetchPendings
}

// Return sets up results that will be returned by RequestFetcher.FetchPendings
func (mmFetchPendings *mRequestFetcherMockFetchPendings) Return(ch1 <-chan *common.Transcript) *RequestFetcherMock {
	if mmFetchPendings.mock.funcFetchPendings != nil {
		mmFetchPendings.mock.t.Fatalf("RequestFetcherMock.FetchPendings mock is already set by Set")
	}

	if mmFetchPendings.defaultExpectation == nil {
		mmFetchPendings.defaultExpectation = &RequestFetcherMockFetchPendingsExpectation{mock: mmFetchPendings.mock}
	}
	mmFetchPendings.defaultExpectation.results = &RequestFetcherMockFetchPendingsResults{ch1}
	return mmFetchPendings.mock
}

//Set uses given function f to mock the RequestFetcher.FetchPendings method
func (mmFetchPendings *mRequestFetcherMockFetchPendings) Set(f func(ctx context.Context) (ch1 <-chan *common.Transcript)) *RequestFetcherMock {
	if mmFetchPendings.defaultExpectation != nil {
		mmFetchPendings.mock.t.Fatalf("Default expectation is already set for the RequestFetcher.FetchPendings method")
	}

	if len(mmFetchPendings.expectations) > 0 {
		mmFetchPendings.mock.t.Fatalf("Some expectations are already set for the RequestFetcher.FetchPendings method")
	}

	mmFetchPendings.mock.funcFetchPendings = f
	return mmFetchPendings.mock
}

// When sets expectation for the RequestFetcher.FetchPendings which will trigger the result defined by the following
// Then helper
func (mmFetchPendings *mRequestFetcherMockFetchPendings) When(ctx context.Context) *RequestFetcherMockFetchPendingsExpectation {
	if mmFetchPendings.mock.funcFetchPendings != nil {
		mmFetchPendings.mock.t.Fatalf("RequestFetcherMock.FetchPendings mock is already set by Set")
	}

	expectation := &RequestFetcherMockFetchPendingsExpectation{
		mock:   mmFetchPendings.mock,
		params: &RequestFetcherMockFetchPendingsParams{ctx},
	}
	mmFetchPendings.expectations = append(mmFetchPendings.expectations, expectation)
	return expectation
}

// Then sets up RequestFetcher.FetchPendings return parameters for the expectation previously defined by the When method
func (e *RequestFetcherMockFetchPendingsExpectation) Then(ch1 <-chan *common.Transcript) *RequestFetcherMock {
	e.results = &RequestFetcherMockFetchPendingsResults{ch1}
	return e.mock
}

// FetchPendings implements RequestFetcher
func (mmFetchPendings *RequestFetcherMock) FetchPendings(ctx context.Context) (ch1 <-chan *common.Transcript) {
	mm_atomic.AddUint64(&mmFetchPendings.beforeFetchPendingsCounter, 1)
	defer mm_atomic.AddUint64(&mmFetchPendings.afterFetchPendingsCounter, 1)

	if mmFetchPendings.inspectFuncFetchPendings != nil {
		mmFetchPendings.inspectFuncFetchPendings(ctx)
	}

	mm_params := &RequestFetcherMockFetchPendingsParams{ctx}

	// Record call args
	mmFetchPendings.FetchPendingsMock.mutex.Lock()
	mmFetchPendings.FetchPendingsMock.callArgs = append(mmFetchPendings.FetchPendingsMock.callArgs, mm_params)
	mmFetchPendings.FetchPendingsMock.mutex.Unlock()

	for _, e := range mmFetchPendings.FetchPendingsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1
		}
	}

	if mmFetchPendings.FetchPendingsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFetchPendings.FetchPendingsMock.defaultExpectation.Counter, 1)
		mm_want := mmFetchPendings.FetchPendingsMock.defaultExpectation.params
		mm_got := RequestFetcherMockFetchPendingsParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFetchPendings.t.Errorf("RequestFetcherMock.FetchPendings got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFetchPendings.FetchPendingsMock.defaultExpectation.results
		if mm_results == nil {
			mmFetchPendings.t.Fatal("No results are set for the RequestFetcherMock.FetchPendings")
		}
		return (*mm_results).ch1
	}
	if mmFetchPendings.funcFetchPendings != nil {
		return mmFetchPendings.funcFetchPendings(ctx)
	}
	mmFetchPendings.t.Fatalf("Unexpected call to RequestFetcherMock.FetchPendings. %v", ctx)
	return
}

// FetchPendingsAfterCounter returns a count of finished RequestFetcherMock.FetchPendings invocations
func (mmFetchPendings *RequestFetcherMock) FetchPendingsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFetchPendings.afterFetchPendingsCounter)
}

// FetchPendingsBeforeCounter returns a count of RequestFetcherMock.FetchPendings invocations
func (mmFetchPendings *RequestFetcherMock) FetchPendingsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFetchPendings.beforeFetchPendingsCounter)
}

// Calls returns a list of arguments used in each call to RequestFetcherMock.FetchPendings.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFetchPendings *mRequestFetcherMockFetchPendings) Calls() []*RequestFetcherMockFetchPendingsParams {
	mmFetchPendings.mutex.RLock()

	argCopy := make([]*RequestFetcherMockFetchPendingsParams, len(mmFetchPendings.callArgs))
	copy(argCopy, mmFetchPendings.callArgs)

	mmFetchPendings.mutex.RUnlock()

	return argCopy
}

// MinimockFetchPendingsDone returns true if the count of the FetchPendings invocations corresponds
// the number of defined expectations
func (m *RequestFetcherMock) MinimockFetchPendingsDone() bool {
	for _, e := range m.FetchPendingsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FetchPendingsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFetchPendingsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFetchPendings != nil && mm_atomic.LoadUint64(&m.afterFetchPendingsCounter) < 1 {
		return false
	}
	return true
}

// MinimockFetchPendingsInspect logs each unmet expectation
func (m *RequestFetcherMock) MinimockFetchPendingsInspect() {
	for _, e := range m.FetchPendingsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestFetcherMock.FetchPendings with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FetchPendingsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFetchPendingsCounter) < 1 {
		if m.FetchPendingsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestFetcherMock.FetchPendings")
		} else {
			m.t.Errorf("Expected call to RequestFetcherMock.FetchPendings with params: %#v", *m.FetchPendingsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFetchPendings != nil && mm_atomic.LoadUint64(&m.afterFetchPendingsCounter) < 1 {
		m.t.Error("Expected call to RequestFetcherMock.FetchPendings")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RequestFetcherMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAbortInspect()

		m.MinimockFetchPendingsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RequestFetcherMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RequestFetcherMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAbortDone() &&
		m.MinimockFetchPendingsDone()
}
