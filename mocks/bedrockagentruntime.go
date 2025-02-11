// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	bedrockagentruntime "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"

	mock "github.com/stretchr/testify/mock"
)

// BedrockagentruntimeClient is an autogenerated mock type for the BedrockagentruntimeClient type
type BedrockagentruntimeClient struct {
	mock.Mock
}

// InvokeAgent provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockagentruntimeClient) InvokeAgent(ctx context.Context, params *bedrockagentruntime.InvokeAgentInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeAgentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeAgent")
	}

	var r0 *bedrockagentruntime.InvokeAgentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.InvokeAgentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.InvokeAgentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.InvokeAgentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.InvokeAgentInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *BedrockagentruntimeClient) Options() bedrockagentruntime.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bedrockagentruntime.Options
	if rf, ok := ret.Get(0).(func() bedrockagentruntime.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bedrockagentruntime.Options)
	}

	return r0
}

// Retrieve provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockagentruntimeClient) Retrieve(ctx context.Context, params *bedrockagentruntime.RetrieveInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Retrieve")
	}

	var r0 *bedrockagentruntime.RetrieveOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RetrieveOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RetrieveOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RetrieveInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAndGenerate provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockagentruntimeClient) RetrieveAndGenerate(ctx context.Context, params *bedrockagentruntime.RetrieveAndGenerateInput, optFns ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAndGenerate")
	}

	var r0 *bedrockagentruntime.RetrieveAndGenerateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) (*bedrockagentruntime.RetrieveAndGenerateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) *bedrockagentruntime.RetrieveAndGenerateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrockagentruntime.RetrieveAndGenerateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrockagentruntime.RetrieveAndGenerateInput, ...func(*bedrockagentruntime.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBedrockagentruntimeClient creates a new instance of BedrockagentruntimeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBedrockagentruntimeClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BedrockagentruntimeClient {
	mock := &BedrockagentruntimeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
