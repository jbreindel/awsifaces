// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

// SfnClient ...
type SfnClient interface {
	Options() sfn.Options
	CreateActivity(ctx context.Context, params *sfn.CreateActivityInput, optFns ...func(*sfn.Options)) (*sfn.CreateActivityOutput, error)
	CreateStateMachine(ctx context.Context, params *sfn.CreateStateMachineInput, optFns ...func(*sfn.Options)) (*sfn.CreateStateMachineOutput, error)
	CreateStateMachineAlias(ctx context.Context, params *sfn.CreateStateMachineAliasInput, optFns ...func(*sfn.Options)) (*sfn.CreateStateMachineAliasOutput, error)
	DeleteActivity(ctx context.Context, params *sfn.DeleteActivityInput, optFns ...func(*sfn.Options)) (*sfn.DeleteActivityOutput, error)
	DeleteStateMachine(ctx context.Context, params *sfn.DeleteStateMachineInput, optFns ...func(*sfn.Options)) (*sfn.DeleteStateMachineOutput, error)
	DeleteStateMachineAlias(ctx context.Context, params *sfn.DeleteStateMachineAliasInput, optFns ...func(*sfn.Options)) (*sfn.DeleteStateMachineAliasOutput, error)
	DeleteStateMachineVersion(ctx context.Context, params *sfn.DeleteStateMachineVersionInput, optFns ...func(*sfn.Options)) (*sfn.DeleteStateMachineVersionOutput, error)
	DescribeActivity(ctx context.Context, params *sfn.DescribeActivityInput, optFns ...func(*sfn.Options)) (*sfn.DescribeActivityOutput, error)
	DescribeExecution(ctx context.Context, params *sfn.DescribeExecutionInput, optFns ...func(*sfn.Options)) (*sfn.DescribeExecutionOutput, error)
	DescribeMapRun(ctx context.Context, params *sfn.DescribeMapRunInput, optFns ...func(*sfn.Options)) (*sfn.DescribeMapRunOutput, error)
	DescribeStateMachine(ctx context.Context, params *sfn.DescribeStateMachineInput, optFns ...func(*sfn.Options)) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineAlias(ctx context.Context, params *sfn.DescribeStateMachineAliasInput, optFns ...func(*sfn.Options)) (*sfn.DescribeStateMachineAliasOutput, error)
	DescribeStateMachineForExecution(ctx context.Context, params *sfn.DescribeStateMachineForExecutionInput, optFns ...func(*sfn.Options)) (*sfn.DescribeStateMachineForExecutionOutput, error)
	GetActivityTask(ctx context.Context, params *sfn.GetActivityTaskInput, optFns ...func(*sfn.Options)) (*sfn.GetActivityTaskOutput, error)
	GetExecutionHistory(ctx context.Context, params *sfn.GetExecutionHistoryInput, optFns ...func(*sfn.Options)) (*sfn.GetExecutionHistoryOutput, error)
	ListActivities(ctx context.Context, params *sfn.ListActivitiesInput, optFns ...func(*sfn.Options)) (*sfn.ListActivitiesOutput, error)
	ListExecutions(ctx context.Context, params *sfn.ListExecutionsInput, optFns ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error)
	ListMapRuns(ctx context.Context, params *sfn.ListMapRunsInput, optFns ...func(*sfn.Options)) (*sfn.ListMapRunsOutput, error)
	ListStateMachineAliases(ctx context.Context, params *sfn.ListStateMachineAliasesInput, optFns ...func(*sfn.Options)) (*sfn.ListStateMachineAliasesOutput, error)
	ListStateMachineVersions(ctx context.Context, params *sfn.ListStateMachineVersionsInput, optFns ...func(*sfn.Options)) (*sfn.ListStateMachineVersionsOutput, error)
	ListStateMachines(ctx context.Context, params *sfn.ListStateMachinesInput, optFns ...func(*sfn.Options)) (*sfn.ListStateMachinesOutput, error)
	ListTagsForResource(ctx context.Context, params *sfn.ListTagsForResourceInput, optFns ...func(*sfn.Options)) (*sfn.ListTagsForResourceOutput, error)
	PublishStateMachineVersion(ctx context.Context, params *sfn.PublishStateMachineVersionInput, optFns ...func(*sfn.Options)) (*sfn.PublishStateMachineVersionOutput, error)
	RedriveExecution(ctx context.Context, params *sfn.RedriveExecutionInput, optFns ...func(*sfn.Options)) (*sfn.RedriveExecutionOutput, error)
	SendTaskFailure(ctx context.Context, params *sfn.SendTaskFailureInput, optFns ...func(*sfn.Options)) (*sfn.SendTaskFailureOutput, error)
	SendTaskHeartbeat(ctx context.Context, params *sfn.SendTaskHeartbeatInput, optFns ...func(*sfn.Options)) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskSuccess(ctx context.Context, params *sfn.SendTaskSuccessInput, optFns ...func(*sfn.Options)) (*sfn.SendTaskSuccessOutput, error)
	StartExecution(ctx context.Context, params *sfn.StartExecutionInput, optFns ...func(*sfn.Options)) (*sfn.StartExecutionOutput, error)
	StartSyncExecution(ctx context.Context, params *sfn.StartSyncExecutionInput, optFns ...func(*sfn.Options)) (*sfn.StartSyncExecutionOutput, error)
	StopExecution(ctx context.Context, params *sfn.StopExecutionInput, optFns ...func(*sfn.Options)) (*sfn.StopExecutionOutput, error)
	TagResource(ctx context.Context, params *sfn.TagResourceInput, optFns ...func(*sfn.Options)) (*sfn.TagResourceOutput, error)
	TestState(ctx context.Context, params *sfn.TestStateInput, optFns ...func(*sfn.Options)) (*sfn.TestStateOutput, error)
	UntagResource(ctx context.Context, params *sfn.UntagResourceInput, optFns ...func(*sfn.Options)) (*sfn.UntagResourceOutput, error)
	UpdateMapRun(ctx context.Context, params *sfn.UpdateMapRunInput, optFns ...func(*sfn.Options)) (*sfn.UpdateMapRunOutput, error)
	UpdateStateMachine(ctx context.Context, params *sfn.UpdateStateMachineInput, optFns ...func(*sfn.Options)) (*sfn.UpdateStateMachineOutput, error)
	UpdateStateMachineAlias(ctx context.Context, params *sfn.UpdateStateMachineAliasInput, optFns ...func(*sfn.Options)) (*sfn.UpdateStateMachineAliasOutput, error)
}
