// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
)

// RedshiftserverlessClient ...
type RedshiftserverlessClient interface {
	Options() redshiftserverless.Options
	ConvertRecoveryPointToSnapshot(ctx context.Context, params *redshiftserverless.ConvertRecoveryPointToSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ConvertRecoveryPointToSnapshotOutput, error)
	CreateCustomDomainAssociation(ctx context.Context, params *redshiftserverless.CreateCustomDomainAssociationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateCustomDomainAssociationOutput, error)
	CreateEndpointAccess(ctx context.Context, params *redshiftserverless.CreateEndpointAccessInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateEndpointAccessOutput, error)
	CreateNamespace(ctx context.Context, params *redshiftserverless.CreateNamespaceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateNamespaceOutput, error)
	CreateScheduledAction(ctx context.Context, params *redshiftserverless.CreateScheduledActionInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateScheduledActionOutput, error)
	CreateSnapshot(ctx context.Context, params *redshiftserverless.CreateSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateSnapshotOutput, error)
	CreateSnapshotCopyConfiguration(ctx context.Context, params *redshiftserverless.CreateSnapshotCopyConfigurationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateSnapshotCopyConfigurationOutput, error)
	CreateUsageLimit(ctx context.Context, params *redshiftserverless.CreateUsageLimitInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateUsageLimitOutput, error)
	CreateWorkgroup(ctx context.Context, params *redshiftserverless.CreateWorkgroupInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.CreateWorkgroupOutput, error)
	DeleteCustomDomainAssociation(ctx context.Context, params *redshiftserverless.DeleteCustomDomainAssociationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteCustomDomainAssociationOutput, error)
	DeleteEndpointAccess(ctx context.Context, params *redshiftserverless.DeleteEndpointAccessInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteEndpointAccessOutput, error)
	DeleteNamespace(ctx context.Context, params *redshiftserverless.DeleteNamespaceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteNamespaceOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *redshiftserverless.DeleteResourcePolicyInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteResourcePolicyOutput, error)
	DeleteScheduledAction(ctx context.Context, params *redshiftserverless.DeleteScheduledActionInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteScheduledActionOutput, error)
	DeleteSnapshot(ctx context.Context, params *redshiftserverless.DeleteSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteSnapshotOutput, error)
	DeleteSnapshotCopyConfiguration(ctx context.Context, params *redshiftserverless.DeleteSnapshotCopyConfigurationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteSnapshotCopyConfigurationOutput, error)
	DeleteUsageLimit(ctx context.Context, params *redshiftserverless.DeleteUsageLimitInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteUsageLimitOutput, error)
	DeleteWorkgroup(ctx context.Context, params *redshiftserverless.DeleteWorkgroupInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.DeleteWorkgroupOutput, error)
	GetCredentials(ctx context.Context, params *redshiftserverless.GetCredentialsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetCredentialsOutput, error)
	GetCustomDomainAssociation(ctx context.Context, params *redshiftserverless.GetCustomDomainAssociationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetCustomDomainAssociationOutput, error)
	GetEndpointAccess(ctx context.Context, params *redshiftserverless.GetEndpointAccessInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetEndpointAccessOutput, error)
	GetNamespace(ctx context.Context, params *redshiftserverless.GetNamespaceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetNamespaceOutput, error)
	GetRecoveryPoint(ctx context.Context, params *redshiftserverless.GetRecoveryPointInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetRecoveryPointOutput, error)
	GetResourcePolicy(ctx context.Context, params *redshiftserverless.GetResourcePolicyInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetResourcePolicyOutput, error)
	GetScheduledAction(ctx context.Context, params *redshiftserverless.GetScheduledActionInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetScheduledActionOutput, error)
	GetSnapshot(ctx context.Context, params *redshiftserverless.GetSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetSnapshotOutput, error)
	GetTableRestoreStatus(ctx context.Context, params *redshiftserverless.GetTableRestoreStatusInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetTableRestoreStatusOutput, error)
	GetUsageLimit(ctx context.Context, params *redshiftserverless.GetUsageLimitInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetUsageLimitOutput, error)
	GetWorkgroup(ctx context.Context, params *redshiftserverless.GetWorkgroupInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.GetWorkgroupOutput, error)
	ListCustomDomainAssociations(ctx context.Context, params *redshiftserverless.ListCustomDomainAssociationsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListCustomDomainAssociationsOutput, error)
	ListEndpointAccess(ctx context.Context, params *redshiftserverless.ListEndpointAccessInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListEndpointAccessOutput, error)
	ListNamespaces(ctx context.Context, params *redshiftserverless.ListNamespacesInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListNamespacesOutput, error)
	ListRecoveryPoints(ctx context.Context, params *redshiftserverless.ListRecoveryPointsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListRecoveryPointsOutput, error)
	ListScheduledActions(ctx context.Context, params *redshiftserverless.ListScheduledActionsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListScheduledActionsOutput, error)
	ListSnapshotCopyConfigurations(ctx context.Context, params *redshiftserverless.ListSnapshotCopyConfigurationsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListSnapshotCopyConfigurationsOutput, error)
	ListSnapshots(ctx context.Context, params *redshiftserverless.ListSnapshotsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListSnapshotsOutput, error)
	ListTableRestoreStatus(ctx context.Context, params *redshiftserverless.ListTableRestoreStatusInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListTableRestoreStatusOutput, error)
	ListTagsForResource(ctx context.Context, params *redshiftserverless.ListTagsForResourceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListTagsForResourceOutput, error)
	ListUsageLimits(ctx context.Context, params *redshiftserverless.ListUsageLimitsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListUsageLimitsOutput, error)
	ListWorkgroups(ctx context.Context, params *redshiftserverless.ListWorkgroupsInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.ListWorkgroupsOutput, error)
	PutResourcePolicy(ctx context.Context, params *redshiftserverless.PutResourcePolicyInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.PutResourcePolicyOutput, error)
	RestoreFromRecoveryPoint(ctx context.Context, params *redshiftserverless.RestoreFromRecoveryPointInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.RestoreFromRecoveryPointOutput, error)
	RestoreFromSnapshot(ctx context.Context, params *redshiftserverless.RestoreFromSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.RestoreFromSnapshotOutput, error)
	RestoreTableFromRecoveryPoint(ctx context.Context, params *redshiftserverless.RestoreTableFromRecoveryPointInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.RestoreTableFromRecoveryPointOutput, error)
	RestoreTableFromSnapshot(ctx context.Context, params *redshiftserverless.RestoreTableFromSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.RestoreTableFromSnapshotOutput, error)
	TagResource(ctx context.Context, params *redshiftserverless.TagResourceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *redshiftserverless.UntagResourceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UntagResourceOutput, error)
	UpdateCustomDomainAssociation(ctx context.Context, params *redshiftserverless.UpdateCustomDomainAssociationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateCustomDomainAssociationOutput, error)
	UpdateEndpointAccess(ctx context.Context, params *redshiftserverless.UpdateEndpointAccessInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateEndpointAccessOutput, error)
	UpdateNamespace(ctx context.Context, params *redshiftserverless.UpdateNamespaceInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateNamespaceOutput, error)
	UpdateScheduledAction(ctx context.Context, params *redshiftserverless.UpdateScheduledActionInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateScheduledActionOutput, error)
	UpdateSnapshot(ctx context.Context, params *redshiftserverless.UpdateSnapshotInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateSnapshotOutput, error)
	UpdateSnapshotCopyConfiguration(ctx context.Context, params *redshiftserverless.UpdateSnapshotCopyConfigurationInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateSnapshotCopyConfigurationOutput, error)
	UpdateUsageLimit(ctx context.Context, params *redshiftserverless.UpdateUsageLimitInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateUsageLimitOutput, error)
	UpdateWorkgroup(ctx context.Context, params *redshiftserverless.UpdateWorkgroupInput, optFns ...func(*redshiftserverless.Options)) (*redshiftserverless.UpdateWorkgroupOutput, error)
}
