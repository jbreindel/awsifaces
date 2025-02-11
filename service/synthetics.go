// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/synthetics"
)

// SyntheticsClient ...
type SyntheticsClient interface {
	Options() synthetics.Options
	AssociateResource(ctx context.Context, params *synthetics.AssociateResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.AssociateResourceOutput, error)
	CreateCanary(ctx context.Context, params *synthetics.CreateCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.CreateCanaryOutput, error)
	CreateGroup(ctx context.Context, params *synthetics.CreateGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.CreateGroupOutput, error)
	DeleteCanary(ctx context.Context, params *synthetics.DeleteCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.DeleteCanaryOutput, error)
	DeleteGroup(ctx context.Context, params *synthetics.DeleteGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.DeleteGroupOutput, error)
	DescribeCanaries(ctx context.Context, params *synthetics.DescribeCanariesInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeCanariesOutput, error)
	DescribeCanariesLastRun(ctx context.Context, params *synthetics.DescribeCanariesLastRunInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeCanariesLastRunOutput, error)
	DescribeRuntimeVersions(ctx context.Context, params *synthetics.DescribeRuntimeVersionsInput, optFns ...func(*synthetics.Options)) (*synthetics.DescribeRuntimeVersionsOutput, error)
	DisassociateResource(ctx context.Context, params *synthetics.DisassociateResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.DisassociateResourceOutput, error)
	GetCanary(ctx context.Context, params *synthetics.GetCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.GetCanaryOutput, error)
	GetCanaryRuns(ctx context.Context, params *synthetics.GetCanaryRunsInput, optFns ...func(*synthetics.Options)) (*synthetics.GetCanaryRunsOutput, error)
	GetGroup(ctx context.Context, params *synthetics.GetGroupInput, optFns ...func(*synthetics.Options)) (*synthetics.GetGroupOutput, error)
	ListAssociatedGroups(ctx context.Context, params *synthetics.ListAssociatedGroupsInput, optFns ...func(*synthetics.Options)) (*synthetics.ListAssociatedGroupsOutput, error)
	ListGroupResources(ctx context.Context, params *synthetics.ListGroupResourcesInput, optFns ...func(*synthetics.Options)) (*synthetics.ListGroupResourcesOutput, error)
	ListGroups(ctx context.Context, params *synthetics.ListGroupsInput, optFns ...func(*synthetics.Options)) (*synthetics.ListGroupsOutput, error)
	ListTagsForResource(ctx context.Context, params *synthetics.ListTagsForResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.ListTagsForResourceOutput, error)
	StartCanary(ctx context.Context, params *synthetics.StartCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.StartCanaryOutput, error)
	StopCanary(ctx context.Context, params *synthetics.StopCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.StopCanaryOutput, error)
	TagResource(ctx context.Context, params *synthetics.TagResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *synthetics.UntagResourceInput, optFns ...func(*synthetics.Options)) (*synthetics.UntagResourceOutput, error)
	UpdateCanary(ctx context.Context, params *synthetics.UpdateCanaryInput, optFns ...func(*synthetics.Options)) (*synthetics.UpdateCanaryOutput, error)
}
