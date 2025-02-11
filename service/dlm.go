// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

// DlmClient ...
type DlmClient interface {
	Options() dlm.Options
	CreateLifecyclePolicy(ctx context.Context, params *dlm.CreateLifecyclePolicyInput, optFns ...func(*dlm.Options)) (*dlm.CreateLifecyclePolicyOutput, error)
	DeleteLifecyclePolicy(ctx context.Context, params *dlm.DeleteLifecyclePolicyInput, optFns ...func(*dlm.Options)) (*dlm.DeleteLifecyclePolicyOutput, error)
	GetLifecyclePolicies(ctx context.Context, params *dlm.GetLifecyclePoliciesInput, optFns ...func(*dlm.Options)) (*dlm.GetLifecyclePoliciesOutput, error)
	GetLifecyclePolicy(ctx context.Context, params *dlm.GetLifecyclePolicyInput, optFns ...func(*dlm.Options)) (*dlm.GetLifecyclePolicyOutput, error)
	ListTagsForResource(ctx context.Context, params *dlm.ListTagsForResourceInput, optFns ...func(*dlm.Options)) (*dlm.ListTagsForResourceOutput, error)
	TagResource(ctx context.Context, params *dlm.TagResourceInput, optFns ...func(*dlm.Options)) (*dlm.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *dlm.UntagResourceInput, optFns ...func(*dlm.Options)) (*dlm.UntagResourceOutput, error)
	UpdateLifecyclePolicy(ctx context.Context, params *dlm.UpdateLifecyclePolicyInput, optFns ...func(*dlm.Options)) (*dlm.UpdateLifecyclePolicyOutput, error)
}
