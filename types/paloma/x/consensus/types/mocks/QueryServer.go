// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/palomachain/pigeon/types/paloma/x/consensus/types"
)

// QueryServer is an autogenerated mock type for the QueryServer type
type QueryServer struct {
	mock.Mock
}

// GetAllQueueNames provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) GetAllQueueNames(_a0 context.Context, _a1 *types.QueryGetAllQueueNamesRequest) (*types.QueryGetAllQueueNamesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryGetAllQueueNamesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetAllQueueNamesRequest) *types.QueryGetAllQueueNamesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetAllQueueNamesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetAllQueueNamesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessagesInQueue provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) MessagesInQueue(_a0 context.Context, _a1 *types.QueryMessagesInQueueRequest) (*types.QueryMessagesInQueueResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryMessagesInQueueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMessagesInQueueRequest) *types.QueryMessagesInQueueResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMessagesInQueueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMessagesInQueueRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) Params(_a0 context.Context, _a1 *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest) *types.QueryParamsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueuedMessagesForSigning provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) QueuedMessagesForSigning(_a0 context.Context, _a1 *types.QueryQueuedMessagesForSigningRequest) (*types.QueryQueuedMessagesForSigningResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.QueryQueuedMessagesForSigningResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryQueuedMessagesForSigningRequest) *types.QueryQueuedMessagesForSigningResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryQueuedMessagesForSigningResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryQueuedMessagesForSigningRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryServer creates a new instance of QueryServer. It also registers a cleanup function to assert the mocks expectations.
func NewQueryServer(t testing.TB) *QueryServer {
	mock := &QueryServer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
