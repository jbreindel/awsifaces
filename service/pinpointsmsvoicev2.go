// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
)

// Pinpointsmsvoicev2Client ...
type Pinpointsmsvoicev2Client interface {
	Options() pinpointsmsvoicev2.Options
	AssociateOriginationIdentity(ctx context.Context, params *pinpointsmsvoicev2.AssociateOriginationIdentityInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.AssociateOriginationIdentityOutput, error)
	CreateConfigurationSet(ctx context.Context, params *pinpointsmsvoicev2.CreateConfigurationSetInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateConfigurationSetOutput, error)
	CreateEventDestination(ctx context.Context, params *pinpointsmsvoicev2.CreateEventDestinationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateEventDestinationOutput, error)
	CreateOptOutList(ctx context.Context, params *pinpointsmsvoicev2.CreateOptOutListInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateOptOutListOutput, error)
	CreatePool(ctx context.Context, params *pinpointsmsvoicev2.CreatePoolInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreatePoolOutput, error)
	CreateRegistration(ctx context.Context, params *pinpointsmsvoicev2.CreateRegistrationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateRegistrationOutput, error)
	CreateRegistrationAssociation(ctx context.Context, params *pinpointsmsvoicev2.CreateRegistrationAssociationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateRegistrationAssociationOutput, error)
	CreateRegistrationAttachment(ctx context.Context, params *pinpointsmsvoicev2.CreateRegistrationAttachmentInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateRegistrationAttachmentOutput, error)
	CreateRegistrationVersion(ctx context.Context, params *pinpointsmsvoicev2.CreateRegistrationVersionInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateRegistrationVersionOutput, error)
	CreateVerifiedDestinationNumber(ctx context.Context, params *pinpointsmsvoicev2.CreateVerifiedDestinationNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.CreateVerifiedDestinationNumberOutput, error)
	DeleteConfigurationSet(ctx context.Context, params *pinpointsmsvoicev2.DeleteConfigurationSetInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteConfigurationSetOutput, error)
	DeleteDefaultMessageType(ctx context.Context, params *pinpointsmsvoicev2.DeleteDefaultMessageTypeInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteDefaultMessageTypeOutput, error)
	DeleteDefaultSenderId(ctx context.Context, params *pinpointsmsvoicev2.DeleteDefaultSenderIdInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteDefaultSenderIdOutput, error)
	DeleteEventDestination(ctx context.Context, params *pinpointsmsvoicev2.DeleteEventDestinationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteEventDestinationOutput, error)
	DeleteKeyword(ctx context.Context, params *pinpointsmsvoicev2.DeleteKeywordInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteKeywordOutput, error)
	DeleteOptOutList(ctx context.Context, params *pinpointsmsvoicev2.DeleteOptOutListInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteOptOutListOutput, error)
	DeleteOptedOutNumber(ctx context.Context, params *pinpointsmsvoicev2.DeleteOptedOutNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteOptedOutNumberOutput, error)
	DeletePool(ctx context.Context, params *pinpointsmsvoicev2.DeletePoolInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeletePoolOutput, error)
	DeleteRegistration(ctx context.Context, params *pinpointsmsvoicev2.DeleteRegistrationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteRegistrationOutput, error)
	DeleteRegistrationAttachment(ctx context.Context, params *pinpointsmsvoicev2.DeleteRegistrationAttachmentInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteRegistrationAttachmentOutput, error)
	DeleteRegistrationFieldValue(ctx context.Context, params *pinpointsmsvoicev2.DeleteRegistrationFieldValueInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteRegistrationFieldValueOutput, error)
	DeleteTextMessageSpendLimitOverride(ctx context.Context, params *pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideOutput, error)
	DeleteVerifiedDestinationNumber(ctx context.Context, params *pinpointsmsvoicev2.DeleteVerifiedDestinationNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteVerifiedDestinationNumberOutput, error)
	DeleteVoiceMessageSpendLimitOverride(ctx context.Context, params *pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideOutput, error)
	DescribeAccountAttributes(ctx context.Context, params *pinpointsmsvoicev2.DescribeAccountAttributesInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeAccountAttributesOutput, error)
	DescribeAccountLimits(ctx context.Context, params *pinpointsmsvoicev2.DescribeAccountLimitsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeAccountLimitsOutput, error)
	DescribeConfigurationSets(ctx context.Context, params *pinpointsmsvoicev2.DescribeConfigurationSetsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, error)
	DescribeKeywords(ctx context.Context, params *pinpointsmsvoicev2.DescribeKeywordsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeKeywordsOutput, error)
	DescribeOptOutLists(ctx context.Context, params *pinpointsmsvoicev2.DescribeOptOutListsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeOptOutListsOutput, error)
	DescribeOptedOutNumbers(ctx context.Context, params *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, error)
	DescribePhoneNumbers(ctx context.Context, params *pinpointsmsvoicev2.DescribePhoneNumbersInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribePhoneNumbersOutput, error)
	DescribePools(ctx context.Context, params *pinpointsmsvoicev2.DescribePoolsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribePoolsOutput, error)
	DescribeRegistrationAttachments(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationAttachmentsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationAttachmentsOutput, error)
	DescribeRegistrationFieldDefinitions(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsOutput, error)
	DescribeRegistrationFieldValues(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationFieldValuesInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationFieldValuesOutput, error)
	DescribeRegistrationSectionDefinitions(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsOutput, error)
	DescribeRegistrationTypeDefinitions(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsOutput, error)
	DescribeRegistrationVersions(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationVersionsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationVersionsOutput, error)
	DescribeRegistrations(ctx context.Context, params *pinpointsmsvoicev2.DescribeRegistrationsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeRegistrationsOutput, error)
	DescribeSenderIds(ctx context.Context, params *pinpointsmsvoicev2.DescribeSenderIdsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeSenderIdsOutput, error)
	DescribeSpendLimits(ctx context.Context, params *pinpointsmsvoicev2.DescribeSpendLimitsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeSpendLimitsOutput, error)
	DescribeVerifiedDestinationNumbers(ctx context.Context, params *pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersOutput, error)
	DisassociateOriginationIdentity(ctx context.Context, params *pinpointsmsvoicev2.DisassociateOriginationIdentityInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DisassociateOriginationIdentityOutput, error)
	DiscardRegistrationVersion(ctx context.Context, params *pinpointsmsvoicev2.DiscardRegistrationVersionInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.DiscardRegistrationVersionOutput, error)
	ListPoolOriginationIdentities(ctx context.Context, params *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, error)
	ListRegistrationAssociations(ctx context.Context, params *pinpointsmsvoicev2.ListRegistrationAssociationsInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.ListRegistrationAssociationsOutput, error)
	ListTagsForResource(ctx context.Context, params *pinpointsmsvoicev2.ListTagsForResourceInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.ListTagsForResourceOutput, error)
	PutKeyword(ctx context.Context, params *pinpointsmsvoicev2.PutKeywordInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.PutKeywordOutput, error)
	PutOptedOutNumber(ctx context.Context, params *pinpointsmsvoicev2.PutOptedOutNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.PutOptedOutNumberOutput, error)
	PutRegistrationFieldValue(ctx context.Context, params *pinpointsmsvoicev2.PutRegistrationFieldValueInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.PutRegistrationFieldValueOutput, error)
	ReleasePhoneNumber(ctx context.Context, params *pinpointsmsvoicev2.ReleasePhoneNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.ReleasePhoneNumberOutput, error)
	ReleaseSenderId(ctx context.Context, params *pinpointsmsvoicev2.ReleaseSenderIdInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.ReleaseSenderIdOutput, error)
	RequestPhoneNumber(ctx context.Context, params *pinpointsmsvoicev2.RequestPhoneNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.RequestPhoneNumberOutput, error)
	RequestSenderId(ctx context.Context, params *pinpointsmsvoicev2.RequestSenderIdInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.RequestSenderIdOutput, error)
	SendDestinationNumberVerificationCode(ctx context.Context, params *pinpointsmsvoicev2.SendDestinationNumberVerificationCodeInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SendDestinationNumberVerificationCodeOutput, error)
	SendTextMessage(ctx context.Context, params *pinpointsmsvoicev2.SendTextMessageInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SendTextMessageOutput, error)
	SendVoiceMessage(ctx context.Context, params *pinpointsmsvoicev2.SendVoiceMessageInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SendVoiceMessageOutput, error)
	SetDefaultMessageType(ctx context.Context, params *pinpointsmsvoicev2.SetDefaultMessageTypeInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SetDefaultMessageTypeOutput, error)
	SetDefaultSenderId(ctx context.Context, params *pinpointsmsvoicev2.SetDefaultSenderIdInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SetDefaultSenderIdOutput, error)
	SetTextMessageSpendLimitOverride(ctx context.Context, params *pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideOutput, error)
	SetVoiceMessageSpendLimitOverride(ctx context.Context, params *pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideOutput, error)
	SubmitRegistrationVersion(ctx context.Context, params *pinpointsmsvoicev2.SubmitRegistrationVersionInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.SubmitRegistrationVersionOutput, error)
	TagResource(ctx context.Context, params *pinpointsmsvoicev2.TagResourceInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *pinpointsmsvoicev2.UntagResourceInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.UntagResourceOutput, error)
	UpdateEventDestination(ctx context.Context, params *pinpointsmsvoicev2.UpdateEventDestinationInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.UpdateEventDestinationOutput, error)
	UpdatePhoneNumber(ctx context.Context, params *pinpointsmsvoicev2.UpdatePhoneNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.UpdatePhoneNumberOutput, error)
	UpdatePool(ctx context.Context, params *pinpointsmsvoicev2.UpdatePoolInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.UpdatePoolOutput, error)
	UpdateSenderId(ctx context.Context, params *pinpointsmsvoicev2.UpdateSenderIdInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.UpdateSenderIdOutput, error)
	VerifyDestinationNumber(ctx context.Context, params *pinpointsmsvoicev2.VerifyDestinationNumberInput, optFns ...func(*pinpointsmsvoicev2.Options)) (*pinpointsmsvoicev2.VerifyDestinationNumberOutput, error)
}
