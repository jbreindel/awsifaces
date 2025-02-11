// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

// Iot1clickdevicesserviceClient ...
type Iot1clickdevicesserviceClient interface {
	Options() iot1clickdevicesservice.Options
	ClaimDevicesByClaimCode(ctx context.Context, params *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error)
	DescribeDevice(ctx context.Context, params *iot1clickdevicesservice.DescribeDeviceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.DescribeDeviceOutput, error)
	FinalizeDeviceClaim(ctx context.Context, params *iot1clickdevicesservice.FinalizeDeviceClaimInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error)
	GetDeviceMethods(ctx context.Context, params *iot1clickdevicesservice.GetDeviceMethodsInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error)
	InitiateDeviceClaim(ctx context.Context, params *iot1clickdevicesservice.InitiateDeviceClaimInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error)
	InvokeDeviceMethod(ctx context.Context, params *iot1clickdevicesservice.InvokeDeviceMethodInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error)
	ListDeviceEvents(ctx context.Context, params *iot1clickdevicesservice.ListDeviceEventsInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDeviceEventsOutput, error)
	ListDevices(ctx context.Context, params *iot1clickdevicesservice.ListDevicesInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListDevicesOutput, error)
	ListTagsForResource(ctx context.Context, params *iot1clickdevicesservice.ListTagsForResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.ListTagsForResourceOutput, error)
	TagResource(ctx context.Context, params *iot1clickdevicesservice.TagResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.TagResourceOutput, error)
	UnclaimDevice(ctx context.Context, params *iot1clickdevicesservice.UnclaimDeviceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UnclaimDeviceOutput, error)
	UntagResource(ctx context.Context, params *iot1clickdevicesservice.UntagResourceInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UntagResourceOutput, error)
	UpdateDeviceState(ctx context.Context, params *iot1clickdevicesservice.UpdateDeviceStateInput, optFns ...func(*iot1clickdevicesservice.Options)) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error)
}
