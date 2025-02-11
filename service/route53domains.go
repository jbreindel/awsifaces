// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
)

// Route53domainsClient ...
type Route53domainsClient interface {
	Options() route53domains.Options
	AcceptDomainTransferFromAnotherAwsAccount(ctx context.Context, params *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput, optFns ...func(*route53domains.Options)) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error)
	AssociateDelegationSignerToDomain(ctx context.Context, params *route53domains.AssociateDelegationSignerToDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.AssociateDelegationSignerToDomainOutput, error)
	CancelDomainTransferToAnotherAwsAccount(ctx context.Context, params *route53domains.CancelDomainTransferToAnotherAwsAccountInput, optFns ...func(*route53domains.Options)) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error)
	CheckDomainAvailability(ctx context.Context, params *route53domains.CheckDomainAvailabilityInput, optFns ...func(*route53domains.Options)) (*route53domains.CheckDomainAvailabilityOutput, error)
	CheckDomainTransferability(ctx context.Context, params *route53domains.CheckDomainTransferabilityInput, optFns ...func(*route53domains.Options)) (*route53domains.CheckDomainTransferabilityOutput, error)
	DeleteDomain(ctx context.Context, params *route53domains.DeleteDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.DeleteDomainOutput, error)
	DeleteTagsForDomain(ctx context.Context, params *route53domains.DeleteTagsForDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.DeleteTagsForDomainOutput, error)
	DisableDomainAutoRenew(ctx context.Context, params *route53domains.DisableDomainAutoRenewInput, optFns ...func(*route53domains.Options)) (*route53domains.DisableDomainAutoRenewOutput, error)
	DisableDomainTransferLock(ctx context.Context, params *route53domains.DisableDomainTransferLockInput, optFns ...func(*route53domains.Options)) (*route53domains.DisableDomainTransferLockOutput, error)
	DisassociateDelegationSignerFromDomain(ctx context.Context, params *route53domains.DisassociateDelegationSignerFromDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.DisassociateDelegationSignerFromDomainOutput, error)
	EnableDomainAutoRenew(ctx context.Context, params *route53domains.EnableDomainAutoRenewInput, optFns ...func(*route53domains.Options)) (*route53domains.EnableDomainAutoRenewOutput, error)
	EnableDomainTransferLock(ctx context.Context, params *route53domains.EnableDomainTransferLockInput, optFns ...func(*route53domains.Options)) (*route53domains.EnableDomainTransferLockOutput, error)
	GetContactReachabilityStatus(ctx context.Context, params *route53domains.GetContactReachabilityStatusInput, optFns ...func(*route53domains.Options)) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetDomainDetail(ctx context.Context, params *route53domains.GetDomainDetailInput, optFns ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error)
	GetDomainSuggestions(ctx context.Context, params *route53domains.GetDomainSuggestionsInput, optFns ...func(*route53domains.Options)) (*route53domains.GetDomainSuggestionsOutput, error)
	GetOperationDetail(ctx context.Context, params *route53domains.GetOperationDetailInput, optFns ...func(*route53domains.Options)) (*route53domains.GetOperationDetailOutput, error)
	ListDomains(ctx context.Context, params *route53domains.ListDomainsInput, optFns ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error)
	ListOperations(ctx context.Context, params *route53domains.ListOperationsInput, optFns ...func(*route53domains.Options)) (*route53domains.ListOperationsOutput, error)
	ListPrices(ctx context.Context, params *route53domains.ListPricesInput, optFns ...func(*route53domains.Options)) (*route53domains.ListPricesOutput, error)
	ListTagsForDomain(ctx context.Context, params *route53domains.ListTagsForDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error)
	PushDomain(ctx context.Context, params *route53domains.PushDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.PushDomainOutput, error)
	RegisterDomain(ctx context.Context, params *route53domains.RegisterDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.RegisterDomainOutput, error)
	RejectDomainTransferFromAnotherAwsAccount(ctx context.Context, params *route53domains.RejectDomainTransferFromAnotherAwsAccountInput, optFns ...func(*route53domains.Options)) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error)
	RenewDomain(ctx context.Context, params *route53domains.RenewDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.RenewDomainOutput, error)
	ResendContactReachabilityEmail(ctx context.Context, params *route53domains.ResendContactReachabilityEmailInput, optFns ...func(*route53domains.Options)) (*route53domains.ResendContactReachabilityEmailOutput, error)
	ResendOperationAuthorization(ctx context.Context, params *route53domains.ResendOperationAuthorizationInput, optFns ...func(*route53domains.Options)) (*route53domains.ResendOperationAuthorizationOutput, error)
	RetrieveDomainAuthCode(ctx context.Context, params *route53domains.RetrieveDomainAuthCodeInput, optFns ...func(*route53domains.Options)) (*route53domains.RetrieveDomainAuthCodeOutput, error)
	TransferDomain(ctx context.Context, params *route53domains.TransferDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.TransferDomainOutput, error)
	TransferDomainToAnotherAwsAccount(ctx context.Context, params *route53domains.TransferDomainToAnotherAwsAccountInput, optFns ...func(*route53domains.Options)) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error)
	UpdateDomainContact(ctx context.Context, params *route53domains.UpdateDomainContactInput, optFns ...func(*route53domains.Options)) (*route53domains.UpdateDomainContactOutput, error)
	UpdateDomainContactPrivacy(ctx context.Context, params *route53domains.UpdateDomainContactPrivacyInput, optFns ...func(*route53domains.Options)) (*route53domains.UpdateDomainContactPrivacyOutput, error)
	UpdateDomainNameservers(ctx context.Context, params *route53domains.UpdateDomainNameserversInput, optFns ...func(*route53domains.Options)) (*route53domains.UpdateDomainNameserversOutput, error)
	UpdateTagsForDomain(ctx context.Context, params *route53domains.UpdateTagsForDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.UpdateTagsForDomainOutput, error)
	ViewBilling(ctx context.Context, params *route53domains.ViewBillingInput, optFns ...func(*route53domains.Options)) (*route53domains.ViewBillingOutput, error)
}
