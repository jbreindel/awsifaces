// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	qldbsession "github.com/aws/aws-sdk-go-v2/service/qldbsession"
	mock "github.com/stretchr/testify/mock"
)

// QldbsessionClient is an autogenerated mock type for the QldbsessionClient type
type QldbsessionClient struct {
	mock.Mock
}

// Options provides a mock function with given fields:
func (_m *QldbsessionClient) Options() qldbsession.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 qldbsession.Options
	if rf, ok := ret.Get(0).(func() qldbsession.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(qldbsession.Options)
	}

	return r0
}

// SendCommand provides a mock function with given fields: ctx, params, optFns
func (_m *QldbsessionClient) SendCommand(ctx context.Context, params *qldbsession.SendCommandInput, optFns ...func(*qldbsession.Options)) (*qldbsession.SendCommandOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendCommand")
	}

	var r0 *qldbsession.SendCommandOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *qldbsession.SendCommandInput, ...func(*qldbsession.Options)) (*qldbsession.SendCommandOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *qldbsession.SendCommandInput, ...func(*qldbsession.Options)) *qldbsession.SendCommandOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*qldbsession.SendCommandOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *qldbsession.SendCommandInput, ...func(*qldbsession.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQldbsessionClient creates a new instance of QldbsessionClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQldbsessionClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *QldbsessionClient {
	mock := &QldbsessionClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
