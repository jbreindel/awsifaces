// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
)

// ConnectcampaignsClient ...
type ConnectcampaignsClient interface {
	Options() connectcampaigns.Options
	CreateCampaign(ctx context.Context, params *connectcampaigns.CreateCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.CreateCampaignOutput, error)
	DeleteCampaign(ctx context.Context, params *connectcampaigns.DeleteCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteCampaignOutput, error)
	DeleteConnectInstanceConfig(ctx context.Context, params *connectcampaigns.DeleteConnectInstanceConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteConnectInstanceConfigOutput, error)
	DeleteInstanceOnboardingJob(ctx context.Context, params *connectcampaigns.DeleteInstanceOnboardingJobInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteInstanceOnboardingJobOutput, error)
	DescribeCampaign(ctx context.Context, params *connectcampaigns.DescribeCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DescribeCampaignOutput, error)
	GetCampaignState(ctx context.Context, params *connectcampaigns.GetCampaignStateInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateOutput, error)
	GetCampaignStateBatch(ctx context.Context, params *connectcampaigns.GetCampaignStateBatchInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateBatchOutput, error)
	GetConnectInstanceConfig(ctx context.Context, params *connectcampaigns.GetConnectInstanceConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetConnectInstanceConfigOutput, error)
	GetInstanceOnboardingJobStatus(ctx context.Context, params *connectcampaigns.GetInstanceOnboardingJobStatusInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error)
	ListCampaigns(ctx context.Context, params *connectcampaigns.ListCampaignsInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ListCampaignsOutput, error)
	ListTagsForResource(ctx context.Context, params *connectcampaigns.ListTagsForResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ListTagsForResourceOutput, error)
	PauseCampaign(ctx context.Context, params *connectcampaigns.PauseCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.PauseCampaignOutput, error)
	PutDialRequestBatch(ctx context.Context, params *connectcampaigns.PutDialRequestBatchInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.PutDialRequestBatchOutput, error)
	ResumeCampaign(ctx context.Context, params *connectcampaigns.ResumeCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ResumeCampaignOutput, error)
	StartCampaign(ctx context.Context, params *connectcampaigns.StartCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StartCampaignOutput, error)
	StartInstanceOnboardingJob(ctx context.Context, params *connectcampaigns.StartInstanceOnboardingJobInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StartInstanceOnboardingJobOutput, error)
	StopCampaign(ctx context.Context, params *connectcampaigns.StopCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StopCampaignOutput, error)
	TagResource(ctx context.Context, params *connectcampaigns.TagResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *connectcampaigns.UntagResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UntagResourceOutput, error)
	UpdateCampaignDialerConfig(ctx context.Context, params *connectcampaigns.UpdateCampaignDialerConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignDialerConfigOutput, error)
	UpdateCampaignName(ctx context.Context, params *connectcampaigns.UpdateCampaignNameInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignNameOutput, error)
	UpdateCampaignOutboundCallConfig(ctx context.Context, params *connectcampaigns.UpdateCampaignOutboundCallConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignOutboundCallConfigOutput, error)
}
