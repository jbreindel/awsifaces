// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

// CloudfrontClient ...
type CloudfrontClient interface {
	Options() cloudfront.Options
	AssociateAlias(ctx context.Context, params *cloudfront.AssociateAliasInput, optFns ...func(*cloudfront.Options)) (*cloudfront.AssociateAliasOutput, error)
	CopyDistribution(ctx context.Context, params *cloudfront.CopyDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CopyDistributionOutput, error)
	CreateCachePolicy(ctx context.Context, params *cloudfront.CreateCachePolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateCachePolicyOutput, error)
	CreateCloudFrontOriginAccessIdentity(ctx context.Context, params *cloudfront.CreateCloudFrontOriginAccessIdentityInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)
	CreateContinuousDeploymentPolicy(ctx context.Context, params *cloudfront.CreateContinuousDeploymentPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateContinuousDeploymentPolicyOutput, error)
	CreateDistribution(ctx context.Context, params *cloudfront.CreateDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateDistributionOutput, error)
	CreateDistributionWithTags(ctx context.Context, params *cloudfront.CreateDistributionWithTagsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateDistributionWithTagsOutput, error)
	CreateFieldLevelEncryptionConfig(ctx context.Context, params *cloudfront.CreateFieldLevelEncryptionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error)
	CreateFieldLevelEncryptionProfile(ctx context.Context, params *cloudfront.CreateFieldLevelEncryptionProfileInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error)
	CreateFunction(ctx context.Context, params *cloudfront.CreateFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateFunctionOutput, error)
	CreateInvalidation(ctx context.Context, params *cloudfront.CreateInvalidationInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateInvalidationOutput, error)
	CreateKeyGroup(ctx context.Context, params *cloudfront.CreateKeyGroupInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateKeyGroupOutput, error)
	CreateKeyValueStore(ctx context.Context, params *cloudfront.CreateKeyValueStoreInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateKeyValueStoreOutput, error)
	CreateMonitoringSubscription(ctx context.Context, params *cloudfront.CreateMonitoringSubscriptionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateMonitoringSubscriptionOutput, error)
	CreateOriginAccessControl(ctx context.Context, params *cloudfront.CreateOriginAccessControlInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateOriginAccessControlOutput, error)
	CreateOriginRequestPolicy(ctx context.Context, params *cloudfront.CreateOriginRequestPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateOriginRequestPolicyOutput, error)
	CreatePublicKey(ctx context.Context, params *cloudfront.CreatePublicKeyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreatePublicKeyOutput, error)
	CreateRealtimeLogConfig(ctx context.Context, params *cloudfront.CreateRealtimeLogConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateRealtimeLogConfigOutput, error)
	CreateResponseHeadersPolicy(ctx context.Context, params *cloudfront.CreateResponseHeadersPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateResponseHeadersPolicyOutput, error)
	CreateStreamingDistribution(ctx context.Context, params *cloudfront.CreateStreamingDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateStreamingDistributionOutput, error)
	CreateStreamingDistributionWithTags(ctx context.Context, params *cloudfront.CreateStreamingDistributionWithTagsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)
	DeleteCachePolicy(ctx context.Context, params *cloudfront.DeleteCachePolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteCachePolicyOutput, error)
	DeleteCloudFrontOriginAccessIdentity(ctx context.Context, params *cloudfront.DeleteCloudFrontOriginAccessIdentityInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)
	DeleteContinuousDeploymentPolicy(ctx context.Context, params *cloudfront.DeleteContinuousDeploymentPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteContinuousDeploymentPolicyOutput, error)
	DeleteDistribution(ctx context.Context, params *cloudfront.DeleteDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteDistributionOutput, error)
	DeleteFieldLevelEncryptionConfig(ctx context.Context, params *cloudfront.DeleteFieldLevelEncryptionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error)
	DeleteFieldLevelEncryptionProfile(ctx context.Context, params *cloudfront.DeleteFieldLevelEncryptionProfileInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error)
	DeleteFunction(ctx context.Context, params *cloudfront.DeleteFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteFunctionOutput, error)
	DeleteKeyGroup(ctx context.Context, params *cloudfront.DeleteKeyGroupInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteKeyGroupOutput, error)
	DeleteKeyValueStore(ctx context.Context, params *cloudfront.DeleteKeyValueStoreInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteKeyValueStoreOutput, error)
	DeleteMonitoringSubscription(ctx context.Context, params *cloudfront.DeleteMonitoringSubscriptionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteMonitoringSubscriptionOutput, error)
	DeleteOriginAccessControl(ctx context.Context, params *cloudfront.DeleteOriginAccessControlInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteOriginAccessControlOutput, error)
	DeleteOriginRequestPolicy(ctx context.Context, params *cloudfront.DeleteOriginRequestPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteOriginRequestPolicyOutput, error)
	DeletePublicKey(ctx context.Context, params *cloudfront.DeletePublicKeyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeletePublicKeyOutput, error)
	DeleteRealtimeLogConfig(ctx context.Context, params *cloudfront.DeleteRealtimeLogConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteRealtimeLogConfigOutput, error)
	DeleteResponseHeadersPolicy(ctx context.Context, params *cloudfront.DeleteResponseHeadersPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteResponseHeadersPolicyOutput, error)
	DeleteStreamingDistribution(ctx context.Context, params *cloudfront.DeleteStreamingDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DeleteStreamingDistributionOutput, error)
	DescribeFunction(ctx context.Context, params *cloudfront.DescribeFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DescribeFunctionOutput, error)
	DescribeKeyValueStore(ctx context.Context, params *cloudfront.DescribeKeyValueStoreInput, optFns ...func(*cloudfront.Options)) (*cloudfront.DescribeKeyValueStoreOutput, error)
	GetCachePolicy(ctx context.Context, params *cloudfront.GetCachePolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyConfig(ctx context.Context, params *cloudfront.GetCachePolicyConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCloudFrontOriginAccessIdentity(ctx context.Context, params *cloudfront.GetCloudFrontOriginAccessIdentityInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityConfig(ctx context.Context, params *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetContinuousDeploymentPolicy(ctx context.Context, params *cloudfront.GetContinuousDeploymentPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetContinuousDeploymentPolicyOutput, error)
	GetContinuousDeploymentPolicyConfig(ctx context.Context, params *cloudfront.GetContinuousDeploymentPolicyConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetContinuousDeploymentPolicyConfigOutput, error)
	GetDistribution(ctx context.Context, params *cloudfront.GetDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error)
	GetDistributionConfig(ctx context.Context, params *cloudfront.GetDistributionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetDistributionConfigOutput, error)
	GetFieldLevelEncryption(ctx context.Context, params *cloudfront.GetFieldLevelEncryptionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionConfig(ctx context.Context, params *cloudfront.GetFieldLevelEncryptionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionProfile(ctx context.Context, params *cloudfront.GetFieldLevelEncryptionProfileInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileConfig(ctx context.Context, params *cloudfront.GetFieldLevelEncryptionProfileConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFunction(ctx context.Context, params *cloudfront.GetFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetFunctionOutput, error)
	GetInvalidation(ctx context.Context, params *cloudfront.GetInvalidationInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetInvalidationOutput, error)
	GetKeyGroup(ctx context.Context, params *cloudfront.GetKeyGroupInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupOutput, error)
	GetKeyGroupConfig(ctx context.Context, params *cloudfront.GetKeyGroupConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupConfigOutput, error)
	GetMonitoringSubscription(ctx context.Context, params *cloudfront.GetMonitoringSubscriptionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetOriginAccessControl(ctx context.Context, params *cloudfront.GetOriginAccessControlInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlOutput, error)
	GetOriginAccessControlConfig(ctx context.Context, params *cloudfront.GetOriginAccessControlConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlConfigOutput, error)
	GetOriginRequestPolicy(ctx context.Context, params *cloudfront.GetOriginRequestPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyConfig(ctx context.Context, params *cloudfront.GetOriginRequestPolicyConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetPublicKey(ctx context.Context, params *cloudfront.GetPublicKeyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyConfig(ctx context.Context, params *cloudfront.GetPublicKeyConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetRealtimeLogConfig(ctx context.Context, params *cloudfront.GetRealtimeLogConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetResponseHeadersPolicy(ctx context.Context, params *cloudfront.GetResponseHeadersPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyOutput, error)
	GetResponseHeadersPolicyConfig(ctx context.Context, params *cloudfront.GetResponseHeadersPolicyConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error)
	GetStreamingDistribution(ctx context.Context, params *cloudfront.GetStreamingDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionConfig(ctx context.Context, params *cloudfront.GetStreamingDistributionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	ListCachePolicies(ctx context.Context, params *cloudfront.ListCachePoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error)
	ListCloudFrontOriginAccessIdentities(ctx context.Context, params *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListConflictingAliases(ctx context.Context, params *cloudfront.ListConflictingAliasesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListConflictingAliasesOutput, error)
	ListContinuousDeploymentPolicies(ctx context.Context, params *cloudfront.ListContinuousDeploymentPoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListContinuousDeploymentPoliciesOutput, error)
	ListDistributions(ctx context.Context, params *cloudfront.ListDistributionsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsByCachePolicyId(ctx context.Context, params *cloudfront.ListDistributionsByCachePolicyIdInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByKeyGroup(ctx context.Context, params *cloudfront.ListDistributionsByKeyGroupInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByKeyGroupOutput, error)
	ListDistributionsByOriginRequestPolicyId(ctx context.Context, params *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByRealtimeLogConfig(ctx context.Context, params *cloudfront.ListDistributionsByRealtimeLogConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByResponseHeadersPolicyId(ctx context.Context, params *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error)
	ListDistributionsByWebACLId(ctx context.Context, params *cloudfront.ListDistributionsByWebACLIdInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListFieldLevelEncryptionConfigs(ctx context.Context, params *cloudfront.ListFieldLevelEncryptionConfigsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionProfiles(ctx context.Context, params *cloudfront.ListFieldLevelEncryptionProfilesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFunctions(ctx context.Context, params *cloudfront.ListFunctionsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListFunctionsOutput, error)
	ListInvalidations(ctx context.Context, params *cloudfront.ListInvalidationsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListInvalidationsOutput, error)
	ListKeyGroups(ctx context.Context, params *cloudfront.ListKeyGroupsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListKeyGroupsOutput, error)
	ListKeyValueStores(ctx context.Context, params *cloudfront.ListKeyValueStoresInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListKeyValueStoresOutput, error)
	ListOriginAccessControls(ctx context.Context, params *cloudfront.ListOriginAccessControlsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListOriginAccessControlsOutput, error)
	ListOriginRequestPolicies(ctx context.Context, params *cloudfront.ListOriginRequestPoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListPublicKeys(ctx context.Context, params *cloudfront.ListPublicKeysInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListPublicKeysOutput, error)
	ListRealtimeLogConfigs(ctx context.Context, params *cloudfront.ListRealtimeLogConfigsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListResponseHeadersPolicies(ctx context.Context, params *cloudfront.ListResponseHeadersPoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListResponseHeadersPoliciesOutput, error)
	ListStreamingDistributions(ctx context.Context, params *cloudfront.ListStreamingDistributionsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListTagsForResource(ctx context.Context, params *cloudfront.ListTagsForResourceInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error)
	PublishFunction(ctx context.Context, params *cloudfront.PublishFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.PublishFunctionOutput, error)
	TagResource(ctx context.Context, params *cloudfront.TagResourceInput, optFns ...func(*cloudfront.Options)) (*cloudfront.TagResourceOutput, error)
	TestFunction(ctx context.Context, params *cloudfront.TestFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.TestFunctionOutput, error)
	UntagResource(ctx context.Context, params *cloudfront.UntagResourceInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UntagResourceOutput, error)
	UpdateCachePolicy(ctx context.Context, params *cloudfront.UpdateCachePolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateCachePolicyOutput, error)
	UpdateCloudFrontOriginAccessIdentity(ctx context.Context, params *cloudfront.UpdateCloudFrontOriginAccessIdentityInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)
	UpdateContinuousDeploymentPolicy(ctx context.Context, params *cloudfront.UpdateContinuousDeploymentPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateContinuousDeploymentPolicyOutput, error)
	UpdateDistribution(ctx context.Context, params *cloudfront.UpdateDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateDistributionOutput, error)
	UpdateDistributionWithStagingConfig(ctx context.Context, params *cloudfront.UpdateDistributionWithStagingConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateDistributionWithStagingConfigOutput, error)
	UpdateFieldLevelEncryptionConfig(ctx context.Context, params *cloudfront.UpdateFieldLevelEncryptionConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error)
	UpdateFieldLevelEncryptionProfile(ctx context.Context, params *cloudfront.UpdateFieldLevelEncryptionProfileInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error)
	UpdateFunction(ctx context.Context, params *cloudfront.UpdateFunctionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateFunctionOutput, error)
	UpdateKeyGroup(ctx context.Context, params *cloudfront.UpdateKeyGroupInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateKeyGroupOutput, error)
	UpdateKeyValueStore(ctx context.Context, params *cloudfront.UpdateKeyValueStoreInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateKeyValueStoreOutput, error)
	UpdateOriginAccessControl(ctx context.Context, params *cloudfront.UpdateOriginAccessControlInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateOriginAccessControlOutput, error)
	UpdateOriginRequestPolicy(ctx context.Context, params *cloudfront.UpdateOriginRequestPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateOriginRequestPolicyOutput, error)
	UpdatePublicKey(ctx context.Context, params *cloudfront.UpdatePublicKeyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdatePublicKeyOutput, error)
	UpdateRealtimeLogConfig(ctx context.Context, params *cloudfront.UpdateRealtimeLogConfigInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateRealtimeLogConfigOutput, error)
	UpdateResponseHeadersPolicy(ctx context.Context, params *cloudfront.UpdateResponseHeadersPolicyInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateResponseHeadersPolicyOutput, error)
	UpdateStreamingDistribution(ctx context.Context, params *cloudfront.UpdateStreamingDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.UpdateStreamingDistributionOutput, error)
}
