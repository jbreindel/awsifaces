// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

// AutoscalingClient ...
type AutoscalingClient interface {
	Options() autoscaling.Options
	AttachInstances(ctx context.Context, params *autoscaling.AttachInstancesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.AttachInstancesOutput, error)
	AttachLoadBalancerTargetGroups(ctx context.Context, params *autoscaling.AttachLoadBalancerTargetGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error)
	AttachLoadBalancers(ctx context.Context, params *autoscaling.AttachLoadBalancersInput, optFns ...func(*autoscaling.Options)) (*autoscaling.AttachLoadBalancersOutput, error)
	AttachTrafficSources(ctx context.Context, params *autoscaling.AttachTrafficSourcesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.AttachTrafficSourcesOutput, error)
	BatchDeleteScheduledAction(ctx context.Context, params *autoscaling.BatchDeleteScheduledActionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.BatchDeleteScheduledActionOutput, error)
	BatchPutScheduledUpdateGroupAction(ctx context.Context, params *autoscaling.BatchPutScheduledUpdateGroupActionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error)
	CancelInstanceRefresh(ctx context.Context, params *autoscaling.CancelInstanceRefreshInput, optFns ...func(*autoscaling.Options)) (*autoscaling.CancelInstanceRefreshOutput, error)
	CompleteLifecycleAction(ctx context.Context, params *autoscaling.CompleteLifecycleActionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.CompleteLifecycleActionOutput, error)
	CreateAutoScalingGroup(ctx context.Context, params *autoscaling.CreateAutoScalingGroupInput, optFns ...func(*autoscaling.Options)) (*autoscaling.CreateAutoScalingGroupOutput, error)
	CreateLaunchConfiguration(ctx context.Context, params *autoscaling.CreateLaunchConfigurationInput, optFns ...func(*autoscaling.Options)) (*autoscaling.CreateLaunchConfigurationOutput, error)
	CreateOrUpdateTags(ctx context.Context, params *autoscaling.CreateOrUpdateTagsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.CreateOrUpdateTagsOutput, error)
	DeleteAutoScalingGroup(ctx context.Context, params *autoscaling.DeleteAutoScalingGroupInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteAutoScalingGroupOutput, error)
	DeleteLaunchConfiguration(ctx context.Context, params *autoscaling.DeleteLaunchConfigurationInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteLaunchConfigurationOutput, error)
	DeleteLifecycleHook(ctx context.Context, params *autoscaling.DeleteLifecycleHookInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteLifecycleHookOutput, error)
	DeleteNotificationConfiguration(ctx context.Context, params *autoscaling.DeleteNotificationConfigurationInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteNotificationConfigurationOutput, error)
	DeletePolicy(ctx context.Context, params *autoscaling.DeletePolicyInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeletePolicyOutput, error)
	DeleteScheduledAction(ctx context.Context, params *autoscaling.DeleteScheduledActionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteScheduledActionOutput, error)
	DeleteTags(ctx context.Context, params *autoscaling.DeleteTagsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteTagsOutput, error)
	DeleteWarmPool(ctx context.Context, params *autoscaling.DeleteWarmPoolInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DeleteWarmPoolOutput, error)
	DescribeAccountLimits(ctx context.Context, params *autoscaling.DescribeAccountLimitsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAccountLimitsOutput, error)
	DescribeAdjustmentTypes(ctx context.Context, params *autoscaling.DescribeAdjustmentTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAdjustmentTypesOutput, error)
	DescribeAutoScalingGroups(ctx context.Context, params *autoscaling.DescribeAutoScalingGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribeAutoScalingInstances(ctx context.Context, params *autoscaling.DescribeAutoScalingInstancesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
	DescribeAutoScalingNotificationTypes(ctx context.Context, params *autoscaling.DescribeAutoScalingNotificationTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
	DescribeInstanceRefreshes(ctx context.Context, params *autoscaling.DescribeInstanceRefreshesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeInstanceRefreshesOutput, error)
	DescribeLaunchConfigurations(ctx context.Context, params *autoscaling.DescribeLaunchConfigurationsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeLifecycleHookTypes(ctx context.Context, params *autoscaling.DescribeLifecycleHookTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
	DescribeLifecycleHooks(ctx context.Context, params *autoscaling.DescribeLifecycleHooksInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeLoadBalancerTargetGroups(ctx context.Context, params *autoscaling.DescribeLoadBalancerTargetGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLoadBalancers(ctx context.Context, params *autoscaling.DescribeLoadBalancersInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeMetricCollectionTypes(ctx context.Context, params *autoscaling.DescribeMetricCollectionTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
	DescribeNotificationConfigurations(ctx context.Context, params *autoscaling.DescribeNotificationConfigurationsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribePolicies(ctx context.Context, params *autoscaling.DescribePoliciesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error)
	DescribeScalingActivities(ctx context.Context, params *autoscaling.DescribeScalingActivitiesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingProcessTypes(ctx context.Context, params *autoscaling.DescribeScalingProcessTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingProcessTypesOutput, error)
	DescribeScheduledActions(ctx context.Context, params *autoscaling.DescribeScheduledActionsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error)
	DescribeTags(ctx context.Context, params *autoscaling.DescribeTagsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error)
	DescribeTerminationPolicyTypes(ctx context.Context, params *autoscaling.DescribeTerminationPolicyTypesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
	DescribeTrafficSources(ctx context.Context, params *autoscaling.DescribeTrafficSourcesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeTrafficSourcesOutput, error)
	DescribeWarmPool(ctx context.Context, params *autoscaling.DescribeWarmPoolInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeWarmPoolOutput, error)
	DetachInstances(ctx context.Context, params *autoscaling.DetachInstancesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DetachInstancesOutput, error)
	DetachLoadBalancerTargetGroups(ctx context.Context, params *autoscaling.DetachLoadBalancerTargetGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error)
	DetachLoadBalancers(ctx context.Context, params *autoscaling.DetachLoadBalancersInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DetachLoadBalancersOutput, error)
	DetachTrafficSources(ctx context.Context, params *autoscaling.DetachTrafficSourcesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DetachTrafficSourcesOutput, error)
	DisableMetricsCollection(ctx context.Context, params *autoscaling.DisableMetricsCollectionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DisableMetricsCollectionOutput, error)
	EnableMetricsCollection(ctx context.Context, params *autoscaling.EnableMetricsCollectionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.EnableMetricsCollectionOutput, error)
	EnterStandby(ctx context.Context, params *autoscaling.EnterStandbyInput, optFns ...func(*autoscaling.Options)) (*autoscaling.EnterStandbyOutput, error)
	ExecutePolicy(ctx context.Context, params *autoscaling.ExecutePolicyInput, optFns ...func(*autoscaling.Options)) (*autoscaling.ExecutePolicyOutput, error)
	ExitStandby(ctx context.Context, params *autoscaling.ExitStandbyInput, optFns ...func(*autoscaling.Options)) (*autoscaling.ExitStandbyOutput, error)
	GetPredictiveScalingForecast(ctx context.Context, params *autoscaling.GetPredictiveScalingForecastInput, optFns ...func(*autoscaling.Options)) (*autoscaling.GetPredictiveScalingForecastOutput, error)
	PutLifecycleHook(ctx context.Context, params *autoscaling.PutLifecycleHookInput, optFns ...func(*autoscaling.Options)) (*autoscaling.PutLifecycleHookOutput, error)
	PutNotificationConfiguration(ctx context.Context, params *autoscaling.PutNotificationConfigurationInput, optFns ...func(*autoscaling.Options)) (*autoscaling.PutNotificationConfigurationOutput, error)
	PutScalingPolicy(ctx context.Context, params *autoscaling.PutScalingPolicyInput, optFns ...func(*autoscaling.Options)) (*autoscaling.PutScalingPolicyOutput, error)
	PutScheduledUpdateGroupAction(ctx context.Context, params *autoscaling.PutScheduledUpdateGroupActionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)
	PutWarmPool(ctx context.Context, params *autoscaling.PutWarmPoolInput, optFns ...func(*autoscaling.Options)) (*autoscaling.PutWarmPoolOutput, error)
	RecordLifecycleActionHeartbeat(ctx context.Context, params *autoscaling.RecordLifecycleActionHeartbeatInput, optFns ...func(*autoscaling.Options)) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)
	ResumeProcesses(ctx context.Context, params *autoscaling.ResumeProcessesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.ResumeProcessesOutput, error)
	RollbackInstanceRefresh(ctx context.Context, params *autoscaling.RollbackInstanceRefreshInput, optFns ...func(*autoscaling.Options)) (*autoscaling.RollbackInstanceRefreshOutput, error)
	SetDesiredCapacity(ctx context.Context, params *autoscaling.SetDesiredCapacityInput, optFns ...func(*autoscaling.Options)) (*autoscaling.SetDesiredCapacityOutput, error)
	SetInstanceHealth(ctx context.Context, params *autoscaling.SetInstanceHealthInput, optFns ...func(*autoscaling.Options)) (*autoscaling.SetInstanceHealthOutput, error)
	SetInstanceProtection(ctx context.Context, params *autoscaling.SetInstanceProtectionInput, optFns ...func(*autoscaling.Options)) (*autoscaling.SetInstanceProtectionOutput, error)
	StartInstanceRefresh(ctx context.Context, params *autoscaling.StartInstanceRefreshInput, optFns ...func(*autoscaling.Options)) (*autoscaling.StartInstanceRefreshOutput, error)
	SuspendProcesses(ctx context.Context, params *autoscaling.SuspendProcessesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.SuspendProcessesOutput, error)
	TerminateInstanceInAutoScalingGroup(ctx context.Context, params *autoscaling.TerminateInstanceInAutoScalingGroupInput, optFns ...func(*autoscaling.Options)) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
	UpdateAutoScalingGroup(ctx context.Context, params *autoscaling.UpdateAutoScalingGroupInput, optFns ...func(*autoscaling.Options)) (*autoscaling.UpdateAutoScalingGroupOutput, error)
}
