// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

// WorklinkClient ...
type WorklinkClient interface {
	Options() worklink.Options
	AssociateDomain(ctx context.Context, params *worklink.AssociateDomainInput, optFns ...func(*worklink.Options)) (*worklink.AssociateDomainOutput, error)
	AssociateWebsiteAuthorizationProvider(ctx context.Context, params *worklink.AssociateWebsiteAuthorizationProviderInput, optFns ...func(*worklink.Options)) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error)
	AssociateWebsiteCertificateAuthority(ctx context.Context, params *worklink.AssociateWebsiteCertificateAuthorityInput, optFns ...func(*worklink.Options)) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error)
	CreateFleet(ctx context.Context, params *worklink.CreateFleetInput, optFns ...func(*worklink.Options)) (*worklink.CreateFleetOutput, error)
	DeleteFleet(ctx context.Context, params *worklink.DeleteFleetInput, optFns ...func(*worklink.Options)) (*worklink.DeleteFleetOutput, error)
	DescribeAuditStreamConfiguration(ctx context.Context, params *worklink.DescribeAuditStreamConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.DescribeAuditStreamConfigurationOutput, error)
	DescribeCompanyNetworkConfiguration(ctx context.Context, params *worklink.DescribeCompanyNetworkConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.DescribeCompanyNetworkConfigurationOutput, error)
	DescribeDevice(ctx context.Context, params *worklink.DescribeDeviceInput, optFns ...func(*worklink.Options)) (*worklink.DescribeDeviceOutput, error)
	DescribeDevicePolicyConfiguration(ctx context.Context, params *worklink.DescribeDevicePolicyConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.DescribeDevicePolicyConfigurationOutput, error)
	DescribeDomain(ctx context.Context, params *worklink.DescribeDomainInput, optFns ...func(*worklink.Options)) (*worklink.DescribeDomainOutput, error)
	DescribeFleetMetadata(ctx context.Context, params *worklink.DescribeFleetMetadataInput, optFns ...func(*worklink.Options)) (*worklink.DescribeFleetMetadataOutput, error)
	DescribeIdentityProviderConfiguration(ctx context.Context, params *worklink.DescribeIdentityProviderConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.DescribeIdentityProviderConfigurationOutput, error)
	DescribeWebsiteCertificateAuthority(ctx context.Context, params *worklink.DescribeWebsiteCertificateAuthorityInput, optFns ...func(*worklink.Options)) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error)
	DisassociateDomain(ctx context.Context, params *worklink.DisassociateDomainInput, optFns ...func(*worklink.Options)) (*worklink.DisassociateDomainOutput, error)
	DisassociateWebsiteAuthorizationProvider(ctx context.Context, params *worklink.DisassociateWebsiteAuthorizationProviderInput, optFns ...func(*worklink.Options)) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error)
	DisassociateWebsiteCertificateAuthority(ctx context.Context, params *worklink.DisassociateWebsiteCertificateAuthorityInput, optFns ...func(*worklink.Options)) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error)
	ListDevices(ctx context.Context, params *worklink.ListDevicesInput, optFns ...func(*worklink.Options)) (*worklink.ListDevicesOutput, error)
	ListDomains(ctx context.Context, params *worklink.ListDomainsInput, optFns ...func(*worklink.Options)) (*worklink.ListDomainsOutput, error)
	ListFleets(ctx context.Context, params *worklink.ListFleetsInput, optFns ...func(*worklink.Options)) (*worklink.ListFleetsOutput, error)
	ListTagsForResource(ctx context.Context, params *worklink.ListTagsForResourceInput, optFns ...func(*worklink.Options)) (*worklink.ListTagsForResourceOutput, error)
	ListWebsiteAuthorizationProviders(ctx context.Context, params *worklink.ListWebsiteAuthorizationProvidersInput, optFns ...func(*worklink.Options)) (*worklink.ListWebsiteAuthorizationProvidersOutput, error)
	ListWebsiteCertificateAuthorities(ctx context.Context, params *worklink.ListWebsiteCertificateAuthoritiesInput, optFns ...func(*worklink.Options)) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error)
	RestoreDomainAccess(ctx context.Context, params *worklink.RestoreDomainAccessInput, optFns ...func(*worklink.Options)) (*worklink.RestoreDomainAccessOutput, error)
	RevokeDomainAccess(ctx context.Context, params *worklink.RevokeDomainAccessInput, optFns ...func(*worklink.Options)) (*worklink.RevokeDomainAccessOutput, error)
	SignOutUser(ctx context.Context, params *worklink.SignOutUserInput, optFns ...func(*worklink.Options)) (*worklink.SignOutUserOutput, error)
	TagResource(ctx context.Context, params *worklink.TagResourceInput, optFns ...func(*worklink.Options)) (*worklink.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *worklink.UntagResourceInput, optFns ...func(*worklink.Options)) (*worklink.UntagResourceOutput, error)
	UpdateAuditStreamConfiguration(ctx context.Context, params *worklink.UpdateAuditStreamConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.UpdateAuditStreamConfigurationOutput, error)
	UpdateCompanyNetworkConfiguration(ctx context.Context, params *worklink.UpdateCompanyNetworkConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.UpdateCompanyNetworkConfigurationOutput, error)
	UpdateDevicePolicyConfiguration(ctx context.Context, params *worklink.UpdateDevicePolicyConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.UpdateDevicePolicyConfigurationOutput, error)
	UpdateDomainMetadata(ctx context.Context, params *worklink.UpdateDomainMetadataInput, optFns ...func(*worklink.Options)) (*worklink.UpdateDomainMetadataOutput, error)
	UpdateFleetMetadata(ctx context.Context, params *worklink.UpdateFleetMetadataInput, optFns ...func(*worklink.Options)) (*worklink.UpdateFleetMetadataOutput, error)
	UpdateIdentityProviderConfiguration(ctx context.Context, params *worklink.UpdateIdentityProviderConfigurationInput, optFns ...func(*worklink.Options)) (*worklink.UpdateIdentityProviderConfigurationOutput, error)
}
