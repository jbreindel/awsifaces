// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
)

// NetworkfirewallClient ...
type NetworkfirewallClient interface {
	Options() networkfirewall.Options
	AssociateFirewallPolicy(ctx context.Context, params *networkfirewall.AssociateFirewallPolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.AssociateFirewallPolicyOutput, error)
	AssociateSubnets(ctx context.Context, params *networkfirewall.AssociateSubnetsInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.AssociateSubnetsOutput, error)
	CreateFirewall(ctx context.Context, params *networkfirewall.CreateFirewallInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.CreateFirewallOutput, error)
	CreateFirewallPolicy(ctx context.Context, params *networkfirewall.CreateFirewallPolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.CreateFirewallPolicyOutput, error)
	CreateRuleGroup(ctx context.Context, params *networkfirewall.CreateRuleGroupInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.CreateRuleGroupOutput, error)
	CreateTLSInspectionConfiguration(ctx context.Context, params *networkfirewall.CreateTLSInspectionConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.CreateTLSInspectionConfigurationOutput, error)
	DeleteFirewall(ctx context.Context, params *networkfirewall.DeleteFirewallInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DeleteFirewallOutput, error)
	DeleteFirewallPolicy(ctx context.Context, params *networkfirewall.DeleteFirewallPolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DeleteFirewallPolicyOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *networkfirewall.DeleteResourcePolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DeleteResourcePolicyOutput, error)
	DeleteRuleGroup(ctx context.Context, params *networkfirewall.DeleteRuleGroupInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DeleteRuleGroupOutput, error)
	DeleteTLSInspectionConfiguration(ctx context.Context, params *networkfirewall.DeleteTLSInspectionConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DeleteTLSInspectionConfigurationOutput, error)
	DescribeFirewall(ctx context.Context, params *networkfirewall.DescribeFirewallInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeFirewallOutput, error)
	DescribeFirewallPolicy(ctx context.Context, params *networkfirewall.DescribeFirewallPolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeFirewallPolicyOutput, error)
	DescribeLoggingConfiguration(ctx context.Context, params *networkfirewall.DescribeLoggingConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeLoggingConfigurationOutput, error)
	DescribeResourcePolicy(ctx context.Context, params *networkfirewall.DescribeResourcePolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeResourcePolicyOutput, error)
	DescribeRuleGroup(ctx context.Context, params *networkfirewall.DescribeRuleGroupInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeRuleGroupOutput, error)
	DescribeRuleGroupMetadata(ctx context.Context, params *networkfirewall.DescribeRuleGroupMetadataInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeRuleGroupMetadataOutput, error)
	DescribeTLSInspectionConfiguration(ctx context.Context, params *networkfirewall.DescribeTLSInspectionConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DescribeTLSInspectionConfigurationOutput, error)
	DisassociateSubnets(ctx context.Context, params *networkfirewall.DisassociateSubnetsInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.DisassociateSubnetsOutput, error)
	ListFirewallPolicies(ctx context.Context, params *networkfirewall.ListFirewallPoliciesInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.ListFirewallPoliciesOutput, error)
	ListFirewalls(ctx context.Context, params *networkfirewall.ListFirewallsInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.ListFirewallsOutput, error)
	ListRuleGroups(ctx context.Context, params *networkfirewall.ListRuleGroupsInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.ListRuleGroupsOutput, error)
	ListTLSInspectionConfigurations(ctx context.Context, params *networkfirewall.ListTLSInspectionConfigurationsInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.ListTLSInspectionConfigurationsOutput, error)
	ListTagsForResource(ctx context.Context, params *networkfirewall.ListTagsForResourceInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.ListTagsForResourceOutput, error)
	PutResourcePolicy(ctx context.Context, params *networkfirewall.PutResourcePolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.PutResourcePolicyOutput, error)
	TagResource(ctx context.Context, params *networkfirewall.TagResourceInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *networkfirewall.UntagResourceInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UntagResourceOutput, error)
	UpdateFirewallDeleteProtection(ctx context.Context, params *networkfirewall.UpdateFirewallDeleteProtectionInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateFirewallDeleteProtectionOutput, error)
	UpdateFirewallDescription(ctx context.Context, params *networkfirewall.UpdateFirewallDescriptionInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateFirewallDescriptionOutput, error)
	UpdateFirewallEncryptionConfiguration(ctx context.Context, params *networkfirewall.UpdateFirewallEncryptionConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateFirewallEncryptionConfigurationOutput, error)
	UpdateFirewallPolicy(ctx context.Context, params *networkfirewall.UpdateFirewallPolicyInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateFirewallPolicyOutput, error)
	UpdateFirewallPolicyChangeProtection(ctx context.Context, params *networkfirewall.UpdateFirewallPolicyChangeProtectionInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateFirewallPolicyChangeProtectionOutput, error)
	UpdateLoggingConfiguration(ctx context.Context, params *networkfirewall.UpdateLoggingConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateLoggingConfigurationOutput, error)
	UpdateRuleGroup(ctx context.Context, params *networkfirewall.UpdateRuleGroupInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateRuleGroupOutput, error)
	UpdateSubnetChangeProtection(ctx context.Context, params *networkfirewall.UpdateSubnetChangeProtectionInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateSubnetChangeProtectionOutput, error)
	UpdateTLSInspectionConfiguration(ctx context.Context, params *networkfirewall.UpdateTLSInspectionConfigurationInput, optFns ...func(*networkfirewall.Options)) (*networkfirewall.UpdateTLSInspectionConfigurationOutput, error)
}
