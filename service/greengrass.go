// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/greengrass"
)

// GreengrassClient ...
type GreengrassClient interface {
	Options() greengrass.Options
	AssociateRoleToGroup(ctx context.Context, params *greengrass.AssociateRoleToGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.AssociateRoleToGroupOutput, error)
	AssociateServiceRoleToAccount(ctx context.Context, params *greengrass.AssociateServiceRoleToAccountInput, optFns ...func(*greengrass.Options)) (*greengrass.AssociateServiceRoleToAccountOutput, error)
	CreateConnectorDefinition(ctx context.Context, params *greengrass.CreateConnectorDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateConnectorDefinitionOutput, error)
	CreateConnectorDefinitionVersion(ctx context.Context, params *greengrass.CreateConnectorDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateConnectorDefinitionVersionOutput, error)
	CreateCoreDefinition(ctx context.Context, params *greengrass.CreateCoreDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateCoreDefinitionOutput, error)
	CreateCoreDefinitionVersion(ctx context.Context, params *greengrass.CreateCoreDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateCoreDefinitionVersionOutput, error)
	CreateDeployment(ctx context.Context, params *greengrass.CreateDeploymentInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateDeploymentOutput, error)
	CreateDeviceDefinition(ctx context.Context, params *greengrass.CreateDeviceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateDeviceDefinitionOutput, error)
	CreateDeviceDefinitionVersion(ctx context.Context, params *greengrass.CreateDeviceDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateDeviceDefinitionVersionOutput, error)
	CreateFunctionDefinition(ctx context.Context, params *greengrass.CreateFunctionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateFunctionDefinitionOutput, error)
	CreateFunctionDefinitionVersion(ctx context.Context, params *greengrass.CreateFunctionDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateFunctionDefinitionVersionOutput, error)
	CreateGroup(ctx context.Context, params *greengrass.CreateGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateGroupOutput, error)
	CreateGroupCertificateAuthority(ctx context.Context, params *greengrass.CreateGroupCertificateAuthorityInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateGroupCertificateAuthorityOutput, error)
	CreateGroupVersion(ctx context.Context, params *greengrass.CreateGroupVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateGroupVersionOutput, error)
	CreateLoggerDefinition(ctx context.Context, params *greengrass.CreateLoggerDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateLoggerDefinitionOutput, error)
	CreateLoggerDefinitionVersion(ctx context.Context, params *greengrass.CreateLoggerDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateLoggerDefinitionVersionOutput, error)
	CreateResourceDefinition(ctx context.Context, params *greengrass.CreateResourceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateResourceDefinitionOutput, error)
	CreateResourceDefinitionVersion(ctx context.Context, params *greengrass.CreateResourceDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateResourceDefinitionVersionOutput, error)
	CreateSoftwareUpdateJob(ctx context.Context, params *greengrass.CreateSoftwareUpdateJobInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateSoftwareUpdateJobOutput, error)
	CreateSubscriptionDefinition(ctx context.Context, params *greengrass.CreateSubscriptionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateSubscriptionDefinitionOutput, error)
	CreateSubscriptionDefinitionVersion(ctx context.Context, params *greengrass.CreateSubscriptionDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error)
	DeleteConnectorDefinition(ctx context.Context, params *greengrass.DeleteConnectorDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteConnectorDefinitionOutput, error)
	DeleteCoreDefinition(ctx context.Context, params *greengrass.DeleteCoreDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteCoreDefinitionOutput, error)
	DeleteDeviceDefinition(ctx context.Context, params *greengrass.DeleteDeviceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteDeviceDefinitionOutput, error)
	DeleteFunctionDefinition(ctx context.Context, params *greengrass.DeleteFunctionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteFunctionDefinitionOutput, error)
	DeleteGroup(ctx context.Context, params *greengrass.DeleteGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteGroupOutput, error)
	DeleteLoggerDefinition(ctx context.Context, params *greengrass.DeleteLoggerDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteLoggerDefinitionOutput, error)
	DeleteResourceDefinition(ctx context.Context, params *greengrass.DeleteResourceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteResourceDefinitionOutput, error)
	DeleteSubscriptionDefinition(ctx context.Context, params *greengrass.DeleteSubscriptionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.DeleteSubscriptionDefinitionOutput, error)
	DisassociateRoleFromGroup(ctx context.Context, params *greengrass.DisassociateRoleFromGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.DisassociateRoleFromGroupOutput, error)
	DisassociateServiceRoleFromAccount(ctx context.Context, params *greengrass.DisassociateServiceRoleFromAccountInput, optFns ...func(*greengrass.Options)) (*greengrass.DisassociateServiceRoleFromAccountOutput, error)
	GetAssociatedRole(ctx context.Context, params *greengrass.GetAssociatedRoleInput, optFns ...func(*greengrass.Options)) (*greengrass.GetAssociatedRoleOutput, error)
	GetBulkDeploymentStatus(ctx context.Context, params *greengrass.GetBulkDeploymentStatusInput, optFns ...func(*greengrass.Options)) (*greengrass.GetBulkDeploymentStatusOutput, error)
	GetConnectivityInfo(ctx context.Context, params *greengrass.GetConnectivityInfoInput, optFns ...func(*greengrass.Options)) (*greengrass.GetConnectivityInfoOutput, error)
	GetConnectorDefinition(ctx context.Context, params *greengrass.GetConnectorDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetConnectorDefinitionOutput, error)
	GetConnectorDefinitionVersion(ctx context.Context, params *greengrass.GetConnectorDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetConnectorDefinitionVersionOutput, error)
	GetCoreDefinition(ctx context.Context, params *greengrass.GetCoreDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetCoreDefinitionOutput, error)
	GetCoreDefinitionVersion(ctx context.Context, params *greengrass.GetCoreDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetCoreDefinitionVersionOutput, error)
	GetDeploymentStatus(ctx context.Context, params *greengrass.GetDeploymentStatusInput, optFns ...func(*greengrass.Options)) (*greengrass.GetDeploymentStatusOutput, error)
	GetDeviceDefinition(ctx context.Context, params *greengrass.GetDeviceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetDeviceDefinitionOutput, error)
	GetDeviceDefinitionVersion(ctx context.Context, params *greengrass.GetDeviceDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetDeviceDefinitionVersionOutput, error)
	GetFunctionDefinition(ctx context.Context, params *greengrass.GetFunctionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetFunctionDefinitionOutput, error)
	GetFunctionDefinitionVersion(ctx context.Context, params *greengrass.GetFunctionDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetFunctionDefinitionVersionOutput, error)
	GetGroup(ctx context.Context, params *greengrass.GetGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.GetGroupOutput, error)
	GetGroupCertificateAuthority(ctx context.Context, params *greengrass.GetGroupCertificateAuthorityInput, optFns ...func(*greengrass.Options)) (*greengrass.GetGroupCertificateAuthorityOutput, error)
	GetGroupCertificateConfiguration(ctx context.Context, params *greengrass.GetGroupCertificateConfigurationInput, optFns ...func(*greengrass.Options)) (*greengrass.GetGroupCertificateConfigurationOutput, error)
	GetGroupVersion(ctx context.Context, params *greengrass.GetGroupVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetGroupVersionOutput, error)
	GetLoggerDefinition(ctx context.Context, params *greengrass.GetLoggerDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetLoggerDefinitionOutput, error)
	GetLoggerDefinitionVersion(ctx context.Context, params *greengrass.GetLoggerDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetLoggerDefinitionVersionOutput, error)
	GetResourceDefinition(ctx context.Context, params *greengrass.GetResourceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetResourceDefinitionOutput, error)
	GetResourceDefinitionVersion(ctx context.Context, params *greengrass.GetResourceDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetResourceDefinitionVersionOutput, error)
	GetServiceRoleForAccount(ctx context.Context, params *greengrass.GetServiceRoleForAccountInput, optFns ...func(*greengrass.Options)) (*greengrass.GetServiceRoleForAccountOutput, error)
	GetSubscriptionDefinition(ctx context.Context, params *greengrass.GetSubscriptionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetSubscriptionDefinitionOutput, error)
	GetSubscriptionDefinitionVersion(ctx context.Context, params *greengrass.GetSubscriptionDefinitionVersionInput, optFns ...func(*greengrass.Options)) (*greengrass.GetSubscriptionDefinitionVersionOutput, error)
	GetThingRuntimeConfiguration(ctx context.Context, params *greengrass.GetThingRuntimeConfigurationInput, optFns ...func(*greengrass.Options)) (*greengrass.GetThingRuntimeConfigurationOutput, error)
	ListBulkDeploymentDetailedReports(ctx context.Context, params *greengrass.ListBulkDeploymentDetailedReportsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error)
	ListBulkDeployments(ctx context.Context, params *greengrass.ListBulkDeploymentsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListBulkDeploymentsOutput, error)
	ListConnectorDefinitionVersions(ctx context.Context, params *greengrass.ListConnectorDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListConnectorDefinitionVersionsOutput, error)
	ListConnectorDefinitions(ctx context.Context, params *greengrass.ListConnectorDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListConnectorDefinitionsOutput, error)
	ListCoreDefinitionVersions(ctx context.Context, params *greengrass.ListCoreDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListCoreDefinitionVersionsOutput, error)
	ListCoreDefinitions(ctx context.Context, params *greengrass.ListCoreDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListCoreDefinitionsOutput, error)
	ListDeployments(ctx context.Context, params *greengrass.ListDeploymentsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListDeploymentsOutput, error)
	ListDeviceDefinitionVersions(ctx context.Context, params *greengrass.ListDeviceDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListDeviceDefinitionVersionsOutput, error)
	ListDeviceDefinitions(ctx context.Context, params *greengrass.ListDeviceDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListDeviceDefinitionsOutput, error)
	ListFunctionDefinitionVersions(ctx context.Context, params *greengrass.ListFunctionDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListFunctionDefinitionVersionsOutput, error)
	ListFunctionDefinitions(ctx context.Context, params *greengrass.ListFunctionDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListFunctionDefinitionsOutput, error)
	ListGroupCertificateAuthorities(ctx context.Context, params *greengrass.ListGroupCertificateAuthoritiesInput, optFns ...func(*greengrass.Options)) (*greengrass.ListGroupCertificateAuthoritiesOutput, error)
	ListGroupVersions(ctx context.Context, params *greengrass.ListGroupVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListGroupVersionsOutput, error)
	ListGroups(ctx context.Context, params *greengrass.ListGroupsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListGroupsOutput, error)
	ListLoggerDefinitionVersions(ctx context.Context, params *greengrass.ListLoggerDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListLoggerDefinitionVersionsOutput, error)
	ListLoggerDefinitions(ctx context.Context, params *greengrass.ListLoggerDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListLoggerDefinitionsOutput, error)
	ListResourceDefinitionVersions(ctx context.Context, params *greengrass.ListResourceDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListResourceDefinitionVersionsOutput, error)
	ListResourceDefinitions(ctx context.Context, params *greengrass.ListResourceDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListResourceDefinitionsOutput, error)
	ListSubscriptionDefinitionVersions(ctx context.Context, params *greengrass.ListSubscriptionDefinitionVersionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error)
	ListSubscriptionDefinitions(ctx context.Context, params *greengrass.ListSubscriptionDefinitionsInput, optFns ...func(*greengrass.Options)) (*greengrass.ListSubscriptionDefinitionsOutput, error)
	ListTagsForResource(ctx context.Context, params *greengrass.ListTagsForResourceInput, optFns ...func(*greengrass.Options)) (*greengrass.ListTagsForResourceOutput, error)
	ResetDeployments(ctx context.Context, params *greengrass.ResetDeploymentsInput, optFns ...func(*greengrass.Options)) (*greengrass.ResetDeploymentsOutput, error)
	StartBulkDeployment(ctx context.Context, params *greengrass.StartBulkDeploymentInput, optFns ...func(*greengrass.Options)) (*greengrass.StartBulkDeploymentOutput, error)
	StopBulkDeployment(ctx context.Context, params *greengrass.StopBulkDeploymentInput, optFns ...func(*greengrass.Options)) (*greengrass.StopBulkDeploymentOutput, error)
	TagResource(ctx context.Context, params *greengrass.TagResourceInput, optFns ...func(*greengrass.Options)) (*greengrass.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *greengrass.UntagResourceInput, optFns ...func(*greengrass.Options)) (*greengrass.UntagResourceOutput, error)
	UpdateConnectivityInfo(ctx context.Context, params *greengrass.UpdateConnectivityInfoInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateConnectivityInfoOutput, error)
	UpdateConnectorDefinition(ctx context.Context, params *greengrass.UpdateConnectorDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateConnectorDefinitionOutput, error)
	UpdateCoreDefinition(ctx context.Context, params *greengrass.UpdateCoreDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateCoreDefinitionOutput, error)
	UpdateDeviceDefinition(ctx context.Context, params *greengrass.UpdateDeviceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateDeviceDefinitionOutput, error)
	UpdateFunctionDefinition(ctx context.Context, params *greengrass.UpdateFunctionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateFunctionDefinitionOutput, error)
	UpdateGroup(ctx context.Context, params *greengrass.UpdateGroupInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateGroupOutput, error)
	UpdateGroupCertificateConfiguration(ctx context.Context, params *greengrass.UpdateGroupCertificateConfigurationInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateGroupCertificateConfigurationOutput, error)
	UpdateLoggerDefinition(ctx context.Context, params *greengrass.UpdateLoggerDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateLoggerDefinitionOutput, error)
	UpdateResourceDefinition(ctx context.Context, params *greengrass.UpdateResourceDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateResourceDefinitionOutput, error)
	UpdateSubscriptionDefinition(ctx context.Context, params *greengrass.UpdateSubscriptionDefinitionInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateSubscriptionDefinitionOutput, error)
	UpdateThingRuntimeConfiguration(ctx context.Context, params *greengrass.UpdateThingRuntimeConfigurationInput, optFns ...func(*greengrass.Options)) (*greengrass.UpdateThingRuntimeConfigurationOutput, error)
}
