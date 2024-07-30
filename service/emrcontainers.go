// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emrcontainers"
)

// EmrcontainersClient ...
type EmrcontainersClient interface {
	Options() emrcontainers.Options
	CancelJobRun(ctx context.Context, params *emrcontainers.CancelJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CancelJobRunOutput, error)
	CreateJobTemplate(ctx context.Context, params *emrcontainers.CreateJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateJobTemplateOutput, error)
	CreateManagedEndpoint(ctx context.Context, params *emrcontainers.CreateManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateManagedEndpointOutput, error)
	CreateVirtualCluster(ctx context.Context, params *emrcontainers.CreateVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateVirtualClusterOutput, error)
	DeleteJobTemplate(ctx context.Context, params *emrcontainers.DeleteJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteJobTemplateOutput, error)
	DeleteManagedEndpoint(ctx context.Context, params *emrcontainers.DeleteManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteManagedEndpointOutput, error)
	DeleteVirtualCluster(ctx context.Context, params *emrcontainers.DeleteVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteVirtualClusterOutput, error)
	DescribeJobRun(ctx context.Context, params *emrcontainers.DescribeJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobRunOutput, error)
	DescribeJobTemplate(ctx context.Context, params *emrcontainers.DescribeJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobTemplateOutput, error)
	DescribeManagedEndpoint(ctx context.Context, params *emrcontainers.DescribeManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeManagedEndpointOutput, error)
	DescribeVirtualCluster(ctx context.Context, params *emrcontainers.DescribeVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeVirtualClusterOutput, error)
	GetManagedEndpointSessionCredentials(ctx context.Context, params *emrcontainers.GetManagedEndpointSessionCredentialsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.GetManagedEndpointSessionCredentialsOutput, error)
	ListJobRuns(ctx context.Context, params *emrcontainers.ListJobRunsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListJobRunsOutput, error)
	ListJobTemplates(ctx context.Context, params *emrcontainers.ListJobTemplatesInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListJobTemplatesOutput, error)
	ListManagedEndpoints(ctx context.Context, params *emrcontainers.ListManagedEndpointsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListManagedEndpointsOutput, error)
	ListTagsForResource(ctx context.Context, params *emrcontainers.ListTagsForResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListTagsForResourceOutput, error)
	ListVirtualClusters(ctx context.Context, params *emrcontainers.ListVirtualClustersInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListVirtualClustersOutput, error)
	StartJobRun(ctx context.Context, params *emrcontainers.StartJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.StartJobRunOutput, error)
	TagResource(ctx context.Context, params *emrcontainers.TagResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *emrcontainers.UntagResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.UntagResourceOutput, error)
}
