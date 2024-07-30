// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	bedrock "github.com/aws/aws-sdk-go-v2/service/bedrock"

	mock "github.com/stretchr/testify/mock"
)

// BedrockClient is an autogenerated mock type for the BedrockClient type
type BedrockClient struct {
	mock.Mock
}

// CreateModelCustomizationJob provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) CreateModelCustomizationJob(ctx context.Context, params *bedrock.CreateModelCustomizationJobInput, optFns ...func(*bedrock.Options)) (*bedrock.CreateModelCustomizationJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateModelCustomizationJob")
	}

	var r0 *bedrock.CreateModelCustomizationJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.CreateModelCustomizationJobInput, ...func(*bedrock.Options)) (*bedrock.CreateModelCustomizationJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.CreateModelCustomizationJobInput, ...func(*bedrock.Options)) *bedrock.CreateModelCustomizationJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.CreateModelCustomizationJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.CreateModelCustomizationJobInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProvisionedModelThroughput provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) CreateProvisionedModelThroughput(ctx context.Context, params *bedrock.CreateProvisionedModelThroughputInput, optFns ...func(*bedrock.Options)) (*bedrock.CreateProvisionedModelThroughputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProvisionedModelThroughput")
	}

	var r0 *bedrock.CreateProvisionedModelThroughputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.CreateProvisionedModelThroughputInput, ...func(*bedrock.Options)) (*bedrock.CreateProvisionedModelThroughputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.CreateProvisionedModelThroughputInput, ...func(*bedrock.Options)) *bedrock.CreateProvisionedModelThroughputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.CreateProvisionedModelThroughputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.CreateProvisionedModelThroughputInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomModel provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) DeleteCustomModel(ctx context.Context, params *bedrock.DeleteCustomModelInput, optFns ...func(*bedrock.Options)) (*bedrock.DeleteCustomModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomModel")
	}

	var r0 *bedrock.DeleteCustomModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteCustomModelInput, ...func(*bedrock.Options)) (*bedrock.DeleteCustomModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteCustomModelInput, ...func(*bedrock.Options)) *bedrock.DeleteCustomModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.DeleteCustomModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.DeleteCustomModelInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteModelInvocationLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) DeleteModelInvocationLoggingConfiguration(ctx context.Context, params *bedrock.DeleteModelInvocationLoggingConfigurationInput, optFns ...func(*bedrock.Options)) (*bedrock.DeleteModelInvocationLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteModelInvocationLoggingConfiguration")
	}

	var r0 *bedrock.DeleteModelInvocationLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) (*bedrock.DeleteModelInvocationLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) *bedrock.DeleteModelInvocationLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.DeleteModelInvocationLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.DeleteModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProvisionedModelThroughput provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) DeleteProvisionedModelThroughput(ctx context.Context, params *bedrock.DeleteProvisionedModelThroughputInput, optFns ...func(*bedrock.Options)) (*bedrock.DeleteProvisionedModelThroughputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProvisionedModelThroughput")
	}

	var r0 *bedrock.DeleteProvisionedModelThroughputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteProvisionedModelThroughputInput, ...func(*bedrock.Options)) (*bedrock.DeleteProvisionedModelThroughputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.DeleteProvisionedModelThroughputInput, ...func(*bedrock.Options)) *bedrock.DeleteProvisionedModelThroughputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.DeleteProvisionedModelThroughputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.DeleteProvisionedModelThroughputInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomModel provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) GetCustomModel(ctx context.Context, params *bedrock.GetCustomModelInput, optFns ...func(*bedrock.Options)) (*bedrock.GetCustomModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomModel")
	}

	var r0 *bedrock.GetCustomModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetCustomModelInput, ...func(*bedrock.Options)) (*bedrock.GetCustomModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetCustomModelInput, ...func(*bedrock.Options)) *bedrock.GetCustomModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.GetCustomModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.GetCustomModelInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFoundationModel provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) GetFoundationModel(ctx context.Context, params *bedrock.GetFoundationModelInput, optFns ...func(*bedrock.Options)) (*bedrock.GetFoundationModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFoundationModel")
	}

	var r0 *bedrock.GetFoundationModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetFoundationModelInput, ...func(*bedrock.Options)) (*bedrock.GetFoundationModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetFoundationModelInput, ...func(*bedrock.Options)) *bedrock.GetFoundationModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.GetFoundationModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.GetFoundationModelInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModelCustomizationJob provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) GetModelCustomizationJob(ctx context.Context, params *bedrock.GetModelCustomizationJobInput, optFns ...func(*bedrock.Options)) (*bedrock.GetModelCustomizationJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetModelCustomizationJob")
	}

	var r0 *bedrock.GetModelCustomizationJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetModelCustomizationJobInput, ...func(*bedrock.Options)) (*bedrock.GetModelCustomizationJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetModelCustomizationJobInput, ...func(*bedrock.Options)) *bedrock.GetModelCustomizationJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.GetModelCustomizationJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.GetModelCustomizationJobInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModelInvocationLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) GetModelInvocationLoggingConfiguration(ctx context.Context, params *bedrock.GetModelInvocationLoggingConfigurationInput, optFns ...func(*bedrock.Options)) (*bedrock.GetModelInvocationLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetModelInvocationLoggingConfiguration")
	}

	var r0 *bedrock.GetModelInvocationLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) (*bedrock.GetModelInvocationLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) *bedrock.GetModelInvocationLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.GetModelInvocationLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.GetModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedModelThroughput provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) GetProvisionedModelThroughput(ctx context.Context, params *bedrock.GetProvisionedModelThroughputInput, optFns ...func(*bedrock.Options)) (*bedrock.GetProvisionedModelThroughputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProvisionedModelThroughput")
	}

	var r0 *bedrock.GetProvisionedModelThroughputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetProvisionedModelThroughputInput, ...func(*bedrock.Options)) (*bedrock.GetProvisionedModelThroughputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.GetProvisionedModelThroughputInput, ...func(*bedrock.Options)) *bedrock.GetProvisionedModelThroughputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.GetProvisionedModelThroughputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.GetProvisionedModelThroughputInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCustomModels provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) ListCustomModels(ctx context.Context, params *bedrock.ListCustomModelsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListCustomModelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCustomModels")
	}

	var r0 *bedrock.ListCustomModelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListCustomModelsInput, ...func(*bedrock.Options)) (*bedrock.ListCustomModelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListCustomModelsInput, ...func(*bedrock.Options)) *bedrock.ListCustomModelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.ListCustomModelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.ListCustomModelsInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFoundationModels provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) ListFoundationModels(ctx context.Context, params *bedrock.ListFoundationModelsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListFoundationModelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFoundationModels")
	}

	var r0 *bedrock.ListFoundationModelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListFoundationModelsInput, ...func(*bedrock.Options)) (*bedrock.ListFoundationModelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListFoundationModelsInput, ...func(*bedrock.Options)) *bedrock.ListFoundationModelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.ListFoundationModelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.ListFoundationModelsInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModelCustomizationJobs provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) ListModelCustomizationJobs(ctx context.Context, params *bedrock.ListModelCustomizationJobsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListModelCustomizationJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListModelCustomizationJobs")
	}

	var r0 *bedrock.ListModelCustomizationJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListModelCustomizationJobsInput, ...func(*bedrock.Options)) (*bedrock.ListModelCustomizationJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListModelCustomizationJobsInput, ...func(*bedrock.Options)) *bedrock.ListModelCustomizationJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.ListModelCustomizationJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.ListModelCustomizationJobsInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProvisionedModelThroughputs provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) ListProvisionedModelThroughputs(ctx context.Context, params *bedrock.ListProvisionedModelThroughputsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListProvisionedModelThroughputsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProvisionedModelThroughputs")
	}

	var r0 *bedrock.ListProvisionedModelThroughputsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListProvisionedModelThroughputsInput, ...func(*bedrock.Options)) (*bedrock.ListProvisionedModelThroughputsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListProvisionedModelThroughputsInput, ...func(*bedrock.Options)) *bedrock.ListProvisionedModelThroughputsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.ListProvisionedModelThroughputsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.ListProvisionedModelThroughputsInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) ListTagsForResource(ctx context.Context, params *bedrock.ListTagsForResourceInput, optFns ...func(*bedrock.Options)) (*bedrock.ListTagsForResourceOutput, error) {
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

	var r0 *bedrock.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListTagsForResourceInput, ...func(*bedrock.Options)) (*bedrock.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.ListTagsForResourceInput, ...func(*bedrock.Options)) *bedrock.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.ListTagsForResourceInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *BedrockClient) Options() bedrock.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 bedrock.Options
	if rf, ok := ret.Get(0).(func() bedrock.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bedrock.Options)
	}

	return r0
}

// PutModelInvocationLoggingConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) PutModelInvocationLoggingConfiguration(ctx context.Context, params *bedrock.PutModelInvocationLoggingConfigurationInput, optFns ...func(*bedrock.Options)) (*bedrock.PutModelInvocationLoggingConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutModelInvocationLoggingConfiguration")
	}

	var r0 *bedrock.PutModelInvocationLoggingConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.PutModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) (*bedrock.PutModelInvocationLoggingConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.PutModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) *bedrock.PutModelInvocationLoggingConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.PutModelInvocationLoggingConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.PutModelInvocationLoggingConfigurationInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopModelCustomizationJob provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) StopModelCustomizationJob(ctx context.Context, params *bedrock.StopModelCustomizationJobInput, optFns ...func(*bedrock.Options)) (*bedrock.StopModelCustomizationJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopModelCustomizationJob")
	}

	var r0 *bedrock.StopModelCustomizationJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.StopModelCustomizationJobInput, ...func(*bedrock.Options)) (*bedrock.StopModelCustomizationJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.StopModelCustomizationJobInput, ...func(*bedrock.Options)) *bedrock.StopModelCustomizationJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.StopModelCustomizationJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.StopModelCustomizationJobInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) TagResource(ctx context.Context, params *bedrock.TagResourceInput, optFns ...func(*bedrock.Options)) (*bedrock.TagResourceOutput, error) {
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

	var r0 *bedrock.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.TagResourceInput, ...func(*bedrock.Options)) (*bedrock.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.TagResourceInput, ...func(*bedrock.Options)) *bedrock.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.TagResourceInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) UntagResource(ctx context.Context, params *bedrock.UntagResourceInput, optFns ...func(*bedrock.Options)) (*bedrock.UntagResourceOutput, error) {
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

	var r0 *bedrock.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.UntagResourceInput, ...func(*bedrock.Options)) (*bedrock.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.UntagResourceInput, ...func(*bedrock.Options)) *bedrock.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.UntagResourceInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProvisionedModelThroughput provides a mock function with given fields: ctx, params, optFns
func (_m *BedrockClient) UpdateProvisionedModelThroughput(ctx context.Context, params *bedrock.UpdateProvisionedModelThroughputInput, optFns ...func(*bedrock.Options)) (*bedrock.UpdateProvisionedModelThroughputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProvisionedModelThroughput")
	}

	var r0 *bedrock.UpdateProvisionedModelThroughputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.UpdateProvisionedModelThroughputInput, ...func(*bedrock.Options)) (*bedrock.UpdateProvisionedModelThroughputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *bedrock.UpdateProvisionedModelThroughputInput, ...func(*bedrock.Options)) *bedrock.UpdateProvisionedModelThroughputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bedrock.UpdateProvisionedModelThroughputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *bedrock.UpdateProvisionedModelThroughputInput, ...func(*bedrock.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBedrockClient creates a new instance of BedrockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBedrockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BedrockClient {
	mock := &BedrockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
