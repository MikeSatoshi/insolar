package longbits

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"io"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// FixedReaderMock implements FixedReader
type FixedReaderMock struct {
	t minimock.Tester

	funcAsByteString          func() (b1 ByteString)
	inspectFuncAsByteString   func()
	afterAsByteStringCounter  uint64
	beforeAsByteStringCounter uint64
	AsByteStringMock          mFixedReaderMockAsByteString

	funcAsBytes          func() (ba1 []byte)
	inspectFuncAsBytes   func()
	afterAsBytesCounter  uint64
	beforeAsBytesCounter uint64
	AsBytesMock          mFixedReaderMockAsBytes

	funcFixedByteSize          func() (i1 int)
	inspectFuncFixedByteSize   func()
	afterFixedByteSizeCounter  uint64
	beforeFixedByteSizeCounter uint64
	FixedByteSizeMock          mFixedReaderMockFixedByteSize

	funcRead          func(p []byte) (n int, err error)
	inspectFuncRead   func(p []byte)
	afterReadCounter  uint64
	beforeReadCounter uint64
	ReadMock          mFixedReaderMockRead

	funcWriteTo          func(w io.Writer) (n int64, err error)
	inspectFuncWriteTo   func(w io.Writer)
	afterWriteToCounter  uint64
	beforeWriteToCounter uint64
	WriteToMock          mFixedReaderMockWriteTo
}

// NewFixedReaderMock returns a mock for FixedReader
func NewFixedReaderMock(t minimock.Tester) *FixedReaderMock {
	m := &FixedReaderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AsByteStringMock = mFixedReaderMockAsByteString{mock: m}

	m.AsBytesMock = mFixedReaderMockAsBytes{mock: m}

	m.FixedByteSizeMock = mFixedReaderMockFixedByteSize{mock: m}

	m.ReadMock = mFixedReaderMockRead{mock: m}
	m.ReadMock.callArgs = []*FixedReaderMockReadParams{}

	m.WriteToMock = mFixedReaderMockWriteTo{mock: m}
	m.WriteToMock.callArgs = []*FixedReaderMockWriteToParams{}

	return m
}

type mFixedReaderMockAsByteString struct {
	mock               *FixedReaderMock
	defaultExpectation *FixedReaderMockAsByteStringExpectation
	expectations       []*FixedReaderMockAsByteStringExpectation
}

// FixedReaderMockAsByteStringExpectation specifies expectation struct of the FixedReader.AsByteString
type FixedReaderMockAsByteStringExpectation struct {
	mock *FixedReaderMock

	results *FixedReaderMockAsByteStringResults
	Counter uint64
}

// FixedReaderMockAsByteStringResults contains results of the FixedReader.AsByteString
type FixedReaderMockAsByteStringResults struct {
	b1 ByteString
}

// Expect sets up expected params for FixedReader.AsByteString
func (mmAsByteString *mFixedReaderMockAsByteString) Expect() *mFixedReaderMockAsByteString {
	if mmAsByteString.mock.funcAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("FixedReaderMock.AsByteString mock is already set by Set")
	}

	if mmAsByteString.defaultExpectation == nil {
		mmAsByteString.defaultExpectation = &FixedReaderMockAsByteStringExpectation{}
	}

	return mmAsByteString
}

// Inspect accepts an inspector function that has same arguments as the FixedReader.AsByteString
func (mmAsByteString *mFixedReaderMockAsByteString) Inspect(f func()) *mFixedReaderMockAsByteString {
	if mmAsByteString.mock.inspectFuncAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("Inspect function is already set for FixedReaderMock.AsByteString")
	}

	mmAsByteString.mock.inspectFuncAsByteString = f

	return mmAsByteString
}

// Return sets up results that will be returned by FixedReader.AsByteString
func (mmAsByteString *mFixedReaderMockAsByteString) Return(b1 ByteString) *FixedReaderMock {
	if mmAsByteString.mock.funcAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("FixedReaderMock.AsByteString mock is already set by Set")
	}

	if mmAsByteString.defaultExpectation == nil {
		mmAsByteString.defaultExpectation = &FixedReaderMockAsByteStringExpectation{mock: mmAsByteString.mock}
	}
	mmAsByteString.defaultExpectation.results = &FixedReaderMockAsByteStringResults{b1}
	return mmAsByteString.mock
}

