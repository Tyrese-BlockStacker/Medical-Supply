// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	medicalsupply "github.com/hyperledger/fabric-samples/medical-supply/regulators/chaincode/medical-supply"
)

type TransactionContext struct {
	GetClientIdentityStub        func() cid.ClientIdentity
	getClientIdentityMutex       sync.RWMutex
	getClientIdentityArgsForCall []struct {
	}
	getClientIdentityReturns struct {
		result1 cid.ClientIdentity
	}
	getClientIdentityReturnsOnCall map[int]struct {
		result1 cid.ClientIdentity
	}
	GetMedicineListStub        func() medicalsupply.ListInterface
	getMedicineListMutex       sync.RWMutex
	getMedicineListArgsForCall []struct {
	}
	getMedicineListReturns struct {
		result1 medicalsupply.ListInterface
	}
	getMedicineListReturnsOnCall map[int]struct {
		result1 medicalsupply.ListInterface
	}
	GetStubStub        func() shim.ChaincodeStubInterface
	getStubMutex       sync.RWMutex
	getStubArgsForCall []struct {
	}
	getStubReturns struct {
		result1 shim.ChaincodeStubInterface
	}
	getStubReturnsOnCall map[int]struct {
		result1 shim.ChaincodeStubInterface
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TransactionContext) GetClientIdentity() cid.ClientIdentity {
	fake.getClientIdentityMutex.Lock()
	ret, specificReturn := fake.getClientIdentityReturnsOnCall[len(fake.getClientIdentityArgsForCall)]
	fake.getClientIdentityArgsForCall = append(fake.getClientIdentityArgsForCall, struct {
	}{})
	stub := fake.GetClientIdentityStub
	fakeReturns := fake.getClientIdentityReturns
	fake.recordInvocation("GetClientIdentity", []interface{}{})
	fake.getClientIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransactionContext) GetClientIdentityCallCount() int {
	fake.getClientIdentityMutex.RLock()
	defer fake.getClientIdentityMutex.RUnlock()
	return len(fake.getClientIdentityArgsForCall)
}

func (fake *TransactionContext) GetClientIdentityCalls(stub func() cid.ClientIdentity) {
	fake.getClientIdentityMutex.Lock()
	defer fake.getClientIdentityMutex.Unlock()
	fake.GetClientIdentityStub = stub
}

func (fake *TransactionContext) GetClientIdentityReturns(result1 cid.ClientIdentity) {
	fake.getClientIdentityMutex.Lock()
	defer fake.getClientIdentityMutex.Unlock()
	fake.GetClientIdentityStub = nil
	fake.getClientIdentityReturns = struct {
		result1 cid.ClientIdentity
	}{result1}
}

func (fake *TransactionContext) GetClientIdentityReturnsOnCall(i int, result1 cid.ClientIdentity) {
	fake.getClientIdentityMutex.Lock()
	defer fake.getClientIdentityMutex.Unlock()
	fake.GetClientIdentityStub = nil
	if fake.getClientIdentityReturnsOnCall == nil {
		fake.getClientIdentityReturnsOnCall = make(map[int]struct {
			result1 cid.ClientIdentity
		})
	}
	fake.getClientIdentityReturnsOnCall[i] = struct {
		result1 cid.ClientIdentity
	}{result1}
}

func (fake *TransactionContext) GetMedicineList() medicalsupply.ListInterface {
	fake.getMedicineListMutex.Lock()
	ret, specificReturn := fake.getMedicineListReturnsOnCall[len(fake.getMedicineListArgsForCall)]
	fake.getMedicineListArgsForCall = append(fake.getMedicineListArgsForCall, struct {
	}{})
	stub := fake.GetMedicineListStub
	fakeReturns := fake.getMedicineListReturns
	fake.recordInvocation("GetMedicineList", []interface{}{})
	fake.getMedicineListMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransactionContext) GetMedicineListCallCount() int {
	fake.getMedicineListMutex.RLock()
	defer fake.getMedicineListMutex.RUnlock()
	return len(fake.getMedicineListArgsForCall)
}

func (fake *TransactionContext) GetMedicineListCalls(stub func() medicalsupply.ListInterface) {
	fake.getMedicineListMutex.Lock()
	defer fake.getMedicineListMutex.Unlock()
	fake.GetMedicineListStub = stub
}

func (fake *TransactionContext) GetMedicineListReturns(result1 medicalsupply.ListInterface) {
	fake.getMedicineListMutex.Lock()
	defer fake.getMedicineListMutex.Unlock()
	fake.GetMedicineListStub = nil
	fake.getMedicineListReturns = struct {
		result1 medicalsupply.ListInterface
	}{result1}
}

func (fake *TransactionContext) GetMedicineListReturnsOnCall(i int, result1 medicalsupply.ListInterface) {
	fake.getMedicineListMutex.Lock()
	defer fake.getMedicineListMutex.Unlock()
	fake.GetMedicineListStub = nil
	if fake.getMedicineListReturnsOnCall == nil {
		fake.getMedicineListReturnsOnCall = make(map[int]struct {
			result1 medicalsupply.ListInterface
		})
	}
	fake.getMedicineListReturnsOnCall[i] = struct {
		result1 medicalsupply.ListInterface
	}{result1}
}

func (fake *TransactionContext) GetStub() shim.ChaincodeStubInterface {
	fake.getStubMutex.Lock()
	ret, specificReturn := fake.getStubReturnsOnCall[len(fake.getStubArgsForCall)]
	fake.getStubArgsForCall = append(fake.getStubArgsForCall, struct {
	}{})
	stub := fake.GetStubStub
	fakeReturns := fake.getStubReturns
	fake.recordInvocation("GetStub", []interface{}{})
	fake.getStubMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransactionContext) GetStubCallCount() int {
	fake.getStubMutex.RLock()
	defer fake.getStubMutex.RUnlock()
	return len(fake.getStubArgsForCall)
}

func (fake *TransactionContext) GetStubCalls(stub func() shim.ChaincodeStubInterface) {
	fake.getStubMutex.Lock()
	defer fake.getStubMutex.Unlock()
	fake.GetStubStub = stub
}

func (fake *TransactionContext) GetStubReturns(result1 shim.ChaincodeStubInterface) {
	fake.getStubMutex.Lock()
	defer fake.getStubMutex.Unlock()
	fake.GetStubStub = nil
	fake.getStubReturns = struct {
		result1 shim.ChaincodeStubInterface
	}{result1}
}

func (fake *TransactionContext) GetStubReturnsOnCall(i int, result1 shim.ChaincodeStubInterface) {
	fake.getStubMutex.Lock()
	defer fake.getStubMutex.Unlock()
	fake.GetStubStub = nil
	if fake.getStubReturnsOnCall == nil {
		fake.getStubReturnsOnCall = make(map[int]struct {
			result1 shim.ChaincodeStubInterface
		})
	}
	fake.getStubReturnsOnCall[i] = struct {
		result1 shim.ChaincodeStubInterface
	}{result1}
}

func (fake *TransactionContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClientIdentityMutex.RLock()
	defer fake.getClientIdentityMutex.RUnlock()
	fake.getMedicineListMutex.RLock()
	defer fake.getMedicineListMutex.RUnlock()
	fake.getStubMutex.RLock()
	defer fake.getStubMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TransactionContext) recordInvocation(key string, args []interface{}) {
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
