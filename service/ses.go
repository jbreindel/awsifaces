// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

// SesClient ...
type SesClient interface {
	Options() ses.Options
	CloneReceiptRuleSet(ctx context.Context, params *ses.CloneReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.CloneReceiptRuleSetOutput, error)
	CreateConfigurationSet(ctx context.Context, params *ses.CreateConfigurationSetInput, optFns ...func(*ses.Options)) (*ses.CreateConfigurationSetOutput, error)
	CreateConfigurationSetEventDestination(ctx context.Context, params *ses.CreateConfigurationSetEventDestinationInput, optFns ...func(*ses.Options)) (*ses.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetTrackingOptions(ctx context.Context, params *ses.CreateConfigurationSetTrackingOptionsInput, optFns ...func(*ses.Options)) (*ses.CreateConfigurationSetTrackingOptionsOutput, error)
	CreateCustomVerificationEmailTemplate(ctx context.Context, params *ses.CreateCustomVerificationEmailTemplateInput, optFns ...func(*ses.Options)) (*ses.CreateCustomVerificationEmailTemplateOutput, error)
	CreateReceiptFilter(ctx context.Context, params *ses.CreateReceiptFilterInput, optFns ...func(*ses.Options)) (*ses.CreateReceiptFilterOutput, error)
	CreateReceiptRule(ctx context.Context, params *ses.CreateReceiptRuleInput, optFns ...func(*ses.Options)) (*ses.CreateReceiptRuleOutput, error)
	CreateReceiptRuleSet(ctx context.Context, params *ses.CreateReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.CreateReceiptRuleSetOutput, error)
	CreateTemplate(ctx context.Context, params *ses.CreateTemplateInput, optFns ...func(*ses.Options)) (*ses.CreateTemplateOutput, error)
	DeleteConfigurationSet(ctx context.Context, params *ses.DeleteConfigurationSetInput, optFns ...func(*ses.Options)) (*ses.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetEventDestination(ctx context.Context, params *ses.DeleteConfigurationSetEventDestinationInput, optFns ...func(*ses.Options)) (*ses.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetTrackingOptions(ctx context.Context, params *ses.DeleteConfigurationSetTrackingOptionsInput, optFns ...func(*ses.Options)) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error)
	DeleteCustomVerificationEmailTemplate(ctx context.Context, params *ses.DeleteCustomVerificationEmailTemplateInput, optFns ...func(*ses.Options)) (*ses.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteIdentity(ctx context.Context, params *ses.DeleteIdentityInput, optFns ...func(*ses.Options)) (*ses.DeleteIdentityOutput, error)
	DeleteIdentityPolicy(ctx context.Context, params *ses.DeleteIdentityPolicyInput, optFns ...func(*ses.Options)) (*ses.DeleteIdentityPolicyOutput, error)
	DeleteReceiptFilter(ctx context.Context, params *ses.DeleteReceiptFilterInput, optFns ...func(*ses.Options)) (*ses.DeleteReceiptFilterOutput, error)
	DeleteReceiptRule(ctx context.Context, params *ses.DeleteReceiptRuleInput, optFns ...func(*ses.Options)) (*ses.DeleteReceiptRuleOutput, error)
	DeleteReceiptRuleSet(ctx context.Context, params *ses.DeleteReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.DeleteReceiptRuleSetOutput, error)
	DeleteTemplate(ctx context.Context, params *ses.DeleteTemplateInput, optFns ...func(*ses.Options)) (*ses.DeleteTemplateOutput, error)
	DeleteVerifiedEmailAddress(ctx context.Context, params *ses.DeleteVerifiedEmailAddressInput, optFns ...func(*ses.Options)) (*ses.DeleteVerifiedEmailAddressOutput, error)
	DescribeActiveReceiptRuleSet(ctx context.Context, params *ses.DescribeActiveReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.DescribeActiveReceiptRuleSetOutput, error)
	DescribeConfigurationSet(ctx context.Context, params *ses.DescribeConfigurationSetInput, optFns ...func(*ses.Options)) (*ses.DescribeConfigurationSetOutput, error)
	DescribeReceiptRule(ctx context.Context, params *ses.DescribeReceiptRuleInput, optFns ...func(*ses.Options)) (*ses.DescribeReceiptRuleOutput, error)
	DescribeReceiptRuleSet(ctx context.Context, params *ses.DescribeReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.DescribeReceiptRuleSetOutput, error)
	GetAccountSendingEnabled(ctx context.Context, params *ses.GetAccountSendingEnabledInput, optFns ...func(*ses.Options)) (*ses.GetAccountSendingEnabledOutput, error)
	GetCustomVerificationEmailTemplate(ctx context.Context, params *ses.GetCustomVerificationEmailTemplateInput, optFns ...func(*ses.Options)) (*ses.GetCustomVerificationEmailTemplateOutput, error)
	GetIdentityDkimAttributes(ctx context.Context, params *ses.GetIdentityDkimAttributesInput, optFns ...func(*ses.Options)) (*ses.GetIdentityDkimAttributesOutput, error)
	GetIdentityMailFromDomainAttributes(ctx context.Context, params *ses.GetIdentityMailFromDomainAttributesInput, optFns ...func(*ses.Options)) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
	GetIdentityNotificationAttributes(ctx context.Context, params *ses.GetIdentityNotificationAttributesInput, optFns ...func(*ses.Options)) (*ses.GetIdentityNotificationAttributesOutput, error)
	GetIdentityPolicies(ctx context.Context, params *ses.GetIdentityPoliciesInput, optFns ...func(*ses.Options)) (*ses.GetIdentityPoliciesOutput, error)
	GetIdentityVerificationAttributes(ctx context.Context, params *ses.GetIdentityVerificationAttributesInput, optFns ...func(*ses.Options)) (*ses.GetIdentityVerificationAttributesOutput, error)
	GetSendQuota(ctx context.Context, params *ses.GetSendQuotaInput, optFns ...func(*ses.Options)) (*ses.GetSendQuotaOutput, error)
	GetSendStatistics(ctx context.Context, params *ses.GetSendStatisticsInput, optFns ...func(*ses.Options)) (*ses.GetSendStatisticsOutput, error)
	GetTemplate(ctx context.Context, params *ses.GetTemplateInput, optFns ...func(*ses.Options)) (*ses.GetTemplateOutput, error)
	ListConfigurationSets(ctx context.Context, params *ses.ListConfigurationSetsInput, optFns ...func(*ses.Options)) (*ses.ListConfigurationSetsOutput, error)
	ListCustomVerificationEmailTemplates(ctx context.Context, params *ses.ListCustomVerificationEmailTemplatesInput, optFns ...func(*ses.Options)) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
	ListIdentities(ctx context.Context, params *ses.ListIdentitiesInput, optFns ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error)
	ListIdentityPolicies(ctx context.Context, params *ses.ListIdentityPoliciesInput, optFns ...func(*ses.Options)) (*ses.ListIdentityPoliciesOutput, error)
	ListReceiptFilters(ctx context.Context, params *ses.ListReceiptFiltersInput, optFns ...func(*ses.Options)) (*ses.ListReceiptFiltersOutput, error)
	ListReceiptRuleSets(ctx context.Context, params *ses.ListReceiptRuleSetsInput, optFns ...func(*ses.Options)) (*ses.ListReceiptRuleSetsOutput, error)
	ListTemplates(ctx context.Context, params *ses.ListTemplatesInput, optFns ...func(*ses.Options)) (*ses.ListTemplatesOutput, error)
	ListVerifiedEmailAddresses(ctx context.Context, params *ses.ListVerifiedEmailAddressesInput, optFns ...func(*ses.Options)) (*ses.ListVerifiedEmailAddressesOutput, error)
	PutConfigurationSetDeliveryOptions(ctx context.Context, params *ses.PutConfigurationSetDeliveryOptionsInput, optFns ...func(*ses.Options)) (*ses.PutConfigurationSetDeliveryOptionsOutput, error)
	PutIdentityPolicy(ctx context.Context, params *ses.PutIdentityPolicyInput, optFns ...func(*ses.Options)) (*ses.PutIdentityPolicyOutput, error)
	ReorderReceiptRuleSet(ctx context.Context, params *ses.ReorderReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.ReorderReceiptRuleSetOutput, error)
	SendBounce(ctx context.Context, params *ses.SendBounceInput, optFns ...func(*ses.Options)) (*ses.SendBounceOutput, error)
	SendBulkTemplatedEmail(ctx context.Context, params *ses.SendBulkTemplatedEmailInput, optFns ...func(*ses.Options)) (*ses.SendBulkTemplatedEmailOutput, error)
	SendCustomVerificationEmail(ctx context.Context, params *ses.SendCustomVerificationEmailInput, optFns ...func(*ses.Options)) (*ses.SendCustomVerificationEmailOutput, error)
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
	SendRawEmail(ctx context.Context, params *ses.SendRawEmailInput, optFns ...func(*ses.Options)) (*ses.SendRawEmailOutput, error)
	SendTemplatedEmail(ctx context.Context, params *ses.SendTemplatedEmailInput, optFns ...func(*ses.Options)) (*ses.SendTemplatedEmailOutput, error)
	SetActiveReceiptRuleSet(ctx context.Context, params *ses.SetActiveReceiptRuleSetInput, optFns ...func(*ses.Options)) (*ses.SetActiveReceiptRuleSetOutput, error)
	SetIdentityDkimEnabled(ctx context.Context, params *ses.SetIdentityDkimEnabledInput, optFns ...func(*ses.Options)) (*ses.SetIdentityDkimEnabledOutput, error)
	SetIdentityFeedbackForwardingEnabled(ctx context.Context, params *ses.SetIdentityFeedbackForwardingEnabledInput, optFns ...func(*ses.Options)) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)
	SetIdentityHeadersInNotificationsEnabled(ctx context.Context, params *ses.SetIdentityHeadersInNotificationsEnabledInput, optFns ...func(*ses.Options)) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)
	SetIdentityMailFromDomain(ctx context.Context, params *ses.SetIdentityMailFromDomainInput, optFns ...func(*ses.Options)) (*ses.SetIdentityMailFromDomainOutput, error)
	SetIdentityNotificationTopic(ctx context.Context, params *ses.SetIdentityNotificationTopicInput, optFns ...func(*ses.Options)) (*ses.SetIdentityNotificationTopicOutput, error)
	SetReceiptRulePosition(ctx context.Context, params *ses.SetReceiptRulePositionInput, optFns ...func(*ses.Options)) (*ses.SetReceiptRulePositionOutput, error)
	TestRenderTemplate(ctx context.Context, params *ses.TestRenderTemplateInput, optFns ...func(*ses.Options)) (*ses.TestRenderTemplateOutput, error)
	UpdateAccountSendingEnabled(ctx context.Context, params *ses.UpdateAccountSendingEnabledInput, optFns ...func(*ses.Options)) (*ses.UpdateAccountSendingEnabledOutput, error)
	UpdateConfigurationSetEventDestination(ctx context.Context, params *ses.UpdateConfigurationSetEventDestinationInput, optFns ...func(*ses.Options)) (*ses.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, params *ses.UpdateConfigurationSetReputationMetricsEnabledInput, optFns ...func(*ses.Options)) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error)
	UpdateConfigurationSetSendingEnabled(ctx context.Context, params *ses.UpdateConfigurationSetSendingEnabledInput, optFns ...func(*ses.Options)) (*ses.UpdateConfigurationSetSendingEnabledOutput, error)
	UpdateConfigurationSetTrackingOptions(ctx context.Context, params *ses.UpdateConfigurationSetTrackingOptionsInput, optFns ...func(*ses.Options)) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error)
	UpdateCustomVerificationEmailTemplate(ctx context.Context, params *ses.UpdateCustomVerificationEmailTemplateInput, optFns ...func(*ses.Options)) (*ses.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateReceiptRule(ctx context.Context, params *ses.UpdateReceiptRuleInput, optFns ...func(*ses.Options)) (*ses.UpdateReceiptRuleOutput, error)
	UpdateTemplate(ctx context.Context, params *ses.UpdateTemplateInput, optFns ...func(*ses.Options)) (*ses.UpdateTemplateOutput, error)
	VerifyDomainDkim(ctx context.Context, params *ses.VerifyDomainDkimInput, optFns ...func(*ses.Options)) (*ses.VerifyDomainDkimOutput, error)
	VerifyDomainIdentity(ctx context.Context, params *ses.VerifyDomainIdentityInput, optFns ...func(*ses.Options)) (*ses.VerifyDomainIdentityOutput, error)
	VerifyEmailAddress(ctx context.Context, params *ses.VerifyEmailAddressInput, optFns ...func(*ses.Options)) (*ses.VerifyEmailAddressOutput, error)
	VerifyEmailIdentity(ctx context.Context, params *ses.VerifyEmailIdentityInput, optFns ...func(*ses.Options)) (*ses.VerifyEmailIdentityOutput, error)
}