//Set uses given function f to mock the FixedReader.AsByteString method
func (mmAsByteString *mFixedReaderMockAsByteString) Set(f func() (b1 ByteString)) *FixedReaderMock {
	if mmAsByteString.defaultExpectation != nil {
		mmAsByteString.mock.t.Fatalf("Default expectation is already set for the FixedReader.AsByteString method")
	}

	if len(mmAsByteString.expectations) > 0 {
		mmAsByteString.mock.t.Fatalf("Some expectations are already set for the FixedReader.AsByteString method")
	}

	mmAsByteString.mock.funcAsByteString = f
	return mmAsByteString.mock
}

// AsByteString implements FixedReader
func (mmAsByteString *FixedReaderMock) AsByteString() (b1 ByteString) {
	mm_atomic.AddUint64(&mmAsByteString.beforeAsByteStringCounter, 1)
	defer mm_atomic.AddUint64(&mmAsByteString.afterAsByteStringCounter, 1)

	if mmAsByteString.inspectFuncAsByteString != nil {
		mmAsByteString.inspectFuncAsByteString()
	}

	if mmAsByteString.AsByteStringMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAsByteString.AsByteStringMock.defaultExpectation.Counter, 1)

		mm_results := mmAsByteString.AsByteStringMock.defaultExpectation.results
		if mm_results == nil {
			mmAsByteString.t.Fatal("No results are set for the FixedReaderMock.AsByteString")
		}
		return (*mm_results).b1
	}
	if mmAsByteString.funcAsByteString != nil {
		return mmAsByteString.funcAsByteString()
	}
	mmAsByteString.t.Fatalf("Unexpected call to FixedReaderMock.AsByteString.")
	return
}

// AsByteStringAfterCounter returns a count of finished FixedReaderMock.AsByteString invocations
func (mmAsByteString *FixedReaderMock) AsByteStringAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsByteString.afterAsByteStringCounter)
}

// AsByteStringBeforeCounter returns a count of FixedReaderMock.AsByteString invocations
func (mmAsByteString *FixedReaderMock) AsByteStringBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsByteString.beforeAsByteStringCounter)
}

// MinimockAsByteStringDone returns true if the count of the AsByteString invocations corresponds
// the number of defined expectations
func (m *FixedReaderMock) MinimockAsByteStringDone() bool {
	for _, e := range m.AsByteStringMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsByteStringMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsByteString != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		return false
	}
	return true
}

// MinimockAsByteStringInspect logs each unmet expectation
func (m *FixedReaderMock) MinimockAsByteStringInspect() {
	for _, e := range m.AsByteStringMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FixedReaderMock.AsByteString")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsByteStringMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.AsByteString")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsByteString != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.AsByteString")
	}
}

type mFixedReaderMockAsBytes struct {
	mock               *FixedReaderMock
	defaultExpectation *FixedReaderMockAsBytesExpectation
	expectations       []*FixedReaderMockAsBytesExpectation
}

// FixedReaderMockAsBytesExpectation specifies expectation struct of the FixedReader.AsBytes
type FixedReaderMockAsBytesExpectation struct {
	mock *FixedReaderMock

	results *FixedReaderMockAsBytesResults
	Counter uint64
}

// FixedReaderMockAsBytesResults contains results of the FixedReader.AsBytes
type FixedReaderMockAsBytesResults struct {
	ba1 []byte
}

// Expect sets up expected params for FixedReader.AsBytes
func (mmAsBytes *mFixedReaderMockAsBytes) Expect() *mFixedReaderMockAsBytes {
	if mmAsBytes.mock.funcAsBytes != nil {
		mmAsBytes.mock.t.Fatalf("FixedReaderMock.AsBytes mock is already set by Set")
	}

	if mmAsBytes.defaultExpectation == nil {
		mmAsBytes.defaultExpectation = &FixedReaderMockAsBytesExpectation{}
	}

	return mmAsBytes
}

