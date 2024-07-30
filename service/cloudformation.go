// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// CloudformationClient ...
type CloudformationClient interface {
	Options() cloudformation.Options
	ActivateOrganizationsAccess(ctx context.Context, params *cloudformation.ActivateOrganizationsAccessInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ActivateOrganizationsAccessOutput, error)
	ActivateType(ctx context.Context, params *cloudformation.ActivateTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ActivateTypeOutput, error)
	BatchDescribeTypeConfigurations(ctx context.Context, params *cloudformation.BatchDescribeTypeConfigurationsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.BatchDescribeTypeConfigurationsOutput, error)
	CancelUpdateStack(ctx context.Context, params *cloudformation.CancelUpdateStackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CancelUpdateStackOutput, error)
	ContinueUpdateRollback(ctx context.Context, params *cloudformation.ContinueUpdateRollbackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ContinueUpdateRollbackOutput, error)
	CreateChangeSet(ctx context.Context, params *cloudformation.CreateChangeSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CreateChangeSetOutput, error)
	CreateGeneratedTemplate(ctx context.Context, params *cloudformation.CreateGeneratedTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CreateGeneratedTemplateOutput, error)
	CreateStack(ctx context.Context, params *cloudformation.CreateStackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CreateStackOutput, error)
	CreateStackInstances(ctx context.Context, params *cloudformation.CreateStackInstancesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackSet(ctx context.Context, params *cloudformation.CreateStackSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.CreateStackSetOutput, error)
	DeactivateOrganizationsAccess(ctx context.Context, params *cloudformation.DeactivateOrganizationsAccessInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeactivateOrganizationsAccessOutput, error)
	DeactivateType(ctx context.Context, params *cloudformation.DeactivateTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeactivateTypeOutput, error)
	DeleteChangeSet(ctx context.Context, params *cloudformation.DeleteChangeSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteGeneratedTemplate(ctx context.Context, params *cloudformation.DeleteGeneratedTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeleteGeneratedTemplateOutput, error)
	DeleteStack(ctx context.Context, params *cloudformation.DeleteStackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeleteStackOutput, error)
	DeleteStackInstances(ctx context.Context, params *cloudformation.DeleteStackInstancesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackSet(ctx context.Context, params *cloudformation.DeleteStackSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeleteStackSetOutput, error)
	DeregisterType(ctx context.Context, params *cloudformation.DeregisterTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DeregisterTypeOutput, error)
	DescribeAccountLimits(ctx context.Context, params *cloudformation.DescribeAccountLimitsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeChangeSet(ctx context.Context, params *cloudformation.DescribeChangeSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetHooks(ctx context.Context, params *cloudformation.DescribeChangeSetHooksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetHooksOutput, error)
	DescribeGeneratedTemplate(ctx context.Context, params *cloudformation.DescribeGeneratedTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeGeneratedTemplateOutput, error)
	DescribeOrganizationsAccess(ctx context.Context, params *cloudformation.DescribeOrganizationsAccessInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeOrganizationsAccessOutput, error)
	DescribePublisher(ctx context.Context, params *cloudformation.DescribePublisherInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribePublisherOutput, error)
	DescribeResourceScan(ctx context.Context, params *cloudformation.DescribeResourceScanInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeResourceScanOutput, error)
	DescribeStackDriftDetectionStatus(ctx context.Context, params *cloudformation.DescribeStackDriftDetectionStatusInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackEvents(ctx context.Context, params *cloudformation.DescribeStackEventsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackInstance(ctx context.Context, params *cloudformation.DescribeStackInstanceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackResource(ctx context.Context, params *cloudformation.DescribeStackResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceDrifts(ctx context.Context, params *cloudformation.DescribeStackResourceDriftsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResources(ctx context.Context, params *cloudformation.DescribeStackResourcesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackSet(ctx context.Context, params *cloudformation.DescribeStackSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetOperation(ctx context.Context, params *cloudformation.DescribeStackSetOperationInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStacks(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)
	DescribeType(ctx context.Context, params *cloudformation.DescribeTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeRegistration(ctx context.Context, params *cloudformation.DescribeTypeRegistrationInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeRegistrationOutput, error)
	DetectStackDrift(ctx context.Context, params *cloudformation.DetectStackDriftInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DetectStackDriftOutput, error)
	DetectStackResourceDrift(ctx context.Context, params *cloudformation.DetectStackResourceDriftInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DetectStackResourceDriftOutput, error)
	DetectStackSetDrift(ctx context.Context, params *cloudformation.DetectStackSetDriftInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DetectStackSetDriftOutput, error)
	EstimateTemplateCost(ctx context.Context, params *cloudformation.EstimateTemplateCostInput, optFns ...func(*cloudformation.Options)) (*cloudformation.EstimateTemplateCostOutput, error)
	ExecuteChangeSet(ctx context.Context, params *cloudformation.ExecuteChangeSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ExecuteChangeSetOutput, error)
	GetGeneratedTemplate(ctx context.Context, params *cloudformation.GetGeneratedTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetGeneratedTemplateOutput, error)
	GetStackPolicy(ctx context.Context, params *cloudformation.GetStackPolicyInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetStackPolicyOutput, error)
	GetTemplate(ctx context.Context, params *cloudformation.GetTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error)
	GetTemplateSummary(ctx context.Context, params *cloudformation.GetTemplateSummaryInput, optFns ...func(*cloudformation.Options)) (*cloudformation.GetTemplateSummaryOutput, error)
	ImportStacksToStackSet(ctx context.Context, params *cloudformation.ImportStacksToStackSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ImportStacksToStackSetOutput, error)
	ListChangeSets(ctx context.Context, params *cloudformation.ListChangeSetsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListChangeSetsOutput, error)
	ListExports(ctx context.Context, params *cloudformation.ListExportsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListExportsOutput, error)
	ListGeneratedTemplates(ctx context.Context, params *cloudformation.ListGeneratedTemplatesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListGeneratedTemplatesOutput, error)
	ListImports(ctx context.Context, params *cloudformation.ListImportsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListImportsOutput, error)
	ListResourceScanRelatedResources(ctx context.Context, params *cloudformation.ListResourceScanRelatedResourcesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListResourceScanRelatedResourcesOutput, error)
	ListResourceScanResources(ctx context.Context, params *cloudformation.ListResourceScanResourcesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListResourceScanResourcesOutput, error)
	ListResourceScans(ctx context.Context, params *cloudformation.ListResourceScansInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListResourceScansOutput, error)
	ListStackInstanceResourceDrifts(ctx context.Context, params *cloudformation.ListStackInstanceResourceDriftsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackInstanceResourceDriftsOutput, error)
	ListStackInstances(ctx context.Context, params *cloudformation.ListStackInstancesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackInstancesOutput, error)
	ListStackResources(ctx context.Context, params *cloudformation.ListStackResourcesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error)
	ListStackSetAutoDeploymentTargets(ctx context.Context, params *cloudformation.ListStackSetAutoDeploymentTargetsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackSetAutoDeploymentTargetsOutput, error)
	ListStackSetOperationResults(ctx context.Context, params *cloudformation.ListStackSetOperationResultsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperations(ctx context.Context, params *cloudformation.ListStackSetOperationsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSets(ctx context.Context, params *cloudformation.ListStackSetsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStackSetsOutput, error)
	ListStacks(ctx context.Context, params *cloudformation.ListStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error)
	ListTypeRegistrations(ctx context.Context, params *cloudformation.ListTypeRegistrationsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeVersions(ctx context.Context, params *cloudformation.ListTypeVersionsInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypes(ctx context.Context, params *cloudformation.ListTypesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ListTypesOutput, error)
	PublishType(ctx context.Context, params *cloudformation.PublishTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.PublishTypeOutput, error)
	RecordHandlerProgress(ctx context.Context, params *cloudformation.RecordHandlerProgressInput, optFns ...func(*cloudformation.Options)) (*cloudformation.RecordHandlerProgressOutput, error)
	RegisterPublisher(ctx context.Context, params *cloudformation.RegisterPublisherInput, optFns ...func(*cloudformation.Options)) (*cloudformation.RegisterPublisherOutput, error)
	RegisterType(ctx context.Context, params *cloudformation.RegisterTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.RegisterTypeOutput, error)
	RollbackStack(ctx context.Context, params *cloudformation.RollbackStackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.RollbackStackOutput, error)
	SetStackPolicy(ctx context.Context, params *cloudformation.SetStackPolicyInput, optFns ...func(*cloudformation.Options)) (*cloudformation.SetStackPolicyOutput, error)
	SetTypeConfiguration(ctx context.Context, params *cloudformation.SetTypeConfigurationInput, optFns ...func(*cloudformation.Options)) (*cloudformation.SetTypeConfigurationOutput, error)
	SetTypeDefaultVersion(ctx context.Context, params *cloudformation.SetTypeDefaultVersionInput, optFns ...func(*cloudformation.Options)) (*cloudformation.SetTypeDefaultVersionOutput, error)
	SignalResource(ctx context.Context, params *cloudformation.SignalResourceInput, optFns ...func(*cloudformation.Options)) (*cloudformation.SignalResourceOutput, error)
	StartResourceScan(ctx context.Context, params *cloudformation.StartResourceScanInput, optFns ...func(*cloudformation.Options)) (*cloudformation.StartResourceScanOutput, error)
	StopStackSetOperation(ctx context.Context, params *cloudformation.StopStackSetOperationInput, optFns ...func(*cloudformation.Options)) (*cloudformation.StopStackSetOperationOutput, error)
	TestType(ctx context.Context, params *cloudformation.TestTypeInput, optFns ...func(*cloudformation.Options)) (*cloudformation.TestTypeOutput, error)
	UpdateGeneratedTemplate(ctx context.Context, params *cloudformation.UpdateGeneratedTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.UpdateGeneratedTemplateOutput, error)
	UpdateStack(ctx context.Context, params *cloudformation.UpdateStackInput, optFns ...func(*cloudformation.Options)) (*cloudformation.UpdateStackOutput, error)
	UpdateStackInstances(ctx context.Context, params *cloudformation.UpdateStackInstancesInput, optFns ...func(*cloudformation.Options)) (*cloudformation.UpdateStackInstancesOutput, error)
	UpdateStackSet(ctx context.Context, params *cloudformation.UpdateStackSetInput, optFns ...func(*cloudformation.Options)) (*cloudformation.UpdateStackSetOutput, error)
	UpdateTerminationProtection(ctx context.Context, params *cloudformation.UpdateTerminationProtectionInput, optFns ...func(*cloudformation.Options)) (*cloudformation.UpdateTerminationProtectionOutput, error)
	ValidateTemplate(ctx context.Context, params *cloudformation.ValidateTemplateInput, optFns ...func(*cloudformation.Options)) (*cloudformation.ValidateTemplateOutput, error)
}
