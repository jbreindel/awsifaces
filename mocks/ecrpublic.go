// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ecrpublic "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	mock "github.com/stretchr/testify/mock"
)

// EcrpublicClient is an autogenerated mock type for the EcrpublicClient type
type EcrpublicClient struct {
	mock.Mock
}

// BatchCheckLayerAvailability provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) BatchCheckLayerAvailability(ctx context.Context, params *ecrpublic.BatchCheckLayerAvailabilityInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchCheckLayerAvailability")
	}

	var r0 *ecrpublic.BatchCheckLayerAvailabilityOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.BatchCheckLayerAvailabilityInput, ...func(*ecrpublic.Options)) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.BatchCheckLayerAvailabilityInput, ...func(*ecrpublic.Options)) *ecrpublic.BatchCheckLayerAvailabilityOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.BatchCheckLayerAvailabilityOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.BatchCheckLayerAvailabilityInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchDeleteImage provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) BatchDeleteImage(ctx context.Context, params *ecrpublic.BatchDeleteImageInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.BatchDeleteImageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeleteImage")
	}

	var r0 *ecrpublic.BatchDeleteImageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.BatchDeleteImageInput, ...func(*ecrpublic.Options)) (*ecrpublic.BatchDeleteImageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.BatchDeleteImageInput, ...func(*ecrpublic.Options)) *ecrpublic.BatchDeleteImageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.BatchDeleteImageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.BatchDeleteImageInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteLayerUpload provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) CompleteLayerUpload(ctx context.Context, params *ecrpublic.CompleteLayerUploadInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.CompleteLayerUploadOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CompleteLayerUpload")
	}

	var r0 *ecrpublic.CompleteLayerUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.CompleteLayerUploadInput, ...func(*ecrpublic.Options)) (*ecrpublic.CompleteLayerUploadOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.CompleteLayerUploadInput, ...func(*ecrpublic.Options)) *ecrpublic.CompleteLayerUploadOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.CompleteLayerUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.CompleteLayerUploadInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) CreateRepository(ctx context.Context, params *ecrpublic.CreateRepositoryInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.CreateRepositoryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRepository")
	}

	var r0 *ecrpublic.CreateRepositoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.CreateRepositoryInput, ...func(*ecrpublic.Options)) (*ecrpublic.CreateRepositoryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.CreateRepositoryInput, ...func(*ecrpublic.Options)) *ecrpublic.CreateRepositoryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.CreateRepositoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.CreateRepositoryInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepository provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DeleteRepository(ctx context.Context, params *ecrpublic.DeleteRepositoryInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DeleteRepositoryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRepository")
	}

	var r0 *ecrpublic.DeleteRepositoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DeleteRepositoryInput, ...func(*ecrpublic.Options)) (*ecrpublic.DeleteRepositoryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DeleteRepositoryInput, ...func(*ecrpublic.Options)) *ecrpublic.DeleteRepositoryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DeleteRepositoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DeleteRepositoryInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepositoryPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DeleteRepositoryPolicy(ctx context.Context, params *ecrpublic.DeleteRepositoryPolicyInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DeleteRepositoryPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRepositoryPolicy")
	}

	var r0 *ecrpublic.DeleteRepositoryPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DeleteRepositoryPolicyInput, ...func(*ecrpublic.Options)) (*ecrpublic.DeleteRepositoryPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DeleteRepositoryPolicyInput, ...func(*ecrpublic.Options)) *ecrpublic.DeleteRepositoryPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DeleteRepositoryPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DeleteRepositoryPolicyInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImageTags provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DescribeImageTags(ctx context.Context, params *ecrpublic.DescribeImageTagsInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeImageTags")
	}

	var r0 *ecrpublic.DescribeImageTagsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeImageTagsInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeImageTagsInput, ...func(*ecrpublic.Options)) *ecrpublic.DescribeImageTagsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DescribeImageTagsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DescribeImageTagsInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImages provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DescribeImages(ctx context.Context, params *ecrpublic.DescribeImagesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeImages")
	}

	var r0 *ecrpublic.DescribeImagesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeImagesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeImagesInput, ...func(*ecrpublic.Options)) *ecrpublic.DescribeImagesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DescribeImagesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DescribeImagesInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRegistries provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DescribeRegistries(ctx context.Context, params *ecrpublic.DescribeRegistriesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRegistries")
	}

	var r0 *ecrpublic.DescribeRegistriesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeRegistriesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeRegistriesInput, ...func(*ecrpublic.Options)) *ecrpublic.DescribeRegistriesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DescribeRegistriesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DescribeRegistriesInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRepositories provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) DescribeRepositories(ctx context.Context, params *ecrpublic.DescribeRepositoriesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRepositories")
	}

	var r0 *ecrpublic.DescribeRepositoriesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeRepositoriesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.DescribeRepositoriesInput, ...func(*ecrpublic.Options)) *ecrpublic.DescribeRepositoriesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.DescribeRepositoriesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.DescribeRepositoriesInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorizationToken provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) GetAuthorizationToken(ctx context.Context, params *ecrpublic.GetAuthorizationTokenInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationToken")
	}

	var r0 *ecrpublic.GetAuthorizationTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetAuthorizationTokenInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetAuthorizationTokenInput, ...func(*ecrpublic.Options)) *ecrpublic.GetAuthorizationTokenOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.GetAuthorizationTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.GetAuthorizationTokenInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistryCatalogData provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) GetRegistryCatalogData(ctx context.Context, params *ecrpublic.GetRegistryCatalogDataInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRegistryCatalogData")
	}

	var r0 *ecrpublic.GetRegistryCatalogDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRegistryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRegistryCatalogDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRegistryCatalogDataInput, ...func(*ecrpublic.Options)) *ecrpublic.GetRegistryCatalogDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.GetRegistryCatalogDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.GetRegistryCatalogDataInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryCatalogData provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) GetRepositoryCatalogData(ctx context.Context, params *ecrpublic.GetRepositoryCatalogDataInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRepositoryCatalogData")
	}

	var r0 *ecrpublic.GetRepositoryCatalogDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryCatalogDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) *ecrpublic.GetRepositoryCatalogDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.GetRepositoryCatalogDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.GetRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) GetRepositoryPolicy(ctx context.Context, params *ecrpublic.GetRepositoryPolicyInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRepositoryPolicy")
	}

	var r0 *ecrpublic.GetRepositoryPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRepositoryPolicyInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.GetRepositoryPolicyInput, ...func(*ecrpublic.Options)) *ecrpublic.GetRepositoryPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.GetRepositoryPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.GetRepositoryPolicyInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateLayerUpload provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) InitiateLayerUpload(ctx context.Context, params *ecrpublic.InitiateLayerUploadInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.InitiateLayerUploadOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InitiateLayerUpload")
	}

	var r0 *ecrpublic.InitiateLayerUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.InitiateLayerUploadInput, ...func(*ecrpublic.Options)) (*ecrpublic.InitiateLayerUploadOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.InitiateLayerUploadInput, ...func(*ecrpublic.Options)) *ecrpublic.InitiateLayerUploadOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.InitiateLayerUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.InitiateLayerUploadInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) ListTagsForResource(ctx context.Context, params *ecrpublic.ListTagsForResourceInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error) {
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

	var r0 *ecrpublic.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.ListTagsForResourceInput, ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.ListTagsForResourceInput, ...func(*ecrpublic.Options)) *ecrpublic.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.ListTagsForResourceInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *EcrpublicClient) Options() ecrpublic.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 ecrpublic.Options
	if rf, ok := ret.Get(0).(func() ecrpublic.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ecrpublic.Options)
	}

	return r0
}