// Inspect accepts an inspector function that has same arguments as the FixedReader.AsBytes
func (mmAsBytes *mFixedReaderMockAsBytes) Inspect(f func()) *mFixedReaderMockAsBytes {
	if mmAsBytes.mock.inspectFuncAsBytes != nil {
		mmAsBytes.mock.t.Fatalf("Inspect function is already set for FixedReaderMock.AsBytes")
	}

	mmAsBytes.mock.inspectFuncAsBytes = f

	return mmAsBytes
}

// Return sets up results that will be returned by FixedReader.AsBytes
func (mmAsBytes *mFixedReaderMockAsBytes) Return(ba1 []byte) *FixedReaderMock {
	if mmAsBytes.mock.funcAsBytes != nil {
		mmAsBytes.mock.t.Fatalf("FixedReaderMock.AsBytes mock is already set by Set")
	}

	if mmAsBytes.defaultExpectation == nil {
		mmAsBytes.defaultExpectation = &FixedReaderMockAsBytesExpectation{mock: mmAsBytes.mock}
	}
	mmAsBytes.defaultExpectation.results = &FixedReaderMockAsBytesResults{ba1}
	return mmAsBytes.mock
}

//Set uses given function f to mock the FixedReader.AsBytes method
func (mmAsBytes *mFixedReaderMockAsBytes) Set(f func() (ba1 []byte)) *FixedReaderMock {
	if mmAsBytes.defaultExpectation != nil {
		mmAsBytes.mock.t.Fatalf("Default expectation is already set for the FixedReader.AsBytes method")
	}

	if len(mmAsBytes.expectations) > 0 {
		mmAsBytes.mock.t.Fatalf("Some expectations are already set for the FixedReader.AsBytes method")
	}

	mmAsBytes.mock.funcAsBytes = f
	return mmAsBytes.mock
}

// AsBytes implements FixedReader
func (mmAsBytes *FixedReaderMock) AsBytes() (ba1 []byte) {
	mm_atomic.AddUint64(&mmAsBytes.beforeAsBytesCounter, 1)
	defer mm_atomic.AddUint64(&mmAsBytes.afterAsBytesCounter, 1)

	if mmAsBytes.inspectFuncAsBytes != nil {
		mmAsBytes.inspectFuncAsBytes()
	}

	if mmAsBytes.AsBytesMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAsBytes.AsBytesMock.defaultExpectation.Counter, 1)

		mm_results := mmAsBytes.AsBytesMock.defaultExpectation.results
		if mm_results == nil {
			mmAsBytes.t.Fatal("No results are set for the FixedReaderMock.AsBytes")
		}
		return (*mm_results).ba1
	}
	if mmAsBytes.funcAsBytes != nil {
		return mmAsBytes.funcAsBytes()
	}
	mmAsBytes.t.Fatalf("Unexpected call to FixedReaderMock.AsBytes.")
	return
}

// AsBytesAfterCounter returns a count of finished FixedReaderMock.AsBytes invocations
func (mmAsBytes *FixedReaderMock) AsBytesAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsBytes.afterAsBytesCounter)
}

// AsBytesBeforeCounter returns a count of FixedReaderMock.AsBytes invocations
func (mmAsBytes *FixedReaderMock) AsBytesBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsBytes.beforeAsBytesCounter)
}

// MinimockAsBytesDone returns true if the count of the AsBytes invocations corresponds
// the number of defined expectations
func (m *FixedReaderMock) MinimockAsBytesDone() bool {
	for _, e := range m.AsBytesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsBytesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsBytesCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsBytes != nil && mm_atomic.LoadUint64(&m.afterAsBytesCounter) < 1 {
		return false
	}
	return true
}

// MinimockAsBytesInspect logs each unmet expectation
func (m *FixedReaderMock) MinimockAsBytesInspect() {
	for _, e := range m.AsBytesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FixedReaderMock.AsBytes")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsBytesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsBytesCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.AsBytes")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsBytes != nil && mm_atomic.LoadUint64(&m.afterAsBytesCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.AsBytes")
	}
}

type mFixedReaderMockFixedByteSize struct {
	mock               *FixedReaderMock
	defaultExpectation *FixedReaderMockFixedByteSizeExpectation
	expectations       []*FixedReaderMockFixedByteSizeExpectation
}

// FixedReaderMockFixedByteSizeExpectation specifies expectation struct of the FixedReader.FixedByteSize
type FixedReaderMockFixedByteSizeExpectation struct {
	mock *FixedReaderMock

	results *FixedReaderMockFixedByteSizeResults
	Counter uint64
}

