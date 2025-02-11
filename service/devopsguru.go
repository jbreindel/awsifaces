// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/devopsguru"
)

// DevopsguruClient ...
type DevopsguruClient interface {
	Options() devopsguru.Options
	AddNotificationChannel(ctx context.Context, params *devopsguru.AddNotificationChannelInput, optFns ...func(*devopsguru.Options)) (*devopsguru.AddNotificationChannelOutput, error)
	DeleteInsight(ctx context.Context, params *devopsguru.DeleteInsightInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DeleteInsightOutput, error)
	DescribeAccountHealth(ctx context.Context, params *devopsguru.DescribeAccountHealthInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeAccountHealthOutput, error)
	DescribeAccountOverview(ctx context.Context, params *devopsguru.DescribeAccountOverviewInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeAccountOverviewOutput, error)
	DescribeAnomaly(ctx context.Context, params *devopsguru.DescribeAnomalyInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeAnomalyOutput, error)
	DescribeEventSourcesConfig(ctx context.Context, params *devopsguru.DescribeEventSourcesConfigInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeEventSourcesConfigOutput, error)
	DescribeFeedback(ctx context.Context, params *devopsguru.DescribeFeedbackInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeFeedbackOutput, error)
	DescribeInsight(ctx context.Context, params *devopsguru.DescribeInsightInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeInsightOutput, error)
	DescribeOrganizationHealth(ctx context.Context, params *devopsguru.DescribeOrganizationHealthInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeOrganizationHealthOutput, error)
	DescribeOrganizationOverview(ctx context.Context, params *devopsguru.DescribeOrganizationOverviewInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeOrganizationOverviewOutput, error)
	DescribeOrganizationResourceCollectionHealth(ctx context.Context, params *devopsguru.DescribeOrganizationResourceCollectionHealthInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, error)
	DescribeResourceCollectionHealth(ctx context.Context, params *devopsguru.DescribeResourceCollectionHealthInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeResourceCollectionHealthOutput, error)
	DescribeServiceIntegration(ctx context.Context, params *devopsguru.DescribeServiceIntegrationInput, optFns ...func(*devopsguru.Options)) (*devopsguru.DescribeServiceIntegrationOutput, error)
	GetCostEstimation(ctx context.Context, params *devopsguru.GetCostEstimationInput, optFns ...func(*devopsguru.Options)) (*devopsguru.GetCostEstimationOutput, error)
	GetResourceCollection(ctx context.Context, params *devopsguru.GetResourceCollectionInput, optFns ...func(*devopsguru.Options)) (*devopsguru.GetResourceCollectionOutput, error)
	ListAnomaliesForInsight(ctx context.Context, params *devopsguru.ListAnomaliesForInsightInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListAnomaliesForInsightOutput, error)
	ListAnomalousLogGroups(ctx context.Context, params *devopsguru.ListAnomalousLogGroupsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListAnomalousLogGroupsOutput, error)
	ListEvents(ctx context.Context, params *devopsguru.ListEventsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListEventsOutput, error)
	ListInsights(ctx context.Context, params *devopsguru.ListInsightsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListInsightsOutput, error)
	ListMonitoredResources(ctx context.Context, params *devopsguru.ListMonitoredResourcesInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListMonitoredResourcesOutput, error)
	ListNotificationChannels(ctx context.Context, params *devopsguru.ListNotificationChannelsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListNotificationChannelsOutput, error)
	ListOrganizationInsights(ctx context.Context, params *devopsguru.ListOrganizationInsightsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListOrganizationInsightsOutput, error)
	ListRecommendations(ctx context.Context, params *devopsguru.ListRecommendationsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.ListRecommendationsOutput, error)
	PutFeedback(ctx context.Context, params *devopsguru.PutFeedbackInput, optFns ...func(*devopsguru.Options)) (*devopsguru.PutFeedbackOutput, error)
	RemoveNotificationChannel(ctx context.Context, params *devopsguru.RemoveNotificationChannelInput, optFns ...func(*devopsguru.Options)) (*devopsguru.RemoveNotificationChannelOutput, error)
	SearchInsights(ctx context.Context, params *devopsguru.SearchInsightsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.SearchInsightsOutput, error)
	SearchOrganizationInsights(ctx context.Context, params *devopsguru.SearchOrganizationInsightsInput, optFns ...func(*devopsguru.Options)) (*devopsguru.SearchOrganizationInsightsOutput, error)
	StartCostEstimation(ctx context.Context, params *devopsguru.StartCostEstimationInput, optFns ...func(*devopsguru.Options)) (*devopsguru.StartCostEstimationOutput, error)
	UpdateEventSourcesConfig(ctx context.Context, params *devopsguru.UpdateEventSourcesConfigInput, optFns ...func(*devopsguru.Options)) (*devopsguru.UpdateEventSourcesConfigOutput, error)
	UpdateResourceCollection(ctx context.Context, params *devopsguru.UpdateResourceCollectionInput, optFns ...func(*devopsguru.Options)) (*devopsguru.UpdateResourceCollectionOutput, error)
	UpdateServiceIntegration(ctx context.Context, params *devopsguru.UpdateServiceIntegrationInput, optFns ...func(*devopsguru.Options)) (*devopsguru.UpdateServiceIntegrationOutput, error)
}
