// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import flow "github.com/dapperlabs/flow-go/model/flow"

import mock "github.com/stretchr/testify/mock"

// Blocks is an autogenerated mock type for the Blocks type
type Blocks struct {
	mock.Mock
}

// Add provides a mock function with given fields: block
func (_m *Blocks) Add(block *flow.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *Blocks) All() []*flow.Block {
	ret := _m.Called()

	var r0 []*flow.Block
	if rf, ok := ret.Get(0).(func() []*flow.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*flow.Block)
		}
	}

	return r0
}

// ByID provides a mock function with given fields: blockID
func (_m *Blocks) ByID(blockID flow.Identifier) (*flow.Block, error) {
	ret := _m.Called(blockID)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Block); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Has provides a mock function with given fields: blockID
func (_m *Blocks) Has(blockID flow.Identifier) bool {
	ret := _m.Called(blockID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(blockID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Hash provides a mock function with given fields:
func (_m *Blocks) Hash() flow.Identifier {
	ret := _m.Called()

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	return r0
}

// Rem provides a mock function with given fields: blockID
func (_m *Blocks) Rem(blockID flow.Identifier) bool {
	ret := _m.Called(blockID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(blockID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *Blocks) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
