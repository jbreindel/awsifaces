// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/oam"
)

// OamClient ...
type OamClient interface {
	Options() oam.Options
	CreateLink(ctx context.Context, params *oam.CreateLinkInput, optFns ...func(*oam.Options)) (*oam.CreateLinkOutput, error)
	CreateSink(ctx context.Context, params *oam.CreateSinkInput, optFns ...func(*oam.Options)) (*oam.CreateSinkOutput, error)
	DeleteLink(ctx context.Context, params *oam.DeleteLinkInput, optFns ...func(*oam.Options)) (*oam.DeleteLinkOutput, error)
	DeleteSink(ctx context.Context, params *oam.DeleteSinkInput, optFns ...func(*oam.Options)) (*oam.DeleteSinkOutput, error)
	GetLink(ctx context.Context, params *oam.GetLinkInput, optFns ...func(*oam.Options)) (*oam.GetLinkOutput, error)
	GetSink(ctx context.Context, params *oam.GetSinkInput, optFns ...func(*oam.Options)) (*oam.GetSinkOutput, error)
	GetSinkPolicy(ctx context.Context, params *oam.GetSinkPolicyInput, optFns ...func(*oam.Options)) (*oam.GetSinkPolicyOutput, error)
	ListAttachedLinks(ctx context.Context, params *oam.ListAttachedLinksInput, optFns ...func(*oam.Options)) (*oam.ListAttachedLinksOutput, error)
	ListLinks(ctx context.Context, params *oam.ListLinksInput, optFns ...func(*oam.Options)) (*oam.ListLinksOutput, error)
	ListSinks(ctx context.Context, params *oam.ListSinksInput, optFns ...func(*oam.Options)) (*oam.ListSinksOutput, error)
	ListTagsForResource(ctx context.Context, params *oam.ListTagsForResourceInput, optFns ...func(*oam.Options)) (*oam.ListTagsForResourceOutput, error)
	PutSinkPolicy(ctx context.Context, params *oam.PutSinkPolicyInput, optFns ...func(*oam.Options)) (*oam.PutSinkPolicyOutput, error)
	TagResource(ctx context.Context, params *oam.TagResourceInput, optFns ...func(*oam.Options)) (*oam.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *oam.UntagResourceInput, optFns ...func(*oam.Options)) (*oam.UntagResourceOutput, error)
	UpdateLink(ctx context.Context, params *oam.UpdateLinkInput, optFns ...func(*oam.Options)) (*oam.UpdateLinkOutput, error)
}
