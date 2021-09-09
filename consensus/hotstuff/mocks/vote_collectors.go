// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	hotstuff "github.com/onflow/flow-go/consensus/hotstuff"
	mock "github.com/stretchr/testify/mock"
)

// VoteCollectors is an autogenerated mock type for the VoteCollectors type
type VoteCollectors struct {
	mock.Mock
}

// GetOrCreateCollector provides a mock function with given fields: view
func (_m *VoteCollectors) GetOrCreateCollector(view uint64) (hotstuff.VoteCollector, bool, error) {
	ret := _m.Called(view)

	var r0 hotstuff.VoteCollector
	if rf, ok := ret.Get(0).(func(uint64) hotstuff.VoteCollector); ok {
		r0 = rf(view)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hotstuff.VoteCollector)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uint64) bool); ok {
		r1 = rf(view)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(view)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PruneUpToView provides a mock function with given fields: view
func (_m *VoteCollectors) PruneUpToView(view uint64) error {
	ret := _m.Called(view)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(view)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}