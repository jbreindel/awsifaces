// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

// DatasyncClient ...
type DatasyncClient interface {
	Options() datasync.Options
	AddStorageSystem(ctx context.Context, params *datasync.AddStorageSystemInput, optFns ...func(*datasync.Options)) (*datasync.AddStorageSystemOutput, error)
	CancelTaskExecution(ctx context.Context, params *datasync.CancelTaskExecutionInput, optFns ...func(*datasync.Options)) (*datasync.CancelTaskExecutionOutput, error)
	CreateAgent(ctx context.Context, params *datasync.CreateAgentInput, optFns ...func(*datasync.Options)) (*datasync.CreateAgentOutput, error)
	CreateLocationAzureBlob(ctx context.Context, params *datasync.CreateLocationAzureBlobInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationAzureBlobOutput, error)
	CreateLocationEfs(ctx context.Context, params *datasync.CreateLocationEfsInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationFsxLustre(ctx context.Context, params *datasync.CreateLocationFsxLustreInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationFsxLustreOutput, error)
	CreateLocationFsxOntap(ctx context.Context, params *datasync.CreateLocationFsxOntapInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationFsxOntapOutput, error)
	CreateLocationFsxOpenZfs(ctx context.Context, params *datasync.CreateLocationFsxOpenZfsInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationFsxOpenZfsOutput, error)
	CreateLocationFsxWindows(ctx context.Context, params *datasync.CreateLocationFsxWindowsInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationHdfs(ctx context.Context, params *datasync.CreateLocationHdfsInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationHdfsOutput, error)
	CreateLocationNfs(ctx context.Context, params *datasync.CreateLocationNfsInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationObjectStorage(ctx context.Context, params *datasync.CreateLocationObjectStorageInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationS3(ctx context.Context, params *datasync.CreateLocationS3Input, optFns ...func(*datasync.Options)) (*datasync.CreateLocationS3Output, error)
	CreateLocationSmb(ctx context.Context, params *datasync.CreateLocationSmbInput, optFns ...func(*datasync.Options)) (*datasync.CreateLocationSmbOutput, error)
	CreateTask(ctx context.Context, params *datasync.CreateTaskInput, optFns ...func(*datasync.Options)) (*datasync.CreateTaskOutput, error)
	DeleteAgent(ctx context.Context, params *datasync.DeleteAgentInput, optFns ...func(*datasync.Options)) (*datasync.DeleteAgentOutput, error)
	DeleteLocation(ctx context.Context, params *datasync.DeleteLocationInput, optFns ...func(*datasync.Options)) (*datasync.DeleteLocationOutput, error)
	DeleteTask(ctx context.Context, params *datasync.DeleteTaskInput, optFns ...func(*datasync.Options)) (*datasync.DeleteTaskOutput, error)
	DescribeAgent(ctx context.Context, params *datasync.DescribeAgentInput, optFns ...func(*datasync.Options)) (*datasync.DescribeAgentOutput, error)
	DescribeDiscoveryJob(ctx context.Context, params *datasync.DescribeDiscoveryJobInput, optFns ...func(*datasync.Options)) (*datasync.DescribeDiscoveryJobOutput, error)
	DescribeLocationAzureBlob(ctx context.Context, params *datasync.DescribeLocationAzureBlobInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationAzureBlobOutput, error)
	DescribeLocationEfs(ctx context.Context, params *datasync.DescribeLocationEfsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationFsxLustre(ctx context.Context, params *datasync.DescribeLocationFsxLustreInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationFsxLustreOutput, error)
	DescribeLocationFsxOntap(ctx context.Context, params *datasync.DescribeLocationFsxOntapInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationFsxOntapOutput, error)
	DescribeLocationFsxOpenZfs(ctx context.Context, params *datasync.DescribeLocationFsxOpenZfsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationFsxOpenZfsOutput, error)
	DescribeLocationFsxWindows(ctx context.Context, params *datasync.DescribeLocationFsxWindowsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationHdfs(ctx context.Context, params *datasync.DescribeLocationHdfsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationHdfsOutput, error)
	DescribeLocationNfs(ctx context.Context, params *datasync.DescribeLocationNfsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationObjectStorage(ctx context.Context, params *datasync.DescribeLocationObjectStorageInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationS3(ctx context.Context, params *datasync.DescribeLocationS3Input, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationSmb(ctx context.Context, params *datasync.DescribeLocationSmbInput, optFns ...func(*datasync.Options)) (*datasync.DescribeLocationSmbOutput, error)
	DescribeStorageSystem(ctx context.Context, params *datasync.DescribeStorageSystemInput, optFns ...func(*datasync.Options)) (*datasync.DescribeStorageSystemOutput, error)
	DescribeStorageSystemResourceMetrics(ctx context.Context, params *datasync.DescribeStorageSystemResourceMetricsInput, optFns ...func(*datasync.Options)) (*datasync.DescribeStorageSystemResourceMetricsOutput, error)
	DescribeStorageSystemResources(ctx context.Context, params *datasync.DescribeStorageSystemResourcesInput, optFns ...func(*datasync.Options)) (*datasync.DescribeStorageSystemResourcesOutput, error)
	DescribeTask(ctx context.Context, params *datasync.DescribeTaskInput, optFns ...func(*datasync.Options)) (*datasync.DescribeTaskOutput, error)
	DescribeTaskExecution(ctx context.Context, params *datasync.DescribeTaskExecutionInput, optFns ...func(*datasync.Options)) (*datasync.DescribeTaskExecutionOutput, error)
	GenerateRecommendations(ctx context.Context, params *datasync.GenerateRecommendationsInput, optFns ...func(*datasync.Options)) (*datasync.GenerateRecommendationsOutput, error)
	ListAgents(ctx context.Context, params *datasync.ListAgentsInput, optFns ...func(*datasync.Options)) (*datasync.ListAgentsOutput, error)
	ListDiscoveryJobs(ctx context.Context, params *datasync.ListDiscoveryJobsInput, optFns ...func(*datasync.Options)) (*datasync.ListDiscoveryJobsOutput, error)
	ListLocations(ctx context.Context, params *datasync.ListLocationsInput, optFns ...func(*datasync.Options)) (*datasync.ListLocationsOutput, error)
	ListStorageSystems(ctx context.Context, params *datasync.ListStorageSystemsInput, optFns ...func(*datasync.Options)) (*datasync.ListStorageSystemsOutput, error)
	ListTagsForResource(ctx context.Context, params *datasync.ListTagsForResourceInput, optFns ...func(*datasync.Options)) (*datasync.ListTagsForResourceOutput, error)
	ListTaskExecutions(ctx context.Context, params *datasync.ListTaskExecutionsInput, optFns ...func(*datasync.Options)) (*datasync.ListTaskExecutionsOutput, error)
	ListTasks(ctx context.Context, params *datasync.ListTasksInput, optFns ...func(*datasync.Options)) (*datasync.ListTasksOutput, error)
	RemoveStorageSystem(ctx context.Context, params *datasync.RemoveStorageSystemInput, optFns ...func(*datasync.Options)) (*datasync.RemoveStorageSystemOutput, error)
	StartDiscoveryJob(ctx context.Context, params *datasync.StartDiscoveryJobInput, optFns ...func(*datasync.Options)) (*datasync.StartDiscoveryJobOutput, error)
	StartTaskExecution(ctx context.Context, params *datasync.StartTaskExecutionInput, optFns ...func(*datasync.Options)) (*datasync.StartTaskExecutionOutput, error)
	StopDiscoveryJob(ctx context.Context, params *datasync.StopDiscoveryJobInput, optFns ...func(*datasync.Options)) (*datasync.StopDiscoveryJobOutput, error)
	TagResource(ctx context.Context, params *datasync.TagResourceInput, optFns ...func(*datasync.Options)) (*datasync.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *datasync.UntagResourceInput, optFns ...func(*datasync.Options)) (*datasync.UntagResourceOutput, error)
	UpdateAgent(ctx context.Context, params *datasync.UpdateAgentInput, optFns ...func(*datasync.Options)) (*datasync.UpdateAgentOutput, error)
	UpdateDiscoveryJob(ctx context.Context, params *datasync.UpdateDiscoveryJobInput, optFns ...func(*datasync.Options)) (*datasync.UpdateDiscoveryJobOutput, error)
	UpdateLocationAzureBlob(ctx context.Context, params *datasync.UpdateLocationAzureBlobInput, optFns ...func(*datasync.Options)) (*datasync.UpdateLocationAzureBlobOutput, error)
	UpdateLocationHdfs(ctx context.Context, params *datasync.UpdateLocationHdfsInput, optFns ...func(*datasync.Options)) (*datasync.UpdateLocationHdfsOutput, error)
	UpdateLocationNfs(ctx context.Context, params *datasync.UpdateLocationNfsInput, optFns ...func(*datasync.Options)) (*datasync.UpdateLocationNfsOutput, error)
	UpdateLocationObjectStorage(ctx context.Context, params *datasync.UpdateLocationObjectStorageInput, optFns ...func(*datasync.Options)) (*datasync.UpdateLocationObjectStorageOutput, error)
	UpdateLocationSmb(ctx context.Context, params *datasync.UpdateLocationSmbInput, optFns ...func(*datasync.Options)) (*datasync.UpdateLocationSmbOutput, error)
	UpdateStorageSystem(ctx context.Context, params *datasync.UpdateStorageSystemInput, optFns ...func(*datasync.Options)) (*datasync.UpdateStorageSystemOutput, error)
	UpdateTask(ctx context.Context, params *datasync.UpdateTaskInput, optFns ...func(*datasync.Options)) (*datasync.UpdateTaskOutput, error)
	UpdateTaskExecution(ctx context.Context, params *datasync.UpdateTaskExecutionInput, optFns ...func(*datasync.Options)) (*datasync.UpdateTaskExecutionOutput, error)
}
