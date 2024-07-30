// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	comprehendmedical "github.com/aws/aws-sdk-go-v2/service/comprehendmedical"

	mock "github.com/stretchr/testify/mock"
)

// ComprehendmedicalClient is an autogenerated mock type for the ComprehendmedicalClient type
type ComprehendmedicalClient struct {
	mock.Mock
}

// DescribeEntitiesDetectionV2Job provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DescribeEntitiesDetectionV2Job(ctx context.Context, params *comprehendmedical.DescribeEntitiesDetectionV2JobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeEntitiesDetectionV2Job")
	}

	var r0 *comprehendmedical.DescribeEntitiesDetectionV2JobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DescribeEntitiesDetectionV2JobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DescribeEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeICD10CMInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DescribeICD10CMInferenceJob(ctx context.Context, params *comprehendmedical.DescribeICD10CMInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeICD10CMInferenceJob")
	}

	var r0 *comprehendmedical.DescribeICD10CMInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DescribeICD10CMInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DescribeICD10CMInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DescribeICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePHIDetectionJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DescribePHIDetectionJob(ctx context.Context, params *comprehendmedical.DescribePHIDetectionJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribePHIDetectionJob")
	}

	var r0 *comprehendmedical.DescribePHIDetectionJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribePHIDetectionJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribePHIDetectionJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribePHIDetectionJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DescribePHIDetectionJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DescribePHIDetectionJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DescribePHIDetectionJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRxNormInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DescribeRxNormInferenceJob(ctx context.Context, params *comprehendmedical.DescribeRxNormInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRxNormInferenceJob")
	}

	var r0 *comprehendmedical.DescribeRxNormInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DescribeRxNormInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DescribeRxNormInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DescribeRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSNOMEDCTInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DescribeSNOMEDCTInferenceJob(ctx context.Context, params *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSNOMEDCTInferenceJob")
	}

	var r0 *comprehendmedical.DescribeSNOMEDCTInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DescribeSNOMEDCTInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetectEntities provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DetectEntities(ctx context.Context, params *comprehendmedical.DetectEntitiesInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectEntitiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DetectEntities")
	}

	var r0 *comprehendmedical.DetectEntitiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectEntitiesInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectEntitiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectEntitiesInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DetectEntitiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DetectEntitiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DetectEntitiesInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetectEntitiesV2 provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DetectEntitiesV2(ctx context.Context, params *comprehendmedical.DetectEntitiesV2Input, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectEntitiesV2Output, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DetectEntitiesV2")
	}

	var r0 *comprehendmedical.DetectEntitiesV2Output
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectEntitiesV2Input, ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectEntitiesV2Output, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectEntitiesV2Input, ...func(*comprehendmedical.Options)) *comprehendmedical.DetectEntitiesV2Output); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DetectEntitiesV2Output)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DetectEntitiesV2Input, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetectPHI provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) DetectPHI(ctx context.Context, params *comprehendmedical.DetectPHIInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectPHIOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DetectPHI")
	}

	var r0 *comprehendmedical.DetectPHIOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectPHIInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.DetectPHIOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.DetectPHIInput, ...func(*comprehendmedical.Options)) *comprehendmedical.DetectPHIOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.DetectPHIOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.DetectPHIInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InferICD10CM provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) InferICD10CM(ctx context.Context, params *comprehendmedical.InferICD10CMInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.InferICD10CMOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InferICD10CM")
	}

	var r0 *comprehendmedical.InferICD10CMOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferICD10CMInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.InferICD10CMOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferICD10CMInput, ...func(*comprehendmedical.Options)) *comprehendmedical.InferICD10CMOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.InferICD10CMOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.InferICD10CMInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InferRxNorm provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) InferRxNorm(ctx context.Context, params *comprehendmedical.InferRxNormInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.InferRxNormOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InferRxNorm")
	}

	var r0 *comprehendmedical.InferRxNormOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferRxNormInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.InferRxNormOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferRxNormInput, ...func(*comprehendmedical.Options)) *comprehendmedical.InferRxNormOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.InferRxNormOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.InferRxNormInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InferSNOMEDCT provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) InferSNOMEDCT(ctx context.Context, params *comprehendmedical.InferSNOMEDCTInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.InferSNOMEDCTOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InferSNOMEDCT")
	}

	var r0 *comprehendmedical.InferSNOMEDCTOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferSNOMEDCTInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.InferSNOMEDCTOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.InferSNOMEDCTInput, ...func(*comprehendmedical.Options)) *comprehendmedical.InferSNOMEDCTOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.InferSNOMEDCTOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.InferSNOMEDCTInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEntitiesDetectionV2Jobs provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) ListEntitiesDetectionV2Jobs(ctx context.Context, params *comprehendmedical.ListEntitiesDetectionV2JobsInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEntitiesDetectionV2Jobs")
	}

	var r0 *comprehendmedical.ListEntitiesDetectionV2JobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListEntitiesDetectionV2JobsInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListEntitiesDetectionV2JobsInput, ...func(*comprehendmedical.Options)) *comprehendmedical.ListEntitiesDetectionV2JobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.ListEntitiesDetectionV2JobsInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListICD10CMInferenceJobs provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) ListICD10CMInferenceJobs(ctx context.Context, params *comprehendmedical.ListICD10CMInferenceJobsInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListICD10CMInferenceJobs")
	}

	var r0 *comprehendmedical.ListICD10CMInferenceJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListICD10CMInferenceJobsInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListICD10CMInferenceJobsInput, ...func(*comprehendmedical.Options)) *comprehendmedical.ListICD10CMInferenceJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.ListICD10CMInferenceJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.ListICD10CMInferenceJobsInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPHIDetectionJobs provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) ListPHIDetectionJobs(ctx context.Context, params *comprehendmedical.ListPHIDetectionJobsInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPHIDetectionJobs")
	}

	var r0 *comprehendmedical.ListPHIDetectionJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListPHIDetectionJobsInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.ListPHIDetectionJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListPHIDetectionJobsInput, ...func(*comprehendmedical.Options)) *comprehendmedical.ListPHIDetectionJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.ListPHIDetectionJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.ListPHIDetectionJobsInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRxNormInferenceJobs provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) ListRxNormInferenceJobs(ctx context.Context, params *comprehendmedical.ListRxNormInferenceJobsInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRxNormInferenceJobs")
	}

	var r0 *comprehendmedical.ListRxNormInferenceJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListRxNormInferenceJobsInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListRxNormInferenceJobsInput, ...func(*comprehendmedical.Options)) *comprehendmedical.ListRxNormInferenceJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.ListRxNormInferenceJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.ListRxNormInferenceJobsInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSNOMEDCTInferenceJobs provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) ListSNOMEDCTInferenceJobs(ctx context.Context, params *comprehendmedical.ListSNOMEDCTInferenceJobsInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSNOMEDCTInferenceJobs")
	}

	var r0 *comprehendmedical.ListSNOMEDCTInferenceJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListSNOMEDCTInferenceJobsInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.ListSNOMEDCTInferenceJobsInput, ...func(*comprehendmedical.Options)) *comprehendmedical.ListSNOMEDCTInferenceJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.ListSNOMEDCTInferenceJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.ListSNOMEDCTInferenceJobsInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *ComprehendmedicalClient) Options() comprehendmedical.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 comprehendmedical.Options
	if rf, ok := ret.Get(0).(func() comprehendmedical.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(comprehendmedical.Options)
	}

	return r0
}

