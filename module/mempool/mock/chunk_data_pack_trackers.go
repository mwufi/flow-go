// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import flow "github.com/dapperlabs/flow-go/model/flow"

import mock "github.com/stretchr/testify/mock"
import tracker "github.com/dapperlabs/flow-go/model/verification/tracker"

// ChunkDataPackTrackers is an autogenerated mock type for the ChunkDataPackTrackers type
type ChunkDataPackTrackers struct {
	mock.Mock
}

// Add provides a mock function with given fields: cdpt
func (_m *ChunkDataPackTrackers) Add(cdpt *tracker.ChunkDataPackTracker) bool {
	ret := _m.Called(cdpt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*tracker.ChunkDataPackTracker) bool); ok {
		r0 = rf(cdpt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *ChunkDataPackTrackers) All() []*tracker.ChunkDataPackTracker {
	ret := _m.Called()

	var r0 []*tracker.ChunkDataPackTracker
	if rf, ok := ret.Get(0).(func() []*tracker.ChunkDataPackTracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*tracker.ChunkDataPackTracker)
		}
	}

	return r0
}

// ByChunkID provides a mock function with given fields: chunkID
func (_m *ChunkDataPackTrackers) ByChunkID(chunkID flow.Identifier) (*tracker.ChunkDataPackTracker, error) {
	ret := _m.Called(chunkID)

	var r0 *tracker.ChunkDataPackTracker
	if rf, ok := ret.Get(0).(func(flow.Identifier) *tracker.ChunkDataPackTracker); ok {
		r0 = rf(chunkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tracker.ChunkDataPackTracker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(chunkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Has provides a mock function with given fields: chunkID
func (_m *ChunkDataPackTrackers) Has(chunkID flow.Identifier) bool {
	ret := _m.Called(chunkID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(chunkID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Inc provides a mock function with given fields: chunkID
func (_m *ChunkDataPackTrackers) Inc(chunkID flow.Identifier) (*tracker.ChunkDataPackTracker, error) {
	ret := _m.Called(chunkID)

	var r0 *tracker.ChunkDataPackTracker
	if rf, ok := ret.Get(0).(func(flow.Identifier) *tracker.ChunkDataPackTracker); ok {
		r0 = rf(chunkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tracker.ChunkDataPackTracker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(chunkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rem provides a mock function with given fields: chunkID
func (_m *ChunkDataPackTrackers) Rem(chunkID flow.Identifier) bool {
	ret := _m.Called(chunkID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(chunkID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
