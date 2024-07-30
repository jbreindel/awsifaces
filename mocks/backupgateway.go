// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	backupgateway "github.com/aws/aws-sdk-go-v2/service/backupgateway"

	mock "github.com/stretchr/testify/mock"
)

// BackupgatewayClient is an autogenerated mock type for the BackupgatewayClient type
type BackupgatewayClient struct {
	mock.Mock
}

// AssociateGatewayToServer provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) AssociateGatewayToServer(ctx context.Context, params *backupgateway.AssociateGatewayToServerInput, optFns ...func(*backupgateway.Options)) (*backupgateway.AssociateGatewayToServerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateGatewayToServer")
	}

	var r0 *backupgateway.AssociateGatewayToServerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.AssociateGatewayToServerInput, ...func(*backupgateway.Options)) (*backupgateway.AssociateGatewayToServerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.AssociateGatewayToServerInput, ...func(*backupgateway.Options)) *backupgateway.AssociateGatewayToServerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.AssociateGatewayToServerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.AssociateGatewayToServerInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGateway provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) CreateGateway(ctx context.Context, params *backupgateway.CreateGatewayInput, optFns ...func(*backupgateway.Options)) (*backupgateway.CreateGatewayOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateGateway")
	}

	var r0 *backupgateway.CreateGatewayOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.CreateGatewayInput, ...func(*backupgateway.Options)) (*backupgateway.CreateGatewayOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.CreateGatewayInput, ...func(*backupgateway.Options)) *backupgateway.CreateGatewayOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.CreateGatewayOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.CreateGatewayInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGateway provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) DeleteGateway(ctx context.Context, params *backupgateway.DeleteGatewayInput, optFns ...func(*backupgateway.Options)) (*backupgateway.DeleteGatewayOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGateway")
	}

	var r0 *backupgateway.DeleteGatewayOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DeleteGatewayInput, ...func(*backupgateway.Options)) (*backupgateway.DeleteGatewayOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DeleteGatewayInput, ...func(*backupgateway.Options)) *backupgateway.DeleteGatewayOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.DeleteGatewayOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.DeleteGatewayInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHypervisor provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) DeleteHypervisor(ctx context.Context, params *backupgateway.DeleteHypervisorInput, optFns ...func(*backupgateway.Options)) (*backupgateway.DeleteHypervisorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteHypervisor")
	}

	var r0 *backupgateway.DeleteHypervisorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DeleteHypervisorInput, ...func(*backupgateway.Options)) (*backupgateway.DeleteHypervisorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DeleteHypervisorInput, ...func(*backupgateway.Options)) *backupgateway.DeleteHypervisorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.DeleteHypervisorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.DeleteHypervisorInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateGatewayFromServer provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) DisassociateGatewayFromServer(ctx context.Context, params *backupgateway.DisassociateGatewayFromServerInput, optFns ...func(*backupgateway.Options)) (*backupgateway.DisassociateGatewayFromServerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateGatewayFromServer")
	}

	var r0 *backupgateway.DisassociateGatewayFromServerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DisassociateGatewayFromServerInput, ...func(*backupgateway.Options)) (*backupgateway.DisassociateGatewayFromServerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.DisassociateGatewayFromServerInput, ...func(*backupgateway.Options)) *backupgateway.DisassociateGatewayFromServerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.DisassociateGatewayFromServerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.DisassociateGatewayFromServerInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBandwidthRateLimitSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) GetBandwidthRateLimitSchedule(ctx context.Context, params *backupgateway.GetBandwidthRateLimitScheduleInput, optFns ...func(*backupgateway.Options)) (*backupgateway.GetBandwidthRateLimitScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBandwidthRateLimitSchedule")
	}

	var r0 *backupgateway.GetBandwidthRateLimitScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) (*backupgateway.GetBandwidthRateLimitScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) *backupgateway.GetBandwidthRateLimitScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.GetBandwidthRateLimitScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.GetBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGateway provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) GetGateway(ctx context.Context, params *backupgateway.GetGatewayInput, optFns ...func(*backupgateway.Options)) (*backupgateway.GetGatewayOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGateway")
	}

	var r0 *backupgateway.GetGatewayOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetGatewayInput, ...func(*backupgateway.Options)) (*backupgateway.GetGatewayOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetGatewayInput, ...func(*backupgateway.Options)) *backupgateway.GetGatewayOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.GetGatewayOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.GetGatewayInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHypervisor provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) GetHypervisor(ctx context.Context, params *backupgateway.GetHypervisorInput, optFns ...func(*backupgateway.Options)) (*backupgateway.GetHypervisorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetHypervisor")
	}

	var r0 *backupgateway.GetHypervisorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetHypervisorInput, ...func(*backupgateway.Options)) (*backupgateway.GetHypervisorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetHypervisorInput, ...func(*backupgateway.Options)) *backupgateway.GetHypervisorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.GetHypervisorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.GetHypervisorInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHypervisorPropertyMappings provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) GetHypervisorPropertyMappings(ctx context.Context, params *backupgateway.GetHypervisorPropertyMappingsInput, optFns ...func(*backupgateway.Options)) (*backupgateway.GetHypervisorPropertyMappingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetHypervisorPropertyMappings")
	}

	var r0 *backupgateway.GetHypervisorPropertyMappingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) (*backupgateway.GetHypervisorPropertyMappingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) *backupgateway.GetHypervisorPropertyMappingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.GetHypervisorPropertyMappingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.GetHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVirtualMachine provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) GetVirtualMachine(ctx context.Context, params *backupgateway.GetVirtualMachineInput, optFns ...func(*backupgateway.Options)) (*backupgateway.GetVirtualMachineOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetVirtualMachine")
	}

	var r0 *backupgateway.GetVirtualMachineOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetVirtualMachineInput, ...func(*backupgateway.Options)) (*backupgateway.GetVirtualMachineOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.GetVirtualMachineInput, ...func(*backupgateway.Options)) *backupgateway.GetVirtualMachineOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.GetVirtualMachineOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.GetVirtualMachineInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportHypervisorConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) ImportHypervisorConfiguration(ctx context.Context, params *backupgateway.ImportHypervisorConfigurationInput, optFns ...func(*backupgateway.Options)) (*backupgateway.ImportHypervisorConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ImportHypervisorConfiguration")
	}

	var r0 *backupgateway.ImportHypervisorConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ImportHypervisorConfigurationInput, ...func(*backupgateway.Options)) (*backupgateway.ImportHypervisorConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ImportHypervisorConfigurationInput, ...func(*backupgateway.Options)) *backupgateway.ImportHypervisorConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.ImportHypervisorConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.ImportHypervisorConfigurationInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGateways provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) ListGateways(ctx context.Context, params *backupgateway.ListGatewaysInput, optFns ...func(*backupgateway.Options)) (*backupgateway.ListGatewaysOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGateways")
	}

	var r0 *backupgateway.ListGatewaysOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListGatewaysInput, ...func(*backupgateway.Options)) (*backupgateway.ListGatewaysOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListGatewaysInput, ...func(*backupgateway.Options)) *backupgateway.ListGatewaysOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.ListGatewaysOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.ListGatewaysInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHypervisors provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) ListHypervisors(ctx context.Context, params *backupgateway.ListHypervisorsInput, optFns ...func(*backupgateway.Options)) (*backupgateway.ListHypervisorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListHypervisors")
	}

	var r0 *backupgateway.ListHypervisorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListHypervisorsInput, ...func(*backupgateway.Options)) (*backupgateway.ListHypervisorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListHypervisorsInput, ...func(*backupgateway.Options)) *backupgateway.ListHypervisorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.ListHypervisorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.ListHypervisorsInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) ListTagsForResource(ctx context.Context, params *backupgateway.ListTagsForResourceInput, optFns ...func(*backupgateway.Options)) (*backupgateway.ListTagsForResourceOutput, error) {
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

	var r0 *backupgateway.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListTagsForResourceInput, ...func(*backupgateway.Options)) (*backupgateway.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListTagsForResourceInput, ...func(*backupgateway.Options)) *backupgateway.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.ListTagsForResourceInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVirtualMachines provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) ListVirtualMachines(ctx context.Context, params *backupgateway.ListVirtualMachinesInput, optFns ...func(*backupgateway.Options)) (*backupgateway.ListVirtualMachinesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListVirtualMachines")
	}

	var r0 *backupgateway.ListVirtualMachinesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListVirtualMachinesInput, ...func(*backupgateway.Options)) (*backupgateway.ListVirtualMachinesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.ListVirtualMachinesInput, ...func(*backupgateway.Options)) *backupgateway.ListVirtualMachinesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.ListVirtualMachinesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.ListVirtualMachinesInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *BackupgatewayClient) Options() backupgateway.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 backupgateway.Options
	if rf, ok := ret.Get(0).(func() backupgateway.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(backupgateway.Options)
	}

	return r0
}

