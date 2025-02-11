// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

// Apigatewayv2Client ...
type Apigatewayv2Client interface {
	Options() apigatewayv2.Options
	CreateApi(ctx context.Context, params *apigatewayv2.CreateApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateApiOutput, error)
	CreateApiMapping(ctx context.Context, params *apigatewayv2.CreateApiMappingInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateApiMappingOutput, error)
	CreateAuthorizer(ctx context.Context, params *apigatewayv2.CreateAuthorizerInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateAuthorizerOutput, error)
	CreateDeployment(ctx context.Context, params *apigatewayv2.CreateDeploymentInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateDeploymentOutput, error)
	CreateDomainName(ctx context.Context, params *apigatewayv2.CreateDomainNameInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateDomainNameOutput, error)
	CreateIntegration(ctx context.Context, params *apigatewayv2.CreateIntegrationInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateIntegrationOutput, error)
	CreateIntegrationResponse(ctx context.Context, params *apigatewayv2.CreateIntegrationResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateIntegrationResponseOutput, error)
	CreateModel(ctx context.Context, params *apigatewayv2.CreateModelInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateModelOutput, error)
	CreateRoute(ctx context.Context, params *apigatewayv2.CreateRouteInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateRouteOutput, error)
	CreateRouteResponse(ctx context.Context, params *apigatewayv2.CreateRouteResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateRouteResponseOutput, error)
	CreateStage(ctx context.Context, params *apigatewayv2.CreateStageInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateStageOutput, error)
	CreateVpcLink(ctx context.Context, params *apigatewayv2.CreateVpcLinkInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.CreateVpcLinkOutput, error)
	DeleteAccessLogSettings(ctx context.Context, params *apigatewayv2.DeleteAccessLogSettingsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteAccessLogSettingsOutput, error)
	DeleteApi(ctx context.Context, params *apigatewayv2.DeleteApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteApiOutput, error)
	DeleteApiMapping(ctx context.Context, params *apigatewayv2.DeleteApiMappingInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteApiMappingOutput, error)
	DeleteAuthorizer(ctx context.Context, params *apigatewayv2.DeleteAuthorizerInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteAuthorizerOutput, error)
	DeleteCorsConfiguration(ctx context.Context, params *apigatewayv2.DeleteCorsConfigurationInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteCorsConfigurationOutput, error)
	DeleteDeployment(ctx context.Context, params *apigatewayv2.DeleteDeploymentInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteDeploymentOutput, error)
	DeleteDomainName(ctx context.Context, params *apigatewayv2.DeleteDomainNameInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteDomainNameOutput, error)
	DeleteIntegration(ctx context.Context, params *apigatewayv2.DeleteIntegrationInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteIntegrationOutput, error)
	DeleteIntegrationResponse(ctx context.Context, params *apigatewayv2.DeleteIntegrationResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteIntegrationResponseOutput, error)
	DeleteModel(ctx context.Context, params *apigatewayv2.DeleteModelInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteModelOutput, error)
	DeleteRoute(ctx context.Context, params *apigatewayv2.DeleteRouteInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteRouteOutput, error)
	DeleteRouteRequestParameter(ctx context.Context, params *apigatewayv2.DeleteRouteRequestParameterInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteRouteRequestParameterOutput, error)
	DeleteRouteResponse(ctx context.Context, params *apigatewayv2.DeleteRouteResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteRouteResponseOutput, error)
	DeleteRouteSettings(ctx context.Context, params *apigatewayv2.DeleteRouteSettingsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteRouteSettingsOutput, error)
	DeleteStage(ctx context.Context, params *apigatewayv2.DeleteStageInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteStageOutput, error)
	DeleteVpcLink(ctx context.Context, params *apigatewayv2.DeleteVpcLinkInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.DeleteVpcLinkOutput, error)
	ExportApi(ctx context.Context, params *apigatewayv2.ExportApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.ExportApiOutput, error)
	GetApi(ctx context.Context, params *apigatewayv2.GetApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiOutput, error)
	GetApiMapping(ctx context.Context, params *apigatewayv2.GetApiMappingInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappings(ctx context.Context, params *apigatewayv2.GetApiMappingsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApis(ctx context.Context, params *apigatewayv2.GetApisInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error)
	GetAuthorizer(ctx context.Context, params *apigatewayv2.GetAuthorizerInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizers(ctx context.Context, params *apigatewayv2.GetAuthorizersInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error)
	GetDeployment(ctx context.Context, params *apigatewayv2.GetDeploymentInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeployments(ctx context.Context, params *apigatewayv2.GetDeploymentsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDomainName(ctx context.Context, params *apigatewayv2.GetDomainNameInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNames(ctx context.Context, params *apigatewayv2.GetDomainNamesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error)
	GetIntegration(ctx context.Context, params *apigatewayv2.GetIntegrationInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationResponse(ctx context.Context, params *apigatewayv2.GetIntegrationResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponses(ctx context.Context, params *apigatewayv2.GetIntegrationResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrations(ctx context.Context, params *apigatewayv2.GetIntegrationsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error)
	GetModel(ctx context.Context, params *apigatewayv2.GetModelInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelOutput, error)
	GetModelTemplate(ctx context.Context, params *apigatewayv2.GetModelTemplateInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModels(ctx context.Context, params *apigatewayv2.GetModelsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error)
	GetRoute(ctx context.Context, params *apigatewayv2.GetRouteInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteOutput, error)
	GetRouteResponse(ctx context.Context, params *apigatewayv2.GetRouteResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponses(ctx context.Context, params *apigatewayv2.GetRouteResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRoutes(ctx context.Context, params *apigatewayv2.GetRoutesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error)
	GetStage(ctx context.Context, params *apigatewayv2.GetStageInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStageOutput, error)
	GetStages(ctx context.Context, params *apigatewayv2.GetStagesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error)
	GetTags(ctx context.Context, params *apigatewayv2.GetTagsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error)
	GetVpcLink(ctx context.Context, params *apigatewayv2.GetVpcLinkInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinkOutput, error)
	GetVpcLinks(ctx context.Context, params *apigatewayv2.GetVpcLinksInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error)
	ImportApi(ctx context.Context, params *apigatewayv2.ImportApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.ImportApiOutput, error)
	ReimportApi(ctx context.Context, params *apigatewayv2.ReimportApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.ReimportApiOutput, error)
	ResetAuthorizersCache(ctx context.Context, params *apigatewayv2.ResetAuthorizersCacheInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.ResetAuthorizersCacheOutput, error)
	TagResource(ctx context.Context, params *apigatewayv2.TagResourceInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *apigatewayv2.UntagResourceInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UntagResourceOutput, error)
	UpdateApi(ctx context.Context, params *apigatewayv2.UpdateApiInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateApiOutput, error)
	UpdateApiMapping(ctx context.Context, params *apigatewayv2.UpdateApiMappingInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateApiMappingOutput, error)
	UpdateAuthorizer(ctx context.Context, params *apigatewayv2.UpdateAuthorizerInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateAuthorizerOutput, error)
	UpdateDeployment(ctx context.Context, params *apigatewayv2.UpdateDeploymentInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateDeploymentOutput, error)
	UpdateDomainName(ctx context.Context, params *apigatewayv2.UpdateDomainNameInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateDomainNameOutput, error)
	UpdateIntegration(ctx context.Context, params *apigatewayv2.UpdateIntegrationInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateIntegrationOutput, error)
	UpdateIntegrationResponse(ctx context.Context, params *apigatewayv2.UpdateIntegrationResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateIntegrationResponseOutput, error)
	UpdateModel(ctx context.Context, params *apigatewayv2.UpdateModelInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateModelOutput, error)
	UpdateRoute(ctx context.Context, params *apigatewayv2.UpdateRouteInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateRouteOutput, error)
	UpdateRouteResponse(ctx context.Context, params *apigatewayv2.UpdateRouteResponseInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateRouteResponseOutput, error)
	UpdateStage(ctx context.Context, params *apigatewayv2.UpdateStageInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateStageOutput, error)
	UpdateVpcLink(ctx context.Context, params *apigatewayv2.UpdateVpcLinkInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.UpdateVpcLinkOutput, error)
}
