// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	migrationhubstrategy "github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
	mock "github.com/stretchr/testify/mock"
)

// MigrationhubstrategyClient is an autogenerated mock type for the MigrationhubstrategyClient type
type MigrationhubstrategyClient struct {
	mock.Mock
}

// GetApplicationComponentDetails provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetApplicationComponentDetails(ctx context.Context, params *migrationhubstrategy.GetApplicationComponentDetailsInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetApplicationComponentDetailsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetApplicationComponentDetails")
	}

	var r0 *migrationhubstrategy.GetApplicationComponentDetailsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetApplicationComponentDetailsInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetApplicationComponentDetailsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetApplicationComponentDetailsInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetApplicationComponentDetailsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetApplicationComponentDetailsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetApplicationComponentDetailsInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplicationComponentStrategies provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetApplicationComponentStrategies(ctx context.Context, params *migrationhubstrategy.GetApplicationComponentStrategiesInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetApplicationComponentStrategiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetApplicationComponentStrategies")
	}

	var r0 *migrationhubstrategy.GetApplicationComponentStrategiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetApplicationComponentStrategiesInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetApplicationComponentStrategiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetApplicationComponentStrategiesInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetApplicationComponentStrategiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetApplicationComponentStrategiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetApplicationComponentStrategiesInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssessment provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetAssessment(ctx context.Context, params *migrationhubstrategy.GetAssessmentInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetAssessmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAssessment")
	}

	var r0 *migrationhubstrategy.GetAssessmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetAssessmentInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetAssessmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetAssessmentInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetAssessmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetAssessmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetAssessmentInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImportFileTask provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetImportFileTask(ctx context.Context, params *migrationhubstrategy.GetImportFileTaskInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetImportFileTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetImportFileTask")
	}

	var r0 *migrationhubstrategy.GetImportFileTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetImportFileTaskInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetImportFileTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetImportFileTaskInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetImportFileTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetImportFileTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetImportFileTaskInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestAssessmentId provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetLatestAssessmentId(ctx context.Context, params *migrationhubstrategy.GetLatestAssessmentIdInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetLatestAssessmentIdOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestAssessmentId")
	}

	var r0 *migrationhubstrategy.GetLatestAssessmentIdOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetLatestAssessmentIdInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetLatestAssessmentIdOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetLatestAssessmentIdInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetLatestAssessmentIdOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetLatestAssessmentIdOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetLatestAssessmentIdInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPortfolioPreferences provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetPortfolioPreferences(ctx context.Context, params *migrationhubstrategy.GetPortfolioPreferencesInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetPortfolioPreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPortfolioPreferences")
	}

	var r0 *migrationhubstrategy.GetPortfolioPreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetPortfolioPreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetPortfolioPreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetPortfolioPreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPortfolioSummary provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetPortfolioSummary(ctx context.Context, params *migrationhubstrategy.GetPortfolioSummaryInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetPortfolioSummaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPortfolioSummary")
	}

	var r0 *migrationhubstrategy.GetPortfolioSummaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetPortfolioSummaryInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetPortfolioSummaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetPortfolioSummaryInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetPortfolioSummaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetPortfolioSummaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetPortfolioSummaryInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecommendationReportDetails provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetRecommendationReportDetails(ctx context.Context, params *migrationhubstrategy.GetRecommendationReportDetailsInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetRecommendationReportDetailsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRecommendationReportDetails")
	}

	var r0 *migrationhubstrategy.GetRecommendationReportDetailsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetRecommendationReportDetailsInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetRecommendationReportDetailsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetRecommendationReportDetailsInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetRecommendationReportDetailsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetRecommendationReportDetailsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetRecommendationReportDetailsInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerDetails provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetServerDetails(ctx context.Context, params *migrationhubstrategy.GetServerDetailsInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetServerDetailsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetServerDetails")
	}

	var r0 *migrationhubstrategy.GetServerDetailsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetServerDetailsInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetServerDetailsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetServerDetailsInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetServerDetailsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetServerDetailsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetServerDetailsInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerStrategies provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) GetServerStrategies(ctx context.Context, params *migrationhubstrategy.GetServerStrategiesInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetServerStrategiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetServerStrategies")
	}

	var r0 *migrationhubstrategy.GetServerStrategiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetServerStrategiesInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.GetServerStrategiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.GetServerStrategiesInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.GetServerStrategiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.GetServerStrategiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.GetServerStrategiesInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAnalyzableServers provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) ListAnalyzableServers(ctx context.Context, params *migrationhubstrategy.ListAnalyzableServersInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListAnalyzableServersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAnalyzableServers")
	}

	var r0 *migrationhubstrategy.ListAnalyzableServersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListAnalyzableServersInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListAnalyzableServersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListAnalyzableServersInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.ListAnalyzableServersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.ListAnalyzableServersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.ListAnalyzableServersInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplicationComponents provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) ListApplicationComponents(ctx context.Context, params *migrationhubstrategy.ListApplicationComponentsInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListApplicationComponentsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListApplicationComponents")
	}

	var r0 *migrationhubstrategy.ListApplicationComponentsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListApplicationComponentsInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListApplicationComponentsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListApplicationComponentsInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.ListApplicationComponentsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.ListApplicationComponentsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.ListApplicationComponentsInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCollectors provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) ListCollectors(ctx context.Context, params *migrationhubstrategy.ListCollectorsInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListCollectorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCollectors")
	}

	var r0 *migrationhubstrategy.ListCollectorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListCollectorsInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListCollectorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListCollectorsInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.ListCollectorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.ListCollectorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.ListCollectorsInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImportFileTask provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) ListImportFileTask(ctx context.Context, params *migrationhubstrategy.ListImportFileTaskInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListImportFileTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListImportFileTask")
	}

	var r0 *migrationhubstrategy.ListImportFileTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListImportFileTaskInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListImportFileTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListImportFileTaskInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.ListImportFileTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.ListImportFileTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.ListImportFileTaskInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServers provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) ListServers(ctx context.Context, params *migrationhubstrategy.ListServersInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListServersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListServers")
	}

	var r0 *migrationhubstrategy.ListServersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListServersInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.ListServersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.ListServersInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.ListServersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.ListServersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.ListServersInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *MigrationhubstrategyClient) Options() migrationhubstrategy.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 migrationhubstrategy.Options
	if rf, ok := ret.Get(0).(func() migrationhubstrategy.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(migrationhubstrategy.Options)
	}

	return r0
}