// FixedReaderMockFixedByteSizeResults contains results of the FixedReader.FixedByteSize
type FixedReaderMockFixedByteSizeResults struct {
	i1 int
}

// Expect sets up expected params for FixedReader.FixedByteSize
func (mmFixedByteSize *mFixedReaderMockFixedByteSize) Expect() *mFixedReaderMockFixedByteSize {
	if mmFixedByteSize.mock.funcFixedByteSize != nil {
		mmFixedByteSize.mock.t.Fatalf("FixedReaderMock.FixedByteSize mock is already set by Set")
	}

	if mmFixedByteSize.defaultExpectation == nil {
		mmFixedByteSize.defaultExpectation = &FixedReaderMockFixedByteSizeExpectation{}
	}

	return mmFixedByteSize
}

// Inspect accepts an inspector function that has same arguments as the FixedReader.FixedByteSize
func (mmFixedByteSize *mFixedReaderMockFixedByteSize) Inspect(f func()) *mFixedReaderMockFixedByteSize {
	if mmFixedByteSize.mock.inspectFuncFixedByteSize != nil {
		mmFixedByteSize.mock.t.Fatalf("Inspect function is already set for FixedReaderMock.FixedByteSize")
	}

	mmFixedByteSize.mock.inspectFuncFixedByteSize = f

	return mmFixedByteSize
}

// Return sets up results that will be returned by FixedReader.FixedByteSize
func (mmFixedByteSize *mFixedReaderMockFixedByteSize) Return(i1 int) *FixedReaderMock {
	if mmFixedByteSize.mock.funcFixedByteSize != nil {
		mmFixedByteSize.mock.t.Fatalf("FixedReaderMock.FixedByteSize mock is already set by Set")
	}

	if mmFixedByteSize.defaultExpectation == nil {
		mmFixedByteSize.defaultExpectation = &FixedReaderMockFixedByteSizeExpectation{mock: mmFixedByteSize.mock}
	}
	mmFixedByteSize.defaultExpectation.results = &FixedReaderMockFixedByteSizeResults{i1}
	return mmFixedByteSize.mock
}

//Set uses given function f to mock the FixedReader.FixedByteSize method
func (mmFixedByteSize *mFixedReaderMockFixedByteSize) Set(f func() (i1 int)) *FixedReaderMock {
	if mmFixedByteSize.defaultExpectation != nil {
		mmFixedByteSize.mock.t.Fatalf("Default expectation is already set for the FixedReader.FixedByteSize method")
	}

	if len(mmFixedByteSize.expectations) > 0 {
		mmFixedByteSize.mock.t.Fatalf("Some expectations are already set for the FixedReader.FixedByteSize method")
	}

	mmFixedByteSize.mock.funcFixedByteSize = f
	return mmFixedByteSize.mock
}

// FixedByteSize implements FixedReader
func (mmFixedByteSize *FixedReaderMock) FixedByteSize() (i1 int) {
	mm_atomic.AddUint64(&mmFixedByteSize.beforeFixedByteSizeCounter, 1)
	defer mm_atomic.AddUint64(&mmFixedByteSize.afterFixedByteSizeCounter, 1)

	if mmFixedByteSize.inspectFuncFixedByteSize != nil {
		mmFixedByteSize.inspectFuncFixedByteSize()
	}

	if mmFixedByteSize.FixedByteSizeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFixedByteSize.FixedByteSizeMock.defaultExpectation.Counter, 1)

		mm_results := mmFixedByteSize.FixedByteSizeMock.defaultExpectation.results
		if mm_results == nil {
			mmFixedByteSize.t.Fatal("No results are set for the FixedReaderMock.FixedByteSize")
		}
		return (*mm_results).i1
	}
	if mmFixedByteSize.funcFixedByteSize != nil {
		return mmFixedByteSize.funcFixedByteSize()
	}
	mmFixedByteSize.t.Fatalf("Unexpected call to FixedReaderMock.FixedByteSize.")
	return
}

// FixedByteSizeAfterCounter returns a count of finished FixedReaderMock.FixedByteSize invocations
func (mmFixedByteSize *FixedReaderMock) FixedByteSizeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFixedByteSize.afterFixedByteSizeCounter)
}

