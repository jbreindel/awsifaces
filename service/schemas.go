// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/schemas"
)

// SchemasClient ...
type SchemasClient interface {
	Options() schemas.Options
	CreateDiscoverer(ctx context.Context, params *schemas.CreateDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.CreateDiscovererOutput, error)
	CreateRegistry(ctx context.Context, params *schemas.CreateRegistryInput, optFns ...func(*schemas.Options)) (*schemas.CreateRegistryOutput, error)
	CreateSchema(ctx context.Context, params *schemas.CreateSchemaInput, optFns ...func(*schemas.Options)) (*schemas.CreateSchemaOutput, error)
	DeleteDiscoverer(ctx context.Context, params *schemas.DeleteDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.DeleteDiscovererOutput, error)
	DeleteRegistry(ctx context.Context, params *schemas.DeleteRegistryInput, optFns ...func(*schemas.Options)) (*schemas.DeleteRegistryOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *schemas.DeleteResourcePolicyInput, optFns ...func(*schemas.Options)) (*schemas.DeleteResourcePolicyOutput, error)
	DeleteSchema(ctx context.Context, params *schemas.DeleteSchemaInput, optFns ...func(*schemas.Options)) (*schemas.DeleteSchemaOutput, error)
	DeleteSchemaVersion(ctx context.Context, params *schemas.DeleteSchemaVersionInput, optFns ...func(*schemas.Options)) (*schemas.DeleteSchemaVersionOutput, error)
	DescribeCodeBinding(ctx context.Context, params *schemas.DescribeCodeBindingInput, optFns ...func(*schemas.Options)) (*schemas.DescribeCodeBindingOutput, error)
	DescribeDiscoverer(ctx context.Context, params *schemas.DescribeDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.DescribeDiscovererOutput, error)
	DescribeRegistry(ctx context.Context, params *schemas.DescribeRegistryInput, optFns ...func(*schemas.Options)) (*schemas.DescribeRegistryOutput, error)
	DescribeSchema(ctx context.Context, params *schemas.DescribeSchemaInput, optFns ...func(*schemas.Options)) (*schemas.DescribeSchemaOutput, error)
	ExportSchema(ctx context.Context, params *schemas.ExportSchemaInput, optFns ...func(*schemas.Options)) (*schemas.ExportSchemaOutput, error)
	GetCodeBindingSource(ctx context.Context, params *schemas.GetCodeBindingSourceInput, optFns ...func(*schemas.Options)) (*schemas.GetCodeBindingSourceOutput, error)
	GetDiscoveredSchema(ctx context.Context, params *schemas.GetDiscoveredSchemaInput, optFns ...func(*schemas.Options)) (*schemas.GetDiscoveredSchemaOutput, error)
	GetResourcePolicy(ctx context.Context, params *schemas.GetResourcePolicyInput, optFns ...func(*schemas.Options)) (*schemas.GetResourcePolicyOutput, error)
	ListDiscoverers(ctx context.Context, params *schemas.ListDiscoverersInput, optFns ...func(*schemas.Options)) (*schemas.ListDiscoverersOutput, error)
	ListRegistries(ctx context.Context, params *schemas.ListRegistriesInput, optFns ...func(*schemas.Options)) (*schemas.ListRegistriesOutput, error)
	ListSchemaVersions(ctx context.Context, params *schemas.ListSchemaVersionsInput, optFns ...func(*schemas.Options)) (*schemas.ListSchemaVersionsOutput, error)
	ListSchemas(ctx context.Context, params *schemas.ListSchemasInput, optFns ...func(*schemas.Options)) (*schemas.ListSchemasOutput, error)
	ListTagsForResource(ctx context.Context, params *schemas.ListTagsForResourceInput, optFns ...func(*schemas.Options)) (*schemas.ListTagsForResourceOutput, error)
	PutCodeBinding(ctx context.Context, params *schemas.PutCodeBindingInput, optFns ...func(*schemas.Options)) (*schemas.PutCodeBindingOutput, error)
	PutResourcePolicy(ctx context.Context, params *schemas.PutResourcePolicyInput, optFns ...func(*schemas.Options)) (*schemas.PutResourcePolicyOutput, error)
	SearchSchemas(ctx context.Context, params *schemas.SearchSchemasInput, optFns ...func(*schemas.Options)) (*schemas.SearchSchemasOutput, error)
	StartDiscoverer(ctx context.Context, params *schemas.StartDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.StartDiscovererOutput, error)
	StopDiscoverer(ctx context.Context, params *schemas.StopDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.StopDiscovererOutput, error)
	TagResource(ctx context.Context, params *schemas.TagResourceInput, optFns ...func(*schemas.Options)) (*schemas.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *schemas.UntagResourceInput, optFns ...func(*schemas.Options)) (*schemas.UntagResourceOutput, error)
	UpdateDiscoverer(ctx context.Context, params *schemas.UpdateDiscovererInput, optFns ...func(*schemas.Options)) (*schemas.UpdateDiscovererOutput, error)
	UpdateRegistry(ctx context.Context, params *schemas.UpdateRegistryInput, optFns ...func(*schemas.Options)) (*schemas.UpdateRegistryOutput, error)
	UpdateSchema(ctx context.Context, params *schemas.UpdateSchemaInput, optFns ...func(*schemas.Options)) (*schemas.UpdateSchemaOutput, error)
}
