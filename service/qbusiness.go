// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/qbusiness"
)

// QbusinessClient ...
type QbusinessClient interface {
	Options() qbusiness.Options
	BatchDeleteDocument(ctx context.Context, params *qbusiness.BatchDeleteDocumentInput, optFns ...func(*qbusiness.Options)) (*qbusiness.BatchDeleteDocumentOutput, error)
	BatchPutDocument(ctx context.Context, params *qbusiness.BatchPutDocumentInput, optFns ...func(*qbusiness.Options)) (*qbusiness.BatchPutDocumentOutput, error)
	ChatSync(ctx context.Context, params *qbusiness.ChatSyncInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ChatSyncOutput, error)
	CreateApplication(ctx context.Context, params *qbusiness.CreateApplicationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateApplicationOutput, error)
	CreateDataSource(ctx context.Context, params *qbusiness.CreateDataSourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateDataSourceOutput, error)
	CreateIndex(ctx context.Context, params *qbusiness.CreateIndexInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateIndexOutput, error)
	CreatePlugin(ctx context.Context, params *qbusiness.CreatePluginInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreatePluginOutput, error)
	CreateRetriever(ctx context.Context, params *qbusiness.CreateRetrieverInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateRetrieverOutput, error)
	CreateUser(ctx context.Context, params *qbusiness.CreateUserInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateUserOutput, error)
	CreateWebExperience(ctx context.Context, params *qbusiness.CreateWebExperienceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.CreateWebExperienceOutput, error)
	DeleteApplication(ctx context.Context, params *qbusiness.DeleteApplicationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteApplicationOutput, error)
	DeleteChatControlsConfiguration(ctx context.Context, params *qbusiness.DeleteChatControlsConfigurationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteChatControlsConfigurationOutput, error)
	DeleteConversation(ctx context.Context, params *qbusiness.DeleteConversationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteConversationOutput, error)
	DeleteDataSource(ctx context.Context, params *qbusiness.DeleteDataSourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteDataSourceOutput, error)
	DeleteGroup(ctx context.Context, params *qbusiness.DeleteGroupInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteGroupOutput, error)
	DeleteIndex(ctx context.Context, params *qbusiness.DeleteIndexInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteIndexOutput, error)
	DeletePlugin(ctx context.Context, params *qbusiness.DeletePluginInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeletePluginOutput, error)
	DeleteRetriever(ctx context.Context, params *qbusiness.DeleteRetrieverInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteRetrieverOutput, error)
	DeleteUser(ctx context.Context, params *qbusiness.DeleteUserInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteUserOutput, error)
	DeleteWebExperience(ctx context.Context, params *qbusiness.DeleteWebExperienceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.DeleteWebExperienceOutput, error)
	GetApplication(ctx context.Context, params *qbusiness.GetApplicationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetApplicationOutput, error)
	GetChatControlsConfiguration(ctx context.Context, params *qbusiness.GetChatControlsConfigurationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetChatControlsConfigurationOutput, error)
	GetDataSource(ctx context.Context, params *qbusiness.GetDataSourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetDataSourceOutput, error)
	GetGroup(ctx context.Context, params *qbusiness.GetGroupInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetGroupOutput, error)
	GetIndex(ctx context.Context, params *qbusiness.GetIndexInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetIndexOutput, error)
	GetPlugin(ctx context.Context, params *qbusiness.GetPluginInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetPluginOutput, error)
	GetRetriever(ctx context.Context, params *qbusiness.GetRetrieverInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetRetrieverOutput, error)
	GetUser(ctx context.Context, params *qbusiness.GetUserInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetUserOutput, error)
	GetWebExperience(ctx context.Context, params *qbusiness.GetWebExperienceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.GetWebExperienceOutput, error)
	ListApplications(ctx context.Context, params *qbusiness.ListApplicationsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListApplicationsOutput, error)
	ListConversations(ctx context.Context, params *qbusiness.ListConversationsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListConversationsOutput, error)
	ListDataSourceSyncJobs(ctx context.Context, params *qbusiness.ListDataSourceSyncJobsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListDataSourceSyncJobsOutput, error)
	ListDataSources(ctx context.Context, params *qbusiness.ListDataSourcesInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListDataSourcesOutput, error)
	ListDocuments(ctx context.Context, params *qbusiness.ListDocumentsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListDocumentsOutput, error)
	ListGroups(ctx context.Context, params *qbusiness.ListGroupsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListGroupsOutput, error)
	ListIndices(ctx context.Context, params *qbusiness.ListIndicesInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListIndicesOutput, error)
	ListMessages(ctx context.Context, params *qbusiness.ListMessagesInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListMessagesOutput, error)
	ListPlugins(ctx context.Context, params *qbusiness.ListPluginsInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListPluginsOutput, error)
	ListRetrievers(ctx context.Context, params *qbusiness.ListRetrieversInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListRetrieversOutput, error)
	ListTagsForResource(ctx context.Context, params *qbusiness.ListTagsForResourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListTagsForResourceOutput, error)
	ListWebExperiences(ctx context.Context, params *qbusiness.ListWebExperiencesInput, optFns ...func(*qbusiness.Options)) (*qbusiness.ListWebExperiencesOutput, error)
	PutFeedback(ctx context.Context, params *qbusiness.PutFeedbackInput, optFns ...func(*qbusiness.Options)) (*qbusiness.PutFeedbackOutput, error)
	PutGroup(ctx context.Context, params *qbusiness.PutGroupInput, optFns ...func(*qbusiness.Options)) (*qbusiness.PutGroupOutput, error)
	StartDataSourceSyncJob(ctx context.Context, params *qbusiness.StartDataSourceSyncJobInput, optFns ...func(*qbusiness.Options)) (*qbusiness.StartDataSourceSyncJobOutput, error)
	StopDataSourceSyncJob(ctx context.Context, params *qbusiness.StopDataSourceSyncJobInput, optFns ...func(*qbusiness.Options)) (*qbusiness.StopDataSourceSyncJobOutput, error)
	TagResource(ctx context.Context, params *qbusiness.TagResourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *qbusiness.UntagResourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UntagResourceOutput, error)
	UpdateApplication(ctx context.Context, params *qbusiness.UpdateApplicationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateApplicationOutput, error)
	UpdateChatControlsConfiguration(ctx context.Context, params *qbusiness.UpdateChatControlsConfigurationInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateChatControlsConfigurationOutput, error)
	UpdateDataSource(ctx context.Context, params *qbusiness.UpdateDataSourceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateDataSourceOutput, error)
	UpdateIndex(ctx context.Context, params *qbusiness.UpdateIndexInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateIndexOutput, error)
	UpdatePlugin(ctx context.Context, params *qbusiness.UpdatePluginInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdatePluginOutput, error)
	UpdateRetriever(ctx context.Context, params *qbusiness.UpdateRetrieverInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateRetrieverOutput, error)
	UpdateUser(ctx context.Context, params *qbusiness.UpdateUserInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateUserOutput, error)
	UpdateWebExperience(ctx context.Context, params *qbusiness.UpdateWebExperienceInput, optFns ...func(*qbusiness.Options)) (*qbusiness.UpdateWebExperienceOutput, error)
}
