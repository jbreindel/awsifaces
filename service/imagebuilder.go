// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
)

// ImagebuilderClient ...
type ImagebuilderClient interface {
	Options() imagebuilder.Options
	CancelImageCreation(ctx context.Context, params *imagebuilder.CancelImageCreationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CancelImageCreationOutput, error)
	CancelLifecycleExecution(ctx context.Context, params *imagebuilder.CancelLifecycleExecutionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CancelLifecycleExecutionOutput, error)
	CreateComponent(ctx context.Context, params *imagebuilder.CreateComponentInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateComponentOutput, error)
	CreateContainerRecipe(ctx context.Context, params *imagebuilder.CreateContainerRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateContainerRecipeOutput, error)
	CreateDistributionConfiguration(ctx context.Context, params *imagebuilder.CreateDistributionConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateDistributionConfigurationOutput, error)
	CreateImage(ctx context.Context, params *imagebuilder.CreateImageInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateImageOutput, error)
	CreateImagePipeline(ctx context.Context, params *imagebuilder.CreateImagePipelineInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateImagePipelineOutput, error)
	CreateImageRecipe(ctx context.Context, params *imagebuilder.CreateImageRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateImageRecipeOutput, error)
	CreateInfrastructureConfiguration(ctx context.Context, params *imagebuilder.CreateInfrastructureConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateInfrastructureConfigurationOutput, error)
	CreateLifecyclePolicy(ctx context.Context, params *imagebuilder.CreateLifecyclePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateLifecyclePolicyOutput, error)
	CreateWorkflow(ctx context.Context, params *imagebuilder.CreateWorkflowInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.CreateWorkflowOutput, error)
	DeleteComponent(ctx context.Context, params *imagebuilder.DeleteComponentInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteComponentOutput, error)
	DeleteContainerRecipe(ctx context.Context, params *imagebuilder.DeleteContainerRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteContainerRecipeOutput, error)
	DeleteDistributionConfiguration(ctx context.Context, params *imagebuilder.DeleteDistributionConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteDistributionConfigurationOutput, error)
	DeleteImage(ctx context.Context, params *imagebuilder.DeleteImageInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteImageOutput, error)
	DeleteImagePipeline(ctx context.Context, params *imagebuilder.DeleteImagePipelineInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteImagePipelineOutput, error)
	DeleteImageRecipe(ctx context.Context, params *imagebuilder.DeleteImageRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteImageRecipeOutput, error)
	DeleteInfrastructureConfiguration(ctx context.Context, params *imagebuilder.DeleteInfrastructureConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error)
	DeleteLifecyclePolicy(ctx context.Context, params *imagebuilder.DeleteLifecyclePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteLifecyclePolicyOutput, error)
	DeleteWorkflow(ctx context.Context, params *imagebuilder.DeleteWorkflowInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.DeleteWorkflowOutput, error)
	GetComponent(ctx context.Context, params *imagebuilder.GetComponentInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetComponentOutput, error)
	GetComponentPolicy(ctx context.Context, params *imagebuilder.GetComponentPolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetComponentPolicyOutput, error)
	GetContainerRecipe(ctx context.Context, params *imagebuilder.GetContainerRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetContainerRecipeOutput, error)
	GetContainerRecipePolicy(ctx context.Context, params *imagebuilder.GetContainerRecipePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetContainerRecipePolicyOutput, error)
	GetDistributionConfiguration(ctx context.Context, params *imagebuilder.GetDistributionConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetDistributionConfigurationOutput, error)
	GetImage(ctx context.Context, params *imagebuilder.GetImageInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetImageOutput, error)
	GetImagePipeline(ctx context.Context, params *imagebuilder.GetImagePipelineInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetImagePipelineOutput, error)
	GetImagePolicy(ctx context.Context, params *imagebuilder.GetImagePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetImagePolicyOutput, error)
	GetImageRecipe(ctx context.Context, params *imagebuilder.GetImageRecipeInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetImageRecipeOutput, error)
	GetImageRecipePolicy(ctx context.Context, params *imagebuilder.GetImageRecipePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetImageRecipePolicyOutput, error)
	GetInfrastructureConfiguration(ctx context.Context, params *imagebuilder.GetInfrastructureConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetInfrastructureConfigurationOutput, error)
	GetLifecycleExecution(ctx context.Context, params *imagebuilder.GetLifecycleExecutionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetLifecycleExecutionOutput, error)
	GetLifecyclePolicy(ctx context.Context, params *imagebuilder.GetLifecyclePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetLifecyclePolicyOutput, error)
	GetWorkflow(ctx context.Context, params *imagebuilder.GetWorkflowInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetWorkflowOutput, error)
	GetWorkflowExecution(ctx context.Context, params *imagebuilder.GetWorkflowExecutionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetWorkflowExecutionOutput, error)
	GetWorkflowStepExecution(ctx context.Context, params *imagebuilder.GetWorkflowStepExecutionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.GetWorkflowStepExecutionOutput, error)
	ImportComponent(ctx context.Context, params *imagebuilder.ImportComponentInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ImportComponentOutput, error)
	ImportVmImage(ctx context.Context, params *imagebuilder.ImportVmImageInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ImportVmImageOutput, error)
	ListComponentBuildVersions(ctx context.Context, params *imagebuilder.ListComponentBuildVersionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListComponentBuildVersionsOutput, error)
	ListComponents(ctx context.Context, params *imagebuilder.ListComponentsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListComponentsOutput, error)
	ListContainerRecipes(ctx context.Context, params *imagebuilder.ListContainerRecipesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListContainerRecipesOutput, error)
	ListDistributionConfigurations(ctx context.Context, params *imagebuilder.ListDistributionConfigurationsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListDistributionConfigurationsOutput, error)
	ListImageBuildVersions(ctx context.Context, params *imagebuilder.ListImageBuildVersionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImageBuildVersionsOutput, error)
	ListImagePackages(ctx context.Context, params *imagebuilder.ListImagePackagesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImagePackagesOutput, error)
	ListImagePipelineImages(ctx context.Context, params *imagebuilder.ListImagePipelineImagesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImagePipelineImagesOutput, error)
	ListImagePipelines(ctx context.Context, params *imagebuilder.ListImagePipelinesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImagePipelinesOutput, error)
	ListImageRecipes(ctx context.Context, params *imagebuilder.ListImageRecipesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImageRecipesOutput, error)
	ListImageScanFindingAggregations(ctx context.Context, params *imagebuilder.ListImageScanFindingAggregationsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImageScanFindingAggregationsOutput, error)
	ListImageScanFindings(ctx context.Context, params *imagebuilder.ListImageScanFindingsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImageScanFindingsOutput, error)
	ListImages(ctx context.Context, params *imagebuilder.ListImagesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListImagesOutput, error)
	ListInfrastructureConfigurations(ctx context.Context, params *imagebuilder.ListInfrastructureConfigurationsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListInfrastructureConfigurationsOutput, error)
	ListLifecycleExecutionResources(ctx context.Context, params *imagebuilder.ListLifecycleExecutionResourcesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListLifecycleExecutionResourcesOutput, error)
	ListLifecycleExecutions(ctx context.Context, params *imagebuilder.ListLifecycleExecutionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListLifecycleExecutionsOutput, error)
	ListLifecyclePolicies(ctx context.Context, params *imagebuilder.ListLifecyclePoliciesInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListLifecyclePoliciesOutput, error)
	ListTagsForResource(ctx context.Context, params *imagebuilder.ListTagsForResourceInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListTagsForResourceOutput, error)
	ListWaitingWorkflowSteps(ctx context.Context, params *imagebuilder.ListWaitingWorkflowStepsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListWaitingWorkflowStepsOutput, error)
	ListWorkflowBuildVersions(ctx context.Context, params *imagebuilder.ListWorkflowBuildVersionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListWorkflowBuildVersionsOutput, error)
	ListWorkflowExecutions(ctx context.Context, params *imagebuilder.ListWorkflowExecutionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListWorkflowExecutionsOutput, error)
	ListWorkflowStepExecutions(ctx context.Context, params *imagebuilder.ListWorkflowStepExecutionsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListWorkflowStepExecutionsOutput, error)
	ListWorkflows(ctx context.Context, params *imagebuilder.ListWorkflowsInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.ListWorkflowsOutput, error)
	PutComponentPolicy(ctx context.Context, params *imagebuilder.PutComponentPolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.PutComponentPolicyOutput, error)
	PutContainerRecipePolicy(ctx context.Context, params *imagebuilder.PutContainerRecipePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.PutContainerRecipePolicyOutput, error)
	PutImagePolicy(ctx context.Context, params *imagebuilder.PutImagePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.PutImagePolicyOutput, error)
	PutImageRecipePolicy(ctx context.Context, params *imagebuilder.PutImageRecipePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.PutImageRecipePolicyOutput, error)
	SendWorkflowStepAction(ctx context.Context, params *imagebuilder.SendWorkflowStepActionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.SendWorkflowStepActionOutput, error)
	StartImagePipelineExecution(ctx context.Context, params *imagebuilder.StartImagePipelineExecutionInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.StartImagePipelineExecutionOutput, error)
	StartResourceStateUpdate(ctx context.Context, params *imagebuilder.StartResourceStateUpdateInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.StartResourceStateUpdateOutput, error)
	TagResource(ctx context.Context, params *imagebuilder.TagResourceInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *imagebuilder.UntagResourceInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.UntagResourceOutput, error)
	UpdateDistributionConfiguration(ctx context.Context, params *imagebuilder.UpdateDistributionConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.UpdateDistributionConfigurationOutput, error)
	UpdateImagePipeline(ctx context.Context, params *imagebuilder.UpdateImagePipelineInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.UpdateImagePipelineOutput, error)
	UpdateInfrastructureConfiguration(ctx context.Context, params *imagebuilder.UpdateInfrastructureConfigurationInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error)
	UpdateLifecyclePolicy(ctx context.Context, params *imagebuilder.UpdateLifecyclePolicyInput, optFns ...func(*imagebuilder.Options)) (*imagebuilder.UpdateLifecyclePolicyOutput, error)
}
