// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

// AppstreamClient ...
type AppstreamClient interface {
	Options() appstream.Options
	AssociateAppBlockBuilderAppBlock(ctx context.Context, params *appstream.AssociateAppBlockBuilderAppBlockInput, optFns ...func(*appstream.Options)) (*appstream.AssociateAppBlockBuilderAppBlockOutput, error)
	AssociateApplicationFleet(ctx context.Context, params *appstream.AssociateApplicationFleetInput, optFns ...func(*appstream.Options)) (*appstream.AssociateApplicationFleetOutput, error)
	AssociateApplicationToEntitlement(ctx context.Context, params *appstream.AssociateApplicationToEntitlementInput, optFns ...func(*appstream.Options)) (*appstream.AssociateApplicationToEntitlementOutput, error)
	AssociateFleet(ctx context.Context, params *appstream.AssociateFleetInput, optFns ...func(*appstream.Options)) (*appstream.AssociateFleetOutput, error)
	BatchAssociateUserStack(ctx context.Context, params *appstream.BatchAssociateUserStackInput, optFns ...func(*appstream.Options)) (*appstream.BatchAssociateUserStackOutput, error)
	BatchDisassociateUserStack(ctx context.Context, params *appstream.BatchDisassociateUserStackInput, optFns ...func(*appstream.Options)) (*appstream.BatchDisassociateUserStackOutput, error)
	CopyImage(ctx context.Context, params *appstream.CopyImageInput, optFns ...func(*appstream.Options)) (*appstream.CopyImageOutput, error)
	CreateAppBlock(ctx context.Context, params *appstream.CreateAppBlockInput, optFns ...func(*appstream.Options)) (*appstream.CreateAppBlockOutput, error)
	CreateAppBlockBuilder(ctx context.Context, params *appstream.CreateAppBlockBuilderInput, optFns ...func(*appstream.Options)) (*appstream.CreateAppBlockBuilderOutput, error)
	CreateAppBlockBuilderStreamingURL(ctx context.Context, params *appstream.CreateAppBlockBuilderStreamingURLInput, optFns ...func(*appstream.Options)) (*appstream.CreateAppBlockBuilderStreamingURLOutput, error)
	CreateApplication(ctx context.Context, params *appstream.CreateApplicationInput, optFns ...func(*appstream.Options)) (*appstream.CreateApplicationOutput, error)
	CreateDirectoryConfig(ctx context.Context, params *appstream.CreateDirectoryConfigInput, optFns ...func(*appstream.Options)) (*appstream.CreateDirectoryConfigOutput, error)
	CreateEntitlement(ctx context.Context, params *appstream.CreateEntitlementInput, optFns ...func(*appstream.Options)) (*appstream.CreateEntitlementOutput, error)
	CreateFleet(ctx context.Context, params *appstream.CreateFleetInput, optFns ...func(*appstream.Options)) (*appstream.CreateFleetOutput, error)
	CreateImageBuilder(ctx context.Context, params *appstream.CreateImageBuilderInput, optFns ...func(*appstream.Options)) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderStreamingURL(ctx context.Context, params *appstream.CreateImageBuilderStreamingURLInput, optFns ...func(*appstream.Options)) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateStack(ctx context.Context, params *appstream.CreateStackInput, optFns ...func(*appstream.Options)) (*appstream.CreateStackOutput, error)
	CreateStreamingURL(ctx context.Context, params *appstream.CreateStreamingURLInput, optFns ...func(*appstream.Options)) (*appstream.CreateStreamingURLOutput, error)
	CreateUpdatedImage(ctx context.Context, params *appstream.CreateUpdatedImageInput, optFns ...func(*appstream.Options)) (*appstream.CreateUpdatedImageOutput, error)
	CreateUsageReportSubscription(ctx context.Context, params *appstream.CreateUsageReportSubscriptionInput, optFns ...func(*appstream.Options)) (*appstream.CreateUsageReportSubscriptionOutput, error)
	CreateUser(ctx context.Context, params *appstream.CreateUserInput, optFns ...func(*appstream.Options)) (*appstream.CreateUserOutput, error)
	DeleteAppBlock(ctx context.Context, params *appstream.DeleteAppBlockInput, optFns ...func(*appstream.Options)) (*appstream.DeleteAppBlockOutput, error)
	DeleteAppBlockBuilder(ctx context.Context, params *appstream.DeleteAppBlockBuilderInput, optFns ...func(*appstream.Options)) (*appstream.DeleteAppBlockBuilderOutput, error)
	DeleteApplication(ctx context.Context, params *appstream.DeleteApplicationInput, optFns ...func(*appstream.Options)) (*appstream.DeleteApplicationOutput, error)
	DeleteDirectoryConfig(ctx context.Context, params *appstream.DeleteDirectoryConfigInput, optFns ...func(*appstream.Options)) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteEntitlement(ctx context.Context, params *appstream.DeleteEntitlementInput, optFns ...func(*appstream.Options)) (*appstream.DeleteEntitlementOutput, error)
	DeleteFleet(ctx context.Context, params *appstream.DeleteFleetInput, optFns ...func(*appstream.Options)) (*appstream.DeleteFleetOutput, error)
	DeleteImage(ctx context.Context, params *appstream.DeleteImageInput, optFns ...func(*appstream.Options)) (*appstream.DeleteImageOutput, error)
	DeleteImageBuilder(ctx context.Context, params *appstream.DeleteImageBuilderInput, optFns ...func(*appstream.Options)) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImagePermissions(ctx context.Context, params *appstream.DeleteImagePermissionsInput, optFns ...func(*appstream.Options)) (*appstream.DeleteImagePermissionsOutput, error)
	DeleteStack(ctx context.Context, params *appstream.DeleteStackInput, optFns ...func(*appstream.Options)) (*appstream.DeleteStackOutput, error)
	DeleteUsageReportSubscription(ctx context.Context, params *appstream.DeleteUsageReportSubscriptionInput, optFns ...func(*appstream.Options)) (*appstream.DeleteUsageReportSubscriptionOutput, error)
	DeleteUser(ctx context.Context, params *appstream.DeleteUserInput, optFns ...func(*appstream.Options)) (*appstream.DeleteUserOutput, error)
	DescribeAppBlockBuilderAppBlockAssociations(ctx context.Context, params *appstream.DescribeAppBlockBuilderAppBlockAssociationsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, error)
	DescribeAppBlockBuilders(ctx context.Context, params *appstream.DescribeAppBlockBuildersInput, optFns ...func(*appstream.Options)) (*appstream.DescribeAppBlockBuildersOutput, error)
	DescribeAppBlocks(ctx context.Context, params *appstream.DescribeAppBlocksInput, optFns ...func(*appstream.Options)) (*appstream.DescribeAppBlocksOutput, error)
	DescribeApplicationFleetAssociations(ctx context.Context, params *appstream.DescribeApplicationFleetAssociationsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeApplicationFleetAssociationsOutput, error)
	DescribeApplications(ctx context.Context, params *appstream.DescribeApplicationsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeApplicationsOutput, error)
	DescribeDirectoryConfigs(ctx context.Context, params *appstream.DescribeDirectoryConfigsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeEntitlements(ctx context.Context, params *appstream.DescribeEntitlementsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeEntitlementsOutput, error)
	DescribeFleets(ctx context.Context, params *appstream.DescribeFleetsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeFleetsOutput, error)
	DescribeImageBuilders(ctx context.Context, params *appstream.DescribeImageBuildersInput, optFns ...func(*appstream.Options)) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImagePermissions(ctx context.Context, params *appstream.DescribeImagePermissionsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImages(ctx context.Context, params *appstream.DescribeImagesInput, optFns ...func(*appstream.Options)) (*appstream.DescribeImagesOutput, error)
	DescribeSessions(ctx context.Context, params *appstream.DescribeSessionsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeSessionsOutput, error)
	DescribeStacks(ctx context.Context, params *appstream.DescribeStacksInput, optFns ...func(*appstream.Options)) (*appstream.DescribeStacksOutput, error)
	DescribeUsageReportSubscriptions(ctx context.Context, params *appstream.DescribeUsageReportSubscriptionsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUserStackAssociations(ctx context.Context, params *appstream.DescribeUserStackAssociationsInput, optFns ...func(*appstream.Options)) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUsers(ctx context.Context, params *appstream.DescribeUsersInput, optFns ...func(*appstream.Options)) (*appstream.DescribeUsersOutput, error)
	DisableUser(ctx context.Context, params *appstream.DisableUserInput, optFns ...func(*appstream.Options)) (*appstream.DisableUserOutput, error)
	DisassociateAppBlockBuilderAppBlock(ctx context.Context, params *appstream.DisassociateAppBlockBuilderAppBlockInput, optFns ...func(*appstream.Options)) (*appstream.DisassociateAppBlockBuilderAppBlockOutput, error)
	DisassociateApplicationFleet(ctx context.Context, params *appstream.DisassociateApplicationFleetInput, optFns ...func(*appstream.Options)) (*appstream.DisassociateApplicationFleetOutput, error)
	DisassociateApplicationFromEntitlement(ctx context.Context, params *appstream.DisassociateApplicationFromEntitlementInput, optFns ...func(*appstream.Options)) (*appstream.DisassociateApplicationFromEntitlementOutput, error)
	DisassociateFleet(ctx context.Context, params *appstream.DisassociateFleetInput, optFns ...func(*appstream.Options)) (*appstream.DisassociateFleetOutput, error)
	EnableUser(ctx context.Context, params *appstream.EnableUserInput, optFns ...func(*appstream.Options)) (*appstream.EnableUserOutput, error)
	ExpireSession(ctx context.Context, params *appstream.ExpireSessionInput, optFns ...func(*appstream.Options)) (*appstream.ExpireSessionOutput, error)
	ListAssociatedFleets(ctx context.Context, params *appstream.ListAssociatedFleetsInput, optFns ...func(*appstream.Options)) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedStacks(ctx context.Context, params *appstream.ListAssociatedStacksInput, optFns ...func(*appstream.Options)) (*appstream.ListAssociatedStacksOutput, error)
	ListEntitledApplications(ctx context.Context, params *appstream.ListEntitledApplicationsInput, optFns ...func(*appstream.Options)) (*appstream.ListEntitledApplicationsOutput, error)
	ListTagsForResource(ctx context.Context, params *appstream.ListTagsForResourceInput, optFns ...func(*appstream.Options)) (*appstream.ListTagsForResourceOutput, error)
	StartAppBlockBuilder(ctx context.Context, params *appstream.StartAppBlockBuilderInput, optFns ...func(*appstream.Options)) (*appstream.StartAppBlockBuilderOutput, error)
	StartFleet(ctx context.Context, params *appstream.StartFleetInput, optFns ...func(*appstream.Options)) (*appstream.StartFleetOutput, error)
	StartImageBuilder(ctx context.Context, params *appstream.StartImageBuilderInput, optFns ...func(*appstream.Options)) (*appstream.StartImageBuilderOutput, error)
	StopAppBlockBuilder(ctx context.Context, params *appstream.StopAppBlockBuilderInput, optFns ...func(*appstream.Options)) (*appstream.StopAppBlockBuilderOutput, error)
	StopFleet(ctx context.Context, params *appstream.StopFleetInput, optFns ...func(*appstream.Options)) (*appstream.StopFleetOutput, error)
	StopImageBuilder(ctx context.Context, params *appstream.StopImageBuilderInput, optFns ...func(*appstream.Options)) (*appstream.StopImageBuilderOutput, error)
	TagResource(ctx context.Context, params *appstream.TagResourceInput, optFns ...func(*appstream.Options)) (*appstream.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *appstream.UntagResourceInput, optFns ...func(*appstream.Options)) (*appstream.UntagResourceOutput, error)
	UpdateAppBlockBuilder(ctx context.Context, params *appstream.UpdateAppBlockBuilderInput, optFns ...func(*appstream.Options)) (*appstream.UpdateAppBlockBuilderOutput, error)
	UpdateApplication(ctx context.Context, params *appstream.UpdateApplicationInput, optFns ...func(*appstream.Options)) (*appstream.UpdateApplicationOutput, error)
	UpdateDirectoryConfig(ctx context.Context, params *appstream.UpdateDirectoryConfigInput, optFns ...func(*appstream.Options)) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateEntitlement(ctx context.Context, params *appstream.UpdateEntitlementInput, optFns ...func(*appstream.Options)) (*appstream.UpdateEntitlementOutput, error)
	UpdateFleet(ctx context.Context, params *appstream.UpdateFleetInput, optFns ...func(*appstream.Options)) (*appstream.UpdateFleetOutput, error)
	UpdateImagePermissions(ctx context.Context, params *appstream.UpdateImagePermissionsInput, optFns ...func(*appstream.Options)) (*appstream.UpdateImagePermissionsOutput, error)
	UpdateStack(ctx context.Context, params *appstream.UpdateStackInput, optFns ...func(*appstream.Options)) (*appstream.UpdateStackOutput, error)
}
