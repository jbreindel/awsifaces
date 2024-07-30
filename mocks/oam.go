// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	oam "github.com/aws/aws-sdk-go-v2/service/oam"
	mock "github.com/stretchr/testify/mock"
)

// OamClient is an autogenerated mock type for the OamClient type
type OamClient struct {
	mock.Mock
}

// CreateLink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) CreateLink(ctx context.Context, params *oam.CreateLinkInput, optFns ...func(*oam.Options)) (*oam.CreateLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLink")
	}

	var r0 *oam.CreateLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.CreateLinkInput, ...func(*oam.Options)) (*oam.CreateLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.CreateLinkInput, ...func(*oam.Options)) *oam.CreateLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.CreateLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.CreateLinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) CreateSink(ctx context.Context, params *oam.CreateSinkInput, optFns ...func(*oam.Options)) (*oam.CreateSinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSink")
	}

	var r0 *oam.CreateSinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.CreateSinkInput, ...func(*oam.Options)) (*oam.CreateSinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.CreateSinkInput, ...func(*oam.Options)) *oam.CreateSinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.CreateSinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.CreateSinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) DeleteLink(ctx context.Context, params *oam.DeleteLinkInput, optFns ...func(*oam.Options)) (*oam.DeleteLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLink")
	}

	var r0 *oam.DeleteLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.DeleteLinkInput, ...func(*oam.Options)) (*oam.DeleteLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.DeleteLinkInput, ...func(*oam.Options)) *oam.DeleteLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.DeleteLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.DeleteLinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) DeleteSink(ctx context.Context, params *oam.DeleteSinkInput, optFns ...func(*oam.Options)) (*oam.DeleteSinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSink")
	}

	var r0 *oam.DeleteSinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.DeleteSinkInput, ...func(*oam.Options)) (*oam.DeleteSinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.DeleteSinkInput, ...func(*oam.Options)) *oam.DeleteSinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.DeleteSinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.DeleteSinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) GetLink(ctx context.Context, params *oam.GetLinkInput, optFns ...func(*oam.Options)) (*oam.GetLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLink")
	}

	var r0 *oam.GetLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetLinkInput, ...func(*oam.Options)) (*oam.GetLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetLinkInput, ...func(*oam.Options)) *oam.GetLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.GetLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.GetLinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) GetSink(ctx context.Context, params *oam.GetSinkInput, optFns ...func(*oam.Options)) (*oam.GetSinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSink")
	}

	var r0 *oam.GetSinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetSinkInput, ...func(*oam.Options)) (*oam.GetSinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetSinkInput, ...func(*oam.Options)) *oam.GetSinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.GetSinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.GetSinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSinkPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) GetSinkPolicy(ctx context.Context, params *oam.GetSinkPolicyInput, optFns ...func(*oam.Options)) (*oam.GetSinkPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSinkPolicy")
	}

	var r0 *oam.GetSinkPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetSinkPolicyInput, ...func(*oam.Options)) (*oam.GetSinkPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.GetSinkPolicyInput, ...func(*oam.Options)) *oam.GetSinkPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.GetSinkPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.GetSinkPolicyInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAttachedLinks provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) ListAttachedLinks(ctx context.Context, params *oam.ListAttachedLinksInput, optFns ...func(*oam.Options)) (*oam.ListAttachedLinksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAttachedLinks")
	}

	var r0 *oam.ListAttachedLinksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListAttachedLinksInput, ...func(*oam.Options)) (*oam.ListAttachedLinksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListAttachedLinksInput, ...func(*oam.Options)) *oam.ListAttachedLinksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.ListAttachedLinksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.ListAttachedLinksInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLinks provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) ListLinks(ctx context.Context, params *oam.ListLinksInput, optFns ...func(*oam.Options)) (*oam.ListLinksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLinks")
	}

	var r0 *oam.ListLinksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListLinksInput, ...func(*oam.Options)) (*oam.ListLinksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListLinksInput, ...func(*oam.Options)) *oam.ListLinksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.ListLinksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.ListLinksInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSinks provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) ListSinks(ctx context.Context, params *oam.ListSinksInput, optFns ...func(*oam.Options)) (*oam.ListSinksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSinks")
	}

	var r0 *oam.ListSinksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListSinksInput, ...func(*oam.Options)) (*oam.ListSinksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListSinksInput, ...func(*oam.Options)) *oam.ListSinksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.ListSinksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.ListSinksInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) ListTagsForResource(ctx context.Context, params *oam.ListTagsForResourceInput, optFns ...func(*oam.Options)) (*oam.ListTagsForResourceOutput, error) {
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

	var r0 *oam.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListTagsForResourceInput, ...func(*oam.Options)) (*oam.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.ListTagsForResourceInput, ...func(*oam.Options)) *oam.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.ListTagsForResourceInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *OamClient) Options() oam.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 oam.Options
	if rf, ok := ret.Get(0).(func() oam.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(oam.Options)
	}

	return r0
}

// PutSinkPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) PutSinkPolicy(ctx context.Context, params *oam.PutSinkPolicyInput, optFns ...func(*oam.Options)) (*oam.PutSinkPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutSinkPolicy")
	}

	var r0 *oam.PutSinkPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.PutSinkPolicyInput, ...func(*oam.Options)) (*oam.PutSinkPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.PutSinkPolicyInput, ...func(*oam.Options)) *oam.PutSinkPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.PutSinkPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.PutSinkPolicyInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) TagResource(ctx context.Context, params *oam.TagResourceInput, optFns ...func(*oam.Options)) (*oam.TagResourceOutput, error) {
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

	var r0 *oam.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.TagResourceInput, ...func(*oam.Options)) (*oam.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.TagResourceInput, ...func(*oam.Options)) *oam.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.TagResourceInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) UntagResource(ctx context.Context, params *oam.UntagResourceInput, optFns ...func(*oam.Options)) (*oam.UntagResourceOutput, error) {
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

	var r0 *oam.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.UntagResourceInput, ...func(*oam.Options)) (*oam.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.UntagResourceInput, ...func(*oam.Options)) *oam.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.UntagResourceInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLink provides a mock function with given fields: ctx, params, optFns
func (_m *OamClient) UpdateLink(ctx context.Context, params *oam.UpdateLinkInput, optFns ...func(*oam.Options)) (*oam.UpdateLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLink")
	}

	var r0 *oam.UpdateLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oam.UpdateLinkInput, ...func(*oam.Options)) (*oam.UpdateLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oam.UpdateLinkInput, ...func(*oam.Options)) *oam.UpdateLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oam.UpdateLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oam.UpdateLinkInput, ...func(*oam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOamClient creates a new instance of OamClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOamClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OamClient {
	mock := &OamClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