// PutImage provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) PutImage(ctx context.Context, params *ecrpublic.PutImageInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.PutImageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutImage")
	}

	var r0 *ecrpublic.PutImageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutImageInput, ...func(*ecrpublic.Options)) (*ecrpublic.PutImageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutImageInput, ...func(*ecrpublic.Options)) *ecrpublic.PutImageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.PutImageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.PutImageInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutRegistryCatalogData provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) PutRegistryCatalogData(ctx context.Context, params *ecrpublic.PutRegistryCatalogDataInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.PutRegistryCatalogDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRegistryCatalogData")
	}

	var r0 *ecrpublic.PutRegistryCatalogDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutRegistryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.PutRegistryCatalogDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutRegistryCatalogDataInput, ...func(*ecrpublic.Options)) *ecrpublic.PutRegistryCatalogDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.PutRegistryCatalogDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.PutRegistryCatalogDataInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutRepositoryCatalogData provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) PutRepositoryCatalogData(ctx context.Context, params *ecrpublic.PutRepositoryCatalogDataInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.PutRepositoryCatalogDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutRepositoryCatalogData")
	}

	var r0 *ecrpublic.PutRepositoryCatalogDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.PutRepositoryCatalogDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.PutRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) *ecrpublic.PutRepositoryCatalogDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.PutRepositoryCatalogDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.PutRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetRepositoryPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) SetRepositoryPolicy(ctx context.Context, params *ecrpublic.SetRepositoryPolicyInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.SetRepositoryPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetRepositoryPolicy")
	}

	var r0 *ecrpublic.SetRepositoryPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.SetRepositoryPolicyInput, ...func(*ecrpublic.Options)) (*ecrpublic.SetRepositoryPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.SetRepositoryPolicyInput, ...func(*ecrpublic.Options)) *ecrpublic.SetRepositoryPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.SetRepositoryPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.SetRepositoryPolicyInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) TagResource(ctx context.Context, params *ecrpublic.TagResourceInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.TagResourceOutput, error) {
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

	var r0 *ecrpublic.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.TagResourceInput, ...func(*ecrpublic.Options)) (*ecrpublic.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.TagResourceInput, ...func(*ecrpublic.Options)) *ecrpublic.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.TagResourceInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) UntagResource(ctx context.Context, params *ecrpublic.UntagResourceInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.UntagResourceOutput, error) {
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

	var r0 *ecrpublic.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.UntagResourceInput, ...func(*ecrpublic.Options)) (*ecrpublic.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.UntagResourceInput, ...func(*ecrpublic.Options)) *ecrpublic.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.UntagResourceInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadLayerPart provides a mock function with given fields: ctx, params, optFns
func (_m *EcrpublicClient) UploadLayerPart(ctx context.Context, params *ecrpublic.UploadLayerPartInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.UploadLayerPartOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UploadLayerPart")
	}

	var r0 *ecrpublic.UploadLayerPartOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.UploadLayerPartInput, ...func(*ecrpublic.Options)) (*ecrpublic.UploadLayerPartOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ecrpublic.UploadLayerPartInput, ...func(*ecrpublic.Options)) *ecrpublic.UploadLayerPartOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecrpublic.UploadLayerPartOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ecrpublic.UploadLayerPartInput, ...func(*ecrpublic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEcrpublicClient creates a new instance of EcrpublicClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEcrpublicClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EcrpublicClient {
	mock := &EcrpublicClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
