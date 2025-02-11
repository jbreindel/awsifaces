// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/keyspaces"
)

// KeyspacesClient ...
type KeyspacesClient interface {
	Options() keyspaces.Options
	CreateKeyspace(ctx context.Context, params *keyspaces.CreateKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.CreateKeyspaceOutput, error)
	CreateTable(ctx context.Context, params *keyspaces.CreateTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.CreateTableOutput, error)
	DeleteKeyspace(ctx context.Context, params *keyspaces.DeleteKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.DeleteKeyspaceOutput, error)
	DeleteTable(ctx context.Context, params *keyspaces.DeleteTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.DeleteTableOutput, error)
	GetKeyspace(ctx context.Context, params *keyspaces.GetKeyspaceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetKeyspaceOutput, error)
	GetTable(ctx context.Context, params *keyspaces.GetTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetTableOutput, error)
	GetTableAutoScalingSettings(ctx context.Context, params *keyspaces.GetTableAutoScalingSettingsInput, optFns ...func(*keyspaces.Options)) (*keyspaces.GetTableAutoScalingSettingsOutput, error)
	ListKeyspaces(ctx context.Context, params *keyspaces.ListKeyspacesInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListKeyspacesOutput, error)
	ListTables(ctx context.Context, params *keyspaces.ListTablesInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListTablesOutput, error)
	ListTagsForResource(ctx context.Context, params *keyspaces.ListTagsForResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.ListTagsForResourceOutput, error)
	RestoreTable(ctx context.Context, params *keyspaces.RestoreTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.RestoreTableOutput, error)
	TagResource(ctx context.Context, params *keyspaces.TagResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *keyspaces.UntagResourceInput, optFns ...func(*keyspaces.Options)) (*keyspaces.UntagResourceOutput, error)
	UpdateTable(ctx context.Context, params *keyspaces.UpdateTableInput, optFns ...func(*keyspaces.Options)) (*keyspaces.UpdateTableOutput, error)
}
