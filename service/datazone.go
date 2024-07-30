// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/datazone"
)

// DatazoneClient ...
type DatazoneClient interface {
	Options() datazone.Options
	AcceptPredictions(ctx context.Context, params *datazone.AcceptPredictionsInput, optFns ...func(*datazone.Options)) (*datazone.AcceptPredictionsOutput, error)
	AcceptSubscriptionRequest(ctx context.Context, params *datazone.AcceptSubscriptionRequestInput, optFns ...func(*datazone.Options)) (*datazone.AcceptSubscriptionRequestOutput, error)
	CancelMetadataGenerationRun(ctx context.Context, params *datazone.CancelMetadataGenerationRunInput, optFns ...func(*datazone.Options)) (*datazone.CancelMetadataGenerationRunOutput, error)
	CancelSubscription(ctx context.Context, params *datazone.CancelSubscriptionInput, optFns ...func(*datazone.Options)) (*datazone.CancelSubscriptionOutput, error)
	CreateAsset(ctx context.Context, params *datazone.CreateAssetInput, optFns ...func(*datazone.Options)) (*datazone.CreateAssetOutput, error)
	CreateAssetRevision(ctx context.Context, params *datazone.CreateAssetRevisionInput, optFns ...func(*datazone.Options)) (*datazone.CreateAssetRevisionOutput, error)
	CreateAssetType(ctx context.Context, params *datazone.CreateAssetTypeInput, optFns ...func(*datazone.Options)) (*datazone.CreateAssetTypeOutput, error)
	CreateDataSource(ctx context.Context, params *datazone.CreateDataSourceInput, optFns ...func(*datazone.Options)) (*datazone.CreateDataSourceOutput, error)
	CreateDomain(ctx context.Context, params *datazone.CreateDomainInput, optFns ...func(*datazone.Options)) (*datazone.CreateDomainOutput, error)
	CreateEnvironment(ctx context.Context, params *datazone.CreateEnvironmentInput, optFns ...func(*datazone.Options)) (*datazone.CreateEnvironmentOutput, error)
	CreateEnvironmentProfile(ctx context.Context, params *datazone.CreateEnvironmentProfileInput, optFns ...func(*datazone.Options)) (*datazone.CreateEnvironmentProfileOutput, error)
	CreateFormType(ctx context.Context, params *datazone.CreateFormTypeInput, optFns ...func(*datazone.Options)) (*datazone.CreateFormTypeOutput, error)
	CreateGlossary(ctx context.Context, params *datazone.CreateGlossaryInput, optFns ...func(*datazone.Options)) (*datazone.CreateGlossaryOutput, error)
	CreateGlossaryTerm(ctx context.Context, params *datazone.CreateGlossaryTermInput, optFns ...func(*datazone.Options)) (*datazone.CreateGlossaryTermOutput, error)
	CreateGroupProfile(ctx context.Context, params *datazone.CreateGroupProfileInput, optFns ...func(*datazone.Options)) (*datazone.CreateGroupProfileOutput, error)
	CreateListingChangeSet(ctx context.Context, params *datazone.CreateListingChangeSetInput, optFns ...func(*datazone.Options)) (*datazone.CreateListingChangeSetOutput, error)
	CreateProject(ctx context.Context, params *datazone.CreateProjectInput, optFns ...func(*datazone.Options)) (*datazone.CreateProjectOutput, error)
	CreateProjectMembership(ctx context.Context, params *datazone.CreateProjectMembershipInput, optFns ...func(*datazone.Options)) (*datazone.CreateProjectMembershipOutput, error)
	CreateSubscriptionGrant(ctx context.Context, params *datazone.CreateSubscriptionGrantInput, optFns ...func(*datazone.Options)) (*datazone.CreateSubscriptionGrantOutput, error)
	CreateSubscriptionRequest(ctx context.Context, params *datazone.CreateSubscriptionRequestInput, optFns ...func(*datazone.Options)) (*datazone.CreateSubscriptionRequestOutput, error)
	CreateSubscriptionTarget(ctx context.Context, params *datazone.CreateSubscriptionTargetInput, optFns ...func(*datazone.Options)) (*datazone.CreateSubscriptionTargetOutput, error)
	CreateUserProfile(ctx context.Context, params *datazone.CreateUserProfileInput, optFns ...func(*datazone.Options)) (*datazone.CreateUserProfileOutput, error)
	DeleteAsset(ctx context.Context, params *datazone.DeleteAssetInput, optFns ...func(*datazone.Options)) (*datazone.DeleteAssetOutput, error)
	DeleteAssetType(ctx context.Context, params *datazone.DeleteAssetTypeInput, optFns ...func(*datazone.Options)) (*datazone.DeleteAssetTypeOutput, error)
	DeleteDataSource(ctx context.Context, params *datazone.DeleteDataSourceInput, optFns ...func(*datazone.Options)) (*datazone.DeleteDataSourceOutput, error)
	DeleteDomain(ctx context.Context, params *datazone.DeleteDomainInput, optFns ...func(*datazone.Options)) (*datazone.DeleteDomainOutput, error)
	DeleteEnvironment(ctx context.Context, params *datazone.DeleteEnvironmentInput, optFns ...func(*datazone.Options)) (*datazone.DeleteEnvironmentOutput, error)
	DeleteEnvironmentBlueprintConfiguration(ctx context.Context, params *datazone.DeleteEnvironmentBlueprintConfigurationInput, optFns ...func(*datazone.Options)) (*datazone.DeleteEnvironmentBlueprintConfigurationOutput, error)
	DeleteEnvironmentProfile(ctx context.Context, params *datazone.DeleteEnvironmentProfileInput, optFns ...func(*datazone.Options)) (*datazone.DeleteEnvironmentProfileOutput, error)
	DeleteFormType(ctx context.Context, params *datazone.DeleteFormTypeInput, optFns ...func(*datazone.Options)) (*datazone.DeleteFormTypeOutput, error)
	DeleteGlossary(ctx context.Context, params *datazone.DeleteGlossaryInput, optFns ...func(*datazone.Options)) (*datazone.DeleteGlossaryOutput, error)
	DeleteGlossaryTerm(ctx context.Context, params *datazone.DeleteGlossaryTermInput, optFns ...func(*datazone.Options)) (*datazone.DeleteGlossaryTermOutput, error)
	DeleteListing(ctx context.Context, params *datazone.DeleteListingInput, optFns ...func(*datazone.Options)) (*datazone.DeleteListingOutput, error)
	DeleteProject(ctx context.Context, params *datazone.DeleteProjectInput, optFns ...func(*datazone.Options)) (*datazone.DeleteProjectOutput, error)
	DeleteProjectMembership(ctx context.Context, params *datazone.DeleteProjectMembershipInput, optFns ...func(*datazone.Options)) (*datazone.DeleteProjectMembershipOutput, error)
	DeleteSubscriptionGrant(ctx context.Context, params *datazone.DeleteSubscriptionGrantInput, optFns ...func(*datazone.Options)) (*datazone.DeleteSubscriptionGrantOutput, error)
	DeleteSubscriptionRequest(ctx context.Context, params *datazone.DeleteSubscriptionRequestInput, optFns ...func(*datazone.Options)) (*datazone.DeleteSubscriptionRequestOutput, error)
	DeleteSubscriptionTarget(ctx context.Context, params *datazone.DeleteSubscriptionTargetInput, optFns ...func(*datazone.Options)) (*datazone.DeleteSubscriptionTargetOutput, error)
	DeleteTimeSeriesDataPoints(ctx context.Context, params *datazone.DeleteTimeSeriesDataPointsInput, optFns ...func(*datazone.Options)) (*datazone.DeleteTimeSeriesDataPointsOutput, error)
	GetAsset(ctx context.Context, params *datazone.GetAssetInput, optFns ...func(*datazone.Options)) (*datazone.GetAssetOutput, error)
	GetAssetType(ctx context.Context, params *datazone.GetAssetTypeInput, optFns ...func(*datazone.Options)) (*datazone.GetAssetTypeOutput, error)
	GetDataSource(ctx context.Context, params *datazone.GetDataSourceInput, optFns ...func(*datazone.Options)) (*datazone.GetDataSourceOutput, error)
	GetDataSourceRun(ctx context.Context, params *datazone.GetDataSourceRunInput, optFns ...func(*datazone.Options)) (*datazone.GetDataSourceRunOutput, error)
	GetDomain(ctx context.Context, params *datazone.GetDomainInput, optFns ...func(*datazone.Options)) (*datazone.GetDomainOutput, error)
	GetEnvironment(ctx context.Context, params *datazone.GetEnvironmentInput, optFns ...func(*datazone.Options)) (*datazone.GetEnvironmentOutput, error)
	GetEnvironmentBlueprint(ctx context.Context, params *datazone.GetEnvironmentBlueprintInput, optFns ...func(*datazone.Options)) (*datazone.GetEnvironmentBlueprintOutput, error)
	GetEnvironmentBlueprintConfiguration(ctx context.Context, params *datazone.GetEnvironmentBlueprintConfigurationInput, optFns ...func(*datazone.Options)) (*datazone.GetEnvironmentBlueprintConfigurationOutput, error)
	GetEnvironmentProfile(ctx context.Context, params *datazone.GetEnvironmentProfileInput, optFns ...func(*datazone.Options)) (*datazone.GetEnvironmentProfileOutput, error)
	GetFormType(ctx context.Context, params *datazone.GetFormTypeInput, optFns ...func(*datazone.Options)) (*datazone.GetFormTypeOutput, error)
	GetGlossary(ctx context.Context, params *datazone.GetGlossaryInput, optFns ...func(*datazone.Options)) (*datazone.GetGlossaryOutput, error)
	GetGlossaryTerm(ctx context.Context, params *datazone.GetGlossaryTermInput, optFns ...func(*datazone.Options)) (*datazone.GetGlossaryTermOutput, error)
	GetGroupProfile(ctx context.Context, params *datazone.GetGroupProfileInput, optFns ...func(*datazone.Options)) (*datazone.GetGroupProfileOutput, error)
	GetIamPortalLoginUrl(ctx context.Context, params *datazone.GetIamPortalLoginUrlInput, optFns ...func(*datazone.Options)) (*datazone.GetIamPortalLoginUrlOutput, error)
	GetListing(ctx context.Context, params *datazone.GetListingInput, optFns ...func(*datazone.Options)) (*datazone.GetListingOutput, error)
	GetMetadataGenerationRun(ctx context.Context, params *datazone.GetMetadataGenerationRunInput, optFns ...func(*datazone.Options)) (*datazone.GetMetadataGenerationRunOutput, error)
	GetProject(ctx context.Context, params *datazone.GetProjectInput, optFns ...func(*datazone.Options)) (*datazone.GetProjectOutput, error)
	GetSubscription(ctx context.Context, params *datazone.GetSubscriptionInput, optFns ...func(*datazone.Options)) (*datazone.GetSubscriptionOutput, error)
	GetSubscriptionGrant(ctx context.Context, params *datazone.GetSubscriptionGrantInput, optFns ...func(*datazone.Options)) (*datazone.GetSubscriptionGrantOutput, error)
	GetSubscriptionRequestDetails(ctx context.Context, params *datazone.GetSubscriptionRequestDetailsInput, optFns ...func(*datazone.Options)) (*datazone.GetSubscriptionRequestDetailsOutput, error)
	GetSubscriptionTarget(ctx context.Context, params *datazone.GetSubscriptionTargetInput, optFns ...func(*datazone.Options)) (*datazone.GetSubscriptionTargetOutput, error)
	GetTimeSeriesDataPoint(ctx context.Context, params *datazone.GetTimeSeriesDataPointInput, optFns ...func(*datazone.Options)) (*datazone.GetTimeSeriesDataPointOutput, error)
	GetUserProfile(ctx context.Context, params *datazone.GetUserProfileInput, optFns ...func(*datazone.Options)) (*datazone.GetUserProfileOutput, error)
	ListAssetRevisions(ctx context.Context, params *datazone.ListAssetRevisionsInput, optFns ...func(*datazone.Options)) (*datazone.ListAssetRevisionsOutput, error)
	ListDataSourceRunActivities(ctx context.Context, params *datazone.ListDataSourceRunActivitiesInput, optFns ...func(*datazone.Options)) (*datazone.ListDataSourceRunActivitiesOutput, error)
	ListDataSourceRuns(ctx context.Context, params *datazone.ListDataSourceRunsInput, optFns ...func(*datazone.Options)) (*datazone.ListDataSourceRunsOutput, error)
	ListDataSources(ctx context.Context, params *datazone.ListDataSourcesInput, optFns ...func(*datazone.Options)) (*datazone.ListDataSourcesOutput, error)
	ListDomains(ctx context.Context, params *datazone.ListDomainsInput, optFns ...func(*datazone.Options)) (*datazone.ListDomainsOutput, error)
	ListEnvironmentBlueprintConfigurations(ctx context.Context, params *datazone.ListEnvironmentBlueprintConfigurationsInput, optFns ...func(*datazone.Options)) (*datazone.ListEnvironmentBlueprintConfigurationsOutput, error)
	ListEnvironmentBlueprints(ctx context.Context, params *datazone.ListEnvironmentBlueprintsInput, optFns ...func(*datazone.Options)) (*datazone.ListEnvironmentBlueprintsOutput, error)
	ListEnvironmentProfiles(ctx context.Context, params *datazone.ListEnvironmentProfilesInput, optFns ...func(*datazone.Options)) (*datazone.ListEnvironmentProfilesOutput, error)
	ListEnvironments(ctx context.Context, params *datazone.ListEnvironmentsInput, optFns ...func(*datazone.Options)) (*datazone.ListEnvironmentsOutput, error)
	ListMetadataGenerationRuns(ctx context.Context, params *datazone.ListMetadataGenerationRunsInput, optFns ...func(*datazone.Options)) (*datazone.ListMetadataGenerationRunsOutput, error)
	ListNotifications(ctx context.Context, params *datazone.ListNotificationsInput, optFns ...func(*datazone.Options)) (*datazone.ListNotificationsOutput, error)
	ListProjectMemberships(ctx context.Context, params *datazone.ListProjectMembershipsInput, optFns ...func(*datazone.Options)) (*datazone.ListProjectMembershipsOutput, error)
	ListProjects(ctx context.Context, params *datazone.ListProjectsInput, optFns ...func(*datazone.Options)) (*datazone.ListProjectsOutput, error)
	ListSubscriptionGrants(ctx context.Context, params *datazone.ListSubscriptionGrantsInput, optFns ...func(*datazone.Options)) (*datazone.ListSubscriptionGrantsOutput, error)
	ListSubscriptionRequests(ctx context.Context, params *datazone.ListSubscriptionRequestsInput, optFns ...func(*datazone.Options)) (*datazone.ListSubscriptionRequestsOutput, error)
	ListSubscriptionTargets(ctx context.Context, params *datazone.ListSubscriptionTargetsInput, optFns ...func(*datazone.Options)) (*datazone.ListSubscriptionTargetsOutput, error)
	ListSubscriptions(ctx context.Context, params *datazone.ListSubscriptionsInput, optFns ...func(*datazone.Options)) (*datazone.ListSubscriptionsOutput, error)
	ListTagsForResource(ctx context.Context, params *datazone.ListTagsForResourceInput, optFns ...func(*datazone.Options)) (*datazone.ListTagsForResourceOutput, error)
	ListTimeSeriesDataPoints(ctx context.Context, params *datazone.ListTimeSeriesDataPointsInput, optFns ...func(*datazone.Options)) (*datazone.ListTimeSeriesDataPointsOutput, error)
	PostTimeSeriesDataPoints(ctx context.Context, params *datazone.PostTimeSeriesDataPointsInput, optFns ...func(*datazone.Options)) (*datazone.PostTimeSeriesDataPointsOutput, error)
	PutEnvironmentBlueprintConfiguration(ctx context.Context, params *datazone.PutEnvironmentBlueprintConfigurationInput, optFns ...func(*datazone.Options)) (*datazone.PutEnvironmentBlueprintConfigurationOutput, error)
	RejectPredictions(ctx context.Context, params *datazone.RejectPredictionsInput, optFns ...func(*datazone.Options)) (*datazone.RejectPredictionsOutput, error)
	RejectSubscriptionRequest(ctx context.Context, params *datazone.RejectSubscriptionRequestInput, optFns ...func(*datazone.Options)) (*datazone.RejectSubscriptionRequestOutput, error)
	RevokeSubscription(ctx context.Context, params *datazone.RevokeSubscriptionInput, optFns ...func(*datazone.Options)) (*datazone.RevokeSubscriptionOutput, error)
	Search(ctx context.Context, params *datazone.SearchInput, optFns ...func(*datazone.Options)) (*datazone.SearchOutput, error)
	SearchGroupProfiles(ctx context.Context, params *datazone.SearchGroupProfilesInput, optFns ...func(*datazone.Options)) (*datazone.SearchGroupProfilesOutput, error)
	SearchListings(ctx context.Context, params *datazone.SearchListingsInput, optFns ...func(*datazone.Options)) (*datazone.SearchListingsOutput, error)
	SearchTypes(ctx context.Context, params *datazone.SearchTypesInput, optFns ...func(*datazone.Options)) (*datazone.SearchTypesOutput, error)
	SearchUserProfiles(ctx context.Context, params *datazone.SearchUserProfilesInput, optFns ...func(*datazone.Options)) (*datazone.SearchUserProfilesOutput, error)
	StartDataSourceRun(ctx context.Context, params *datazone.StartDataSourceRunInput, optFns ...func(*datazone.Options)) (*datazone.StartDataSourceRunOutput, error)
	StartMetadataGenerationRun(ctx context.Context, params *datazone.StartMetadataGenerationRunInput, optFns ...func(*datazone.Options)) (*datazone.StartMetadataGenerationRunOutput, error)
	TagResource(ctx context.Context, params *datazone.TagResourceInput, optFns ...func(*datazone.Options)) (*datazone.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *datazone.UntagResourceInput, optFns ...func(*datazone.Options)) (*datazone.UntagResourceOutput, error)
	UpdateDataSource(ctx context.Context, params *datazone.UpdateDataSourceInput, optFns ...func(*datazone.Options)) (*datazone.UpdateDataSourceOutput, error)
	UpdateDomain(ctx context.Context, params *datazone.UpdateDomainInput, optFns ...func(*datazone.Options)) (*datazone.UpdateDomainOutput, error)
	UpdateEnvironment(ctx context.Context, params *datazone.UpdateEnvironmentInput, optFns ...func(*datazone.Options)) (*datazone.UpdateEnvironmentOutput, error)
	UpdateEnvironmentProfile(ctx context.Context, params *datazone.UpdateEnvironmentProfileInput, optFns ...func(*datazone.Options)) (*datazone.UpdateEnvironmentProfileOutput, error)
	UpdateGlossary(ctx context.Context, params *datazone.UpdateGlossaryInput, optFns ...func(*datazone.Options)) (*datazone.UpdateGlossaryOutput, error)
	UpdateGlossaryTerm(ctx context.Context, params *datazone.UpdateGlossaryTermInput, optFns ...func(*datazone.Options)) (*datazone.UpdateGlossaryTermOutput, error)
	UpdateGroupProfile(ctx context.Context, params *datazone.UpdateGroupProfileInput, optFns ...func(*datazone.Options)) (*datazone.UpdateGroupProfileOutput, error)
	UpdateProject(ctx context.Context, params *datazone.UpdateProjectInput, optFns ...func(*datazone.Options)) (*datazone.UpdateProjectOutput, error)
	UpdateSubscriptionGrantStatus(ctx context.Context, params *datazone.UpdateSubscriptionGrantStatusInput, optFns ...func(*datazone.Options)) (*datazone.UpdateSubscriptionGrantStatusOutput, error)
	UpdateSubscriptionRequest(ctx context.Context, params *datazone.UpdateSubscriptionRequestInput, optFns ...func(*datazone.Options)) (*datazone.UpdateSubscriptionRequestOutput, error)
	UpdateSubscriptionTarget(ctx context.Context, params *datazone.UpdateSubscriptionTargetInput, optFns ...func(*datazone.Options)) (*datazone.UpdateSubscriptionTargetOutput, error)
	UpdateUserProfile(ctx context.Context, params *datazone.UpdateUserProfileInput, optFns ...func(*datazone.Options)) (*datazone.UpdateUserProfileOutput, error)
}
