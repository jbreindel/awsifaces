// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iotevents"
)

// IoteventsClient ...
type IoteventsClient interface {
	Options() iotevents.Options
	CreateAlarmModel(ctx context.Context, params *iotevents.CreateAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateAlarmModelOutput, error)
	CreateDetectorModel(ctx context.Context, params *iotevents.CreateDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateDetectorModelOutput, error)
	CreateInput(ctx context.Context, params *iotevents.CreateInputInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateInputOutput, error)
	DeleteAlarmModel(ctx context.Context, params *iotevents.DeleteAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteAlarmModelOutput, error)
	DeleteDetectorModel(ctx context.Context, params *iotevents.DeleteDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteDetectorModelOutput, error)
	DeleteInput(ctx context.Context, params *iotevents.DeleteInputInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteInputOutput, error)
	DescribeAlarmModel(ctx context.Context, params *iotevents.DescribeAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeAlarmModelOutput, error)
	DescribeDetectorModel(ctx context.Context, params *iotevents.DescribeDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelOutput, error)
	DescribeDetectorModelAnalysis(ctx context.Context, params *iotevents.DescribeDetectorModelAnalysisInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelAnalysisOutput, error)
	DescribeInput(ctx context.Context, params *iotevents.DescribeInputInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeInputOutput, error)
	DescribeLoggingOptions(ctx context.Context, params *iotevents.DescribeLoggingOptionsInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeLoggingOptionsOutput, error)
	GetDetectorModelAnalysisResults(ctx context.Context, params *iotevents.GetDetectorModelAnalysisResultsInput, optFns ...func(*iotevents.Options)) (*iotevents.GetDetectorModelAnalysisResultsOutput, error)
	ListAlarmModelVersions(ctx context.Context, params *iotevents.ListAlarmModelVersionsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListAlarmModelVersionsOutput, error)
	ListAlarmModels(ctx context.Context, params *iotevents.ListAlarmModelsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListAlarmModelsOutput, error)
	ListDetectorModelVersions(ctx context.Context, params *iotevents.ListDetectorModelVersionsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListDetectorModelVersionsOutput, error)
	ListDetectorModels(ctx context.Context, params *iotevents.ListDetectorModelsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListDetectorModelsOutput, error)
	ListInputRoutings(ctx context.Context, params *iotevents.ListInputRoutingsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListInputRoutingsOutput, error)
	ListInputs(ctx context.Context, params *iotevents.ListInputsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListInputsOutput, error)
	ListTagsForResource(ctx context.Context, params *iotevents.ListTagsForResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.ListTagsForResourceOutput, error)
	PutLoggingOptions(ctx context.Context, params *iotevents.PutLoggingOptionsInput, optFns ...func(*iotevents.Options)) (*iotevents.PutLoggingOptionsOutput, error)
	StartDetectorModelAnalysis(ctx context.Context, params *iotevents.StartDetectorModelAnalysisInput, optFns ...func(*iotevents.Options)) (*iotevents.StartDetectorModelAnalysisOutput, error)
	TagResource(ctx context.Context, params *iotevents.TagResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *iotevents.UntagResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.UntagResourceOutput, error)
	UpdateAlarmModel(ctx context.Context, params *iotevents.UpdateAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateAlarmModelOutput, error)
	UpdateDetectorModel(ctx context.Context, params *iotevents.UpdateDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateDetectorModelOutput, error)
	UpdateInput(ctx context.Context, params *iotevents.UpdateInputInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateInputOutput, error)
}
