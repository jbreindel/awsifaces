// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/nimble"
)

// NimbleClient ...
type NimbleClient interface {
	Options() nimble.Options
	AcceptEulas(ctx context.Context, params *nimble.AcceptEulasInput, optFns ...func(*nimble.Options)) (*nimble.AcceptEulasOutput, error)
	CreateLaunchProfile(ctx context.Context, params *nimble.CreateLaunchProfileInput, optFns ...func(*nimble.Options)) (*nimble.CreateLaunchProfileOutput, error)
	CreateStreamingImage(ctx context.Context, params *nimble.CreateStreamingImageInput, optFns ...func(*nimble.Options)) (*nimble.CreateStreamingImageOutput, error)
	CreateStreamingSession(ctx context.Context, params *nimble.CreateStreamingSessionInput, optFns ...func(*nimble.Options)) (*nimble.CreateStreamingSessionOutput, error)
	CreateStreamingSessionStream(ctx context.Context, params *nimble.CreateStreamingSessionStreamInput, optFns ...func(*nimble.Options)) (*nimble.CreateStreamingSessionStreamOutput, error)
	CreateStudio(ctx context.Context, params *nimble.CreateStudioInput, optFns ...func(*nimble.Options)) (*nimble.CreateStudioOutput, error)
	CreateStudioComponent(ctx context.Context, params *nimble.CreateStudioComponentInput, optFns ...func(*nimble.Options)) (*nimble.CreateStudioComponentOutput, error)
	DeleteLaunchProfile(ctx context.Context, params *nimble.DeleteLaunchProfileInput, optFns ...func(*nimble.Options)) (*nimble.DeleteLaunchProfileOutput, error)
	DeleteLaunchProfileMember(ctx context.Context, params *nimble.DeleteLaunchProfileMemberInput, optFns ...func(*nimble.Options)) (*nimble.DeleteLaunchProfileMemberOutput, error)
	DeleteStreamingImage(ctx context.Context, params *nimble.DeleteStreamingImageInput, optFns ...func(*nimble.Options)) (*nimble.DeleteStreamingImageOutput, error)
	DeleteStreamingSession(ctx context.Context, params *nimble.DeleteStreamingSessionInput, optFns ...func(*nimble.Options)) (*nimble.DeleteStreamingSessionOutput, error)
	DeleteStudio(ctx context.Context, params *nimble.DeleteStudioInput, optFns ...func(*nimble.Options)) (*nimble.DeleteStudioOutput, error)
	DeleteStudioComponent(ctx context.Context, params *nimble.DeleteStudioComponentInput, optFns ...func(*nimble.Options)) (*nimble.DeleteStudioComponentOutput, error)
	DeleteStudioMember(ctx context.Context, params *nimble.DeleteStudioMemberInput, optFns ...func(*nimble.Options)) (*nimble.DeleteStudioMemberOutput, error)
	GetEula(ctx context.Context, params *nimble.GetEulaInput, optFns ...func(*nimble.Options)) (*nimble.GetEulaOutput, error)
	GetLaunchProfile(ctx context.Context, params *nimble.GetLaunchProfileInput, optFns ...func(*nimble.Options)) (*nimble.GetLaunchProfileOutput, error)
	GetLaunchProfileDetails(ctx context.Context, params *nimble.GetLaunchProfileDetailsInput, optFns ...func(*nimble.Options)) (*nimble.GetLaunchProfileDetailsOutput, error)
	GetLaunchProfileInitialization(ctx context.Context, params *nimble.GetLaunchProfileInitializationInput, optFns ...func(*nimble.Options)) (*nimble.GetLaunchProfileInitializationOutput, error)
	GetLaunchProfileMember(ctx context.Context, params *nimble.GetLaunchProfileMemberInput, optFns ...func(*nimble.Options)) (*nimble.GetLaunchProfileMemberOutput, error)
	GetStreamingImage(ctx context.Context, params *nimble.GetStreamingImageInput, optFns ...func(*nimble.Options)) (*nimble.GetStreamingImageOutput, error)
	GetStreamingSession(ctx context.Context, params *nimble.GetStreamingSessionInput, optFns ...func(*nimble.Options)) (*nimble.GetStreamingSessionOutput, error)
	GetStreamingSessionBackup(ctx context.Context, params *nimble.GetStreamingSessionBackupInput, optFns ...func(*nimble.Options)) (*nimble.GetStreamingSessionBackupOutput, error)
	GetStreamingSessionStream(ctx context.Context, params *nimble.GetStreamingSessionStreamInput, optFns ...func(*nimble.Options)) (*nimble.GetStreamingSessionStreamOutput, error)
	GetStudio(ctx context.Context, params *nimble.GetStudioInput, optFns ...func(*nimble.Options)) (*nimble.GetStudioOutput, error)
	GetStudioComponent(ctx context.Context, params *nimble.GetStudioComponentInput, optFns ...func(*nimble.Options)) (*nimble.GetStudioComponentOutput, error)
	GetStudioMember(ctx context.Context, params *nimble.GetStudioMemberInput, optFns ...func(*nimble.Options)) (*nimble.GetStudioMemberOutput, error)
	ListEulaAcceptances(ctx context.Context, params *nimble.ListEulaAcceptancesInput, optFns ...func(*nimble.Options)) (*nimble.ListEulaAcceptancesOutput, error)
	ListEulas(ctx context.Context, params *nimble.ListEulasInput, optFns ...func(*nimble.Options)) (*nimble.ListEulasOutput, error)
	ListLaunchProfileMembers(ctx context.Context, params *nimble.ListLaunchProfileMembersInput, optFns ...func(*nimble.Options)) (*nimble.ListLaunchProfileMembersOutput, error)
	ListLaunchProfiles(ctx context.Context, params *nimble.ListLaunchProfilesInput, optFns ...func(*nimble.Options)) (*nimble.ListLaunchProfilesOutput, error)
	ListStreamingImages(ctx context.Context, params *nimble.ListStreamingImagesInput, optFns ...func(*nimble.Options)) (*nimble.ListStreamingImagesOutput, error)
	ListStreamingSessionBackups(ctx context.Context, params *nimble.ListStreamingSessionBackupsInput, optFns ...func(*nimble.Options)) (*nimble.ListStreamingSessionBackupsOutput, error)
	ListStreamingSessions(ctx context.Context, params *nimble.ListStreamingSessionsInput, optFns ...func(*nimble.Options)) (*nimble.ListStreamingSessionsOutput, error)
	ListStudioComponents(ctx context.Context, params *nimble.ListStudioComponentsInput, optFns ...func(*nimble.Options)) (*nimble.ListStudioComponentsOutput, error)
	ListStudioMembers(ctx context.Context, params *nimble.ListStudioMembersInput, optFns ...func(*nimble.Options)) (*nimble.ListStudioMembersOutput, error)
	ListStudios(ctx context.Context, params *nimble.ListStudiosInput, optFns ...func(*nimble.Options)) (*nimble.ListStudiosOutput, error)
	ListTagsForResource(ctx context.Context, params *nimble.ListTagsForResourceInput, optFns ...func(*nimble.Options)) (*nimble.ListTagsForResourceOutput, error)
	PutLaunchProfileMembers(ctx context.Context, params *nimble.PutLaunchProfileMembersInput, optFns ...func(*nimble.Options)) (*nimble.PutLaunchProfileMembersOutput, error)
	PutStudioMembers(ctx context.Context, params *nimble.PutStudioMembersInput, optFns ...func(*nimble.Options)) (*nimble.PutStudioMembersOutput, error)
	StartStreamingSession(ctx context.Context, params *nimble.StartStreamingSessionInput, optFns ...func(*nimble.Options)) (*nimble.StartStreamingSessionOutput, error)
	StartStudioSSOConfigurationRepair(ctx context.Context, params *nimble.StartStudioSSOConfigurationRepairInput, optFns ...func(*nimble.Options)) (*nimble.StartStudioSSOConfigurationRepairOutput, error)
	StopStreamingSession(ctx context.Context, params *nimble.StopStreamingSessionInput, optFns ...func(*nimble.Options)) (*nimble.StopStreamingSessionOutput, error)
	TagResource(ctx context.Context, params *nimble.TagResourceInput, optFns ...func(*nimble.Options)) (*nimble.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *nimble.UntagResourceInput, optFns ...func(*nimble.Options)) (*nimble.UntagResourceOutput, error)
	UpdateLaunchProfile(ctx context.Context, params *nimble.UpdateLaunchProfileInput, optFns ...func(*nimble.Options)) (*nimble.UpdateLaunchProfileOutput, error)
	UpdateLaunchProfileMember(ctx context.Context, params *nimble.UpdateLaunchProfileMemberInput, optFns ...func(*nimble.Options)) (*nimble.UpdateLaunchProfileMemberOutput, error)
	UpdateStreamingImage(ctx context.Context, params *nimble.UpdateStreamingImageInput, optFns ...func(*nimble.Options)) (*nimble.UpdateStreamingImageOutput, error)
	UpdateStudio(ctx context.Context, params *nimble.UpdateStudioInput, optFns ...func(*nimble.Options)) (*nimble.UpdateStudioOutput, error)
	UpdateStudioComponent(ctx context.Context, params *nimble.UpdateStudioComponentInput, optFns ...func(*nimble.Options)) (*nimble.UpdateStudioComponentOutput, error)
}
