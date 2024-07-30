// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appflow"
)

// AppflowClient ...
type AppflowClient interface {
	Options() appflow.Options
	CancelFlowExecutions(ctx context.Context, params *appflow.CancelFlowExecutionsInput, optFns ...func(*appflow.Options)) (*appflow.CancelFlowExecutionsOutput, error)
	CreateConnectorProfile(ctx context.Context, params *appflow.CreateConnectorProfileInput, optFns ...func(*appflow.Options)) (*appflow.CreateConnectorProfileOutput, error)
	CreateFlow(ctx context.Context, params *appflow.CreateFlowInput, optFns ...func(*appflow.Options)) (*appflow.CreateFlowOutput, error)
	DeleteConnectorProfile(ctx context.Context, params *appflow.DeleteConnectorProfileInput, optFns ...func(*appflow.Options)) (*appflow.DeleteConnectorProfileOutput, error)
	DeleteFlow(ctx context.Context, params *appflow.DeleteFlowInput, optFns ...func(*appflow.Options)) (*appflow.DeleteFlowOutput, error)
	DescribeConnector(ctx context.Context, params *appflow.DescribeConnectorInput, optFns ...func(*appflow.Options)) (*appflow.DescribeConnectorOutput, error)
	DescribeConnectorEntity(ctx context.Context, params *appflow.DescribeConnectorEntityInput, optFns ...func(*appflow.Options)) (*appflow.DescribeConnectorEntityOutput, error)
	DescribeConnectorProfiles(ctx context.Context, params *appflow.DescribeConnectorProfilesInput, optFns ...func(*appflow.Options)) (*appflow.DescribeConnectorProfilesOutput, error)
	DescribeConnectors(ctx context.Context, params *appflow.DescribeConnectorsInput, optFns ...func(*appflow.Options)) (*appflow.DescribeConnectorsOutput, error)
	DescribeFlow(ctx context.Context, params *appflow.DescribeFlowInput, optFns ...func(*appflow.Options)) (*appflow.DescribeFlowOutput, error)
	DescribeFlowExecutionRecords(ctx context.Context, params *appflow.DescribeFlowExecutionRecordsInput, optFns ...func(*appflow.Options)) (*appflow.DescribeFlowExecutionRecordsOutput, error)
	ListConnectorEntities(ctx context.Context, params *appflow.ListConnectorEntitiesInput, optFns ...func(*appflow.Options)) (*appflow.ListConnectorEntitiesOutput, error)
	ListConnectors(ctx context.Context, params *appflow.ListConnectorsInput, optFns ...func(*appflow.Options)) (*appflow.ListConnectorsOutput, error)
	ListFlows(ctx context.Context, params *appflow.ListFlowsInput, optFns ...func(*appflow.Options)) (*appflow.ListFlowsOutput, error)
	ListTagsForResource(ctx context.Context, params *appflow.ListTagsForResourceInput, optFns ...func(*appflow.Options)) (*appflow.ListTagsForResourceOutput, error)
	RegisterConnector(ctx context.Context, params *appflow.RegisterConnectorInput, optFns ...func(*appflow.Options)) (*appflow.RegisterConnectorOutput, error)
	ResetConnectorMetadataCache(ctx context.Context, params *appflow.ResetConnectorMetadataCacheInput, optFns ...func(*appflow.Options)) (*appflow.ResetConnectorMetadataCacheOutput, error)
	StartFlow(ctx context.Context, params *appflow.StartFlowInput, optFns ...func(*appflow.Options)) (*appflow.StartFlowOutput, error)
	StopFlow(ctx context.Context, params *appflow.StopFlowInput, optFns ...func(*appflow.Options)) (*appflow.StopFlowOutput, error)
	TagResource(ctx context.Context, params *appflow.TagResourceInput, optFns ...func(*appflow.Options)) (*appflow.TagResourceOutput, error)
	UnregisterConnector(ctx context.Context, params *appflow.UnregisterConnectorInput, optFns ...func(*appflow.Options)) (*appflow.UnregisterConnectorOutput, error)
	UntagResource(ctx context.Context, params *appflow.UntagResourceInput, optFns ...func(*appflow.Options)) (*appflow.UntagResourceOutput, error)
	UpdateConnectorProfile(ctx context.Context, params *appflow.UpdateConnectorProfileInput, optFns ...func(*appflow.Options)) (*appflow.UpdateConnectorProfileOutput, error)
	UpdateConnectorRegistration(ctx context.Context, params *appflow.UpdateConnectorRegistrationInput, optFns ...func(*appflow.Options)) (*appflow.UpdateConnectorRegistrationOutput, error)
	UpdateFlow(ctx context.Context, params *appflow.UpdateFlowInput, optFns ...func(*appflow.Options)) (*appflow.UpdateFlowOutput, error)
}
