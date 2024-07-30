// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cognitoidentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"

	mock "github.com/stretchr/testify/mock"
)

// CognitoidentityClient is an autogenerated mock type for the CognitoidentityClient type
type CognitoidentityClient struct {
	mock.Mock
}

// CreateIdentityPool provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) CreateIdentityPool(ctx context.Context, params *cognitoidentity.CreateIdentityPoolInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.CreateIdentityPoolOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateIdentityPool")
	}

	var r0 *cognitoidentity.CreateIdentityPoolOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.CreateIdentityPoolInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.CreateIdentityPoolOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.CreateIdentityPoolInput, ...func(*cognitoidentity.Options)) *cognitoidentity.CreateIdentityPoolOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.CreateIdentityPoolOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.CreateIdentityPoolInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIdentities provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) DeleteIdentities(ctx context.Context, params *cognitoidentity.DeleteIdentitiesInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIdentities")
	}

	var r0 *cognitoidentity.DeleteIdentitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DeleteIdentitiesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DeleteIdentitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DeleteIdentitiesInput, ...func(*cognitoidentity.Options)) *cognitoidentity.DeleteIdentitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.DeleteIdentitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.DeleteIdentitiesInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIdentityPool provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) DeleteIdentityPool(ctx context.Context, params *cognitoidentity.DeleteIdentityPoolInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIdentityPool")
	}

	var r0 *cognitoidentity.DeleteIdentityPoolOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DeleteIdentityPoolInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DeleteIdentityPoolOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DeleteIdentityPoolInput, ...func(*cognitoidentity.Options)) *cognitoidentity.DeleteIdentityPoolOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.DeleteIdentityPoolOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.DeleteIdentityPoolInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) DescribeIdentity(ctx context.Context, params *cognitoidentity.DescribeIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeIdentity")
	}

	var r0 *cognitoidentity.DescribeIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DescribeIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DescribeIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.DescribeIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.DescribeIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.DescribeIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentityPool provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) DescribeIdentityPool(ctx context.Context, params *cognitoidentity.DescribeIdentityPoolInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeIdentityPool")
	}

	var r0 *cognitoidentity.DescribeIdentityPoolOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DescribeIdentityPoolInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.DescribeIdentityPoolInput, ...func(*cognitoidentity.Options)) *cognitoidentity.DescribeIdentityPoolOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.DescribeIdentityPoolOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.DescribeIdentityPoolInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCredentialsForIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetCredentialsForIdentity(ctx context.Context, params *cognitoidentity.GetCredentialsForIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCredentialsForIdentity")
	}

	var r0 *cognitoidentity.GetCredentialsForIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetCredentialsForIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetCredentialsForIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetCredentialsForIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetCredentialsForIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetCredentialsForIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetCredentialsForIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetId(ctx context.Context, params *cognitoidentity.GetIdInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetId")
	}

	var r0 *cognitoidentity.GetIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetIdInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetIdInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetIdInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityPoolRoles provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetIdentityPoolRoles(ctx context.Context, params *cognitoidentity.GetIdentityPoolRolesInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIdentityPoolRoles")
	}

	var r0 *cognitoidentity.GetIdentityPoolRolesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdentityPoolRolesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetIdentityPoolRolesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetIdentityPoolRolesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOpenIdToken provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetOpenIdToken(ctx context.Context, params *cognitoidentity.GetOpenIdTokenInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOpenIdToken")
	}

	var r0 *cognitoidentity.GetOpenIdTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetOpenIdTokenInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetOpenIdTokenInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetOpenIdTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetOpenIdTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetOpenIdTokenInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOpenIdTokenForDeveloperIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetOpenIdTokenForDeveloperIdentity(ctx context.Context, params *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOpenIdTokenForDeveloperIdentity")
	}

	var r0 *cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrincipalTagAttributeMap provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) GetPrincipalTagAttributeMap(ctx context.Context, params *cognitoidentity.GetPrincipalTagAttributeMapInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPrincipalTagAttributeMap")
	}

	var r0 *cognitoidentity.GetPrincipalTagAttributeMapOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.GetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) *cognitoidentity.GetPrincipalTagAttributeMapOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.GetPrincipalTagAttributeMapOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.GetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentities provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) ListIdentities(ctx context.Context, params *cognitoidentity.ListIdentitiesInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdentities")
	}

	var r0 *cognitoidentity.ListIdentitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListIdentitiesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListIdentitiesInput, ...func(*cognitoidentity.Options)) *cognitoidentity.ListIdentitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.ListIdentitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.ListIdentitiesInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentityPools provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) ListIdentityPools(ctx context.Context, params *cognitoidentity.ListIdentityPoolsInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdentityPools")
	}

	var r0 *cognitoidentity.ListIdentityPoolsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListIdentityPoolsInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListIdentityPoolsInput, ...func(*cognitoidentity.Options)) *cognitoidentity.ListIdentityPoolsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.ListIdentityPoolsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.ListIdentityPoolsInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) ListTagsForResource(ctx context.Context, params *cognitoidentity.ListTagsForResourceInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.ListTagsForResourceOutput, error) {
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

	var r0 *cognitoidentity.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListTagsForResourceInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.ListTagsForResourceInput, ...func(*cognitoidentity.Options)) *cognitoidentity.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.ListTagsForResourceInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupDeveloperIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) LookupDeveloperIdentity(ctx context.Context, params *cognitoidentity.LookupDeveloperIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for LookupDeveloperIdentity")
	}

	var r0 *cognitoidentity.LookupDeveloperIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.LookupDeveloperIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.LookupDeveloperIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.LookupDeveloperIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.LookupDeveloperIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.LookupDeveloperIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.LookupDeveloperIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MergeDeveloperIdentities provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) MergeDeveloperIdentities(ctx context.Context, params *cognitoidentity.MergeDeveloperIdentitiesInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MergeDeveloperIdentities")
	}

	var r0 *cognitoidentity.MergeDeveloperIdentitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.MergeDeveloperIdentitiesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.MergeDeveloperIdentitiesInput, ...func(*cognitoidentity.Options)) *cognitoidentity.MergeDeveloperIdentitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.MergeDeveloperIdentitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.MergeDeveloperIdentitiesInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *CognitoidentityClient) Options() cognitoidentity.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 cognitoidentity.Options
	if rf, ok := ret.Get(0).(func() cognitoidentity.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cognitoidentity.Options)
	}

	return r0
}

