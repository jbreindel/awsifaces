// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
)

// EmrClient ...
type EmrClient interface {
	Options() emr.Options
	AddInstanceFleet(ctx context.Context, params *emr.AddInstanceFleetInput, optFns ...func(*emr.Options)) (*emr.AddInstanceFleetOutput, error)
	AddInstanceGroups(ctx context.Context, params *emr.AddInstanceGroupsInput, optFns ...func(*emr.Options)) (*emr.AddInstanceGroupsOutput, error)
	AddJobFlowSteps(ctx context.Context, params *emr.AddJobFlowStepsInput, optFns ...func(*emr.Options)) (*emr.AddJobFlowStepsOutput, error)
	AddTags(ctx context.Context, params *emr.AddTagsInput, optFns ...func(*emr.Options)) (*emr.AddTagsOutput, error)
	CancelSteps(ctx context.Context, params *emr.CancelStepsInput, optFns ...func(*emr.Options)) (*emr.CancelStepsOutput, error)
	CreateSecurityConfiguration(ctx context.Context, params *emr.CreateSecurityConfigurationInput, optFns ...func(*emr.Options)) (*emr.CreateSecurityConfigurationOutput, error)
	CreateStudio(ctx context.Context, params *emr.CreateStudioInput, optFns ...func(*emr.Options)) (*emr.CreateStudioOutput, error)
	CreateStudioSessionMapping(ctx context.Context, params *emr.CreateStudioSessionMappingInput, optFns ...func(*emr.Options)) (*emr.CreateStudioSessionMappingOutput, error)
	DeleteSecurityConfiguration(ctx context.Context, params *emr.DeleteSecurityConfigurationInput, optFns ...func(*emr.Options)) (*emr.DeleteSecurityConfigurationOutput, error)
	DeleteStudio(ctx context.Context, params *emr.DeleteStudioInput, optFns ...func(*emr.Options)) (*emr.DeleteStudioOutput, error)
	DeleteStudioSessionMapping(ctx context.Context, params *emr.DeleteStudioSessionMappingInput, optFns ...func(*emr.Options)) (*emr.DeleteStudioSessionMappingOutput, error)
	DescribeCluster(ctx context.Context, params *emr.DescribeClusterInput, optFns ...func(*emr.Options)) (*emr.DescribeClusterOutput, error)
	DescribeJobFlows(ctx context.Context, params *emr.DescribeJobFlowsInput, optFns ...func(*emr.Options)) (*emr.DescribeJobFlowsOutput, error)
	DescribeNotebookExecution(ctx context.Context, params *emr.DescribeNotebookExecutionInput, optFns ...func(*emr.Options)) (*emr.DescribeNotebookExecutionOutput, error)
	DescribeReleaseLabel(ctx context.Context, params *emr.DescribeReleaseLabelInput, optFns ...func(*emr.Options)) (*emr.DescribeReleaseLabelOutput, error)
	DescribeSecurityConfiguration(ctx context.Context, params *emr.DescribeSecurityConfigurationInput, optFns ...func(*emr.Options)) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeStep(ctx context.Context, params *emr.DescribeStepInput, optFns ...func(*emr.Options)) (*emr.DescribeStepOutput, error)
	DescribeStudio(ctx context.Context, params *emr.DescribeStudioInput, optFns ...func(*emr.Options)) (*emr.DescribeStudioOutput, error)
	GetAutoTerminationPolicy(ctx context.Context, params *emr.GetAutoTerminationPolicyInput, optFns ...func(*emr.Options)) (*emr.GetAutoTerminationPolicyOutput, error)
	GetBlockPublicAccessConfiguration(ctx context.Context, params *emr.GetBlockPublicAccessConfigurationInput, optFns ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	GetClusterSessionCredentials(ctx context.Context, params *emr.GetClusterSessionCredentialsInput, optFns ...func(*emr.Options)) (*emr.GetClusterSessionCredentialsOutput, error)
	GetManagedScalingPolicy(ctx context.Context, params *emr.GetManagedScalingPolicyInput, optFns ...func(*emr.Options)) (*emr.GetManagedScalingPolicyOutput, error)
	GetStudioSessionMapping(ctx context.Context, params *emr.GetStudioSessionMappingInput, optFns ...func(*emr.Options)) (*emr.GetStudioSessionMappingOutput, error)
	ListBootstrapActions(ctx context.Context, params *emr.ListBootstrapActionsInput, optFns ...func(*emr.Options)) (*emr.ListBootstrapActionsOutput, error)
	ListClusters(ctx context.Context, params *emr.ListClustersInput, optFns ...func(*emr.Options)) (*emr.ListClustersOutput, error)
	ListInstanceFleets(ctx context.Context, params *emr.ListInstanceFleetsInput, optFns ...func(*emr.Options)) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceGroups(ctx context.Context, params *emr.ListInstanceGroupsInput, optFns ...func(*emr.Options)) (*emr.ListInstanceGroupsOutput, error)
	ListInstances(ctx context.Context, params *emr.ListInstancesInput, optFns ...func(*emr.Options)) (*emr.ListInstancesOutput, error)
	ListNotebookExecutions(ctx context.Context, params *emr.ListNotebookExecutionsInput, optFns ...func(*emr.Options)) (*emr.ListNotebookExecutionsOutput, error)
	ListReleaseLabels(ctx context.Context, params *emr.ListReleaseLabelsInput, optFns ...func(*emr.Options)) (*emr.ListReleaseLabelsOutput, error)
	ListSecurityConfigurations(ctx context.Context, params *emr.ListSecurityConfigurationsInput, optFns ...func(*emr.Options)) (*emr.ListSecurityConfigurationsOutput, error)
	ListSteps(ctx context.Context, params *emr.ListStepsInput, optFns ...func(*emr.Options)) (*emr.ListStepsOutput, error)
	ListStudioSessionMappings(ctx context.Context, params *emr.ListStudioSessionMappingsInput, optFns ...func(*emr.Options)) (*emr.ListStudioSessionMappingsOutput, error)
	ListStudios(ctx context.Context, params *emr.ListStudiosInput, optFns ...func(*emr.Options)) (*emr.ListStudiosOutput, error)
	ListSupportedInstanceTypes(ctx context.Context, params *emr.ListSupportedInstanceTypesInput, optFns ...func(*emr.Options)) (*emr.ListSupportedInstanceTypesOutput, error)
	ModifyCluster(ctx context.Context, params *emr.ModifyClusterInput, optFns ...func(*emr.Options)) (*emr.ModifyClusterOutput, error)
	ModifyInstanceFleet(ctx context.Context, params *emr.ModifyInstanceFleetInput, optFns ...func(*emr.Options)) (*emr.ModifyInstanceFleetOutput, error)
	ModifyInstanceGroups(ctx context.Context, params *emr.ModifyInstanceGroupsInput, optFns ...func(*emr.Options)) (*emr.ModifyInstanceGroupsOutput, error)
	PutAutoScalingPolicy(ctx context.Context, params *emr.PutAutoScalingPolicyInput, optFns ...func(*emr.Options)) (*emr.PutAutoScalingPolicyOutput, error)
	PutAutoTerminationPolicy(ctx context.Context, params *emr.PutAutoTerminationPolicyInput, optFns ...func(*emr.Options)) (*emr.PutAutoTerminationPolicyOutput, error)
	PutBlockPublicAccessConfiguration(ctx context.Context, params *emr.PutBlockPublicAccessConfigurationInput, optFns ...func(*emr.Options)) (*emr.PutBlockPublicAccessConfigurationOutput, error)
	PutManagedScalingPolicy(ctx context.Context, params *emr.PutManagedScalingPolicyInput, optFns ...func(*emr.Options)) (*emr.PutManagedScalingPolicyOutput, error)
	RemoveAutoScalingPolicy(ctx context.Context, params *emr.RemoveAutoScalingPolicyInput, optFns ...func(*emr.Options)) (*emr.RemoveAutoScalingPolicyOutput, error)
	RemoveAutoTerminationPolicy(ctx context.Context, params *emr.RemoveAutoTerminationPolicyInput, optFns ...func(*emr.Options)) (*emr.RemoveAutoTerminationPolicyOutput, error)
	RemoveManagedScalingPolicy(ctx context.Context, params *emr.RemoveManagedScalingPolicyInput, optFns ...func(*emr.Options)) (*emr.RemoveManagedScalingPolicyOutput, error)
	RemoveTags(ctx context.Context, params *emr.RemoveTagsInput, optFns ...func(*emr.Options)) (*emr.RemoveTagsOutput, error)
	RunJobFlow(ctx context.Context, params *emr.RunJobFlowInput, optFns ...func(*emr.Options)) (*emr.RunJobFlowOutput, error)
	SetKeepJobFlowAliveWhenNoSteps(ctx context.Context, params *emr.SetKeepJobFlowAliveWhenNoStepsInput, optFns ...func(*emr.Options)) (*emr.SetKeepJobFlowAliveWhenNoStepsOutput, error)
	SetTerminationProtection(ctx context.Context, params *emr.SetTerminationProtectionInput, optFns ...func(*emr.Options)) (*emr.SetTerminationProtectionOutput, error)
	SetUnhealthyNodeReplacement(ctx context.Context, params *emr.SetUnhealthyNodeReplacementInput, optFns ...func(*emr.Options)) (*emr.SetUnhealthyNodeReplacementOutput, error)
	SetVisibleToAllUsers(ctx context.Context, params *emr.SetVisibleToAllUsersInput, optFns ...func(*emr.Options)) (*emr.SetVisibleToAllUsersOutput, error)
	StartNotebookExecution(ctx context.Context, params *emr.StartNotebookExecutionInput, optFns ...func(*emr.Options)) (*emr.StartNotebookExecutionOutput, error)
	StopNotebookExecution(ctx context.Context, params *emr.StopNotebookExecutionInput, optFns ...func(*emr.Options)) (*emr.StopNotebookExecutionOutput, error)
	TerminateJobFlows(ctx context.Context, params *emr.TerminateJobFlowsInput, optFns ...func(*emr.Options)) (*emr.TerminateJobFlowsOutput, error)
	UpdateStudio(ctx context.Context, params *emr.UpdateStudioInput, optFns ...func(*emr.Options)) (*emr.UpdateStudioOutput, error)
	UpdateStudioSessionMapping(ctx context.Context, params *emr.UpdateStudioSessionMappingInput, optFns ...func(*emr.Options)) (*emr.UpdateStudioSessionMappingOutput, error)
}