// FixedByteSizeBeforeCounter returns a count of FixedReaderMock.FixedByteSize invocations
func (mmFixedByteSize *FixedReaderMock) FixedByteSizeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFixedByteSize.beforeFixedByteSizeCounter)
}

// MinimockFixedByteSizeDone returns true if the count of the FixedByteSize invocations corresponds
// the number of defined expectations
func (m *FixedReaderMock) MinimockFixedByteSizeDone() bool {
	for _, e := range m.FixedByteSizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FixedByteSizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFixedByteSizeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFixedByteSize != nil && mm_atomic.LoadUint64(&m.afterFixedByteSizeCounter) < 1 {
		return false
	}
	return true
}

// MinimockFixedByteSizeInspect logs each unmet expectation
func (m *FixedReaderMock) MinimockFixedByteSizeInspect() {
	for _, e := range m.FixedByteSizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FixedReaderMock.FixedByteSize")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FixedByteSizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFixedByteSizeCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.FixedByteSize")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFixedByteSize != nil && mm_atomic.LoadUint64(&m.afterFixedByteSizeCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.FixedByteSize")
	}
}

type mFixedReaderMockRead struct {
	mock               *FixedReaderMock
	defaultExpectation *FixedReaderMockReadExpectation
	expectations       []*FixedReaderMockReadExpectation

	callArgs []*FixedReaderMockReadParams
	mutex    sync.RWMutex
}

// FixedReaderMockReadExpectation specifies expectation struct of the FixedReader.Read
type FixedReaderMockReadExpectation struct {
	mock    *FixedReaderMock
	params  *FixedReaderMockReadParams
	results *FixedReaderMockReadResults
	Counter uint64
}

// FixedReaderMockReadParams contains parameters of the FixedReader.Read
type FixedReaderMockReadParams struct {
	p []byte
}

// FixedReaderMockReadResults contains results of the FixedReader.Read
type FixedReaderMockReadResults struct {
	n   int
	err error
}

// Expect sets up expected params for FixedReader.Read
func (mmRead *mFixedReaderMockRead) Expect(p []byte) *mFixedReaderMockRead {
	if mmRead.mock.funcRead != nil {
		mmRead.mock.t.Fatalf("FixedReaderMock.Read mock is already set by Set")
	}

	if mmRead.defaultExpectation == nil {
		mmRead.defaultExpectation = &FixedReaderMockReadExpectation{}
	}

	mmRead.defaultExpectation.params = &FixedReaderMockReadParams{p}
	for _, e := range mmRead.expectations {
		if minimock.Equal(e.params, mmRead.defaultExpectation.params) {
			mmRead.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRead.defaultExpectation.params)
		}
	}

	return mmRead
}

// Inspect accepts an inspector function that has same arguments as the FixedReader.Read
func (mmRead *mFixedReaderMockRead) Inspect(f func(p []byte)) *mFixedReaderMockRead {
	if mmRead.mock.inspectFuncRead != nil {
		mmRead.mock.t.Fatalf("Inspect function is already set for FixedReaderMock.Read")
	}

	mmRead.mock.inspectFuncRead = f

	return mmRead
}

// Return sets up results that will be returned by FixedReader.Read
func (mmRead *mFixedReaderMockRead) Return(n int, err error) *FixedReaderMock {
	if mmRead.mock.funcRead != nil {
		mmRead.mock.t.Fatalf("FixedReaderMock.Read mock is already set by Set")
	}

	if mmRead.defaultExpectation == nil {
		mmRead.defaultExpectation = &FixedReaderMockReadExpectation{mock: mmRead.mock}
	}
	mmRead.defaultExpectation.results = &FixedReaderMockReadResults{n, err}
	return mmRead.mock
}

//Set uses given function f to mock the FixedReader.Read method
func (mmRead *mFixedReaderMockRead) Set(f func(p []byte) (n int, err error)) *FixedReaderMock {
	if mmRead.defaultExpectation != nil {
		mmRead.mock.t.Fatalf("Default expectation is already set for the FixedReader.Read method")
	}

	if len(mmRead.expectations) > 0 {
		mmRead.mock.t.Fatalf("Some expectations are already set for the FixedReader.Read method")
	}

	mmRead.mock.funcRead = f
	return mmRead.mock
}