// SetIdentityPoolRoles provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) SetIdentityPoolRoles(ctx context.Context, params *cognitoidentity.SetIdentityPoolRolesInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetIdentityPoolRoles")
	}

	var r0 *cognitoidentity.SetIdentityPoolRolesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.SetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.SetIdentityPoolRolesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.SetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) *cognitoidentity.SetIdentityPoolRolesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.SetIdentityPoolRolesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.SetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPrincipalTagAttributeMap provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) SetPrincipalTagAttributeMap(ctx context.Context, params *cognitoidentity.SetPrincipalTagAttributeMapInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.SetPrincipalTagAttributeMapOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetPrincipalTagAttributeMap")
	}

	var r0 *cognitoidentity.SetPrincipalTagAttributeMapOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.SetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.SetPrincipalTagAttributeMapOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.SetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) *cognitoidentity.SetPrincipalTagAttributeMapOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.SetPrincipalTagAttributeMapOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.SetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) TagResource(ctx context.Context, params *cognitoidentity.TagResourceInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.TagResourceOutput, error) {
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

	var r0 *cognitoidentity.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.TagResourceInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.TagResourceInput, ...func(*cognitoidentity.Options)) *cognitoidentity.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.TagResourceInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlinkDeveloperIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) UnlinkDeveloperIdentity(ctx context.Context, params *cognitoidentity.UnlinkDeveloperIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnlinkDeveloperIdentity")
	}

	var r0 *cognitoidentity.UnlinkDeveloperIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UnlinkDeveloperIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UnlinkDeveloperIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.UnlinkDeveloperIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.UnlinkDeveloperIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.UnlinkDeveloperIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlinkIdentity provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) UnlinkIdentity(ctx context.Context, params *cognitoidentity.UnlinkIdentityInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.UnlinkIdentityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnlinkIdentity")
	}

	var r0 *cognitoidentity.UnlinkIdentityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UnlinkIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.UnlinkIdentityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UnlinkIdentityInput, ...func(*cognitoidentity.Options)) *cognitoidentity.UnlinkIdentityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.UnlinkIdentityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.UnlinkIdentityInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) UntagResource(ctx context.Context, params *cognitoidentity.UntagResourceInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.UntagResourceOutput, error) {
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

	var r0 *cognitoidentity.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UntagResourceInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UntagResourceInput, ...func(*cognitoidentity.Options)) *cognitoidentity.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.UntagResourceInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentityPool provides a mock function with given fields: ctx, params, optFns
func (_m *CognitoidentityClient) UpdateIdentityPool(ctx context.Context, params *cognitoidentity.UpdateIdentityPoolInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.UpdateIdentityPoolOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIdentityPool")
	}

	var r0 *cognitoidentity.UpdateIdentityPoolOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UpdateIdentityPoolInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.UpdateIdentityPoolOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cognitoidentity.UpdateIdentityPoolInput, ...func(*cognitoidentity.Options)) *cognitoidentity.UpdateIdentityPoolOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentity.UpdateIdentityPoolOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cognitoidentity.UpdateIdentityPoolInput, ...func(*cognitoidentity.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCognitoidentityClient creates a new instance of CognitoidentityClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCognitoidentityClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CognitoidentityClient {
	mock := &CognitoidentityClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
