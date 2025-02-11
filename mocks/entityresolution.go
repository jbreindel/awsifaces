// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entityresolution "github.com/aws/aws-sdk-go-v2/service/entityresolution"
	mock "github.com/stretchr/testify/mock"
)

// EntityresolutionClient is an autogenerated mock type for the EntityresolutionClient type
type EntityresolutionClient struct {
	mock.Mock
}

// CreateIdMappingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) CreateIdMappingWorkflow(ctx context.Context, params *entityresolution.CreateIdMappingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.CreateIdMappingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateIdMappingWorkflow")
	}

	var r0 *entityresolution.CreateIdMappingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateIdMappingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.CreateIdMappingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateIdMappingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.CreateIdMappingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.CreateIdMappingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.CreateIdMappingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMatchingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) CreateMatchingWorkflow(ctx context.Context, params *entityresolution.CreateMatchingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.CreateMatchingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMatchingWorkflow")
	}

	var r0 *entityresolution.CreateMatchingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateMatchingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.CreateMatchingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateMatchingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.CreateMatchingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.CreateMatchingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.CreateMatchingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSchemaMapping provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) CreateSchemaMapping(ctx context.Context, params *entityresolution.CreateSchemaMappingInput, optFns ...func(*entityresolution.Options)) (*entityresolution.CreateSchemaMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSchemaMapping")
	}

	var r0 *entityresolution.CreateSchemaMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateSchemaMappingInput, ...func(*entityresolution.Options)) (*entityresolution.CreateSchemaMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.CreateSchemaMappingInput, ...func(*entityresolution.Options)) *entityresolution.CreateSchemaMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.CreateSchemaMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.CreateSchemaMappingInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIdMappingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) DeleteIdMappingWorkflow(ctx context.Context, params *entityresolution.DeleteIdMappingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.DeleteIdMappingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIdMappingWorkflow")
	}

	var r0 *entityresolution.DeleteIdMappingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteIdMappingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.DeleteIdMappingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteIdMappingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.DeleteIdMappingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.DeleteIdMappingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.DeleteIdMappingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMatchingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) DeleteMatchingWorkflow(ctx context.Context, params *entityresolution.DeleteMatchingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.DeleteMatchingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMatchingWorkflow")
	}

	var r0 *entityresolution.DeleteMatchingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteMatchingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.DeleteMatchingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteMatchingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.DeleteMatchingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.DeleteMatchingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.DeleteMatchingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSchemaMapping provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) DeleteSchemaMapping(ctx context.Context, params *entityresolution.DeleteSchemaMappingInput, optFns ...func(*entityresolution.Options)) (*entityresolution.DeleteSchemaMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSchemaMapping")
	}

	var r0 *entityresolution.DeleteSchemaMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteSchemaMappingInput, ...func(*entityresolution.Options)) (*entityresolution.DeleteSchemaMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.DeleteSchemaMappingInput, ...func(*entityresolution.Options)) *entityresolution.DeleteSchemaMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.DeleteSchemaMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.DeleteSchemaMappingInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdMappingJob provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetIdMappingJob(ctx context.Context, params *entityresolution.GetIdMappingJobInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetIdMappingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIdMappingJob")
	}

	var r0 *entityresolution.GetIdMappingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetIdMappingJobInput, ...func(*entityresolution.Options)) (*entityresolution.GetIdMappingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetIdMappingJobInput, ...func(*entityresolution.Options)) *entityresolution.GetIdMappingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetIdMappingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetIdMappingJobInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdMappingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetIdMappingWorkflow(ctx context.Context, params *entityresolution.GetIdMappingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetIdMappingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetIdMappingWorkflow")
	}

	var r0 *entityresolution.GetIdMappingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetIdMappingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.GetIdMappingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetIdMappingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.GetIdMappingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetIdMappingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetIdMappingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMatchId provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetMatchId(ctx context.Context, params *entityresolution.GetMatchIdInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetMatchIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMatchId")
	}

	var r0 *entityresolution.GetMatchIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchIdInput, ...func(*entityresolution.Options)) (*entityresolution.GetMatchIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchIdInput, ...func(*entityresolution.Options)) *entityresolution.GetMatchIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetMatchIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetMatchIdInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMatchingJob provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetMatchingJob(ctx context.Context, params *entityresolution.GetMatchingJobInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetMatchingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMatchingJob")
	}

	var r0 *entityresolution.GetMatchingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchingJobInput, ...func(*entityresolution.Options)) (*entityresolution.GetMatchingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchingJobInput, ...func(*entityresolution.Options)) *entityresolution.GetMatchingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetMatchingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetMatchingJobInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMatchingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetMatchingWorkflow(ctx context.Context, params *entityresolution.GetMatchingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetMatchingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMatchingWorkflow")
	}

	var r0 *entityresolution.GetMatchingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.GetMatchingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetMatchingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.GetMatchingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetMatchingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetMatchingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProviderService provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetProviderService(ctx context.Context, params *entityresolution.GetProviderServiceInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetProviderServiceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProviderService")
	}

	var r0 *entityresolution.GetProviderServiceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetProviderServiceInput, ...func(*entityresolution.Options)) (*entityresolution.GetProviderServiceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetProviderServiceInput, ...func(*entityresolution.Options)) *entityresolution.GetProviderServiceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetProviderServiceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetProviderServiceInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaMapping provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) GetSchemaMapping(ctx context.Context, params *entityresolution.GetSchemaMappingInput, optFns ...func(*entityresolution.Options)) (*entityresolution.GetSchemaMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaMapping")
	}

	var r0 *entityresolution.GetSchemaMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetSchemaMappingInput, ...func(*entityresolution.Options)) (*entityresolution.GetSchemaMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.GetSchemaMappingInput, ...func(*entityresolution.Options)) *entityresolution.GetSchemaMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.GetSchemaMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.GetSchemaMappingInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdMappingJobs provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListIdMappingJobs(ctx context.Context, params *entityresolution.ListIdMappingJobsInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListIdMappingJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdMappingJobs")
	}

	var r0 *entityresolution.ListIdMappingJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListIdMappingJobsInput, ...func(*entityresolution.Options)) (*entityresolution.ListIdMappingJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListIdMappingJobsInput, ...func(*entityresolution.Options)) *entityresolution.ListIdMappingJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListIdMappingJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListIdMappingJobsInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdMappingWorkflows provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListIdMappingWorkflows(ctx context.Context, params *entityresolution.ListIdMappingWorkflowsInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListIdMappingWorkflowsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListIdMappingWorkflows")
	}

	var r0 *entityresolution.ListIdMappingWorkflowsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListIdMappingWorkflowsInput, ...func(*entityresolution.Options)) (*entityresolution.ListIdMappingWorkflowsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListIdMappingWorkflowsInput, ...func(*entityresolution.Options)) *entityresolution.ListIdMappingWorkflowsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListIdMappingWorkflowsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListIdMappingWorkflowsInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMatchingJobs provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListMatchingJobs(ctx context.Context, params *entityresolution.ListMatchingJobsInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListMatchingJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMatchingJobs")
	}

	var r0 *entityresolution.ListMatchingJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListMatchingJobsInput, ...func(*entityresolution.Options)) (*entityresolution.ListMatchingJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListMatchingJobsInput, ...func(*entityresolution.Options)) *entityresolution.ListMatchingJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListMatchingJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListMatchingJobsInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMatchingWorkflows provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListMatchingWorkflows(ctx context.Context, params *entityresolution.ListMatchingWorkflowsInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListMatchingWorkflowsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMatchingWorkflows")
	}

	var r0 *entityresolution.ListMatchingWorkflowsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListMatchingWorkflowsInput, ...func(*entityresolution.Options)) (*entityresolution.ListMatchingWorkflowsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListMatchingWorkflowsInput, ...func(*entityresolution.Options)) *entityresolution.ListMatchingWorkflowsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListMatchingWorkflowsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListMatchingWorkflowsInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviderServices provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListProviderServices(ctx context.Context, params *entityresolution.ListProviderServicesInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListProviderServicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProviderServices")
	}

	var r0 *entityresolution.ListProviderServicesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListProviderServicesInput, ...func(*entityresolution.Options)) (*entityresolution.ListProviderServicesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListProviderServicesInput, ...func(*entityresolution.Options)) *entityresolution.ListProviderServicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListProviderServicesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListProviderServicesInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchemaMappings provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListSchemaMappings(ctx context.Context, params *entityresolution.ListSchemaMappingsInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListSchemaMappingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSchemaMappings")
	}

	var r0 *entityresolution.ListSchemaMappingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListSchemaMappingsInput, ...func(*entityresolution.Options)) (*entityresolution.ListSchemaMappingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListSchemaMappingsInput, ...func(*entityresolution.Options)) *entityresolution.ListSchemaMappingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListSchemaMappingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListSchemaMappingsInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) ListTagsForResource(ctx context.Context, params *entityresolution.ListTagsForResourceInput, optFns ...func(*entityresolution.Options)) (*entityresolution.ListTagsForResourceOutput, error) {
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

	var r0 *entityresolution.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListTagsForResourceInput, ...func(*entityresolution.Options)) (*entityresolution.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.ListTagsForResourceInput, ...func(*entityresolution.Options)) *entityresolution.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.ListTagsForResourceInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *EntityresolutionClient) Options() entityresolution.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 entityresolution.Options
	if rf, ok := ret.Get(0).(func() entityresolution.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entityresolution.Options)
	}

	return r0
}

