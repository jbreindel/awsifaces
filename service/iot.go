// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

// IotClient ...
type IotClient interface {
	Options() iot.Options
	AcceptCertificateTransfer(ctx context.Context, params *iot.AcceptCertificateTransferInput, optFns ...func(*iot.Options)) (*iot.AcceptCertificateTransferOutput, error)
	AddThingToBillingGroup(ctx context.Context, params *iot.AddThingToBillingGroupInput, optFns ...func(*iot.Options)) (*iot.AddThingToBillingGroupOutput, error)
	AddThingToThingGroup(ctx context.Context, params *iot.AddThingToThingGroupInput, optFns ...func(*iot.Options)) (*iot.AddThingToThingGroupOutput, error)
	AssociateTargetsWithJob(ctx context.Context, params *iot.AssociateTargetsWithJobInput, optFns ...func(*iot.Options)) (*iot.AssociateTargetsWithJobOutput, error)
	AttachPolicy(ctx context.Context, params *iot.AttachPolicyInput, optFns ...func(*iot.Options)) (*iot.AttachPolicyOutput, error)
	AttachPrincipalPolicy(ctx context.Context, params *iot.AttachPrincipalPolicyInput, optFns ...func(*iot.Options)) (*iot.AttachPrincipalPolicyOutput, error)
	AttachSecurityProfile(ctx context.Context, params *iot.AttachSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.AttachSecurityProfileOutput, error)
	AttachThingPrincipal(ctx context.Context, params *iot.AttachThingPrincipalInput, optFns ...func(*iot.Options)) (*iot.AttachThingPrincipalOutput, error)
	CancelAuditMitigationActionsTask(ctx context.Context, params *iot.CancelAuditMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.CancelAuditMitigationActionsTaskOutput, error)
	CancelAuditTask(ctx context.Context, params *iot.CancelAuditTaskInput, optFns ...func(*iot.Options)) (*iot.CancelAuditTaskOutput, error)
	CancelCertificateTransfer(ctx context.Context, params *iot.CancelCertificateTransferInput, optFns ...func(*iot.Options)) (*iot.CancelCertificateTransferOutput, error)
	CancelDetectMitigationActionsTask(ctx context.Context, params *iot.CancelDetectMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.CancelDetectMitigationActionsTaskOutput, error)
	CancelJob(ctx context.Context, params *iot.CancelJobInput, optFns ...func(*iot.Options)) (*iot.CancelJobOutput, error)
	CancelJobExecution(ctx context.Context, params *iot.CancelJobExecutionInput, optFns ...func(*iot.Options)) (*iot.CancelJobExecutionOutput, error)
	ClearDefaultAuthorizer(ctx context.Context, params *iot.ClearDefaultAuthorizerInput, optFns ...func(*iot.Options)) (*iot.ClearDefaultAuthorizerOutput, error)
	ConfirmTopicRuleDestination(ctx context.Context, params *iot.ConfirmTopicRuleDestinationInput, optFns ...func(*iot.Options)) (*iot.ConfirmTopicRuleDestinationOutput, error)
	CreateAuditSuppression(ctx context.Context, params *iot.CreateAuditSuppressionInput, optFns ...func(*iot.Options)) (*iot.CreateAuditSuppressionOutput, error)
	CreateAuthorizer(ctx context.Context, params *iot.CreateAuthorizerInput, optFns ...func(*iot.Options)) (*iot.CreateAuthorizerOutput, error)
	CreateBillingGroup(ctx context.Context, params *iot.CreateBillingGroupInput, optFns ...func(*iot.Options)) (*iot.CreateBillingGroupOutput, error)
	CreateCertificateFromCsr(ctx context.Context, params *iot.CreateCertificateFromCsrInput, optFns ...func(*iot.Options)) (*iot.CreateCertificateFromCsrOutput, error)
	CreateCertificateProvider(ctx context.Context, params *iot.CreateCertificateProviderInput, optFns ...func(*iot.Options)) (*iot.CreateCertificateProviderOutput, error)
	CreateCustomMetric(ctx context.Context, params *iot.CreateCustomMetricInput, optFns ...func(*iot.Options)) (*iot.CreateCustomMetricOutput, error)
	CreateDimension(ctx context.Context, params *iot.CreateDimensionInput, optFns ...func(*iot.Options)) (*iot.CreateDimensionOutput, error)
	CreateDomainConfiguration(ctx context.Context, params *iot.CreateDomainConfigurationInput, optFns ...func(*iot.Options)) (*iot.CreateDomainConfigurationOutput, error)
	CreateDynamicThingGroup(ctx context.Context, params *iot.CreateDynamicThingGroupInput, optFns ...func(*iot.Options)) (*iot.CreateDynamicThingGroupOutput, error)
	CreateFleetMetric(ctx context.Context, params *iot.CreateFleetMetricInput, optFns ...func(*iot.Options)) (*iot.CreateFleetMetricOutput, error)
	CreateJob(ctx context.Context, params *iot.CreateJobInput, optFns ...func(*iot.Options)) (*iot.CreateJobOutput, error)
	CreateJobTemplate(ctx context.Context, params *iot.CreateJobTemplateInput, optFns ...func(*iot.Options)) (*iot.CreateJobTemplateOutput, error)
	CreateKeysAndCertificate(ctx context.Context, params *iot.CreateKeysAndCertificateInput, optFns ...func(*iot.Options)) (*iot.CreateKeysAndCertificateOutput, error)
	CreateMitigationAction(ctx context.Context, params *iot.CreateMitigationActionInput, optFns ...func(*iot.Options)) (*iot.CreateMitigationActionOutput, error)
	CreateOTAUpdate(ctx context.Context, params *iot.CreateOTAUpdateInput, optFns ...func(*iot.Options)) (*iot.CreateOTAUpdateOutput, error)
	CreatePackage(ctx context.Context, params *iot.CreatePackageInput, optFns ...func(*iot.Options)) (*iot.CreatePackageOutput, error)
	CreatePackageVersion(ctx context.Context, params *iot.CreatePackageVersionInput, optFns ...func(*iot.Options)) (*iot.CreatePackageVersionOutput, error)
	CreatePolicy(ctx context.Context, params *iot.CreatePolicyInput, optFns ...func(*iot.Options)) (*iot.CreatePolicyOutput, error)
	CreatePolicyVersion(ctx context.Context, params *iot.CreatePolicyVersionInput, optFns ...func(*iot.Options)) (*iot.CreatePolicyVersionOutput, error)
	CreateProvisioningClaim(ctx context.Context, params *iot.CreateProvisioningClaimInput, optFns ...func(*iot.Options)) (*iot.CreateProvisioningClaimOutput, error)
	CreateProvisioningTemplate(ctx context.Context, params *iot.CreateProvisioningTemplateInput, optFns ...func(*iot.Options)) (*iot.CreateProvisioningTemplateOutput, error)
	CreateProvisioningTemplateVersion(ctx context.Context, params *iot.CreateProvisioningTemplateVersionInput, optFns ...func(*iot.Options)) (*iot.CreateProvisioningTemplateVersionOutput, error)
	CreateRoleAlias(ctx context.Context, params *iot.CreateRoleAliasInput, optFns ...func(*iot.Options)) (*iot.CreateRoleAliasOutput, error)
	CreateScheduledAudit(ctx context.Context, params *iot.CreateScheduledAuditInput, optFns ...func(*iot.Options)) (*iot.CreateScheduledAuditOutput, error)
	CreateSecurityProfile(ctx context.Context, params *iot.CreateSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.CreateSecurityProfileOutput, error)
	CreateStream(ctx context.Context, params *iot.CreateStreamInput, optFns ...func(*iot.Options)) (*iot.CreateStreamOutput, error)
	CreateThing(ctx context.Context, params *iot.CreateThingInput, optFns ...func(*iot.Options)) (*iot.CreateThingOutput, error)
	CreateThingGroup(ctx context.Context, params *iot.CreateThingGroupInput, optFns ...func(*iot.Options)) (*iot.CreateThingGroupOutput, error)
	CreateThingType(ctx context.Context, params *iot.CreateThingTypeInput, optFns ...func(*iot.Options)) (*iot.CreateThingTypeOutput, error)
	CreateTopicRule(ctx context.Context, params *iot.CreateTopicRuleInput, optFns ...func(*iot.Options)) (*iot.CreateTopicRuleOutput, error)
	CreateTopicRuleDestination(ctx context.Context, params *iot.CreateTopicRuleDestinationInput, optFns ...func(*iot.Options)) (*iot.CreateTopicRuleDestinationOutput, error)
	DeleteAccountAuditConfiguration(ctx context.Context, params *iot.DeleteAccountAuditConfigurationInput, optFns ...func(*iot.Options)) (*iot.DeleteAccountAuditConfigurationOutput, error)
	DeleteAuditSuppression(ctx context.Context, params *iot.DeleteAuditSuppressionInput, optFns ...func(*iot.Options)) (*iot.DeleteAuditSuppressionOutput, error)
	DeleteAuthorizer(ctx context.Context, params *iot.DeleteAuthorizerInput, optFns ...func(*iot.Options)) (*iot.DeleteAuthorizerOutput, error)
	DeleteBillingGroup(ctx context.Context, params *iot.DeleteBillingGroupInput, optFns ...func(*iot.Options)) (*iot.DeleteBillingGroupOutput, error)
	DeleteCACertificate(ctx context.Context, params *iot.DeleteCACertificateInput, optFns ...func(*iot.Options)) (*iot.DeleteCACertificateOutput, error)
	DeleteCertificate(ctx context.Context, params *iot.DeleteCertificateInput, optFns ...func(*iot.Options)) (*iot.DeleteCertificateOutput, error)
	DeleteCertificateProvider(ctx context.Context, params *iot.DeleteCertificateProviderInput, optFns ...func(*iot.Options)) (*iot.DeleteCertificateProviderOutput, error)
	DeleteCustomMetric(ctx context.Context, params *iot.DeleteCustomMetricInput, optFns ...func(*iot.Options)) (*iot.DeleteCustomMetricOutput, error)
	DeleteDimension(ctx context.Context, params *iot.DeleteDimensionInput, optFns ...func(*iot.Options)) (*iot.DeleteDimensionOutput, error)
	DeleteDomainConfiguration(ctx context.Context, params *iot.DeleteDomainConfigurationInput, optFns ...func(*iot.Options)) (*iot.DeleteDomainConfigurationOutput, error)
	DeleteDynamicThingGroup(ctx context.Context, params *iot.DeleteDynamicThingGroupInput, optFns ...func(*iot.Options)) (*iot.DeleteDynamicThingGroupOutput, error)
	DeleteFleetMetric(ctx context.Context, params *iot.DeleteFleetMetricInput, optFns ...func(*iot.Options)) (*iot.DeleteFleetMetricOutput, error)
	DeleteJob(ctx context.Context, params *iot.DeleteJobInput, optFns ...func(*iot.Options)) (*iot.DeleteJobOutput, error)
	DeleteJobExecution(ctx context.Context, params *iot.DeleteJobExecutionInput, optFns ...func(*iot.Options)) (*iot.DeleteJobExecutionOutput, error)
	DeleteJobTemplate(ctx context.Context, params *iot.DeleteJobTemplateInput, optFns ...func(*iot.Options)) (*iot.DeleteJobTemplateOutput, error)
	DeleteMitigationAction(ctx context.Context, params *iot.DeleteMitigationActionInput, optFns ...func(*iot.Options)) (*iot.DeleteMitigationActionOutput, error)
	DeleteOTAUpdate(ctx context.Context, params *iot.DeleteOTAUpdateInput, optFns ...func(*iot.Options)) (*iot.DeleteOTAUpdateOutput, error)
	DeletePackage(ctx context.Context, params *iot.DeletePackageInput, optFns ...func(*iot.Options)) (*iot.DeletePackageOutput, error)
	DeletePackageVersion(ctx context.Context, params *iot.DeletePackageVersionInput, optFns ...func(*iot.Options)) (*iot.DeletePackageVersionOutput, error)
	DeletePolicy(ctx context.Context, params *iot.DeletePolicyInput, optFns ...func(*iot.Options)) (*iot.DeletePolicyOutput, error)
	DeletePolicyVersion(ctx context.Context, params *iot.DeletePolicyVersionInput, optFns ...func(*iot.Options)) (*iot.DeletePolicyVersionOutput, error)
	DeleteProvisioningTemplate(ctx context.Context, params *iot.DeleteProvisioningTemplateInput, optFns ...func(*iot.Options)) (*iot.DeleteProvisioningTemplateOutput, error)
	DeleteProvisioningTemplateVersion(ctx context.Context, params *iot.DeleteProvisioningTemplateVersionInput, optFns ...func(*iot.Options)) (*iot.DeleteProvisioningTemplateVersionOutput, error)
	DeleteRegistrationCode(ctx context.Context, params *iot.DeleteRegistrationCodeInput, optFns ...func(*iot.Options)) (*iot.DeleteRegistrationCodeOutput, error)
	DeleteRoleAlias(ctx context.Context, params *iot.DeleteRoleAliasInput, optFns ...func(*iot.Options)) (*iot.DeleteRoleAliasOutput, error)
	DeleteScheduledAudit(ctx context.Context, params *iot.DeleteScheduledAuditInput, optFns ...func(*iot.Options)) (*iot.DeleteScheduledAuditOutput, error)
	DeleteSecurityProfile(ctx context.Context, params *iot.DeleteSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.DeleteSecurityProfileOutput, error)
	DeleteStream(ctx context.Context, params *iot.DeleteStreamInput, optFns ...func(*iot.Options)) (*iot.DeleteStreamOutput, error)
	DeleteThing(ctx context.Context, params *iot.DeleteThingInput, optFns ...func(*iot.Options)) (*iot.DeleteThingOutput, error)
	DeleteThingGroup(ctx context.Context, params *iot.DeleteThingGroupInput, optFns ...func(*iot.Options)) (*iot.DeleteThingGroupOutput, error)
	DeleteThingType(ctx context.Context, params *iot.DeleteThingTypeInput, optFns ...func(*iot.Options)) (*iot.DeleteThingTypeOutput, error)
	DeleteTopicRule(ctx context.Context, params *iot.DeleteTopicRuleInput, optFns ...func(*iot.Options)) (*iot.DeleteTopicRuleOutput, error)
	DeleteTopicRuleDestination(ctx context.Context, params *iot.DeleteTopicRuleDestinationInput, optFns ...func(*iot.Options)) (*iot.DeleteTopicRuleDestinationOutput, error)
	DeleteV2LoggingLevel(ctx context.Context, params *iot.DeleteV2LoggingLevelInput, optFns ...func(*iot.Options)) (*iot.DeleteV2LoggingLevelOutput, error)
	DeprecateThingType(ctx context.Context, params *iot.DeprecateThingTypeInput, optFns ...func(*iot.Options)) (*iot.DeprecateThingTypeOutput, error)
	DescribeAccountAuditConfiguration(ctx context.Context, params *iot.DescribeAccountAuditConfigurationInput, optFns ...func(*iot.Options)) (*iot.DescribeAccountAuditConfigurationOutput, error)
	DescribeAuditFinding(ctx context.Context, params *iot.DescribeAuditFindingInput, optFns ...func(*iot.Options)) (*iot.DescribeAuditFindingOutput, error)
	DescribeAuditMitigationActionsTask(ctx context.Context, params *iot.DescribeAuditMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.DescribeAuditMitigationActionsTaskOutput, error)
	DescribeAuditSuppression(ctx context.Context, params *iot.DescribeAuditSuppressionInput, optFns ...func(*iot.Options)) (*iot.DescribeAuditSuppressionOutput, error)
	DescribeAuditTask(ctx context.Context, params *iot.DescribeAuditTaskInput, optFns ...func(*iot.Options)) (*iot.DescribeAuditTaskOutput, error)
	DescribeAuthorizer(ctx context.Context, params *iot.DescribeAuthorizerInput, optFns ...func(*iot.Options)) (*iot.DescribeAuthorizerOutput, error)
	DescribeBillingGroup(ctx context.Context, params *iot.DescribeBillingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error)
	DescribeCACertificate(ctx context.Context, params *iot.DescribeCACertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error)
	DescribeCertificate(ctx context.Context, params *iot.DescribeCertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error)
	DescribeCertificateProvider(ctx context.Context, params *iot.DescribeCertificateProviderInput, optFns ...func(*iot.Options)) (*iot.DescribeCertificateProviderOutput, error)
	DescribeCustomMetric(ctx context.Context, params *iot.DescribeCustomMetricInput, optFns ...func(*iot.Options)) (*iot.DescribeCustomMetricOutput, error)
	DescribeDefaultAuthorizer(ctx context.Context, params *iot.DescribeDefaultAuthorizerInput, optFns ...func(*iot.Options)) (*iot.DescribeDefaultAuthorizerOutput, error)
	DescribeDetectMitigationActionsTask(ctx context.Context, params *iot.DescribeDetectMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.DescribeDetectMitigationActionsTaskOutput, error)
	DescribeDimension(ctx context.Context, params *iot.DescribeDimensionInput, optFns ...func(*iot.Options)) (*iot.DescribeDimensionOutput, error)
	DescribeDomainConfiguration(ctx context.Context, params *iot.DescribeDomainConfigurationInput, optFns ...func(*iot.Options)) (*iot.DescribeDomainConfigurationOutput, error)
	DescribeEndpoint(ctx context.Context, params *iot.DescribeEndpointInput, optFns ...func(*iot.Options)) (*iot.DescribeEndpointOutput, error)
	DescribeEventConfigurations(ctx context.Context, params *iot.DescribeEventConfigurationsInput, optFns ...func(*iot.Options)) (*iot.DescribeEventConfigurationsOutput, error)
	DescribeFleetMetric(ctx context.Context, params *iot.DescribeFleetMetricInput, optFns ...func(*iot.Options)) (*iot.DescribeFleetMetricOutput, error)
	DescribeIndex(ctx context.Context, params *iot.DescribeIndexInput, optFns ...func(*iot.Options)) (*iot.DescribeIndexOutput, error)
	DescribeJob(ctx context.Context, params *iot.DescribeJobInput, optFns ...func(*iot.Options)) (*iot.DescribeJobOutput, error)
	DescribeJobExecution(ctx context.Context, params *iot.DescribeJobExecutionInput, optFns ...func(*iot.Options)) (*iot.DescribeJobExecutionOutput, error)
	DescribeJobTemplate(ctx context.Context, params *iot.DescribeJobTemplateInput, optFns ...func(*iot.Options)) (*iot.DescribeJobTemplateOutput, error)
	DescribeManagedJobTemplate(ctx context.Context, params *iot.DescribeManagedJobTemplateInput, optFns ...func(*iot.Options)) (*iot.DescribeManagedJobTemplateOutput, error)
	DescribeMitigationAction(ctx context.Context, params *iot.DescribeMitigationActionInput, optFns ...func(*iot.Options)) (*iot.DescribeMitigationActionOutput, error)
	DescribeProvisioningTemplate(ctx context.Context, params *iot.DescribeProvisioningTemplateInput, optFns ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateOutput, error)
	DescribeProvisioningTemplateVersion(ctx context.Context, params *iot.DescribeProvisioningTemplateVersionInput, optFns ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateVersionOutput, error)
	DescribeRoleAlias(ctx context.Context, params *iot.DescribeRoleAliasInput, optFns ...func(*iot.Options)) (*iot.DescribeRoleAliasOutput, error)
	DescribeScheduledAudit(ctx context.Context, params *iot.DescribeScheduledAuditInput, optFns ...func(*iot.Options)) (*iot.DescribeScheduledAuditOutput, error)
	DescribeSecurityProfile(ctx context.Context, params *iot.DescribeSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error)
	DescribeStream(ctx context.Context, params *iot.DescribeStreamInput, optFns ...func(*iot.Options)) (*iot.DescribeStreamOutput, error)
	DescribeThing(ctx context.Context, params *iot.DescribeThingInput, optFns ...func(*iot.Options)) (*iot.DescribeThingOutput, error)
	DescribeThingGroup(ctx context.Context, params *iot.DescribeThingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error)
	DescribeThingRegistrationTask(ctx context.Context, params *iot.DescribeThingRegistrationTaskInput, optFns ...func(*iot.Options)) (*iot.DescribeThingRegistrationTaskOutput, error)
	DescribeThingType(ctx context.Context, params *iot.DescribeThingTypeInput, optFns ...func(*iot.Options)) (*iot.DescribeThingTypeOutput, error)
	DetachPolicy(ctx context.Context, params *iot.DetachPolicyInput, optFns ...func(*iot.Options)) (*iot.DetachPolicyOutput, error)
	DetachPrincipalPolicy(ctx context.Context, params *iot.DetachPrincipalPolicyInput, optFns ...func(*iot.Options)) (*iot.DetachPrincipalPolicyOutput, error)
	DetachSecurityProfile(ctx context.Context, params *iot.DetachSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.DetachSecurityProfileOutput, error)
	DetachThingPrincipal(ctx context.Context, params *iot.DetachThingPrincipalInput, optFns ...func(*iot.Options)) (*iot.DetachThingPrincipalOutput, error)
	DisableTopicRule(ctx context.Context, params *iot.DisableTopicRuleInput, optFns ...func(*iot.Options)) (*iot.DisableTopicRuleOutput, error)
	EnableTopicRule(ctx context.Context, params *iot.EnableTopicRuleInput, optFns ...func(*iot.Options)) (*iot.EnableTopicRuleOutput, error)
	GetBehaviorModelTrainingSummaries(ctx context.Context, params *iot.GetBehaviorModelTrainingSummariesInput, optFns ...func(*iot.Options)) (*iot.GetBehaviorModelTrainingSummariesOutput, error)
	GetBucketsAggregation(ctx context.Context, params *iot.GetBucketsAggregationInput, optFns ...func(*iot.Options)) (*iot.GetBucketsAggregationOutput, error)
	GetCardinality(ctx context.Context, params *iot.GetCardinalityInput, optFns ...func(*iot.Options)) (*iot.GetCardinalityOutput, error)
	GetEffectivePolicies(ctx context.Context, params *iot.GetEffectivePoliciesInput, optFns ...func(*iot.Options)) (*iot.GetEffectivePoliciesOutput, error)
	GetIndexingConfiguration(ctx context.Context, params *iot.GetIndexingConfigurationInput, optFns ...func(*iot.Options)) (*iot.GetIndexingConfigurationOutput, error)
	GetJobDocument(ctx context.Context, params *iot.GetJobDocumentInput, optFns ...func(*iot.Options)) (*iot.GetJobDocumentOutput, error)
	GetLoggingOptions(ctx context.Context, params *iot.GetLoggingOptionsInput, optFns ...func(*iot.Options)) (*iot.GetLoggingOptionsOutput, error)
	GetOTAUpdate(ctx context.Context, params *iot.GetOTAUpdateInput, optFns ...func(*iot.Options)) (*iot.GetOTAUpdateOutput, error)
	GetPackage(ctx context.Context, params *iot.GetPackageInput, optFns ...func(*iot.Options)) (*iot.GetPackageOutput, error)
	GetPackageConfiguration(ctx context.Context, params *iot.GetPackageConfigurationInput, optFns ...func(*iot.Options)) (*iot.GetPackageConfigurationOutput, error)
	GetPackageVersion(ctx context.Context, params *iot.GetPackageVersionInput, optFns ...func(*iot.Options)) (*iot.GetPackageVersionOutput, error)
	GetPercentiles(ctx context.Context, params *iot.GetPercentilesInput, optFns ...func(*iot.Options)) (*iot.GetPercentilesOutput, error)
	GetPolicy(ctx context.Context, params *iot.GetPolicyInput, optFns ...func(*iot.Options)) (*iot.GetPolicyOutput, error)
	GetPolicyVersion(ctx context.Context, params *iot.GetPolicyVersionInput, optFns ...func(*iot.Options)) (*iot.GetPolicyVersionOutput, error)
	GetRegistrationCode(ctx context.Context, params *iot.GetRegistrationCodeInput, optFns ...func(*iot.Options)) (*iot.GetRegistrationCodeOutput, error)
	GetStatistics(ctx context.Context, params *iot.GetStatisticsInput, optFns ...func(*iot.Options)) (*iot.GetStatisticsOutput, error)
	GetTopicRule(ctx context.Context, params *iot.GetTopicRuleInput, optFns ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleDestination(ctx context.Context, params *iot.GetTopicRuleDestinationInput, optFns ...func(*iot.Options)) (*iot.GetTopicRuleDestinationOutput, error)
	GetV2LoggingOptions(ctx context.Context, params *iot.GetV2LoggingOptionsInput, optFns ...func(*iot.Options)) (*iot.GetV2LoggingOptionsOutput, error)
	ListActiveViolations(ctx context.Context, params *iot.ListActiveViolationsInput, optFns ...func(*iot.Options)) (*iot.ListActiveViolationsOutput, error)
	ListAttachedPolicies(ctx context.Context, params *iot.ListAttachedPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error)
	ListAuditFindings(ctx context.Context, params *iot.ListAuditFindingsInput, optFns ...func(*iot.Options)) (*iot.ListAuditFindingsOutput, error)
	ListAuditMitigationActionsExecutions(ctx context.Context, params *iot.ListAuditMitigationActionsExecutionsInput, optFns ...func(*iot.Options)) (*iot.ListAuditMitigationActionsExecutionsOutput, error)
	ListAuditMitigationActionsTasks(ctx context.Context, params *iot.ListAuditMitigationActionsTasksInput, optFns ...func(*iot.Options)) (*iot.ListAuditMitigationActionsTasksOutput, error)
	ListAuditSuppressions(ctx context.Context, params *iot.ListAuditSuppressionsInput, optFns ...func(*iot.Options)) (*iot.ListAuditSuppressionsOutput, error)
	ListAuditTasks(ctx context.Context, params *iot.ListAuditTasksInput, optFns ...func(*iot.Options)) (*iot.ListAuditTasksOutput, error)
	ListAuthorizers(ctx context.Context, params *iot.ListAuthorizersInput, optFns ...func(*iot.Options)) (*iot.ListAuthorizersOutput, error)
	ListBillingGroups(ctx context.Context, params *iot.ListBillingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error)
	ListCACertificates(ctx context.Context, params *iot.ListCACertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error)
	ListCertificateProviders(ctx context.Context, params *iot.ListCertificateProvidersInput, optFns ...func(*iot.Options)) (*iot.ListCertificateProvidersOutput, error)
	ListCertificates(ctx context.Context, params *iot.ListCertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesOutput, error)
	ListCertificatesByCA(ctx context.Context, params *iot.ListCertificatesByCAInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error)
	ListCustomMetrics(ctx context.Context, params *iot.ListCustomMetricsInput, optFns ...func(*iot.Options)) (*iot.ListCustomMetricsOutput, error)
	ListDetectMitigationActionsExecutions(ctx context.Context, params *iot.ListDetectMitigationActionsExecutionsInput, optFns ...func(*iot.Options)) (*iot.ListDetectMitigationActionsExecutionsOutput, error)
	ListDetectMitigationActionsTasks(ctx context.Context, params *iot.ListDetectMitigationActionsTasksInput, optFns ...func(*iot.Options)) (*iot.ListDetectMitigationActionsTasksOutput, error)
	ListDimensions(ctx context.Context, params *iot.ListDimensionsInput, optFns ...func(*iot.Options)) (*iot.ListDimensionsOutput, error)
	ListDomainConfigurations(ctx context.Context, params *iot.ListDomainConfigurationsInput, optFns ...func(*iot.Options)) (*iot.ListDomainConfigurationsOutput, error)
	ListFleetMetrics(ctx context.Context, params *iot.ListFleetMetricsInput, optFns ...func(*iot.Options)) (*iot.ListFleetMetricsOutput, error)
	ListIndices(ctx context.Context, params *iot.ListIndicesInput, optFns ...func(*iot.Options)) (*iot.ListIndicesOutput, error)
	ListJobExecutionsForJob(ctx context.Context, params *iot.ListJobExecutionsForJobInput, optFns ...func(*iot.Options)) (*iot.ListJobExecutionsForJobOutput, error)
	ListJobExecutionsForThing(ctx context.Context, params *iot.ListJobExecutionsForThingInput, optFns ...func(*iot.Options)) (*iot.ListJobExecutionsForThingOutput, error)
	ListJobTemplates(ctx context.Context, params *iot.ListJobTemplatesInput, optFns ...func(*iot.Options)) (*iot.ListJobTemplatesOutput, error)
	ListJobs(ctx context.Context, params *iot.ListJobsInput, optFns ...func(*iot.Options)) (*iot.ListJobsOutput, error)
	ListManagedJobTemplates(ctx context.Context, params *iot.ListManagedJobTemplatesInput, optFns ...func(*iot.Options)) (*iot.ListManagedJobTemplatesOutput, error)
	ListMetricValues(ctx context.Context, params *iot.ListMetricValuesInput, optFns ...func(*iot.Options)) (*iot.ListMetricValuesOutput, error)
	ListMitigationActions(ctx context.Context, params *iot.ListMitigationActionsInput, optFns ...func(*iot.Options)) (*iot.ListMitigationActionsOutput, error)
	ListOTAUpdates(ctx context.Context, params *iot.ListOTAUpdatesInput, optFns ...func(*iot.Options)) (*iot.ListOTAUpdatesOutput, error)
	ListOutgoingCertificates(ctx context.Context, params *iot.ListOutgoingCertificatesInput, optFns ...func(*iot.Options)) (*iot.ListOutgoingCertificatesOutput, error)
	ListPackageVersions(ctx context.Context, params *iot.ListPackageVersionsInput, optFns ...func(*iot.Options)) (*iot.ListPackageVersionsOutput, error)
	ListPackages(ctx context.Context, params *iot.ListPackagesInput, optFns ...func(*iot.Options)) (*iot.ListPackagesOutput, error)
	ListPolicies(ctx context.Context, params *iot.ListPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListPoliciesOutput, error)
	ListPolicyPrincipals(ctx context.Context, params *iot.ListPolicyPrincipalsInput, optFns ...func(*iot.Options)) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyVersions(ctx context.Context, params *iot.ListPolicyVersionsInput, optFns ...func(*iot.Options)) (*iot.ListPolicyVersionsOutput, error)
	ListPrincipalPolicies(ctx context.Context, params *iot.ListPrincipalPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalThings(ctx context.Context, params *iot.ListPrincipalThingsInput, optFns ...func(*iot.Options)) (*iot.ListPrincipalThingsOutput, error)
	ListProvisioningTemplateVersions(ctx context.Context, params *iot.ListProvisioningTemplateVersionsInput, optFns ...func(*iot.Options)) (*iot.ListProvisioningTemplateVersionsOutput, error)
	ListProvisioningTemplates(ctx context.Context, params *iot.ListProvisioningTemplatesInput, optFns ...func(*iot.Options)) (*iot.ListProvisioningTemplatesOutput, error)
	ListRelatedResourcesForAuditFinding(ctx context.Context, params *iot.ListRelatedResourcesForAuditFindingInput, optFns ...func(*iot.Options)) (*iot.ListRelatedResourcesForAuditFindingOutput, error)
	ListRoleAliases(ctx context.Context, params *iot.ListRoleAliasesInput, optFns ...func(*iot.Options)) (*iot.ListRoleAliasesOutput, error)
	ListScheduledAudits(ctx context.Context, params *iot.ListScheduledAuditsInput, optFns ...func(*iot.Options)) (*iot.ListScheduledAuditsOutput, error)
	ListSecurityProfiles(ctx context.Context, params *iot.ListSecurityProfilesInput, optFns ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error)
	ListSecurityProfilesForTarget(ctx context.Context, params *iot.ListSecurityProfilesForTargetInput, optFns ...func(*iot.Options)) (*iot.ListSecurityProfilesForTargetOutput, error)
	ListStreams(ctx context.Context, params *iot.ListStreamsInput, optFns ...func(*iot.Options)) (*iot.ListStreamsOutput, error)
	ListTagsForResource(ctx context.Context, params *iot.ListTagsForResourceInput, optFns ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error)
	ListTargetsForPolicy(ctx context.Context, params *iot.ListTargetsForPolicyInput, optFns ...func(*iot.Options)) (*iot.ListTargetsForPolicyOutput, error)
	ListTargetsForSecurityProfile(ctx context.Context, params *iot.ListTargetsForSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error)
	ListThingGroups(ctx context.Context, params *iot.ListThingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error)
	ListThingGroupsForThing(ctx context.Context, params *iot.ListThingGroupsForThingInput, optFns ...func(*iot.Options)) (*iot.ListThingGroupsForThingOutput, error)
	ListThingPrincipals(ctx context.Context, params *iot.ListThingPrincipalsInput, optFns ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error)
	ListThingRegistrationTaskReports(ctx context.Context, params *iot.ListThingRegistrationTaskReportsInput, optFns ...func(*iot.Options)) (*iot.ListThingRegistrationTaskReportsOutput, error)
	ListThingRegistrationTasks(ctx context.Context, params *iot.ListThingRegistrationTasksInput, optFns ...func(*iot.Options)) (*iot.ListThingRegistrationTasksOutput, error)
	ListThingTypes(ctx context.Context, params *iot.ListThingTypesInput, optFns ...func(*iot.Options)) (*iot.ListThingTypesOutput, error)
	ListThings(ctx context.Context, params *iot.ListThingsInput, optFns ...func(*iot.Options)) (*iot.ListThingsOutput, error)
	ListThingsInBillingGroup(ctx context.Context, params *iot.ListThingsInBillingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error)
	ListThingsInThingGroup(ctx context.Context, params *iot.ListThingsInThingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error)
	ListTopicRuleDestinations(ctx context.Context, params *iot.ListTopicRuleDestinationsInput, optFns ...func(*iot.Options)) (*iot.ListTopicRuleDestinationsOutput, error)
	ListTopicRules(ctx context.Context, params *iot.ListTopicRulesInput, optFns ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error)
	ListV2LoggingLevels(ctx context.Context, params *iot.ListV2LoggingLevelsInput, optFns ...func(*iot.Options)) (*iot.ListV2LoggingLevelsOutput, error)
	ListViolationEvents(ctx context.Context, params *iot.ListViolationEventsInput, optFns ...func(*iot.Options)) (*iot.ListViolationEventsOutput, error)
	PutVerificationStateOnViolation(ctx context.Context, params *iot.PutVerificationStateOnViolationInput, optFns ...func(*iot.Options)) (*iot.PutVerificationStateOnViolationOutput, error)
	RegisterCACertificate(ctx context.Context, params *iot.RegisterCACertificateInput, optFns ...func(*iot.Options)) (*iot.RegisterCACertificateOutput, error)
	RegisterCertificate(ctx context.Context, params *iot.RegisterCertificateInput, optFns ...func(*iot.Options)) (*iot.RegisterCertificateOutput, error)
	RegisterCertificateWithoutCA(ctx context.Context, params *iot.RegisterCertificateWithoutCAInput, optFns ...func(*iot.Options)) (*iot.RegisterCertificateWithoutCAOutput, error)
	RegisterThing(ctx context.Context, params *iot.RegisterThingInput, optFns ...func(*iot.Options)) (*iot.RegisterThingOutput, error)
	RejectCertificateTransfer(ctx context.Context, params *iot.RejectCertificateTransferInput, optFns ...func(*iot.Options)) (*iot.RejectCertificateTransferOutput, error)
	RemoveThingFromBillingGroup(ctx context.Context, params *iot.RemoveThingFromBillingGroupInput, optFns ...func(*iot.Options)) (*iot.RemoveThingFromBillingGroupOutput, error)
	RemoveThingFromThingGroup(ctx context.Context, params *iot.RemoveThingFromThingGroupInput, optFns ...func(*iot.Options)) (*iot.RemoveThingFromThingGroupOutput, error)
	ReplaceTopicRule(ctx context.Context, params *iot.ReplaceTopicRuleInput, optFns ...func(*iot.Options)) (*iot.ReplaceTopicRuleOutput, error)
	SearchIndex(ctx context.Context, params *iot.SearchIndexInput, optFns ...func(*iot.Options)) (*iot.SearchIndexOutput, error)
	SetDefaultAuthorizer(ctx context.Context, params *iot.SetDefaultAuthorizerInput, optFns ...func(*iot.Options)) (*iot.SetDefaultAuthorizerOutput, error)
	SetDefaultPolicyVersion(ctx context.Context, params *iot.SetDefaultPolicyVersionInput, optFns ...func(*iot.Options)) (*iot.SetDefaultPolicyVersionOutput, error)
	SetLoggingOptions(ctx context.Context, params *iot.SetLoggingOptionsInput, optFns ...func(*iot.Options)) (*iot.SetLoggingOptionsOutput, error)
	SetV2LoggingLevel(ctx context.Context, params *iot.SetV2LoggingLevelInput, optFns ...func(*iot.Options)) (*iot.SetV2LoggingLevelOutput, error)
	SetV2LoggingOptions(ctx context.Context, params *iot.SetV2LoggingOptionsInput, optFns ...func(*iot.Options)) (*iot.SetV2LoggingOptionsOutput, error)
	StartAuditMitigationActionsTask(ctx context.Context, params *iot.StartAuditMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.StartAuditMitigationActionsTaskOutput, error)
	StartDetectMitigationActionsTask(ctx context.Context, params *iot.StartDetectMitigationActionsTaskInput, optFns ...func(*iot.Options)) (*iot.StartDetectMitigationActionsTaskOutput, error)
	StartOnDemandAuditTask(ctx context.Context, params *iot.StartOnDemandAuditTaskInput, optFns ...func(*iot.Options)) (*iot.StartOnDemandAuditTaskOutput, error)
	StartThingRegistrationTask(ctx context.Context, params *iot.StartThingRegistrationTaskInput, optFns ...func(*iot.Options)) (*iot.StartThingRegistrationTaskOutput, error)
	StopThingRegistrationTask(ctx context.Context, params *iot.StopThingRegistrationTaskInput, optFns ...func(*iot.Options)) (*iot.StopThingRegistrationTaskOutput, error)
	TagResource(ctx context.Context, params *iot.TagResourceInput, optFns ...func(*iot.Options)) (*iot.TagResourceOutput, error)
	TestAuthorization(ctx context.Context, params *iot.TestAuthorizationInput, optFns ...func(*iot.Options)) (*iot.TestAuthorizationOutput, error)
	TestInvokeAuthorizer(ctx context.Context, params *iot.TestInvokeAuthorizerInput, optFns ...func(*iot.Options)) (*iot.TestInvokeAuthorizerOutput, error)
	TransferCertificate(ctx context.Context, params *iot.TransferCertificateInput, optFns ...func(*iot.Options)) (*iot.TransferCertificateOutput, error)
	UntagResource(ctx context.Context, params *iot.UntagResourceInput, optFns ...func(*iot.Options)) (*iot.UntagResourceOutput, error)
	UpdateAccountAuditConfiguration(ctx context.Context, params *iot.UpdateAccountAuditConfigurationInput, optFns ...func(*iot.Options)) (*iot.UpdateAccountAuditConfigurationOutput, error)
	UpdateAuditSuppression(ctx context.Context, params *iot.UpdateAuditSuppressionInput, optFns ...func(*iot.Options)) (*iot.UpdateAuditSuppressionOutput, error)
	UpdateAuthorizer(ctx context.Context, params *iot.UpdateAuthorizerInput, optFns ...func(*iot.Options)) (*iot.UpdateAuthorizerOutput, error)
	UpdateBillingGroup(ctx context.Context, params *iot.UpdateBillingGroupInput, optFns ...func(*iot.Options)) (*iot.UpdateBillingGroupOutput, error)
	UpdateCACertificate(ctx context.Context, params *iot.UpdateCACertificateInput, optFns ...func(*iot.Options)) (*iot.UpdateCACertificateOutput, error)
	UpdateCertificate(ctx context.Context, params *iot.UpdateCertificateInput, optFns ...func(*iot.Options)) (*iot.UpdateCertificateOutput, error)
	UpdateCertificateProvider(ctx context.Context, params *iot.UpdateCertificateProviderInput, optFns ...func(*iot.Options)) (*iot.UpdateCertificateProviderOutput, error)
	UpdateCustomMetric(ctx context.Context, params *iot.UpdateCustomMetricInput, optFns ...func(*iot.Options)) (*iot.UpdateCustomMetricOutput, error)
	UpdateDimension(ctx context.Context, params *iot.UpdateDimensionInput, optFns ...func(*iot.Options)) (*iot.UpdateDimensionOutput, error)
	UpdateDomainConfiguration(ctx context.Context, params *iot.UpdateDomainConfigurationInput, optFns ...func(*iot.Options)) (*iot.UpdateDomainConfigurationOutput, error)
	UpdateDynamicThingGroup(ctx context.Context, params *iot.UpdateDynamicThingGroupInput, optFns ...func(*iot.Options)) (*iot.UpdateDynamicThingGroupOutput, error)
	UpdateEventConfigurations(ctx context.Context, params *iot.UpdateEventConfigurationsInput, optFns ...func(*iot.Options)) (*iot.UpdateEventConfigurationsOutput, error)
	UpdateFleetMetric(ctx context.Context, params *iot.UpdateFleetMetricInput, optFns ...func(*iot.Options)) (*iot.UpdateFleetMetricOutput, error)
	UpdateIndexingConfiguration(ctx context.Context, params *iot.UpdateIndexingConfigurationInput, optFns ...func(*iot.Options)) (*iot.UpdateIndexingConfigurationOutput, error)
	UpdateJob(ctx context.Context, params *iot.UpdateJobInput, optFns ...func(*iot.Options)) (*iot.UpdateJobOutput, error)
	UpdateMitigationAction(ctx context.Context, params *iot.UpdateMitigationActionInput, optFns ...func(*iot.Options)) (*iot.UpdateMitigationActionOutput, error)
	UpdatePackage(ctx context.Context, params *iot.UpdatePackageInput, optFns ...func(*iot.Options)) (*iot.UpdatePackageOutput, error)
	UpdatePackageConfiguration(ctx context.Context, params *iot.UpdatePackageConfigurationInput, optFns ...func(*iot.Options)) (*iot.UpdatePackageConfigurationOutput, error)
	UpdatePackageVersion(ctx context.Context, params *iot.UpdatePackageVersionInput, optFns ...func(*iot.Options)) (*iot.UpdatePackageVersionOutput, error)
	UpdateProvisioningTemplate(ctx context.Context, params *iot.UpdateProvisioningTemplateInput, optFns ...func(*iot.Options)) (*iot.UpdateProvisioningTemplateOutput, error)
	UpdateRoleAlias(ctx context.Context, params *iot.UpdateRoleAliasInput, optFns ...func(*iot.Options)) (*iot.UpdateRoleAliasOutput, error)
	UpdateScheduledAudit(ctx context.Context, params *iot.UpdateScheduledAuditInput, optFns ...func(*iot.Options)) (*iot.UpdateScheduledAuditOutput, error)
	UpdateSecurityProfile(ctx context.Context, params *iot.UpdateSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.UpdateSecurityProfileOutput, error)
	UpdateStream(ctx context.Context, params *iot.UpdateStreamInput, optFns ...func(*iot.Options)) (*iot.UpdateStreamOutput, error)
	UpdateThing(ctx context.Context, params *iot.UpdateThingInput, optFns ...func(*iot.Options)) (*iot.UpdateThingOutput, error)
	UpdateThingGroup(ctx context.Context, params *iot.UpdateThingGroupInput, optFns ...func(*iot.Options)) (*iot.UpdateThingGroupOutput, error)
	UpdateThingGroupsForThing(ctx context.Context, params *iot.UpdateThingGroupsForThingInput, optFns ...func(*iot.Options)) (*iot.UpdateThingGroupsForThingOutput, error)
	UpdateTopicRuleDestination(ctx context.Context, params *iot.UpdateTopicRuleDestinationInput, optFns ...func(*iot.Options)) (*iot.UpdateTopicRuleDestinationOutput, error)
	ValidateSecurityProfileBehaviors(ctx context.Context, params *iot.ValidateSecurityProfileBehaviorsInput, optFns ...func(*iot.Options)) (*iot.ValidateSecurityProfileBehaviorsOutput, error)
}
