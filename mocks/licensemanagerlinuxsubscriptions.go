// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	licensemanagerlinuxsubscriptions "github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
	mock "github.com/stretchr/testify/mock"
)

// LicensemanagerlinuxsubscriptionsClient is an autogenerated mock type for the LicensemanagerlinuxsubscriptionsClient type
type LicensemanagerlinuxsubscriptionsClient struct {
	mock.Mock
}

// GetServiceSettings provides a mock function with given fields: ctx, params, optFns
func (_m *LicensemanagerlinuxsubscriptionsClient) GetServiceSettings(ctx context.Context, params *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, optFns ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetServiceSettings")
	}

	var r0 *licensemanagerlinuxsubscriptions.GetServiceSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) *licensemanagerlinuxsubscriptions.GetServiceSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLinuxSubscriptionInstances provides a mock function with given fields: ctx, params, optFns
func (_m *LicensemanagerlinuxsubscriptionsClient) ListLinuxSubscriptionInstances(ctx context.Context, params *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, optFns ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLinuxSubscriptionInstances")
	}

	var r0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, ...func(*licensemanagerlinuxsubscriptions.Options)) *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, ...func(*licensemanagerlinuxsubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLinuxSubscriptions provides a mock function with given fields: ctx, params, optFns
func (_m *LicensemanagerlinuxsubscriptionsClient) ListLinuxSubscriptions(ctx context.Context, params *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, optFns ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLinuxSubscriptions")
	}

	var r0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *LicensemanagerlinuxsubscriptionsClient) Options() licensemanagerlinuxsubscriptions.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 licensemanagerlinuxsubscriptions.Options
	if rf, ok := ret.Get(0).(func() licensemanagerlinuxsubscriptions.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(licensemanagerlinuxsubscriptions.Options)
	}

	return r0
}

// UpdateServiceSettings provides a mock function with given fields: ctx, params, optFns
func (_m *LicensemanagerlinuxsubscriptionsClient) UpdateServiceSettings(ctx context.Context, params *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput, optFns ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateServiceSettings")
	}

	var r0 *licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) (*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) *licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput, ...func(*licensemanagerlinuxsubscriptions.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLicensemanagerlinuxsubscriptionsClient creates a new instance of LicensemanagerlinuxsubscriptionsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLicensemanagerlinuxsubscriptionsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *LicensemanagerlinuxsubscriptionsClient {
	mock := &LicensemanagerlinuxsubscriptionsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