// PutBandwidthRateLimitSchedule provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) PutBandwidthRateLimitSchedule(ctx context.Context, params *backupgateway.PutBandwidthRateLimitScheduleInput, optFns ...func(*backupgateway.Options)) (*backupgateway.PutBandwidthRateLimitScheduleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutBandwidthRateLimitSchedule")
	}

	var r0 *backupgateway.PutBandwidthRateLimitScheduleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) (*backupgateway.PutBandwidthRateLimitScheduleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) *backupgateway.PutBandwidthRateLimitScheduleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.PutBandwidthRateLimitScheduleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.PutBandwidthRateLimitScheduleInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutHypervisorPropertyMappings provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) PutHypervisorPropertyMappings(ctx context.Context, params *backupgateway.PutHypervisorPropertyMappingsInput, optFns ...func(*backupgateway.Options)) (*backupgateway.PutHypervisorPropertyMappingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutHypervisorPropertyMappings")
	}

	var r0 *backupgateway.PutHypervisorPropertyMappingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) (*backupgateway.PutHypervisorPropertyMappingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) *backupgateway.PutHypervisorPropertyMappingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.PutHypervisorPropertyMappingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.PutHypervisorPropertyMappingsInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMaintenanceStartTime provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) PutMaintenanceStartTime(ctx context.Context, params *backupgateway.PutMaintenanceStartTimeInput, optFns ...func(*backupgateway.Options)) (*backupgateway.PutMaintenanceStartTimeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutMaintenanceStartTime")
	}

	var r0 *backupgateway.PutMaintenanceStartTimeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutMaintenanceStartTimeInput, ...func(*backupgateway.Options)) (*backupgateway.PutMaintenanceStartTimeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.PutMaintenanceStartTimeInput, ...func(*backupgateway.Options)) *backupgateway.PutMaintenanceStartTimeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.PutMaintenanceStartTimeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.PutMaintenanceStartTimeInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartVirtualMachinesMetadataSync provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) StartVirtualMachinesMetadataSync(ctx context.Context, params *backupgateway.StartVirtualMachinesMetadataSyncInput, optFns ...func(*backupgateway.Options)) (*backupgateway.StartVirtualMachinesMetadataSyncOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartVirtualMachinesMetadataSync")
	}

	var r0 *backupgateway.StartVirtualMachinesMetadataSyncOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.StartVirtualMachinesMetadataSyncInput, ...func(*backupgateway.Options)) (*backupgateway.StartVirtualMachinesMetadataSyncOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.StartVirtualMachinesMetadataSyncInput, ...func(*backupgateway.Options)) *backupgateway.StartVirtualMachinesMetadataSyncOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.StartVirtualMachinesMetadataSyncOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.StartVirtualMachinesMetadataSyncInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) TagResource(ctx context.Context, params *backupgateway.TagResourceInput, optFns ...func(*backupgateway.Options)) (*backupgateway.TagResourceOutput, error) {
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

	var r0 *backupgateway.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.TagResourceInput, ...func(*backupgateway.Options)) (*backupgateway.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.TagResourceInput, ...func(*backupgateway.Options)) *backupgateway.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.TagResourceInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestHypervisorConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) TestHypervisorConfiguration(ctx context.Context, params *backupgateway.TestHypervisorConfigurationInput, optFns ...func(*backupgateway.Options)) (*backupgateway.TestHypervisorConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TestHypervisorConfiguration")
	}

	var r0 *backupgateway.TestHypervisorConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.TestHypervisorConfigurationInput, ...func(*backupgateway.Options)) (*backupgateway.TestHypervisorConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.TestHypervisorConfigurationInput, ...func(*backupgateway.Options)) *backupgateway.TestHypervisorConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.TestHypervisorConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.TestHypervisorConfigurationInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) UntagResource(ctx context.Context, params *backupgateway.UntagResourceInput, optFns ...func(*backupgateway.Options)) (*backupgateway.UntagResourceOutput, error) {
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

	var r0 *backupgateway.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UntagResourceInput, ...func(*backupgateway.Options)) (*backupgateway.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UntagResourceInput, ...func(*backupgateway.Options)) *backupgateway.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.UntagResourceInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGatewayInformation provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) UpdateGatewayInformation(ctx context.Context, params *backupgateway.UpdateGatewayInformationInput, optFns ...func(*backupgateway.Options)) (*backupgateway.UpdateGatewayInformationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGatewayInformation")
	}

	var r0 *backupgateway.UpdateGatewayInformationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateGatewayInformationInput, ...func(*backupgateway.Options)) (*backupgateway.UpdateGatewayInformationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateGatewayInformationInput, ...func(*backupgateway.Options)) *backupgateway.UpdateGatewayInformationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.UpdateGatewayInformationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.UpdateGatewayInformationInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGatewaySoftwareNow provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) UpdateGatewaySoftwareNow(ctx context.Context, params *backupgateway.UpdateGatewaySoftwareNowInput, optFns ...func(*backupgateway.Options)) (*backupgateway.UpdateGatewaySoftwareNowOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGatewaySoftwareNow")
	}

	var r0 *backupgateway.UpdateGatewaySoftwareNowOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateGatewaySoftwareNowInput, ...func(*backupgateway.Options)) (*backupgateway.UpdateGatewaySoftwareNowOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateGatewaySoftwareNowInput, ...func(*backupgateway.Options)) *backupgateway.UpdateGatewaySoftwareNowOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.UpdateGatewaySoftwareNowOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.UpdateGatewaySoftwareNowInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHypervisor provides a mock function with given fields: ctx, params, optFns
func (_m *BackupgatewayClient) UpdateHypervisor(ctx context.Context, params *backupgateway.UpdateHypervisorInput, optFns ...func(*backupgateway.Options)) (*backupgateway.UpdateHypervisorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateHypervisor")
	}

	var r0 *backupgateway.UpdateHypervisorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateHypervisorInput, ...func(*backupgateway.Options)) (*backupgateway.UpdateHypervisorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *backupgateway.UpdateHypervisorInput, ...func(*backupgateway.Options)) *backupgateway.UpdateHypervisorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupgateway.UpdateHypervisorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *backupgateway.UpdateHypervisorInput, ...func(*backupgateway.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBackupgatewayClient creates a new instance of BackupgatewayClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBackupgatewayClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BackupgatewayClient {
	mock := &BackupgatewayClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