// When sets expectation for the FixedReader.Read which will trigger the result defined by the following
// Then helper
func (mmRead *mFixedReaderMockRead) When(p []byte) *FixedReaderMockReadExpectation {
	if mmRead.mock.funcRead != nil {
		mmRead.mock.t.Fatalf("FixedReaderMock.Read mock is already set by Set")
	}

	expectation := &FixedReaderMockReadExpectation{
		mock:   mmRead.mock,
		params: &FixedReaderMockReadParams{p},
	}
	mmRead.expectations = append(mmRead.expectations, expectation)
	return expectation
}

// Then sets up FixedReader.Read return parameters for the expectation previously defined by the When method
func (e *FixedReaderMockReadExpectation) Then(n int, err error) *FixedReaderMock {
	e.results = &FixedReaderMockReadResults{n, err}
	return e.mock
}

// Read implements FixedReader
func (mmRead *FixedReaderMock) Read(p []byte) (n int, err error) {
	mm_atomic.AddUint64(&mmRead.beforeReadCounter, 1)
	defer mm_atomic.AddUint64(&mmRead.afterReadCounter, 1)

	if mmRead.inspectFuncRead != nil {
		mmRead.inspectFuncRead(p)
	}

	mm_params := &FixedReaderMockReadParams{p}

	// Record call args
	mmRead.ReadMock.mutex.Lock()
	mmRead.ReadMock.callArgs = append(mmRead.ReadMock.callArgs, mm_params)
	mmRead.ReadMock.mutex.Unlock()

	for _, e := range mmRead.ReadMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.n, e.results.err
		}
	}

	if mmRead.ReadMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRead.ReadMock.defaultExpectation.Counter, 1)
		mm_want := mmRead.ReadMock.defaultExpectation.params
		mm_got := FixedReaderMockReadParams{p}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRead.t.Errorf("FixedReaderMock.Read got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRead.ReadMock.defaultExpectation.results
		if mm_results == nil {
			mmRead.t.Fatal("No results are set for the FixedReaderMock.Read")
		}
		return (*mm_results).n, (*mm_results).err
	}
	if mmRead.funcRead != nil {
		return mmRead.funcRead(p)
	}
	mmRead.t.Fatalf("Unexpected call to FixedReaderMock.Read. %v", p)
	return
}

// ReadAfterCounter returns a count of finished FixedReaderMock.Read invocations
func (mmRead *FixedReaderMock) ReadAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRead.afterReadCounter)
}

// ReadBeforeCounter returns a count of FixedReaderMock.Read invocations
func (mmRead *FixedReaderMock) ReadBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRead.beforeReadCounter)
}

// Calls returns a list of arguments used in each call to FixedReaderMock.Read.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRead *mFixedReaderMockRead) Calls() []*FixedReaderMockReadParams {
	mmRead.mutex.RLock()

	argCopy := make([]*FixedReaderMockReadParams, len(mmRead.callArgs))
	copy(argCopy, mmRead.callArgs)

	mmRead.mutex.RUnlock()

	return argCopy
}

// MinimockReadDone returns true if the count of the Read invocations corresponds
// the number of defined expectations
func (m *FixedReaderMock) MinimockReadDone() bool {
	for _, e := range m.ReadMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRead != nil && mm_atomic.LoadUint64(&m.afterReadCounter) < 1 {
		return false
	}
	return true
}

