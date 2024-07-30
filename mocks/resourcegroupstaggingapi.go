// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	resourcegroupstaggingapi "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	mock "github.com/stretchr/testify/mock"
)

// ResourcegroupstaggingapiClient is an autogenerated mock type for the ResourcegroupstaggingapiClient type
type ResourcegroupstaggingapiClient struct {
	mock.Mock
}

// DescribeReportCreation provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) DescribeReportCreation(ctx context.Context, params *resourcegroupstaggingapi.DescribeReportCreationInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeReportCreation")
	}

	var r0 *resourcegroupstaggingapi.DescribeReportCreationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.DescribeReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.DescribeReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.DescribeReportCreationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.DescribeReportCreationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.DescribeReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetComplianceSummary provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) GetComplianceSummary(ctx context.Context, params *resourcegroupstaggingapi.GetComplianceSummaryInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetComplianceSummary")
	}

	var r0 *resourcegroupstaggingapi.GetComplianceSummaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetComplianceSummaryInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetComplianceSummaryInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.GetComplianceSummaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.GetComplianceSummaryInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResources provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) GetResources(ctx context.Context, params *resourcegroupstaggingapi.GetResourcesInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResources")
	}

	var r0 *resourcegroupstaggingapi.GetResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.GetResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.GetResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagKeys provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) GetTagKeys(ctx context.Context, params *resourcegroupstaggingapi.GetTagKeysInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTagKeys")
	}

	var r0 *resourcegroupstaggingapi.GetTagKeysOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetTagKeysInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetTagKeysOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetTagKeysInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.GetTagKeysOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.GetTagKeysOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.GetTagKeysInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagValues provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) GetTagValues(ctx context.Context, params *resourcegroupstaggingapi.GetTagValuesInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTagValues")
	}

	var r0 *resourcegroupstaggingapi.GetTagValuesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetTagValuesInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.GetTagValuesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetTagValuesInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.GetTagValuesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.GetTagValuesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.GetTagValuesInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *ResourcegroupstaggingapiClient) Options() resourcegroupstaggingapi.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 resourcegroupstaggingapi.Options
	if rf, ok := ret.Get(0).(func() resourcegroupstaggingapi.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(resourcegroupstaggingapi.Options)
	}

	return r0
}

// StartReportCreation provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) StartReportCreation(ctx context.Context, params *resourcegroupstaggingapi.StartReportCreationInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartReportCreation")
	}

	var r0 *resourcegroupstaggingapi.StartReportCreationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.StartReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.StartReportCreationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.StartReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.StartReportCreationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.StartReportCreationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.StartReportCreationInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResources provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) TagResources(ctx context.Context, params *resourcegroupstaggingapi.TagResourcesInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TagResources")
	}

	var r0 *resourcegroupstaggingapi.TagResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.TagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.TagResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.TagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.TagResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.TagResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.TagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResources provides a mock function with given fields: ctx, params, optFns
func (_m *ResourcegroupstaggingapiClient) UntagResources(ctx context.Context, params *resourcegroupstaggingapi.UntagResourcesInput, optFns ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UntagResources")
	}

	var r0 *resourcegroupstaggingapi.UntagResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.UntagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) (*resourcegroupstaggingapi.UntagResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.UntagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) *resourcegroupstaggingapi.UntagResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.UntagResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.UntagResourcesInput, ...func(*resourcegroupstaggingapi.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResourcegroupstaggingapiClient creates a new instance of ResourcegroupstaggingapiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResourcegroupstaggingapiClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResourcegroupstaggingapiClient {
	mock := &ResourcegroupstaggingapiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