// PutPortfolioPreferences provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) PutPortfolioPreferences(ctx context.Context, params *migrationhubstrategy.PutPortfolioPreferencesInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.PutPortfolioPreferencesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutPortfolioPreferences")
	}

	var r0 *migrationhubstrategy.PutPortfolioPreferencesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.PutPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.PutPortfolioPreferencesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.PutPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.PutPortfolioPreferencesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.PutPortfolioPreferencesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.PutPortfolioPreferencesInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartAssessment provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) StartAssessment(ctx context.Context, params *migrationhubstrategy.StartAssessmentInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartAssessmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartAssessment")
	}

	var r0 *migrationhubstrategy.StartAssessmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartAssessmentInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartAssessmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartAssessmentInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.StartAssessmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.StartAssessmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.StartAssessmentInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartImportFileTask provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) StartImportFileTask(ctx context.Context, params *migrationhubstrategy.StartImportFileTaskInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartImportFileTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartImportFileTask")
	}

	var r0 *migrationhubstrategy.StartImportFileTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartImportFileTaskInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartImportFileTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartImportFileTaskInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.StartImportFileTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.StartImportFileTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.StartImportFileTaskInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartRecommendationReportGeneration provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) StartRecommendationReportGeneration(ctx context.Context, params *migrationhubstrategy.StartRecommendationReportGenerationInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartRecommendationReportGenerationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartRecommendationReportGeneration")
	}

	var r0 *migrationhubstrategy.StartRecommendationReportGenerationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartRecommendationReportGenerationInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StartRecommendationReportGenerationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StartRecommendationReportGenerationInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.StartRecommendationReportGenerationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.StartRecommendationReportGenerationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.StartRecommendationReportGenerationInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopAssessment provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) StopAssessment(ctx context.Context, params *migrationhubstrategy.StopAssessmentInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StopAssessmentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopAssessment")
	}

	var r0 *migrationhubstrategy.StopAssessmentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StopAssessmentInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.StopAssessmentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.StopAssessmentInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.StopAssessmentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.StopAssessmentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.StopAssessmentInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApplicationComponentConfig provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) UpdateApplicationComponentConfig(ctx context.Context, params *migrationhubstrategy.UpdateApplicationComponentConfigInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.UpdateApplicationComponentConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateApplicationComponentConfig")
	}

	var r0 *migrationhubstrategy.UpdateApplicationComponentConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.UpdateApplicationComponentConfigInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.UpdateApplicationComponentConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.UpdateApplicationComponentConfigInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.UpdateApplicationComponentConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.UpdateApplicationComponentConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.UpdateApplicationComponentConfigInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateServerConfig provides a mock function with given fields: ctx, params, optFns
func (_m *MigrationhubstrategyClient) UpdateServerConfig(ctx context.Context, params *migrationhubstrategy.UpdateServerConfigInput, optFns ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.UpdateServerConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateServerConfig")
	}

	var r0 *migrationhubstrategy.UpdateServerConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.UpdateServerConfigInput, ...func(*migrationhubstrategy.Options)) (*migrationhubstrategy.UpdateServerConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhubstrategy.UpdateServerConfigInput, ...func(*migrationhubstrategy.Options)) *migrationhubstrategy.UpdateServerConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhubstrategy.UpdateServerConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhubstrategy.UpdateServerConfigInput, ...func(*migrationhubstrategy.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMigrationhubstrategyClient creates a new instance of MigrationhubstrategyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMigrationhubstrategyClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MigrationhubstrategyClient {
	mock := &MigrationhubstrategyClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