// MinimockReadInspect logs each unmet expectation
func (m *FixedReaderMock) MinimockReadInspect() {
	for _, e := range m.ReadMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FixedReaderMock.Read with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadCounter) < 1 {
		if m.ReadMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FixedReaderMock.Read")
		} else {
			m.t.Errorf("Expected call to FixedReaderMock.Read with params: %#v", *m.ReadMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRead != nil && mm_atomic.LoadUint64(&m.afterReadCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.Read")
	}
}

type mFixedReaderMockWriteTo struct {
	mock               *FixedReaderMock
	defaultExpectation *FixedReaderMockWriteToExpectation
	expectations       []*FixedReaderMockWriteToExpectation

	callArgs []*FixedReaderMockWriteToParams
	mutex    sync.RWMutex
}

// FixedReaderMockWriteToExpectation specifies expectation struct of the FixedReader.WriteTo
type FixedReaderMockWriteToExpectation struct {
	mock    *FixedReaderMock
	params  *FixedReaderMockWriteToParams
	results *FixedReaderMockWriteToResults
	Counter uint64
}

// FixedReaderMockWriteToParams contains parameters of the FixedReader.WriteTo
type FixedReaderMockWriteToParams struct {
	w io.Writer
}

// FixedReaderMockWriteToResults contains results of the FixedReader.WriteTo
type FixedReaderMockWriteToResults struct {
	n   int64
	err error
}

// Expect sets up expected params for FixedReader.WriteTo
func (mmWriteTo *mFixedReaderMockWriteTo) Expect(w io.Writer) *mFixedReaderMockWriteTo {
	if mmWriteTo.mock.funcWriteTo != nil {
		mmWriteTo.mock.t.Fatalf("FixedReaderMock.WriteTo mock is already set by Set")
	}

	if mmWriteTo.defaultExpectation == nil {
		mmWriteTo.defaultExpectation = &FixedReaderMockWriteToExpectation{}
	}

	mmWriteTo.defaultExpectation.params = &FixedReaderMockWriteToParams{w}
	for _, e := range mmWriteTo.expectations {
		if minimock.Equal(e.params, mmWriteTo.defaultExpectation.params) {
			mmWriteTo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmWriteTo.defaultExpectation.params)
		}
	}

	return mmWriteTo
}

// Inspect accepts an inspector function that has same arguments as the FixedReader.WriteTo
func (mmWriteTo *mFixedReaderMockWriteTo) Inspect(f func(w io.Writer)) *mFixedReaderMockWriteTo {
	if mmWriteTo.mock.inspectFuncWriteTo != nil {
		mmWriteTo.mock.t.Fatalf("Inspect function is already set for FixedReaderMock.WriteTo")
	}

	mmWriteTo.mock.inspectFuncWriteTo = f

	return mmWriteTo
}

// Return sets up results that will be returned by FixedReader.WriteTo
func (mmWriteTo *mFixedReaderMockWriteTo) Return(n int64, err error) *FixedReaderMock {
	if mmWriteTo.mock.funcWriteTo != nil {
		mmWriteTo.mock.t.Fatalf("FixedReaderMock.WriteTo mock is already set by Set")
	}

	if mmWriteTo.defaultExpectation == nil {
		mmWriteTo.defaultExpectation = &FixedReaderMockWriteToExpectation{mock: mmWriteTo.mock}
	}
	mmWriteTo.defaultExpectation.results = &FixedReaderMockWriteToResults{n, err}
	return mmWriteTo.mock
}

//Set uses given function f to mock the FixedReader.WriteTo method
func (mmWriteTo *mFixedReaderMockWriteTo) Set(f func(w io.Writer) (n int64, err error)) *FixedReaderMock {
	if mmWriteTo.defaultExpectation != nil {
		mmWriteTo.mock.t.Fatalf("Default expectation is already set for the FixedReader.WriteTo method")
	}

	if len(mmWriteTo.expectations) > 0 {
		mmWriteTo.mock.t.Fatalf("Some expectations are already set for the FixedReader.WriteTo method")
	}

	mmWriteTo.mock.funcWriteTo = f
	return mmWriteTo.mock
}

// When sets expectation for the FixedReader.WriteTo which will trigger the result defined by the following
// Then helper
func (mmWriteTo *mFixedReaderMockWriteTo) When(w io.Writer) *FixedReaderMockWriteToExpectation {
	if mmWriteTo.mock.funcWriteTo != nil {
		mmWriteTo.mock.t.Fatalf("FixedReaderMock.WriteTo mock is already set by Set")
	}

	expectation := &FixedReaderMockWriteToExpectation{
		mock:   mmWriteTo.mock,
		params: &FixedReaderMockWriteToParams{w},
	}
	mmWriteTo.expectations = append(mmWriteTo.expectations, expectation)
	return expectation
}

