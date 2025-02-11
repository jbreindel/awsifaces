// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/repostspace"
)

// RepostspaceClient ...
type RepostspaceClient interface {
	Options() repostspace.Options
	CreateSpace(ctx context.Context, params *repostspace.CreateSpaceInput, optFns ...func(*repostspace.Options)) (*repostspace.CreateSpaceOutput, error)
	DeleteSpace(ctx context.Context, params *repostspace.DeleteSpaceInput, optFns ...func(*repostspace.Options)) (*repostspace.DeleteSpaceOutput, error)
	DeregisterAdmin(ctx context.Context, params *repostspace.DeregisterAdminInput, optFns ...func(*repostspace.Options)) (*repostspace.DeregisterAdminOutput, error)
	GetSpace(ctx context.Context, params *repostspace.GetSpaceInput, optFns ...func(*repostspace.Options)) (*repostspace.GetSpaceOutput, error)
	ListSpaces(ctx context.Context, params *repostspace.ListSpacesInput, optFns ...func(*repostspace.Options)) (*repostspace.ListSpacesOutput, error)
	ListTagsForResource(ctx context.Context, params *repostspace.ListTagsForResourceInput, optFns ...func(*repostspace.Options)) (*repostspace.ListTagsForResourceOutput, error)
	RegisterAdmin(ctx context.Context, params *repostspace.RegisterAdminInput, optFns ...func(*repostspace.Options)) (*repostspace.RegisterAdminOutput, error)
	SendInvites(ctx context.Context, params *repostspace.SendInvitesInput, optFns ...func(*repostspace.Options)) (*repostspace.SendInvitesOutput, error)
	TagResource(ctx context.Context, params *repostspace.TagResourceInput, optFns ...func(*repostspace.Options)) (*repostspace.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *repostspace.UntagResourceInput, optFns ...func(*repostspace.Options)) (*repostspace.UntagResourceOutput, error)
	UpdateSpace(ctx context.Context, params *repostspace.UpdateSpaceInput, optFns ...func(*repostspace.Options)) (*repostspace.UpdateSpaceOutput, error)
}
