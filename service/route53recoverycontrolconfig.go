// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
)

// Route53recoverycontrolconfigClient ...
type Route53recoverycontrolconfigClient interface {
	Options() route53recoverycontrolconfig.Options
	CreateCluster(ctx context.Context, params *route53recoverycontrolconfig.CreateClusterInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.CreateClusterOutput, error)
	CreateControlPanel(ctx context.Context, params *route53recoverycontrolconfig.CreateControlPanelInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.CreateControlPanelOutput, error)
	CreateRoutingControl(ctx context.Context, params *route53recoverycontrolconfig.CreateRoutingControlInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.CreateRoutingControlOutput, error)
	CreateSafetyRule(ctx context.Context, params *route53recoverycontrolconfig.CreateSafetyRuleInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.CreateSafetyRuleOutput, error)
	DeleteCluster(ctx context.Context, params *route53recoverycontrolconfig.DeleteClusterInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DeleteClusterOutput, error)
	DeleteControlPanel(ctx context.Context, params *route53recoverycontrolconfig.DeleteControlPanelInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DeleteControlPanelOutput, error)
	DeleteRoutingControl(ctx context.Context, params *route53recoverycontrolconfig.DeleteRoutingControlInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DeleteRoutingControlOutput, error)
	DeleteSafetyRule(ctx context.Context, params *route53recoverycontrolconfig.DeleteSafetyRuleInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DeleteSafetyRuleOutput, error)
	DescribeCluster(ctx context.Context, params *route53recoverycontrolconfig.DescribeClusterInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeClusterOutput, error)
	DescribeControlPanel(ctx context.Context, params *route53recoverycontrolconfig.DescribeControlPanelInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeControlPanelOutput, error)
	DescribeRoutingControl(ctx context.Context, params *route53recoverycontrolconfig.DescribeRoutingControlInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeRoutingControlOutput, error)
	DescribeSafetyRule(ctx context.Context, params *route53recoverycontrolconfig.DescribeSafetyRuleInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeSafetyRuleOutput, error)
	GetResourcePolicy(ctx context.Context, params *route53recoverycontrolconfig.GetResourcePolicyInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.GetResourcePolicyOutput, error)
	ListAssociatedRoute53HealthChecks(ctx context.Context, params *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, error)
	ListClusters(ctx context.Context, params *route53recoverycontrolconfig.ListClustersInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListClustersOutput, error)
	ListControlPanels(ctx context.Context, params *route53recoverycontrolconfig.ListControlPanelsInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListControlPanelsOutput, error)
	ListRoutingControls(ctx context.Context, params *route53recoverycontrolconfig.ListRoutingControlsInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListRoutingControlsOutput, error)
	ListSafetyRules(ctx context.Context, params *route53recoverycontrolconfig.ListSafetyRulesInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListSafetyRulesOutput, error)
	ListTagsForResource(ctx context.Context, params *route53recoverycontrolconfig.ListTagsForResourceInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListTagsForResourceOutput, error)
	TagResource(ctx context.Context, params *route53recoverycontrolconfig.TagResourceInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *route53recoverycontrolconfig.UntagResourceInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.UntagResourceOutput, error)
	UpdateControlPanel(ctx context.Context, params *route53recoverycontrolconfig.UpdateControlPanelInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.UpdateControlPanelOutput, error)
	UpdateRoutingControl(ctx context.Context, params *route53recoverycontrolconfig.UpdateRoutingControlInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.UpdateRoutingControlOutput, error)
	UpdateSafetyRule(ctx context.Context, params *route53recoverycontrolconfig.UpdateSafetyRuleInput, optFns ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.UpdateSafetyRuleOutput, error)
}
