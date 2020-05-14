// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/hyperledger/fabric-gm/protos/common"
import ledger "github.com/hyperledger/fabric-gm/core/ledger"
import mock "github.com/stretchr/testify/mock"
import peer "github.com/hyperledger/fabric-gm/protos/peer"
import privdata "github.com/hyperledger/fabric-gm/core/common/privdata"

// CollectionStore is an autogenerated mock type for the CollectionStore type
type CollectionStore struct {
	mock.Mock
}

// AccessFilter provides a mock function with given fields: channelName, collectionPolicyConfig
func (_m *CollectionStore) AccessFilter(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (privdata.Filter, error) {
	ret := _m.Called(channelName, collectionPolicyConfig)

	var r0 privdata.Filter
	if rf, ok := ret.Get(0).(func(string, *common.CollectionPolicyConfig) privdata.Filter); ok {
		r0 = rf(channelName, collectionPolicyConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *common.CollectionPolicyConfig) error); ok {
		r1 = rf(channelName, collectionPolicyConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasReadAccess provides a mock function with given fields: _a0, _a1, _a2
func (_m *CollectionStore) HasReadAccess(_a0 common.CollectionCriteria, _a1 *peer.SignedProposal, _a2 ledger.QueryExecutor) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollection provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollection(_a0 common.CollectionCriteria) (privdata.Collection, error) {
	ret := _m.Called(_a0)

	var r0 privdata.Collection
	if rf, ok := ret.Get(0).(func(common.CollectionCriteria) privdata.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionAccessPolicy provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionAccessPolicy(_a0 common.CollectionCriteria) (privdata.CollectionAccessPolicy, error) {
	ret := _m.Called(_a0)

	var r0 privdata.CollectionAccessPolicy
	if rf, ok := ret.Get(0).(func(common.CollectionCriteria) privdata.CollectionAccessPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.CollectionAccessPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionConfigPackage provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionConfigPackage(_a0 common.CollectionCriteria) (*common.CollectionConfigPackage, error) {
	ret := _m.Called(_a0)

	var r0 *common.CollectionConfigPackage
	if rf, ok := ret.Get(0).(func(common.CollectionCriteria) *common.CollectionConfigPackage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CollectionConfigPackage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionPersistenceConfigs provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionPersistenceConfigs(_a0 common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error) {
	ret := _m.Called(_a0)

	var r0 privdata.CollectionPersistenceConfigs
	if rf, ok := ret.Get(0).(func(common.CollectionCriteria) privdata.CollectionPersistenceConfigs); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.CollectionPersistenceConfigs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}