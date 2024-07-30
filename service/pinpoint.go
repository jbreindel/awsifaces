// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

// PinpointClient ...
type PinpointClient interface {
	Options() pinpoint.Options
	CreateApp(ctx context.Context, params *pinpoint.CreateAppInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateAppOutput, error)
	CreateCampaign(ctx context.Context, params *pinpoint.CreateCampaignInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateCampaignOutput, error)
	CreateEmailTemplate(ctx context.Context, params *pinpoint.CreateEmailTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateEmailTemplateOutput, error)
	CreateExportJob(ctx context.Context, params *pinpoint.CreateExportJobInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateExportJobOutput, error)
	CreateImportJob(ctx context.Context, params *pinpoint.CreateImportJobInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateImportJobOutput, error)
	CreateInAppTemplate(ctx context.Context, params *pinpoint.CreateInAppTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateInAppTemplateOutput, error)
	CreateJourney(ctx context.Context, params *pinpoint.CreateJourneyInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateJourneyOutput, error)
	CreatePushTemplate(ctx context.Context, params *pinpoint.CreatePushTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreatePushTemplateOutput, error)
	CreateRecommenderConfiguration(ctx context.Context, params *pinpoint.CreateRecommenderConfigurationInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateRecommenderConfigurationOutput, error)
	CreateSegment(ctx context.Context, params *pinpoint.CreateSegmentInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateSegmentOutput, error)
	CreateSmsTemplate(ctx context.Context, params *pinpoint.CreateSmsTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateSmsTemplateOutput, error)
	CreateVoiceTemplate(ctx context.Context, params *pinpoint.CreateVoiceTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.CreateVoiceTemplateOutput, error)
	DeleteAdmChannel(ctx context.Context, params *pinpoint.DeleteAdmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteAdmChannelOutput, error)
	DeleteApnsChannel(ctx context.Context, params *pinpoint.DeleteApnsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteApnsChannelOutput, error)
	DeleteApnsSandboxChannel(ctx context.Context, params *pinpoint.DeleteApnsSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteApnsSandboxChannelOutput, error)
	DeleteApnsVoipChannel(ctx context.Context, params *pinpoint.DeleteApnsVoipChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteApnsVoipChannelOutput, error)
	DeleteApnsVoipSandboxChannel(ctx context.Context, params *pinpoint.DeleteApnsVoipSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error)
	DeleteApp(ctx context.Context, params *pinpoint.DeleteAppInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteAppOutput, error)
	DeleteBaiduChannel(ctx context.Context, params *pinpoint.DeleteBaiduChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteBaiduChannelOutput, error)
	DeleteCampaign(ctx context.Context, params *pinpoint.DeleteCampaignInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteCampaignOutput, error)
	DeleteEmailChannel(ctx context.Context, params *pinpoint.DeleteEmailChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteEmailChannelOutput, error)
	DeleteEmailTemplate(ctx context.Context, params *pinpoint.DeleteEmailTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteEmailTemplateOutput, error)
	DeleteEndpoint(ctx context.Context, params *pinpoint.DeleteEndpointInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteEndpointOutput, error)
	DeleteEventStream(ctx context.Context, params *pinpoint.DeleteEventStreamInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteEventStreamOutput, error)
	DeleteGcmChannel(ctx context.Context, params *pinpoint.DeleteGcmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteGcmChannelOutput, error)
	DeleteInAppTemplate(ctx context.Context, params *pinpoint.DeleteInAppTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteInAppTemplateOutput, error)
	DeleteJourney(ctx context.Context, params *pinpoint.DeleteJourneyInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteJourneyOutput, error)
	DeletePushTemplate(ctx context.Context, params *pinpoint.DeletePushTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeletePushTemplateOutput, error)
	DeleteRecommenderConfiguration(ctx context.Context, params *pinpoint.DeleteRecommenderConfigurationInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteRecommenderConfigurationOutput, error)
	DeleteSegment(ctx context.Context, params *pinpoint.DeleteSegmentInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteSegmentOutput, error)
	DeleteSmsChannel(ctx context.Context, params *pinpoint.DeleteSmsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteSmsChannelOutput, error)
	DeleteSmsTemplate(ctx context.Context, params *pinpoint.DeleteSmsTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteSmsTemplateOutput, error)
	DeleteUserEndpoints(ctx context.Context, params *pinpoint.DeleteUserEndpointsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteUserEndpointsOutput, error)
	DeleteVoiceChannel(ctx context.Context, params *pinpoint.DeleteVoiceChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteVoiceChannelOutput, error)
	DeleteVoiceTemplate(ctx context.Context, params *pinpoint.DeleteVoiceTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.DeleteVoiceTemplateOutput, error)
	GetAdmChannel(ctx context.Context, params *pinpoint.GetAdmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetAdmChannelOutput, error)
	GetApnsChannel(ctx context.Context, params *pinpoint.GetApnsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApnsChannelOutput, error)
	GetApnsSandboxChannel(ctx context.Context, params *pinpoint.GetApnsSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApnsSandboxChannelOutput, error)
	GetApnsVoipChannel(ctx context.Context, params *pinpoint.GetApnsVoipChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApnsVoipChannelOutput, error)
	GetApnsVoipSandboxChannel(ctx context.Context, params *pinpoint.GetApnsVoipSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApnsVoipSandboxChannelOutput, error)
	GetApp(ctx context.Context, params *pinpoint.GetAppInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetAppOutput, error)
	GetApplicationDateRangeKpi(ctx context.Context, params *pinpoint.GetApplicationDateRangeKpiInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApplicationDateRangeKpiOutput, error)
	GetApplicationSettings(ctx context.Context, params *pinpoint.GetApplicationSettingsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetApplicationSettingsOutput, error)
	GetApps(ctx context.Context, params *pinpoint.GetAppsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetAppsOutput, error)
	GetBaiduChannel(ctx context.Context, params *pinpoint.GetBaiduChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetBaiduChannelOutput, error)
	GetCampaign(ctx context.Context, params *pinpoint.GetCampaignInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignOutput, error)
	GetCampaignActivities(ctx context.Context, params *pinpoint.GetCampaignActivitiesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignActivitiesOutput, error)
	GetCampaignDateRangeKpi(ctx context.Context, params *pinpoint.GetCampaignDateRangeKpiInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignDateRangeKpiOutput, error)
	GetCampaignVersion(ctx context.Context, params *pinpoint.GetCampaignVersionInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignVersionOutput, error)
	GetCampaignVersions(ctx context.Context, params *pinpoint.GetCampaignVersionsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignVersionsOutput, error)
	GetCampaigns(ctx context.Context, params *pinpoint.GetCampaignsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetCampaignsOutput, error)
	GetChannels(ctx context.Context, params *pinpoint.GetChannelsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetChannelsOutput, error)
	GetEmailChannel(ctx context.Context, params *pinpoint.GetEmailChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetEmailChannelOutput, error)
	GetEmailTemplate(ctx context.Context, params *pinpoint.GetEmailTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetEmailTemplateOutput, error)
	GetEndpoint(ctx context.Context, params *pinpoint.GetEndpointInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetEndpointOutput, error)
	GetEventStream(ctx context.Context, params *pinpoint.GetEventStreamInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetEventStreamOutput, error)
	GetExportJob(ctx context.Context, params *pinpoint.GetExportJobInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetExportJobOutput, error)
	GetExportJobs(ctx context.Context, params *pinpoint.GetExportJobsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetExportJobsOutput, error)
	GetGcmChannel(ctx context.Context, params *pinpoint.GetGcmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetGcmChannelOutput, error)
	GetImportJob(ctx context.Context, params *pinpoint.GetImportJobInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetImportJobOutput, error)
	GetImportJobs(ctx context.Context, params *pinpoint.GetImportJobsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetImportJobsOutput, error)
	GetInAppMessages(ctx context.Context, params *pinpoint.GetInAppMessagesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetInAppMessagesOutput, error)
	GetInAppTemplate(ctx context.Context, params *pinpoint.GetInAppTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetInAppTemplateOutput, error)
	GetJourney(ctx context.Context, params *pinpoint.GetJourneyInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyOutput, error)
	GetJourneyDateRangeKpi(ctx context.Context, params *pinpoint.GetJourneyDateRangeKpiInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyDateRangeKpiOutput, error)
	GetJourneyExecutionActivityMetrics(ctx context.Context, params *pinpoint.GetJourneyExecutionActivityMetricsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error)
	GetJourneyExecutionMetrics(ctx context.Context, params *pinpoint.GetJourneyExecutionMetricsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyExecutionMetricsOutput, error)
	GetJourneyRunExecutionActivityMetrics(ctx context.Context, params *pinpoint.GetJourneyRunExecutionActivityMetricsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyRunExecutionActivityMetricsOutput, error)
	GetJourneyRunExecutionMetrics(ctx context.Context, params *pinpoint.GetJourneyRunExecutionMetricsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyRunExecutionMetricsOutput, error)
	GetJourneyRuns(ctx context.Context, params *pinpoint.GetJourneyRunsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetJourneyRunsOutput, error)
	GetPushTemplate(ctx context.Context, params *pinpoint.GetPushTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetPushTemplateOutput, error)
	GetRecommenderConfiguration(ctx context.Context, params *pinpoint.GetRecommenderConfigurationInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetRecommenderConfigurationOutput, error)
	GetRecommenderConfigurations(ctx context.Context, params *pinpoint.GetRecommenderConfigurationsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetRecommenderConfigurationsOutput, error)
	GetSegment(ctx context.Context, params *pinpoint.GetSegmentInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentOutput, error)
	GetSegmentExportJobs(ctx context.Context, params *pinpoint.GetSegmentExportJobsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentExportJobsOutput, error)
	GetSegmentImportJobs(ctx context.Context, params *pinpoint.GetSegmentImportJobsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentImportJobsOutput, error)
	GetSegmentVersion(ctx context.Context, params *pinpoint.GetSegmentVersionInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentVersionOutput, error)
	GetSegmentVersions(ctx context.Context, params *pinpoint.GetSegmentVersionsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentVersionsOutput, error)
	GetSegments(ctx context.Context, params *pinpoint.GetSegmentsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSegmentsOutput, error)
	GetSmsChannel(ctx context.Context, params *pinpoint.GetSmsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSmsChannelOutput, error)
	GetSmsTemplate(ctx context.Context, params *pinpoint.GetSmsTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetSmsTemplateOutput, error)
	GetUserEndpoints(ctx context.Context, params *pinpoint.GetUserEndpointsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetUserEndpointsOutput, error)
	GetVoiceChannel(ctx context.Context, params *pinpoint.GetVoiceChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetVoiceChannelOutput, error)
	GetVoiceTemplate(ctx context.Context, params *pinpoint.GetVoiceTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.GetVoiceTemplateOutput, error)
	ListJourneys(ctx context.Context, params *pinpoint.ListJourneysInput, optFns ...func(*pinpoint.Options)) (*pinpoint.ListJourneysOutput, error)
	ListTagsForResource(ctx context.Context, params *pinpoint.ListTagsForResourceInput, optFns ...func(*pinpoint.Options)) (*pinpoint.ListTagsForResourceOutput, error)
	ListTemplateVersions(ctx context.Context, params *pinpoint.ListTemplateVersionsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.ListTemplateVersionsOutput, error)
	ListTemplates(ctx context.Context, params *pinpoint.ListTemplatesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.ListTemplatesOutput, error)
	PhoneNumberValidate(ctx context.Context, params *pinpoint.PhoneNumberValidateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.PhoneNumberValidateOutput, error)
	PutEventStream(ctx context.Context, params *pinpoint.PutEventStreamInput, optFns ...func(*pinpoint.Options)) (*pinpoint.PutEventStreamOutput, error)
	PutEvents(ctx context.Context, params *pinpoint.PutEventsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.PutEventsOutput, error)
	RemoveAttributes(ctx context.Context, params *pinpoint.RemoveAttributesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.RemoveAttributesOutput, error)
	SendMessages(ctx context.Context, params *pinpoint.SendMessagesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.SendMessagesOutput, error)
	SendOTPMessage(ctx context.Context, params *pinpoint.SendOTPMessageInput, optFns ...func(*pinpoint.Options)) (*pinpoint.SendOTPMessageOutput, error)
	SendUsersMessages(ctx context.Context, params *pinpoint.SendUsersMessagesInput, optFns ...func(*pinpoint.Options)) (*pinpoint.SendUsersMessagesOutput, error)
	TagResource(ctx context.Context, params *pinpoint.TagResourceInput, optFns ...func(*pinpoint.Options)) (*pinpoint.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *pinpoint.UntagResourceInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UntagResourceOutput, error)
	UpdateAdmChannel(ctx context.Context, params *pinpoint.UpdateAdmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateAdmChannelOutput, error)
	UpdateApnsChannel(ctx context.Context, params *pinpoint.UpdateApnsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateApnsChannelOutput, error)
	UpdateApnsSandboxChannel(ctx context.Context, params *pinpoint.UpdateApnsSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateApnsSandboxChannelOutput, error)
	UpdateApnsVoipChannel(ctx context.Context, params *pinpoint.UpdateApnsVoipChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateApnsVoipChannelOutput, error)
	UpdateApnsVoipSandboxChannel(ctx context.Context, params *pinpoint.UpdateApnsVoipSandboxChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error)
	UpdateApplicationSettings(ctx context.Context, params *pinpoint.UpdateApplicationSettingsInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateApplicationSettingsOutput, error)
	UpdateBaiduChannel(ctx context.Context, params *pinpoint.UpdateBaiduChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateBaiduChannelOutput, error)
	UpdateCampaign(ctx context.Context, params *pinpoint.UpdateCampaignInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateCampaignOutput, error)
	UpdateEmailChannel(ctx context.Context, params *pinpoint.UpdateEmailChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateEmailChannelOutput, error)
	UpdateEmailTemplate(ctx context.Context, params *pinpoint.UpdateEmailTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateEmailTemplateOutput, error)
	UpdateEndpoint(ctx context.Context, params *pinpoint.UpdateEndpointInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateEndpointOutput, error)
	UpdateEndpointsBatch(ctx context.Context, params *pinpoint.UpdateEndpointsBatchInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateEndpointsBatchOutput, error)
	UpdateGcmChannel(ctx context.Context, params *pinpoint.UpdateGcmChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateGcmChannelOutput, error)
	UpdateInAppTemplate(ctx context.Context, params *pinpoint.UpdateInAppTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateInAppTemplateOutput, error)
	UpdateJourney(ctx context.Context, params *pinpoint.UpdateJourneyInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateJourneyOutput, error)
	UpdateJourneyState(ctx context.Context, params *pinpoint.UpdateJourneyStateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateJourneyStateOutput, error)
	UpdatePushTemplate(ctx context.Context, params *pinpoint.UpdatePushTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdatePushTemplateOutput, error)
	UpdateRecommenderConfiguration(ctx context.Context, params *pinpoint.UpdateRecommenderConfigurationInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateRecommenderConfigurationOutput, error)
	UpdateSegment(ctx context.Context, params *pinpoint.UpdateSegmentInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateSegmentOutput, error)
	UpdateSmsChannel(ctx context.Context, params *pinpoint.UpdateSmsChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateSmsChannelOutput, error)
	UpdateSmsTemplate(ctx context.Context, params *pinpoint.UpdateSmsTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateSmsTemplateOutput, error)
	UpdateTemplateActiveVersion(ctx context.Context, params *pinpoint.UpdateTemplateActiveVersionInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateTemplateActiveVersionOutput, error)
	UpdateVoiceChannel(ctx context.Context, params *pinpoint.UpdateVoiceChannelInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateVoiceChannelOutput, error)
	UpdateVoiceTemplate(ctx context.Context, params *pinpoint.UpdateVoiceTemplateInput, optFns ...func(*pinpoint.Options)) (*pinpoint.UpdateVoiceTemplateOutput, error)
	VerifyOTPMessage(ctx context.Context, params *pinpoint.VerifyOTPMessageInput, optFns ...func(*pinpoint.Options)) (*pinpoint.VerifyOTPMessageOutput, error)
}
