// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/supportapp"
)

// SupportappClient ...
type SupportappClient interface {
	Options() supportapp.Options
	CreateSlackChannelConfiguration(ctx context.Context, params *supportapp.CreateSlackChannelConfigurationInput, optFns ...func(*supportapp.Options)) (*supportapp.CreateSlackChannelConfigurationOutput, error)
	DeleteAccountAlias(ctx context.Context, params *supportapp.DeleteAccountAliasInput, optFns ...func(*supportapp.Options)) (*supportapp.DeleteAccountAliasOutput, error)
	DeleteSlackChannelConfiguration(ctx context.Context, params *supportapp.DeleteSlackChannelConfigurationInput, optFns ...func(*supportapp.Options)) (*supportapp.DeleteSlackChannelConfigurationOutput, error)
	DeleteSlackWorkspaceConfiguration(ctx context.Context, params *supportapp.DeleteSlackWorkspaceConfigurationInput, optFns ...func(*supportapp.Options)) (*supportapp.DeleteSlackWorkspaceConfigurationOutput, error)
	GetAccountAlias(ctx context.Context, params *supportapp.GetAccountAliasInput, optFns ...func(*supportapp.Options)) (*supportapp.GetAccountAliasOutput, error)
	ListSlackChannelConfigurations(ctx context.Context, params *supportapp.ListSlackChannelConfigurationsInput, optFns ...func(*supportapp.Options)) (*supportapp.ListSlackChannelConfigurationsOutput, error)
	ListSlackWorkspaceConfigurations(ctx context.Context, params *supportapp.ListSlackWorkspaceConfigurationsInput, optFns ...func(*supportapp.Options)) (*supportapp.ListSlackWorkspaceConfigurationsOutput, error)
	PutAccountAlias(ctx context.Context, params *supportapp.PutAccountAliasInput, optFns ...func(*supportapp.Options)) (*supportapp.PutAccountAliasOutput, error)
	RegisterSlackWorkspaceForOrganization(ctx context.Context, params *supportapp.RegisterSlackWorkspaceForOrganizationInput, optFns ...func(*supportapp.Options)) (*supportapp.RegisterSlackWorkspaceForOrganizationOutput, error)
	UpdateSlackChannelConfiguration(ctx context.Context, params *supportapp.UpdateSlackChannelConfigurationInput, optFns ...func(*supportapp.Options)) (*supportapp.UpdateSlackChannelConfigurationOutput, error)
}