// StartIdMappingJob provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) StartIdMappingJob(ctx context.Context, params *entityresolution.StartIdMappingJobInput, optFns ...func(*entityresolution.Options)) (*entityresolution.StartIdMappingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartIdMappingJob")
	}

	var r0 *entityresolution.StartIdMappingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.StartIdMappingJobInput, ...func(*entityresolution.Options)) (*entityresolution.StartIdMappingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.StartIdMappingJobInput, ...func(*entityresolution.Options)) *entityresolution.StartIdMappingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.StartIdMappingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.StartIdMappingJobInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartMatchingJob provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) StartMatchingJob(ctx context.Context, params *entityresolution.StartMatchingJobInput, optFns ...func(*entityresolution.Options)) (*entityresolution.StartMatchingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartMatchingJob")
	}

	var r0 *entityresolution.StartMatchingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.StartMatchingJobInput, ...func(*entityresolution.Options)) (*entityresolution.StartMatchingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.StartMatchingJobInput, ...func(*entityresolution.Options)) *entityresolution.StartMatchingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.StartMatchingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.StartMatchingJobInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) TagResource(ctx context.Context, params *entityresolution.TagResourceInput, optFns ...func(*entityresolution.Options)) (*entityresolution.TagResourceOutput, error) {
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

	var r0 *entityresolution.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.TagResourceInput, ...func(*entityresolution.Options)) (*entityresolution.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.TagResourceInput, ...func(*entityresolution.Options)) *entityresolution.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.TagResourceInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) UntagResource(ctx context.Context, params *entityresolution.UntagResourceInput, optFns ...func(*entityresolution.Options)) (*entityresolution.UntagResourceOutput, error) {
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

	var r0 *entityresolution.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UntagResourceInput, ...func(*entityresolution.Options)) (*entityresolution.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UntagResourceInput, ...func(*entityresolution.Options)) *entityresolution.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.UntagResourceInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdMappingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) UpdateIdMappingWorkflow(ctx context.Context, params *entityresolution.UpdateIdMappingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.UpdateIdMappingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIdMappingWorkflow")
	}

	var r0 *entityresolution.UpdateIdMappingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateIdMappingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.UpdateIdMappingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateIdMappingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.UpdateIdMappingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.UpdateIdMappingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.UpdateIdMappingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMatchingWorkflow provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) UpdateMatchingWorkflow(ctx context.Context, params *entityresolution.UpdateMatchingWorkflowInput, optFns ...func(*entityresolution.Options)) (*entityresolution.UpdateMatchingWorkflowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMatchingWorkflow")
	}

	var r0 *entityresolution.UpdateMatchingWorkflowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateMatchingWorkflowInput, ...func(*entityresolution.Options)) (*entityresolution.UpdateMatchingWorkflowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateMatchingWorkflowInput, ...func(*entityresolution.Options)) *entityresolution.UpdateMatchingWorkflowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.UpdateMatchingWorkflowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.UpdateMatchingWorkflowInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSchemaMapping provides a mock function with given fields: ctx, params, optFns
func (_m *EntityresolutionClient) UpdateSchemaMapping(ctx context.Context, params *entityresolution.UpdateSchemaMappingInput, optFns ...func(*entityresolution.Options)) (*entityresolution.UpdateSchemaMappingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSchemaMapping")
	}

	var r0 *entityresolution.UpdateSchemaMappingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateSchemaMappingInput, ...func(*entityresolution.Options)) (*entityresolution.UpdateSchemaMappingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entityresolution.UpdateSchemaMappingInput, ...func(*entityresolution.Options)) *entityresolution.UpdateSchemaMappingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entityresolution.UpdateSchemaMappingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entityresolution.UpdateSchemaMappingInput, ...func(*entityresolution.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEntityresolutionClient creates a new instance of EntityresolutionClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEntityresolutionClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EntityresolutionClient {
	mock := &EntityresolutionClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