// Then sets up FixedReader.WriteTo return parameters for the expectation previously defined by the When method
func (e *FixedReaderMockWriteToExpectation) Then(n int64, err error) *FixedReaderMock {
	e.results = &FixedReaderMockWriteToResults{n, err}
	return e.mock
}

// WriteTo implements FixedReader
func (mmWriteTo *FixedReaderMock) WriteTo(w io.Writer) (n int64, err error) {
	mm_atomic.AddUint64(&mmWriteTo.beforeWriteToCounter, 1)
	defer mm_atomic.AddUint64(&mmWriteTo.afterWriteToCounter, 1)

	if mmWriteTo.inspectFuncWriteTo != nil {
		mmWriteTo.inspectFuncWriteTo(w)
	}

	mm_params := &FixedReaderMockWriteToParams{w}

	// Record call args
	mmWriteTo.WriteToMock.mutex.Lock()
	mmWriteTo.WriteToMock.callArgs = append(mmWriteTo.WriteToMock.callArgs, mm_params)
	mmWriteTo.WriteToMock.mutex.Unlock()

	for _, e := range mmWriteTo.WriteToMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.n, e.results.err
		}
	}

	if mmWriteTo.WriteToMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmWriteTo.WriteToMock.defaultExpectation.Counter, 1)
		mm_want := mmWriteTo.WriteToMock.defaultExpectation.params
		mm_got := FixedReaderMockWriteToParams{w}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmWriteTo.t.Errorf("FixedReaderMock.WriteTo got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmWriteTo.WriteToMock.defaultExpectation.results
		if mm_results == nil {
			mmWriteTo.t.Fatal("No results are set for the FixedReaderMock.WriteTo")
		}
		return (*mm_results).n, (*mm_results).err
	}
	if mmWriteTo.funcWriteTo != nil {
		return mmWriteTo.funcWriteTo(w)
	}
	mmWriteTo.t.Fatalf("Unexpected call to FixedReaderMock.WriteTo. %v", w)
	return
}

// WriteToAfterCounter returns a count of finished FixedReaderMock.WriteTo invocations
func (mmWriteTo *FixedReaderMock) WriteToAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmWriteTo.afterWriteToCounter)
}

// WriteToBeforeCounter returns a count of FixedReaderMock.WriteTo invocations
func (mmWriteTo *FixedReaderMock) WriteToBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmWriteTo.beforeWriteToCounter)
}

// Calls returns a list of arguments used in each call to FixedReaderMock.WriteTo.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmWriteTo *mFixedReaderMockWriteTo) Calls() []*FixedReaderMockWriteToParams {
	mmWriteTo.mutex.RLock()

	argCopy := make([]*FixedReaderMockWriteToParams, len(mmWriteTo.callArgs))
	copy(argCopy, mmWriteTo.callArgs)

	mmWriteTo.mutex.RUnlock()

	return argCopy
}

// MinimockWriteToDone returns true if the count of the WriteTo invocations corresponds
// the number of defined expectations
func (m *FixedReaderMock) MinimockWriteToDone() bool {
	for _, e := range m.WriteToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.WriteToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterWriteToCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcWriteTo != nil && mm_atomic.LoadUint64(&m.afterWriteToCounter) < 1 {
		return false
	}
	return true
}

// MinimockWriteToInspect logs each unmet expectation
func (m *FixedReaderMock) MinimockWriteToInspect() {
	for _, e := range m.WriteToMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FixedReaderMock.WriteTo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.WriteToMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterWriteToCounter) < 1 {
		if m.WriteToMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FixedReaderMock.WriteTo")
		} else {
			m.t.Errorf("Expected call to FixedReaderMock.WriteTo with params: %#v", *m.WriteToMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcWriteTo != nil && mm_atomic.LoadUint64(&m.afterWriteToCounter) < 1 {
		m.t.Error("Expected call to FixedReaderMock.WriteTo")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FixedReaderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAsByteStringInspect()

		m.MinimockAsBytesInspect()

		m.MinimockFixedByteSizeInspect()

		m.MinimockReadInspect()

		m.MinimockWriteToInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FixedReaderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *FixedReaderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAsByteStringDone() &&
		m.MinimockAsBytesDone() &&
		m.MinimockFixedByteSizeDone() &&
		m.MinimockReadDone() &&
		m.MinimockWriteToDone()
}
