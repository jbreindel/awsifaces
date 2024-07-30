// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/customerprofiles"
)

// CustomerprofilesClient ...
type CustomerprofilesClient interface {
	Options() customerprofiles.Options
	AddProfileKey(ctx context.Context, params *customerprofiles.AddProfileKeyInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.AddProfileKeyOutput, error)
	CreateCalculatedAttributeDefinition(ctx context.Context, params *customerprofiles.CreateCalculatedAttributeDefinitionInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.CreateCalculatedAttributeDefinitionOutput, error)
	CreateDomain(ctx context.Context, params *customerprofiles.CreateDomainInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.CreateDomainOutput, error)
	CreateEventStream(ctx context.Context, params *customerprofiles.CreateEventStreamInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.CreateEventStreamOutput, error)
	CreateIntegrationWorkflow(ctx context.Context, params *customerprofiles.CreateIntegrationWorkflowInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.CreateIntegrationWorkflowOutput, error)
	CreateProfile(ctx context.Context, params *customerprofiles.CreateProfileInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.CreateProfileOutput, error)
	DeleteCalculatedAttributeDefinition(ctx context.Context, params *customerprofiles.DeleteCalculatedAttributeDefinitionInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteCalculatedAttributeDefinitionOutput, error)
	DeleteDomain(ctx context.Context, params *customerprofiles.DeleteDomainInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteDomainOutput, error)
	DeleteEventStream(ctx context.Context, params *customerprofiles.DeleteEventStreamInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteEventStreamOutput, error)
	DeleteIntegration(ctx context.Context, params *customerprofiles.DeleteIntegrationInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteIntegrationOutput, error)
	DeleteProfile(ctx context.Context, params *customerprofiles.DeleteProfileInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteProfileOutput, error)
	DeleteProfileKey(ctx context.Context, params *customerprofiles.DeleteProfileKeyInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteProfileKeyOutput, error)
	DeleteProfileObject(ctx context.Context, params *customerprofiles.DeleteProfileObjectInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteProfileObjectOutput, error)
	DeleteProfileObjectType(ctx context.Context, params *customerprofiles.DeleteProfileObjectTypeInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteProfileObjectTypeOutput, error)
	DeleteWorkflow(ctx context.Context, params *customerprofiles.DeleteWorkflowInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DeleteWorkflowOutput, error)
	DetectProfileObjectType(ctx context.Context, params *customerprofiles.DetectProfileObjectTypeInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.DetectProfileObjectTypeOutput, error)
	GetAutoMergingPreview(ctx context.Context, params *customerprofiles.GetAutoMergingPreviewInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetAutoMergingPreviewOutput, error)
	GetCalculatedAttributeDefinition(ctx context.Context, params *customerprofiles.GetCalculatedAttributeDefinitionInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetCalculatedAttributeDefinitionOutput, error)
	GetCalculatedAttributeForProfile(ctx context.Context, params *customerprofiles.GetCalculatedAttributeForProfileInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetCalculatedAttributeForProfileOutput, error)
	GetDomain(ctx context.Context, params *customerprofiles.GetDomainInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetDomainOutput, error)
	GetEventStream(ctx context.Context, params *customerprofiles.GetEventStreamInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetEventStreamOutput, error)
	GetIdentityResolutionJob(ctx context.Context, params *customerprofiles.GetIdentityResolutionJobInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetIdentityResolutionJobOutput, error)
	GetIntegration(ctx context.Context, params *customerprofiles.GetIntegrationInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetIntegrationOutput, error)
	GetMatches(ctx context.Context, params *customerprofiles.GetMatchesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetMatchesOutput, error)
	GetProfileObjectType(ctx context.Context, params *customerprofiles.GetProfileObjectTypeInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetProfileObjectTypeOutput, error)
	GetProfileObjectTypeTemplate(ctx context.Context, params *customerprofiles.GetProfileObjectTypeTemplateInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error)
	GetSimilarProfiles(ctx context.Context, params *customerprofiles.GetSimilarProfilesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetSimilarProfilesOutput, error)
	GetWorkflow(ctx context.Context, params *customerprofiles.GetWorkflowInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetWorkflowOutput, error)
	GetWorkflowSteps(ctx context.Context, params *customerprofiles.GetWorkflowStepsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.GetWorkflowStepsOutput, error)
	ListAccountIntegrations(ctx context.Context, params *customerprofiles.ListAccountIntegrationsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListAccountIntegrationsOutput, error)
	ListCalculatedAttributeDefinitions(ctx context.Context, params *customerprofiles.ListCalculatedAttributeDefinitionsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListCalculatedAttributeDefinitionsOutput, error)
	ListCalculatedAttributesForProfile(ctx context.Context, params *customerprofiles.ListCalculatedAttributesForProfileInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListCalculatedAttributesForProfileOutput, error)
	ListDomains(ctx context.Context, params *customerprofiles.ListDomainsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListDomainsOutput, error)
	ListEventStreams(ctx context.Context, params *customerprofiles.ListEventStreamsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListEventStreamsOutput, error)
	ListIdentityResolutionJobs(ctx context.Context, params *customerprofiles.ListIdentityResolutionJobsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListIdentityResolutionJobsOutput, error)
	ListIntegrations(ctx context.Context, params *customerprofiles.ListIntegrationsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListIntegrationsOutput, error)
	ListProfileObjectTypeTemplates(ctx context.Context, params *customerprofiles.ListProfileObjectTypeTemplatesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error)
	ListProfileObjectTypes(ctx context.Context, params *customerprofiles.ListProfileObjectTypesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListProfileObjectTypesOutput, error)
	ListProfileObjects(ctx context.Context, params *customerprofiles.ListProfileObjectsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListProfileObjectsOutput, error)
	ListRuleBasedMatches(ctx context.Context, params *customerprofiles.ListRuleBasedMatchesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListRuleBasedMatchesOutput, error)
	ListTagsForResource(ctx context.Context, params *customerprofiles.ListTagsForResourceInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListTagsForResourceOutput, error)
	ListWorkflows(ctx context.Context, params *customerprofiles.ListWorkflowsInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.ListWorkflowsOutput, error)
	MergeProfiles(ctx context.Context, params *customerprofiles.MergeProfilesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.MergeProfilesOutput, error)
	PutIntegration(ctx context.Context, params *customerprofiles.PutIntegrationInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.PutIntegrationOutput, error)
	PutProfileObject(ctx context.Context, params *customerprofiles.PutProfileObjectInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.PutProfileObjectOutput, error)
	PutProfileObjectType(ctx context.Context, params *customerprofiles.PutProfileObjectTypeInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.PutProfileObjectTypeOutput, error)
	SearchProfiles(ctx context.Context, params *customerprofiles.SearchProfilesInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.SearchProfilesOutput, error)
	TagResource(ctx context.Context, params *customerprofiles.TagResourceInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *customerprofiles.UntagResourceInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.UntagResourceOutput, error)
	UpdateCalculatedAttributeDefinition(ctx context.Context, params *customerprofiles.UpdateCalculatedAttributeDefinitionInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.UpdateCalculatedAttributeDefinitionOutput, error)
	UpdateDomain(ctx context.Context, params *customerprofiles.UpdateDomainInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.UpdateDomainOutput, error)
	UpdateProfile(ctx context.Context, params *customerprofiles.UpdateProfileInput, optFns ...func(*customerprofiles.Options)) (*customerprofiles.UpdateProfileOutput, error)
}
