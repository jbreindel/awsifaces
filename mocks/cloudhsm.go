// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cloudhsm "github.com/aws/aws-sdk-go-v2/service/cloudhsm"

	mock "github.com/stretchr/testify/mock"
)

// CloudhsmClient is an autogenerated mock type for the CloudhsmClient type
type CloudhsmClient struct {
	mock.Mock
}

// AddTagsToResource provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) AddTagsToResource(ctx context.Context, params *cloudhsm.AddTagsToResourceInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.AddTagsToResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddTagsToResource")
	}

	var r0 *cloudhsm.AddTagsToResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.AddTagsToResourceInput, ...func(*cloudhsm.Options)) (*cloudhsm.AddTagsToResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.AddTagsToResourceInput, ...func(*cloudhsm.Options)) *cloudhsm.AddTagsToResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.AddTagsToResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.AddTagsToResourceInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHapg provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) CreateHapg(ctx context.Context, params *cloudhsm.CreateHapgInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.CreateHapgOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateHapg")
	}

	var r0 *cloudhsm.CreateHapgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateHapgInput, ...func(*cloudhsm.Options)) (*cloudhsm.CreateHapgOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateHapgInput, ...func(*cloudhsm.Options)) *cloudhsm.CreateHapgOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.CreateHapgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.CreateHapgInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHsm provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) CreateHsm(ctx context.Context, params *cloudhsm.CreateHsmInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.CreateHsmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateHsm")
	}

	var r0 *cloudhsm.CreateHsmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateHsmInput, ...func(*cloudhsm.Options)) (*cloudhsm.CreateHsmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateHsmInput, ...func(*cloudhsm.Options)) *cloudhsm.CreateHsmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.CreateHsmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.CreateHsmInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLunaClient provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) CreateLunaClient(ctx context.Context, params *cloudhsm.CreateLunaClientInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.CreateLunaClientOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLunaClient")
	}

	var r0 *cloudhsm.CreateLunaClientOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateLunaClientInput, ...func(*cloudhsm.Options)) (*cloudhsm.CreateLunaClientOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.CreateLunaClientInput, ...func(*cloudhsm.Options)) *cloudhsm.CreateLunaClientOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.CreateLunaClientOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.CreateLunaClientInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHapg provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DeleteHapg(ctx context.Context, params *cloudhsm.DeleteHapgInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DeleteHapgOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteHapg")
	}

	var r0 *cloudhsm.DeleteHapgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteHapgInput, ...func(*cloudhsm.Options)) (*cloudhsm.DeleteHapgOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteHapgInput, ...func(*cloudhsm.Options)) *cloudhsm.DeleteHapgOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DeleteHapgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DeleteHapgInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHsm provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DeleteHsm(ctx context.Context, params *cloudhsm.DeleteHsmInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DeleteHsmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteHsm")
	}

	var r0 *cloudhsm.DeleteHsmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteHsmInput, ...func(*cloudhsm.Options)) (*cloudhsm.DeleteHsmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteHsmInput, ...func(*cloudhsm.Options)) *cloudhsm.DeleteHsmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DeleteHsmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DeleteHsmInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLunaClient provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DeleteLunaClient(ctx context.Context, params *cloudhsm.DeleteLunaClientInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DeleteLunaClientOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLunaClient")
	}

	var r0 *cloudhsm.DeleteLunaClientOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteLunaClientInput, ...func(*cloudhsm.Options)) (*cloudhsm.DeleteLunaClientOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DeleteLunaClientInput, ...func(*cloudhsm.Options)) *cloudhsm.DeleteLunaClientOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DeleteLunaClientOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DeleteLunaClientInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHapg provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DescribeHapg(ctx context.Context, params *cloudhsm.DescribeHapgInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DescribeHapgOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeHapg")
	}

	var r0 *cloudhsm.DescribeHapgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeHapgInput, ...func(*cloudhsm.Options)) (*cloudhsm.DescribeHapgOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeHapgInput, ...func(*cloudhsm.Options)) *cloudhsm.DescribeHapgOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DescribeHapgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DescribeHapgInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHsm provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DescribeHsm(ctx context.Context, params *cloudhsm.DescribeHsmInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DescribeHsmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeHsm")
	}

	var r0 *cloudhsm.DescribeHsmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeHsmInput, ...func(*cloudhsm.Options)) (*cloudhsm.DescribeHsmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeHsmInput, ...func(*cloudhsm.Options)) *cloudhsm.DescribeHsmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DescribeHsmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DescribeHsmInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLunaClient provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) DescribeLunaClient(ctx context.Context, params *cloudhsm.DescribeLunaClientInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.DescribeLunaClientOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeLunaClient")
	}

	var r0 *cloudhsm.DescribeLunaClientOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeLunaClientInput, ...func(*cloudhsm.Options)) (*cloudhsm.DescribeLunaClientOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.DescribeLunaClientInput, ...func(*cloudhsm.Options)) *cloudhsm.DescribeLunaClientOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.DescribeLunaClientOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.DescribeLunaClientInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) GetConfig(ctx context.Context, params *cloudhsm.GetConfigInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.GetConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConfig")
	}

	var r0 *cloudhsm.GetConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.GetConfigInput, ...func(*cloudhsm.Options)) (*cloudhsm.GetConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.GetConfigInput, ...func(*cloudhsm.Options)) *cloudhsm.GetConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.GetConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.GetConfigInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAvailableZones provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ListAvailableZones(ctx context.Context, params *cloudhsm.ListAvailableZonesInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ListAvailableZonesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAvailableZones")
	}

	var r0 *cloudhsm.ListAvailableZonesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListAvailableZonesInput, ...func(*cloudhsm.Options)) (*cloudhsm.ListAvailableZonesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListAvailableZonesInput, ...func(*cloudhsm.Options)) *cloudhsm.ListAvailableZonesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ListAvailableZonesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ListAvailableZonesInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHapgs provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ListHapgs(ctx context.Context, params *cloudhsm.ListHapgsInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ListHapgsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListHapgs")
	}

	var r0 *cloudhsm.ListHapgsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListHapgsInput, ...func(*cloudhsm.Options)) (*cloudhsm.ListHapgsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListHapgsInput, ...func(*cloudhsm.Options)) *cloudhsm.ListHapgsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ListHapgsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ListHapgsInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHsms provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ListHsms(ctx context.Context, params *cloudhsm.ListHsmsInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ListHsmsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListHsms")
	}

	var r0 *cloudhsm.ListHsmsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListHsmsInput, ...func(*cloudhsm.Options)) (*cloudhsm.ListHsmsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListHsmsInput, ...func(*cloudhsm.Options)) *cloudhsm.ListHsmsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ListHsmsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ListHsmsInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLunaClients provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ListLunaClients(ctx context.Context, params *cloudhsm.ListLunaClientsInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ListLunaClientsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLunaClients")
	}

	var r0 *cloudhsm.ListLunaClientsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListLunaClientsInput, ...func(*cloudhsm.Options)) (*cloudhsm.ListLunaClientsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListLunaClientsInput, ...func(*cloudhsm.Options)) *cloudhsm.ListLunaClientsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ListLunaClientsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ListLunaClientsInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ListTagsForResource(ctx context.Context, params *cloudhsm.ListTagsForResourceInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ListTagsForResourceOutput, error) {
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

	var r0 *cloudhsm.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListTagsForResourceInput, ...func(*cloudhsm.Options)) (*cloudhsm.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ListTagsForResourceInput, ...func(*cloudhsm.Options)) *cloudhsm.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ListTagsForResourceInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyHapg provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ModifyHapg(ctx context.Context, params *cloudhsm.ModifyHapgInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ModifyHapgOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ModifyHapg")
	}

	var r0 *cloudhsm.ModifyHapgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyHapgInput, ...func(*cloudhsm.Options)) (*cloudhsm.ModifyHapgOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyHapgInput, ...func(*cloudhsm.Options)) *cloudhsm.ModifyHapgOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ModifyHapgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ModifyHapgInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyHsm provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ModifyHsm(ctx context.Context, params *cloudhsm.ModifyHsmInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ModifyHsmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ModifyHsm")
	}

	var r0 *cloudhsm.ModifyHsmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyHsmInput, ...func(*cloudhsm.Options)) (*cloudhsm.ModifyHsmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyHsmInput, ...func(*cloudhsm.Options)) *cloudhsm.ModifyHsmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ModifyHsmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ModifyHsmInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyLunaClient provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) ModifyLunaClient(ctx context.Context, params *cloudhsm.ModifyLunaClientInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.ModifyLunaClientOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ModifyLunaClient")
	}

	var r0 *cloudhsm.ModifyLunaClientOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyLunaClientInput, ...func(*cloudhsm.Options)) (*cloudhsm.ModifyLunaClientOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.ModifyLunaClientInput, ...func(*cloudhsm.Options)) *cloudhsm.ModifyLunaClientOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.ModifyLunaClientOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.ModifyLunaClientInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *CloudhsmClient) Options() cloudhsm.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 cloudhsm.Options
	if rf, ok := ret.Get(0).(func() cloudhsm.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cloudhsm.Options)
	}

	return r0
}

// RemoveTagsFromResource provides a mock function with given fields: ctx, params, optFns
func (_m *CloudhsmClient) RemoveTagsFromResource(ctx context.Context, params *cloudhsm.RemoveTagsFromResourceInput, optFns ...func(*cloudhsm.Options)) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTagsFromResource")
	}

	var r0 *cloudhsm.RemoveTagsFromResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.RemoveTagsFromResourceInput, ...func(*cloudhsm.Options)) (*cloudhsm.RemoveTagsFromResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cloudhsm.RemoveTagsFromResourceInput, ...func(*cloudhsm.Options)) *cloudhsm.RemoveTagsFromResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudhsm.RemoveTagsFromResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cloudhsm.RemoveTagsFromResourceInput, ...func(*cloudhsm.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCloudhsmClient creates a new instance of CloudhsmClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudhsmClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudhsmClient {
	mock := &CloudhsmClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
