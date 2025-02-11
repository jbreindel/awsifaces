// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ec2instanceconnect "github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	mock "github.com/stretchr/testify/mock"
)

// Ec2instanceconnectClient is an autogenerated mock type for the Ec2instanceconnectClient type
type Ec2instanceconnectClient struct {
	mock.Mock
}

// Options provides a mock function with given fields:
func (_m *Ec2instanceconnectClient) Options() ec2instanceconnect.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 ec2instanceconnect.Options
	if rf, ok := ret.Get(0).(func() ec2instanceconnect.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ec2instanceconnect.Options)
	}

	return r0
}

// SendSSHPublicKey provides a mock function with given fields: ctx, params, optFns
func (_m *Ec2instanceconnectClient) SendSSHPublicKey(ctx context.Context, params *ec2instanceconnect.SendSSHPublicKeyInput, optFns ...func(*ec2instanceconnect.Options)) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendSSHPublicKey")
	}

	var r0 *ec2instanceconnect.SendSSHPublicKeyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ec2instanceconnect.SendSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) (*ec2instanceconnect.SendSSHPublicKeyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ec2instanceconnect.SendSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) *ec2instanceconnect.SendSSHPublicKeyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2instanceconnect.SendSSHPublicKeyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ec2instanceconnect.SendSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendSerialConsoleSSHPublicKey provides a mock function with given fields: ctx, params, optFns
func (_m *Ec2instanceconnectClient) SendSerialConsoleSSHPublicKey(ctx context.Context, params *ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput, optFns ...func(*ec2instanceconnect.Options)) (*ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendSerialConsoleSSHPublicKey")
	}

	var r0 *ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) (*ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) *ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2instanceconnect.SendSerialConsoleSSHPublicKeyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput, ...func(*ec2instanceconnect.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEc2instanceconnectClient creates a new instance of Ec2instanceconnectClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEc2instanceconnectClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Ec2instanceconnectClient {
	mock := &Ec2instanceconnectClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
