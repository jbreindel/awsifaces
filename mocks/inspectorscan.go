// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	inspectorscan "github.com/aws/aws-sdk-go-v2/service/inspectorscan"
	mock "github.com/stretchr/testify/mock"
)

// InspectorscanClient is an autogenerated mock type for the InspectorscanClient type
type InspectorscanClient struct {
	mock.Mock
}

// Options provides a mock function with given fields:
func (_m *InspectorscanClient) Options() inspectorscan.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 inspectorscan.Options
	if rf, ok := ret.Get(0).(func() inspectorscan.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(inspectorscan.Options)
	}

	return r0
}

// ScanSbom provides a mock function with given fields: ctx, params, optFns
func (_m *InspectorscanClient) ScanSbom(ctx context.Context, params *inspectorscan.ScanSbomInput, optFns ...func(*inspectorscan.Options)) (*inspectorscan.ScanSbomOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ScanSbom")
	}

	var r0 *inspectorscan.ScanSbomOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inspectorscan.ScanSbomInput, ...func(*inspectorscan.Options)) (*inspectorscan.ScanSbomOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inspectorscan.ScanSbomInput, ...func(*inspectorscan.Options)) *inspectorscan.ScanSbomOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inspectorscan.ScanSbomOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inspectorscan.ScanSbomInput, ...func(*inspectorscan.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInspectorscanClient creates a new instance of InspectorscanClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInspectorscanClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *InspectorscanClient {
	mock := &InspectorscanClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
