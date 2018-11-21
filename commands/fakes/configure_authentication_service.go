// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type ConfigureAuthenticationService struct {
	EnsureAvailabilityStub        func(api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error)
	ensureAvailabilityMutex       sync.RWMutex
	ensureAvailabilityArgsForCall []struct {
		arg1 api.EnsureAvailabilityInput
	}
	ensureAvailabilityReturns struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	ensureAvailabilityReturnsOnCall map[int]struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	SetupStub        func(api.SetupInput) (api.SetupOutput, error)
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 api.SetupInput
	}
	setupReturns struct {
		result1 api.SetupOutput
		result2 error
	}
	setupReturnsOnCall map[int]struct {
		result1 api.SetupOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigureAuthenticationService) EnsureAvailability(arg1 api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error) {
	fake.ensureAvailabilityMutex.Lock()
	ret, specificReturn := fake.ensureAvailabilityReturnsOnCall[len(fake.ensureAvailabilityArgsForCall)]
	fake.ensureAvailabilityArgsForCall = append(fake.ensureAvailabilityArgsForCall, struct {
		arg1 api.EnsureAvailabilityInput
	}{arg1})
	fake.recordInvocation("EnsureAvailability", []interface{}{arg1})
	fake.ensureAvailabilityMutex.Unlock()
	if fake.EnsureAvailabilityStub != nil {
		return fake.EnsureAvailabilityStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.ensureAvailabilityReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ConfigureAuthenticationService) EnsureAvailabilityCallCount() int {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return len(fake.ensureAvailabilityArgsForCall)
}

func (fake *ConfigureAuthenticationService) EnsureAvailabilityCalls(stub func(api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error)) {
	fake.ensureAvailabilityMutex.Lock()
	defer fake.ensureAvailabilityMutex.Unlock()
	fake.EnsureAvailabilityStub = stub
}

func (fake *ConfigureAuthenticationService) EnsureAvailabilityArgsForCall(i int) api.EnsureAvailabilityInput {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	argsForCall := fake.ensureAvailabilityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureAuthenticationService) EnsureAvailabilityReturns(result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.ensureAvailabilityMutex.Lock()
	defer fake.ensureAvailabilityMutex.Unlock()
	fake.EnsureAvailabilityStub = nil
	fake.ensureAvailabilityReturns = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureAuthenticationService) EnsureAvailabilityReturnsOnCall(i int, result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.ensureAvailabilityMutex.Lock()
	defer fake.ensureAvailabilityMutex.Unlock()
	fake.EnsureAvailabilityStub = nil
	if fake.ensureAvailabilityReturnsOnCall == nil {
		fake.ensureAvailabilityReturnsOnCall = make(map[int]struct {
			result1 api.EnsureAvailabilityOutput
			result2 error
		})
	}
	fake.ensureAvailabilityReturnsOnCall[i] = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureAuthenticationService) Setup(arg1 api.SetupInput) (api.SetupOutput, error) {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 api.SetupInput
	}{arg1})
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ConfigureAuthenticationService) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *ConfigureAuthenticationService) SetupCalls(stub func(api.SetupInput) (api.SetupOutput, error)) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *ConfigureAuthenticationService) SetupArgsForCall(i int) api.SetupInput {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureAuthenticationService) SetupReturns(result1 api.SetupOutput, result2 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 api.SetupOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureAuthenticationService) SetupReturnsOnCall(i int, result1 api.SetupOutput, result2 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 api.SetupOutput
			result2 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 api.SetupOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureAuthenticationService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigureAuthenticationService) recordInvocation(key string, args []interface{}) {
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