// StartEntitiesDetectionV2Job provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StartEntitiesDetectionV2Job(ctx context.Context, params *comprehendmedical.StartEntitiesDetectionV2JobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartEntitiesDetectionV2Job")
	}

	var r0 *comprehendmedical.StartEntitiesDetectionV2JobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StartEntitiesDetectionV2JobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StartEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartICD10CMInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StartICD10CMInferenceJob(ctx context.Context, params *comprehendmedical.StartICD10CMInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartICD10CMInferenceJob")
	}

	var r0 *comprehendmedical.StartICD10CMInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StartICD10CMInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StartICD10CMInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StartICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartPHIDetectionJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StartPHIDetectionJob(ctx context.Context, params *comprehendmedical.StartPHIDetectionJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartPHIDetectionJob")
	}

	var r0 *comprehendmedical.StartPHIDetectionJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartPHIDetectionJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StartPHIDetectionJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartPHIDetectionJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StartPHIDetectionJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StartPHIDetectionJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StartPHIDetectionJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartRxNormInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StartRxNormInferenceJob(ctx context.Context, params *comprehendmedical.StartRxNormInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartRxNormInferenceJob")
	}

	var r0 *comprehendmedical.StartRxNormInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StartRxNormInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StartRxNormInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StartRxNormInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StartRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartSNOMEDCTInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StartSNOMEDCTInferenceJob(ctx context.Context, params *comprehendmedical.StartSNOMEDCTInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StartSNOMEDCTInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartSNOMEDCTInferenceJob")
	}

	var r0 *comprehendmedical.StartSNOMEDCTInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StartSNOMEDCTInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StartSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StartSNOMEDCTInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StartSNOMEDCTInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StartSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopEntitiesDetectionV2Job provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StopEntitiesDetectionV2Job(ctx context.Context, params *comprehendmedical.StopEntitiesDetectionV2JobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopEntitiesDetectionV2Job")
	}

	var r0 *comprehendmedical.StopEntitiesDetectionV2JobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StopEntitiesDetectionV2JobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StopEntitiesDetectionV2JobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopICD10CMInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StopICD10CMInferenceJob(ctx context.Context, params *comprehendmedical.StopICD10CMInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopICD10CMInferenceJob")
	}

	var r0 *comprehendmedical.StopICD10CMInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StopICD10CMInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StopICD10CMInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StopICD10CMInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopPHIDetectionJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StopPHIDetectionJob(ctx context.Context, params *comprehendmedical.StopPHIDetectionJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopPHIDetectionJob")
	}

	var r0 *comprehendmedical.StopPHIDetectionJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopPHIDetectionJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StopPHIDetectionJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopPHIDetectionJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StopPHIDetectionJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StopPHIDetectionJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StopPHIDetectionJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopRxNormInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StopRxNormInferenceJob(ctx context.Context, params *comprehendmedical.StopRxNormInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopRxNormInferenceJob")
	}

	var r0 *comprehendmedical.StopRxNormInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StopRxNormInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StopRxNormInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StopRxNormInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StopRxNormInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopSNOMEDCTInferenceJob provides a mock function with given fields: ctx, params, optFns
func (_m *ComprehendmedicalClient) StopSNOMEDCTInferenceJob(ctx context.Context, params *comprehendmedical.StopSNOMEDCTInferenceJobInput, optFns ...func(*comprehendmedical.Options)) (*comprehendmedical.StopSNOMEDCTInferenceJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopSNOMEDCTInferenceJob")
	}

	var r0 *comprehendmedical.StopSNOMEDCTInferenceJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) (*comprehendmedical.StopSNOMEDCTInferenceJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *comprehendmedical.StopSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) *comprehendmedical.StopSNOMEDCTInferenceJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*comprehendmedical.StopSNOMEDCTInferenceJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *comprehendmedical.StopSNOMEDCTInferenceJobInput, ...func(*comprehendmedical.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewComprehendmedicalClient creates a new instance of ComprehendmedicalClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComprehendmedicalClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ComprehendmedicalClient {
	mock := &ComprehendmedicalClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
