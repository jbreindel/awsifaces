// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	autoscalingplans "github.com/aws/aws-sdk-go-v2/service/autoscalingplans"

	mock "github.com/stretchr/testify/mock"
)

// AutoscalingplansClient is an autogenerated mock type for the AutoscalingplansClient type
type AutoscalingplansClient struct {
	mock.Mock
}

// CreateScalingPlan provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) CreateScalingPlan(ctx context.Context, params *autoscalingplans.CreateScalingPlanInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.CreateScalingPlanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateScalingPlan")
	}

	var r0 *autoscalingplans.CreateScalingPlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.CreateScalingPlanInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.CreateScalingPlanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.CreateScalingPlanInput, ...func(*autoscalingplans.Options)) *autoscalingplans.CreateScalingPlanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.CreateScalingPlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.CreateScalingPlanInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScalingPlan provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) DeleteScalingPlan(ctx context.Context, params *autoscalingplans.DeleteScalingPlanInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScalingPlan")
	}

	var r0 *autoscalingplans.DeleteScalingPlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DeleteScalingPlanInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.DeleteScalingPlanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DeleteScalingPlanInput, ...func(*autoscalingplans.Options)) *autoscalingplans.DeleteScalingPlanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.DeleteScalingPlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.DeleteScalingPlanInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScalingPlanResources provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) DescribeScalingPlanResources(ctx context.Context, params *autoscalingplans.DescribeScalingPlanResourcesInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScalingPlanResources")
	}

	var r0 *autoscalingplans.DescribeScalingPlanResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DescribeScalingPlanResourcesInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DescribeScalingPlanResourcesInput, ...func(*autoscalingplans.Options)) *autoscalingplans.DescribeScalingPlanResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.DescribeScalingPlanResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.DescribeScalingPlanResourcesInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScalingPlans provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) DescribeScalingPlans(ctx context.Context, params *autoscalingplans.DescribeScalingPlansInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeScalingPlans")
	}

	var r0 *autoscalingplans.DescribeScalingPlansOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DescribeScalingPlansInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.DescribeScalingPlansOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.DescribeScalingPlansInput, ...func(*autoscalingplans.Options)) *autoscalingplans.DescribeScalingPlansOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.DescribeScalingPlansOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.DescribeScalingPlansInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScalingPlanResourceForecastData provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) GetScalingPlanResourceForecastData(ctx context.Context, params *autoscalingplans.GetScalingPlanResourceForecastDataInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetScalingPlanResourceForecastData")
	}

	var r0 *autoscalingplans.GetScalingPlanResourceForecastDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.GetScalingPlanResourceForecastDataInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.GetScalingPlanResourceForecastDataInput, ...func(*autoscalingplans.Options)) *autoscalingplans.GetScalingPlanResourceForecastDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.GetScalingPlanResourceForecastDataInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *AutoscalingplansClient) Options() autoscalingplans.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 autoscalingplans.Options
	if rf, ok := ret.Get(0).(func() autoscalingplans.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(autoscalingplans.Options)
	}

	return r0
}

// UpdateScalingPlan provides a mock function with given fields: ctx, params, optFns
func (_m *AutoscalingplansClient) UpdateScalingPlan(ctx context.Context, params *autoscalingplans.UpdateScalingPlanInput, optFns ...func(*autoscalingplans.Options)) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateScalingPlan")
	}

	var r0 *autoscalingplans.UpdateScalingPlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.UpdateScalingPlanInput, ...func(*autoscalingplans.Options)) (*autoscalingplans.UpdateScalingPlanOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingplans.UpdateScalingPlanInput, ...func(*autoscalingplans.Options)) *autoscalingplans.UpdateScalingPlanOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingplans.UpdateScalingPlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingplans.UpdateScalingPlanInput, ...func(*autoscalingplans.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAutoscalingplansClient creates a new instance of AutoscalingplansClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAutoscalingplansClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AutoscalingplansClient {
	mock := &AutoscalingplansClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
