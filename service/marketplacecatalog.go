// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
)

// MarketplacecatalogClient ...
type MarketplacecatalogClient interface {
	Options() marketplacecatalog.Options
	BatchDescribeEntities(ctx context.Context, params *marketplacecatalog.BatchDescribeEntitiesInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.BatchDescribeEntitiesOutput, error)
	CancelChangeSet(ctx context.Context, params *marketplacecatalog.CancelChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.CancelChangeSetOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *marketplacecatalog.DeleteResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DeleteResourcePolicyOutput, error)
	DescribeChangeSet(ctx context.Context, params *marketplacecatalog.DescribeChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeChangeSetOutput, error)
	DescribeEntity(ctx context.Context, params *marketplacecatalog.DescribeEntityInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.DescribeEntityOutput, error)
	GetResourcePolicy(ctx context.Context, params *marketplacecatalog.GetResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.GetResourcePolicyOutput, error)
	ListChangeSets(ctx context.Context, params *marketplacecatalog.ListChangeSetsInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListChangeSetsOutput, error)
	ListEntities(ctx context.Context, params *marketplacecatalog.ListEntitiesInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListEntitiesOutput, error)
	ListTagsForResource(ctx context.Context, params *marketplacecatalog.ListTagsForResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.ListTagsForResourceOutput, error)
	PutResourcePolicy(ctx context.Context, params *marketplacecatalog.PutResourcePolicyInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.PutResourcePolicyOutput, error)
	StartChangeSet(ctx context.Context, params *marketplacecatalog.StartChangeSetInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.StartChangeSetOutput, error)
	TagResource(ctx context.Context, params *marketplacecatalog.TagResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *marketplacecatalog.UntagResourceInput, optFns ...func(*marketplacecatalog.Options)) (*marketplacecatalog.UntagResourceOutput, error)
}
