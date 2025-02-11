// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
)

// ResourcegroupsClient ...
type ResourcegroupsClient interface {
	Options() resourcegroups.Options
	CreateGroup(ctx context.Context, params *resourcegroups.CreateGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.CreateGroupOutput, error)
	DeleteGroup(ctx context.Context, params *resourcegroups.DeleteGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.DeleteGroupOutput, error)
	GetAccountSettings(ctx context.Context, params *resourcegroups.GetAccountSettingsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetAccountSettingsOutput, error)
	GetGroup(ctx context.Context, params *resourcegroups.GetGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error)
	GetGroupConfiguration(ctx context.Context, params *resourcegroups.GetGroupConfigurationInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error)
	GetGroupQuery(ctx context.Context, params *resourcegroups.GetGroupQueryInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error)
	GetTags(ctx context.Context, params *resourcegroups.GetTagsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error)
	GroupResources(ctx context.Context, params *resourcegroups.GroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GroupResourcesOutput, error)
	ListGroupResources(ctx context.Context, params *resourcegroups.ListGroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error)
	ListGroups(ctx context.Context, params *resourcegroups.ListGroupsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error)
	PutGroupConfiguration(ctx context.Context, params *resourcegroups.PutGroupConfigurationInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.PutGroupConfigurationOutput, error)
	SearchResources(ctx context.Context, params *resourcegroups.SearchResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error)
	Tag(ctx context.Context, params *resourcegroups.TagInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.TagOutput, error)
	UngroupResources(ctx context.Context, params *resourcegroups.UngroupResourcesInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UngroupResourcesOutput, error)
	Untag(ctx context.Context, params *resourcegroups.UntagInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UntagOutput, error)
	UpdateAccountSettings(ctx context.Context, params *resourcegroups.UpdateAccountSettingsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateAccountSettingsOutput, error)
	UpdateGroup(ctx context.Context, params *resourcegroups.UpdateGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupOutput, error)
	UpdateGroupQuery(ctx context.Context, params *resourcegroups.UpdateGroupQueryInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.UpdateGroupQueryOutput, error)
}
