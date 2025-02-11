// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

// MediastoreClient ...
type MediastoreClient interface {
	Options() mediastore.Options
	CreateContainer(ctx context.Context, params *mediastore.CreateContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.CreateContainerOutput, error)
	DeleteContainer(ctx context.Context, params *mediastore.DeleteContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerPolicy(ctx context.Context, params *mediastore.DeleteContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteCorsPolicy(ctx context.Context, params *mediastore.DeleteCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteLifecyclePolicy(ctx context.Context, params *mediastore.DeleteLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteMetricPolicy(ctx context.Context, params *mediastore.DeleteMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteMetricPolicyOutput, error)
	DescribeContainer(ctx context.Context, params *mediastore.DescribeContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.DescribeContainerOutput, error)
	GetContainerPolicy(ctx context.Context, params *mediastore.GetContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetContainerPolicyOutput, error)
	GetCorsPolicy(ctx context.Context, params *mediastore.GetCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetCorsPolicyOutput, error)
	GetLifecyclePolicy(ctx context.Context, params *mediastore.GetLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetLifecyclePolicyOutput, error)
	GetMetricPolicy(ctx context.Context, params *mediastore.GetMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetMetricPolicyOutput, error)
	ListContainers(ctx context.Context, params *mediastore.ListContainersInput, optFns ...func(*mediastore.Options)) (*mediastore.ListContainersOutput, error)
	ListTagsForResource(ctx context.Context, params *mediastore.ListTagsForResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.ListTagsForResourceOutput, error)
	PutContainerPolicy(ctx context.Context, params *mediastore.PutContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutContainerPolicyOutput, error)
	PutCorsPolicy(ctx context.Context, params *mediastore.PutCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutCorsPolicyOutput, error)
	PutLifecyclePolicy(ctx context.Context, params *mediastore.PutLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutLifecyclePolicyOutput, error)
	PutMetricPolicy(ctx context.Context, params *mediastore.PutMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutMetricPolicyOutput, error)
	StartAccessLogging(ctx context.Context, params *mediastore.StartAccessLoggingInput, optFns ...func(*mediastore.Options)) (*mediastore.StartAccessLoggingOutput, error)
	StopAccessLogging(ctx context.Context, params *mediastore.StopAccessLoggingInput, optFns ...func(*mediastore.Options)) (*mediastore.StopAccessLoggingOutput, error)
	TagResource(ctx context.Context, params *mediastore.TagResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *mediastore.UntagResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.UntagResourceOutput, error)
}
