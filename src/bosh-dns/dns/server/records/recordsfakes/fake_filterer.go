// Code generated by counterfeiter. DO NOT EDIT.
package recordsfakes

import (
	"bosh-dns/dns/server/criteria"
	"bosh-dns/dns/server/record"
	"bosh-dns/dns/server/records"
	"sync"
)

type FakeFilterer struct {
	FilterStub        func(criteria.MatchMaker, []record.Record) []record.Record
	filterMutex       sync.RWMutex
	filterArgsForCall []struct {
		arg1 criteria.MatchMaker
		arg2 []record.Record
	}
	filterReturns struct {
		result1 []record.Record
	}
	filterReturnsOnCall map[int]struct {
		result1 []record.Record
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilterer) Filter(arg1 criteria.MatchMaker, arg2 []record.Record) []record.Record {
	var arg2Copy []record.Record
	if arg2 != nil {
		arg2Copy = make([]record.Record, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.filterMutex.Lock()
	ret, specificReturn := fake.filterReturnsOnCall[len(fake.filterArgsForCall)]
	fake.filterArgsForCall = append(fake.filterArgsForCall, struct {
		arg1 criteria.MatchMaker
		arg2 []record.Record
	}{arg1, arg2Copy})
	fake.recordInvocation("Filter", []interface{}{arg1, arg2Copy})
	fake.filterMutex.Unlock()
	if fake.FilterStub != nil {
		return fake.FilterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.filterReturns
	return fakeReturns.result1
}

func (fake *FakeFilterer) FilterCallCount() int {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	return len(fake.filterArgsForCall)
}

func (fake *FakeFilterer) FilterCalls(stub func(criteria.MatchMaker, []record.Record) []record.Record) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = stub
}

func (fake *FakeFilterer) FilterArgsForCall(i int) (criteria.MatchMaker, []record.Record) {
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	argsForCall := fake.filterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFilterer) FilterReturns(result1 []record.Record) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	fake.filterReturns = struct {
		result1 []record.Record
	}{result1}
}

func (fake *FakeFilterer) FilterReturnsOnCall(i int, result1 []record.Record) {
	fake.filterMutex.Lock()
	defer fake.filterMutex.Unlock()
	fake.FilterStub = nil
	if fake.filterReturnsOnCall == nil {
		fake.filterReturnsOnCall = make(map[int]struct {
			result1 []record.Record
		})
	}
	fake.filterReturnsOnCall[i] = struct {
		result1 []record.Record
	}{result1}
}

func (fake *FakeFilterer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterMutex.RLock()
	defer fake.filterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFilterer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ records.Filterer = new(FakeFilterer)
