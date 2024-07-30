// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

// CloudtrailClient ...
type CloudtrailClient interface {
	Options() cloudtrail.Options
	AddTags(ctx context.Context, params *cloudtrail.AddTagsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.AddTagsOutput, error)
	CancelQuery(ctx context.Context, params *cloudtrail.CancelQueryInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.CancelQueryOutput, error)
	CreateChannel(ctx context.Context, params *cloudtrail.CreateChannelInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.CreateChannelOutput, error)
	CreateEventDataStore(ctx context.Context, params *cloudtrail.CreateEventDataStoreInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.CreateEventDataStoreOutput, error)
	CreateTrail(ctx context.Context, params *cloudtrail.CreateTrailInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.CreateTrailOutput, error)
	DeleteChannel(ctx context.Context, params *cloudtrail.DeleteChannelInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DeleteChannelOutput, error)
	DeleteEventDataStore(ctx context.Context, params *cloudtrail.DeleteEventDataStoreInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DeleteEventDataStoreOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *cloudtrail.DeleteResourcePolicyInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DeleteResourcePolicyOutput, error)
	DeleteTrail(ctx context.Context, params *cloudtrail.DeleteTrailInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DeleteTrailOutput, error)
	DeregisterOrganizationDelegatedAdmin(ctx context.Context, params *cloudtrail.DeregisterOrganizationDelegatedAdminInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DeregisterOrganizationDelegatedAdminOutput, error)
	DescribeQuery(ctx context.Context, params *cloudtrail.DescribeQueryInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DescribeQueryOutput, error)
	DescribeTrails(ctx context.Context, params *cloudtrail.DescribeTrailsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error)
	DisableFederation(ctx context.Context, params *cloudtrail.DisableFederationInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DisableFederationOutput, error)
	EnableFederation(ctx context.Context, params *cloudtrail.EnableFederationInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.EnableFederationOutput, error)
	GetChannel(ctx context.Context, params *cloudtrail.GetChannelInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetChannelOutput, error)
	GetEventDataStore(ctx context.Context, params *cloudtrail.GetEventDataStoreInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetEventDataStoreOutput, error)
	GetEventSelectors(ctx context.Context, params *cloudtrail.GetEventSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error)
	GetImport(ctx context.Context, params *cloudtrail.GetImportInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetImportOutput, error)
	GetInsightSelectors(ctx context.Context, params *cloudtrail.GetInsightSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetInsightSelectorsOutput, error)
	GetQueryResults(ctx context.Context, params *cloudtrail.GetQueryResultsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetQueryResultsOutput, error)
	GetResourcePolicy(ctx context.Context, params *cloudtrail.GetResourcePolicyInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetResourcePolicyOutput, error)
	GetTrail(ctx context.Context, params *cloudtrail.GetTrailInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailOutput, error)
	GetTrailStatus(ctx context.Context, params *cloudtrail.GetTrailStatusInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error)
	ListChannels(ctx context.Context, params *cloudtrail.ListChannelsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListChannelsOutput, error)
	ListEventDataStores(ctx context.Context, params *cloudtrail.ListEventDataStoresInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListEventDataStoresOutput, error)
	ListImportFailures(ctx context.Context, params *cloudtrail.ListImportFailuresInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListImportFailuresOutput, error)
	ListImports(ctx context.Context, params *cloudtrail.ListImportsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListImportsOutput, error)
	ListInsightsMetricData(ctx context.Context, params *cloudtrail.ListInsightsMetricDataInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListInsightsMetricDataOutput, error)
	ListPublicKeys(ctx context.Context, params *cloudtrail.ListPublicKeysInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListPublicKeysOutput, error)
	ListQueries(ctx context.Context, params *cloudtrail.ListQueriesInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListQueriesOutput, error)
	ListTags(ctx context.Context, params *cloudtrail.ListTagsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error)
	ListTrails(ctx context.Context, params *cloudtrail.ListTrailsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListTrailsOutput, error)
	LookupEvents(ctx context.Context, params *cloudtrail.LookupEventsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.LookupEventsOutput, error)
	PutEventSelectors(ctx context.Context, params *cloudtrail.PutEventSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.PutEventSelectorsOutput, error)
	PutInsightSelectors(ctx context.Context, params *cloudtrail.PutInsightSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.PutInsightSelectorsOutput, error)
	PutResourcePolicy(ctx context.Context, params *cloudtrail.PutResourcePolicyInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.PutResourcePolicyOutput, error)
	RegisterOrganizationDelegatedAdmin(ctx context.Context, params *cloudtrail.RegisterOrganizationDelegatedAdminInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.RegisterOrganizationDelegatedAdminOutput, error)
	RemoveTags(ctx context.Context, params *cloudtrail.RemoveTagsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.RemoveTagsOutput, error)
	RestoreEventDataStore(ctx context.Context, params *cloudtrail.RestoreEventDataStoreInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.RestoreEventDataStoreOutput, error)
	StartEventDataStoreIngestion(ctx context.Context, params *cloudtrail.StartEventDataStoreIngestionInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StartEventDataStoreIngestionOutput, error)
	StartImport(ctx context.Context, params *cloudtrail.StartImportInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StartImportOutput, error)
	StartLogging(ctx context.Context, params *cloudtrail.StartLoggingInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StartLoggingOutput, error)
	StartQuery(ctx context.Context, params *cloudtrail.StartQueryInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StartQueryOutput, error)
	StopEventDataStoreIngestion(ctx context.Context, params *cloudtrail.StopEventDataStoreIngestionInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StopEventDataStoreIngestionOutput, error)
	StopImport(ctx context.Context, params *cloudtrail.StopImportInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StopImportOutput, error)
	StopLogging(ctx context.Context, params *cloudtrail.StopLoggingInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.StopLoggingOutput, error)
	UpdateChannel(ctx context.Context, params *cloudtrail.UpdateChannelInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.UpdateChannelOutput, error)
	UpdateEventDataStore(ctx context.Context, params *cloudtrail.UpdateEventDataStoreInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.UpdateEventDataStoreOutput, error)
	UpdateTrail(ctx context.Context, params *cloudtrail.UpdateTrailInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.UpdateTrailOutput, error)
}
