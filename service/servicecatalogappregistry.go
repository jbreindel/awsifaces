// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
)

// ServicecatalogappregistryClient ...
type ServicecatalogappregistryClient interface {
	Options() servicecatalogappregistry.Options
	AssociateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.AssociateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateAttributeGroupOutput, error)
	AssociateResource(ctx context.Context, params *servicecatalogappregistry.AssociateResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.AssociateResourceOutput, error)
	CreateApplication(ctx context.Context, params *servicecatalogappregistry.CreateApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateApplicationOutput, error)
	CreateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.CreateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.CreateAttributeGroupOutput, error)
	DeleteApplication(ctx context.Context, params *servicecatalogappregistry.DeleteApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteApplicationOutput, error)
	DeleteAttributeGroup(ctx context.Context, params *servicecatalogappregistry.DeleteAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DeleteAttributeGroupOutput, error)
	DisassociateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.DisassociateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateAttributeGroupOutput, error)
	DisassociateResource(ctx context.Context, params *servicecatalogappregistry.DisassociateResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.DisassociateResourceOutput, error)
	GetApplication(ctx context.Context, params *servicecatalogappregistry.GetApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error)
	GetAssociatedResource(ctx context.Context, params *servicecatalogappregistry.GetAssociatedResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error)
	GetAttributeGroup(ctx context.Context, params *servicecatalogappregistry.GetAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error)
	GetConfiguration(ctx context.Context, params *servicecatalogappregistry.GetConfigurationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error)
	ListApplications(ctx context.Context, params *servicecatalogappregistry.ListApplicationsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error)
	ListAssociatedAttributeGroups(ctx context.Context, params *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error)
	ListAssociatedResources(ctx context.Context, params *servicecatalogappregistry.ListAssociatedResourcesInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error)
	ListAttributeGroups(ctx context.Context, params *servicecatalogappregistry.ListAttributeGroupsInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error)
	ListAttributeGroupsForApplication(ctx context.Context, params *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error)
	ListTagsForResource(ctx context.Context, params *servicecatalogappregistry.ListTagsForResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error)
	PutConfiguration(ctx context.Context, params *servicecatalogappregistry.PutConfigurationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.PutConfigurationOutput, error)
	SyncResource(ctx context.Context, params *servicecatalogappregistry.SyncResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.SyncResourceOutput, error)
	TagResource(ctx context.Context, params *servicecatalogappregistry.TagResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *servicecatalogappregistry.UntagResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UntagResourceOutput, error)
	UpdateApplication(ctx context.Context, params *servicecatalogappregistry.UpdateApplicationInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateApplicationOutput, error)
	UpdateAttributeGroup(ctx context.Context, params *servicecatalogappregistry.UpdateAttributeGroupInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.UpdateAttributeGroupOutput, error)
}
