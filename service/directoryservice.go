// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

// DirectoryserviceClient ...
type DirectoryserviceClient interface {
	Options() directoryservice.Options
	AcceptSharedDirectory(ctx context.Context, params *directoryservice.AcceptSharedDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.AcceptSharedDirectoryOutput, error)
	AddIpRoutes(ctx context.Context, params *directoryservice.AddIpRoutesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.AddIpRoutesOutput, error)
	AddRegion(ctx context.Context, params *directoryservice.AddRegionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.AddRegionOutput, error)
	AddTagsToResource(ctx context.Context, params *directoryservice.AddTagsToResourceInput, optFns ...func(*directoryservice.Options)) (*directoryservice.AddTagsToResourceOutput, error)
	CancelSchemaExtension(ctx context.Context, params *directoryservice.CancelSchemaExtensionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CancelSchemaExtensionOutput, error)
	ConnectDirectory(ctx context.Context, params *directoryservice.ConnectDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ConnectDirectoryOutput, error)
	CreateAlias(ctx context.Context, params *directoryservice.CreateAliasInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateAliasOutput, error)
	CreateComputer(ctx context.Context, params *directoryservice.CreateComputerInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateComputerOutput, error)
	CreateConditionalForwarder(ctx context.Context, params *directoryservice.CreateConditionalForwarderInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateDirectory(ctx context.Context, params *directoryservice.CreateDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateDirectoryOutput, error)
	CreateLogSubscription(ctx context.Context, params *directoryservice.CreateLogSubscriptionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateLogSubscriptionOutput, error)
	CreateMicrosoftAD(ctx context.Context, params *directoryservice.CreateMicrosoftADInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateSnapshot(ctx context.Context, params *directoryservice.CreateSnapshotInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateSnapshotOutput, error)
	CreateTrust(ctx context.Context, params *directoryservice.CreateTrustInput, optFns ...func(*directoryservice.Options)) (*directoryservice.CreateTrustOutput, error)
	DeleteConditionalForwarder(ctx context.Context, params *directoryservice.DeleteConditionalForwarderInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteDirectory(ctx context.Context, params *directoryservice.DeleteDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteLogSubscription(ctx context.Context, params *directoryservice.DeleteLogSubscriptionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeleteLogSubscriptionOutput, error)
	DeleteSnapshot(ctx context.Context, params *directoryservice.DeleteSnapshotInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteTrust(ctx context.Context, params *directoryservice.DeleteTrustInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeleteTrustOutput, error)
	DeregisterCertificate(ctx context.Context, params *directoryservice.DeregisterCertificateInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeregisterCertificateOutput, error)
	DeregisterEventTopic(ctx context.Context, params *directoryservice.DeregisterEventTopicInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DeregisterEventTopicOutput, error)
	DescribeCertificate(ctx context.Context, params *directoryservice.DescribeCertificateInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeCertificateOutput, error)
	DescribeClientAuthenticationSettings(ctx context.Context, params *directoryservice.DescribeClientAuthenticationSettingsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeClientAuthenticationSettingsOutput, error)
	DescribeConditionalForwarders(ctx context.Context, params *directoryservice.DescribeConditionalForwardersInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeDirectories(ctx context.Context, params *directoryservice.DescribeDirectoriesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDomainControllers(ctx context.Context, params *directoryservice.DescribeDomainControllersInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeEventTopics(ctx context.Context, params *directoryservice.DescribeEventTopicsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeLDAPSSettings(ctx context.Context, params *directoryservice.DescribeLDAPSSettingsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeLDAPSSettingsOutput, error)
	DescribeRegions(ctx context.Context, params *directoryservice.DescribeRegionsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeRegionsOutput, error)
	DescribeSettings(ctx context.Context, params *directoryservice.DescribeSettingsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeSettingsOutput, error)
	DescribeSharedDirectories(ctx context.Context, params *directoryservice.DescribeSharedDirectoriesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeSharedDirectoriesOutput, error)
	DescribeSnapshots(ctx context.Context, params *directoryservice.DescribeSnapshotsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeTrusts(ctx context.Context, params *directoryservice.DescribeTrustsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeTrustsOutput, error)
	DescribeUpdateDirectory(ctx context.Context, params *directoryservice.DescribeUpdateDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DescribeUpdateDirectoryOutput, error)
	DisableClientAuthentication(ctx context.Context, params *directoryservice.DisableClientAuthenticationInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DisableClientAuthenticationOutput, error)
	DisableLDAPS(ctx context.Context, params *directoryservice.DisableLDAPSInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DisableLDAPSOutput, error)
	DisableRadius(ctx context.Context, params *directoryservice.DisableRadiusInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DisableRadiusOutput, error)
	DisableSso(ctx context.Context, params *directoryservice.DisableSsoInput, optFns ...func(*directoryservice.Options)) (*directoryservice.DisableSsoOutput, error)
	EnableClientAuthentication(ctx context.Context, params *directoryservice.EnableClientAuthenticationInput, optFns ...func(*directoryservice.Options)) (*directoryservice.EnableClientAuthenticationOutput, error)
	EnableLDAPS(ctx context.Context, params *directoryservice.EnableLDAPSInput, optFns ...func(*directoryservice.Options)) (*directoryservice.EnableLDAPSOutput, error)
	EnableRadius(ctx context.Context, params *directoryservice.EnableRadiusInput, optFns ...func(*directoryservice.Options)) (*directoryservice.EnableRadiusOutput, error)
	EnableSso(ctx context.Context, params *directoryservice.EnableSsoInput, optFns ...func(*directoryservice.Options)) (*directoryservice.EnableSsoOutput, error)
	GetDirectoryLimits(ctx context.Context, params *directoryservice.GetDirectoryLimitsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetSnapshotLimits(ctx context.Context, params *directoryservice.GetSnapshotLimitsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.GetSnapshotLimitsOutput, error)
	ListCertificates(ctx context.Context, params *directoryservice.ListCertificatesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ListCertificatesOutput, error)
	ListIpRoutes(ctx context.Context, params *directoryservice.ListIpRoutesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ListIpRoutesOutput, error)
	ListLogSubscriptions(ctx context.Context, params *directoryservice.ListLogSubscriptionsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ListLogSubscriptionsOutput, error)
	ListSchemaExtensions(ctx context.Context, params *directoryservice.ListSchemaExtensionsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListTagsForResource(ctx context.Context, params *directoryservice.ListTagsForResourceInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ListTagsForResourceOutput, error)
	RegisterCertificate(ctx context.Context, params *directoryservice.RegisterCertificateInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RegisterCertificateOutput, error)
	RegisterEventTopic(ctx context.Context, params *directoryservice.RegisterEventTopicInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RegisterEventTopicOutput, error)
	RejectSharedDirectory(ctx context.Context, params *directoryservice.RejectSharedDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RejectSharedDirectoryOutput, error)
	RemoveIpRoutes(ctx context.Context, params *directoryservice.RemoveIpRoutesInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveRegion(ctx context.Context, params *directoryservice.RemoveRegionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RemoveRegionOutput, error)
	RemoveTagsFromResource(ctx context.Context, params *directoryservice.RemoveTagsFromResourceInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RemoveTagsFromResourceOutput, error)
	ResetUserPassword(ctx context.Context, params *directoryservice.ResetUserPasswordInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ResetUserPasswordOutput, error)
	RestoreFromSnapshot(ctx context.Context, params *directoryservice.RestoreFromSnapshotInput, optFns ...func(*directoryservice.Options)) (*directoryservice.RestoreFromSnapshotOutput, error)
	ShareDirectory(ctx context.Context, params *directoryservice.ShareDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.ShareDirectoryOutput, error)
	StartSchemaExtension(ctx context.Context, params *directoryservice.StartSchemaExtensionInput, optFns ...func(*directoryservice.Options)) (*directoryservice.StartSchemaExtensionOutput, error)
	UnshareDirectory(ctx context.Context, params *directoryservice.UnshareDirectoryInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UnshareDirectoryOutput, error)
	UpdateConditionalForwarder(ctx context.Context, params *directoryservice.UpdateConditionalForwarderInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateDirectorySetup(ctx context.Context, params *directoryservice.UpdateDirectorySetupInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateDirectorySetupOutput, error)
	UpdateNumberOfDomainControllers(ctx context.Context, params *directoryservice.UpdateNumberOfDomainControllersInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateRadius(ctx context.Context, params *directoryservice.UpdateRadiusInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateRadiusOutput, error)
	UpdateSettings(ctx context.Context, params *directoryservice.UpdateSettingsInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateSettingsOutput, error)
	UpdateTrust(ctx context.Context, params *directoryservice.UpdateTrustInput, optFns ...func(*directoryservice.Options)) (*directoryservice.UpdateTrustOutput, error)
	VerifyTrust(ctx context.Context, params *directoryservice.VerifyTrustInput, optFns ...func(*directoryservice.Options)) (*directoryservice.VerifyTrustOutput, error)
}
