// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	iot1clickdevicesservice "github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	mock "github.com/stretchr/testify/mock"
)

// Iot1clickdevicesserviceClient is an autogenerated mock type for the Iot1clickdevicesserviceClient type
type Iot1clickdevicesserviceClient struct {
	mock.Mock
}

// ClaimDevicesByClaimCode provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) ClaimDevicesByClaimCode(ctx context.Context, params *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ClaimDevicesByClaimCode")
	}

	var r0 *iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDevice provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) DescribeDevice(ctx context.Context, params *iot1clickdevicesservice.DescribeDeviceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDevice")
	}

	var r0 *iot1clickdevicesservice.DescribeDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.DescribeDeviceInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.DescribeDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.DescribeDeviceInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.DescribeDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.DescribeDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.DescribeDeviceInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeDeviceClaim provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) FinalizeDeviceClaim(ctx context.Context, params *iot1clickdevicesservice.FinalizeDeviceClaimInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FinalizeDeviceClaim")
	}

	var r0 *iot1clickdevicesservice.FinalizeDeviceClaimOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.FinalizeDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.FinalizeDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.FinalizeDeviceClaimOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.FinalizeDeviceClaimOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.FinalizeDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceMethods provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) GetDeviceMethods(ctx context.Context, params *iot1clickdevicesservice.GetDeviceMethodsInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceMethods")
	}

	var r0 *iot1clickdevicesservice.GetDeviceMethodsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.GetDeviceMethodsInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.GetDeviceMethodsInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.GetDeviceMethodsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.GetDeviceMethodsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.GetDeviceMethodsInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateDeviceClaim provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) InitiateDeviceClaim(ctx context.Context, params *iot1clickdevicesservice.InitiateDeviceClaimInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InitiateDeviceClaim")
	}

	var r0 *iot1clickdevicesservice.InitiateDeviceClaimOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.InitiateDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.InitiateDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.InitiateDeviceClaimOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.InitiateDeviceClaimOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.InitiateDeviceClaimInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvokeDeviceMethod provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) InvokeDeviceMethod(ctx context.Context, params *iot1clickdevicesservice.InvokeDeviceMethodInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InvokeDeviceMethod")
	}

	var r0 *iot1clickdevicesservice.InvokeDeviceMethodOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.InvokeDeviceMethodInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.InvokeDeviceMethodInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.InvokeDeviceMethodOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.InvokeDeviceMethodOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.InvokeDeviceMethodInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeviceEvents provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) ListDeviceEvents(ctx context.Context, params *iot1clickdevicesservice.ListDeviceEventsInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDeviceEvents")
	}

	var r0 *iot1clickdevicesservice.ListDeviceEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListDeviceEventsInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDeviceEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListDeviceEventsInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.ListDeviceEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.ListDeviceEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.ListDeviceEventsInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDevices provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) ListDevices(ctx context.Context, params *iot1clickdevicesservice.ListDevicesInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 *iot1clickdevicesservice.ListDevicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListDevicesInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDevicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListDevicesInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.ListDevicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.ListDevicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.ListDevicesInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) ListTagsForResource(ctx context.Context, params *iot1clickdevicesservice.ListTagsForResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTagsForResource")
	}

	var r0 *iot1clickdevicesservice.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListTagsForResourceInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.ListTagsForResourceInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.ListTagsForResourceInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *Iot1clickdevicesserviceClient) Options() iot1clickdevicesservice.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 iot1clickdevicesservice.Options
	if rf, ok := ret.Get(0).(func() iot1clickdevicesservice.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(iot1clickdevicesservice.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) TagResource(ctx context.Context, params *iot1clickdevicesservice.TagResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.TagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TagResource")
	}

	var r0 *iot1clickdevicesservice.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.TagResourceInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.TagResourceInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.TagResourceInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnclaimDevice provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) UnclaimDevice(ctx context.Context, params *iot1clickdevicesservice.UnclaimDeviceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnclaimDevice")
	}

	var r0 *iot1clickdevicesservice.UnclaimDeviceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UnclaimDeviceInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UnclaimDeviceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UnclaimDeviceInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.UnclaimDeviceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.UnclaimDeviceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.UnclaimDeviceInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) UntagResource(ctx context.Context, params *iot1clickdevicesservice.UntagResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UntagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UntagResource")
	}

	var r0 *iot1clickdevicesservice.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UntagResourceInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UntagResourceInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.UntagResourceInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceState provides a mock function with given fields: ctx, params, optFns
func (_m *Iot1clickdevicesserviceClient) UpdateDeviceState(ctx context.Context, params *iot1clickdevicesservice.UpdateDeviceStateInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeviceState")
	}

	var r0 *iot1clickdevicesservice.UpdateDeviceStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UpdateDeviceStateInput, ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iot1clickdevicesservice.UpdateDeviceStateInput, ...func(*iot1clickdevicesservice.Options)) *iot1clickdevicesservice.UpdateDeviceStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iot1clickdevicesservice.UpdateDeviceStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iot1clickdevicesservice.UpdateDeviceStateInput, ...func(*iot1clickdevicesservice.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIot1clickdevicesserviceClient creates a new instance of Iot1clickdevicesserviceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIot1clickdevicesserviceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Iot1clickdevicesserviceClient {
	mock := &Iot1clickdevicesserviceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
