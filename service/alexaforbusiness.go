// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
)

// AlexaforbusinessClient ...
type AlexaforbusinessClient interface {
	Options() alexaforbusiness.Options
	ApproveSkill(ctx context.Context, params *alexaforbusiness.ApproveSkillInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ApproveSkillOutput, error)
	AssociateContactWithAddressBook(ctx context.Context, params *alexaforbusiness.AssociateContactWithAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateDeviceWithNetworkProfile(ctx context.Context, params *alexaforbusiness.AssociateDeviceWithNetworkProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error)
	AssociateDeviceWithRoom(ctx context.Context, params *alexaforbusiness.AssociateDeviceWithRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateSkillGroupWithRoom(ctx context.Context, params *alexaforbusiness.AssociateSkillGroupWithRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillWithSkillGroup(ctx context.Context, params *alexaforbusiness.AssociateSkillWithSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithUsers(ctx context.Context, params *alexaforbusiness.AssociateSkillWithUsersInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.AssociateSkillWithUsersOutput, error)
	CreateAddressBook(ctx context.Context, params *alexaforbusiness.CreateAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateBusinessReportSchedule(ctx context.Context, params *alexaforbusiness.CreateBusinessReportScheduleInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateConferenceProvider(ctx context.Context, params *alexaforbusiness.CreateConferenceProviderInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateContact(ctx context.Context, params *alexaforbusiness.CreateContactInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateContactOutput, error)
	CreateGatewayGroup(ctx context.Context, params *alexaforbusiness.CreateGatewayGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateGatewayGroupOutput, error)
	CreateNetworkProfile(ctx context.Context, params *alexaforbusiness.CreateNetworkProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateNetworkProfileOutput, error)
	CreateProfile(ctx context.Context, params *alexaforbusiness.CreateProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateProfileOutput, error)
	CreateRoom(ctx context.Context, params *alexaforbusiness.CreateRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateRoomOutput, error)
	CreateSkillGroup(ctx context.Context, params *alexaforbusiness.CreateSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateUser(ctx context.Context, params *alexaforbusiness.CreateUserInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.CreateUserOutput, error)
	DeleteAddressBook(ctx context.Context, params *alexaforbusiness.DeleteAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteBusinessReportSchedule(ctx context.Context, params *alexaforbusiness.DeleteBusinessReportScheduleInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteConferenceProvider(ctx context.Context, params *alexaforbusiness.DeleteConferenceProviderInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteContact(ctx context.Context, params *alexaforbusiness.DeleteContactInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteDevice(ctx context.Context, params *alexaforbusiness.DeleteDeviceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceUsageData(ctx context.Context, params *alexaforbusiness.DeleteDeviceUsageDataInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error)
	DeleteGatewayGroup(ctx context.Context, params *alexaforbusiness.DeleteGatewayGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteGatewayGroupOutput, error)
	DeleteNetworkProfile(ctx context.Context, params *alexaforbusiness.DeleteNetworkProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteNetworkProfileOutput, error)
	DeleteProfile(ctx context.Context, params *alexaforbusiness.DeleteProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteRoom(ctx context.Context, params *alexaforbusiness.DeleteRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomSkillParameter(ctx context.Context, params *alexaforbusiness.DeleteRoomSkillParameterInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteSkillAuthorization(ctx context.Context, params *alexaforbusiness.DeleteSkillAuthorizationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillGroup(ctx context.Context, params *alexaforbusiness.DeleteSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteUser(ctx context.Context, params *alexaforbusiness.DeleteUserInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DeleteUserOutput, error)
	DisassociateContactFromAddressBook(ctx context.Context, params *alexaforbusiness.DisassociateContactFromAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateDeviceFromRoom(ctx context.Context, params *alexaforbusiness.DisassociateDeviceFromRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateSkillFromSkillGroup(ctx context.Context, params *alexaforbusiness.DisassociateSkillFromSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromUsers(ctx context.Context, params *alexaforbusiness.DisassociateSkillFromUsersInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error)
	DisassociateSkillGroupFromRoom(ctx context.Context, params *alexaforbusiness.DisassociateSkillGroupFromRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	ForgetSmartHomeAppliances(ctx context.Context, params *alexaforbusiness.ForgetSmartHomeAppliancesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	GetAddressBook(ctx context.Context, params *alexaforbusiness.GetAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetAddressBookOutput, error)
	GetConferencePreference(ctx context.Context, params *alexaforbusiness.GetConferencePreferenceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferenceProvider(ctx context.Context, params *alexaforbusiness.GetConferenceProviderInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetContact(ctx context.Context, params *alexaforbusiness.GetContactInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetContactOutput, error)
	GetDevice(ctx context.Context, params *alexaforbusiness.GetDeviceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetDeviceOutput, error)
	GetGateway(ctx context.Context, params *alexaforbusiness.GetGatewayInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetGatewayOutput, error)
	GetGatewayGroup(ctx context.Context, params *alexaforbusiness.GetGatewayGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetGatewayGroupOutput, error)
	GetInvitationConfiguration(ctx context.Context, params *alexaforbusiness.GetInvitationConfigurationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetInvitationConfigurationOutput, error)
	GetNetworkProfile(ctx context.Context, params *alexaforbusiness.GetNetworkProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetNetworkProfileOutput, error)
	GetProfile(ctx context.Context, params *alexaforbusiness.GetProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetProfileOutput, error)
	GetRoom(ctx context.Context, params *alexaforbusiness.GetRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomSkillParameter(ctx context.Context, params *alexaforbusiness.GetRoomSkillParameterInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetSkillGroup(ctx context.Context, params *alexaforbusiness.GetSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.GetSkillGroupOutput, error)
	ListBusinessReportSchedules(ctx context.Context, params *alexaforbusiness.ListBusinessReportSchedulesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListConferenceProviders(ctx context.Context, params *alexaforbusiness.ListConferenceProvidersInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListDeviceEvents(ctx context.Context, params *alexaforbusiness.ListDeviceEventsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListGatewayGroups(ctx context.Context, params *alexaforbusiness.ListGatewayGroupsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListGatewayGroupsOutput, error)
	ListGateways(ctx context.Context, params *alexaforbusiness.ListGatewaysInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListGatewaysOutput, error)
	ListSkills(ctx context.Context, params *alexaforbusiness.ListSkillsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsStoreCategories(ctx context.Context, params *alexaforbusiness.ListSkillsStoreCategoriesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreSkillsByCategory(ctx context.Context, params *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSmartHomeAppliances(ctx context.Context, params *alexaforbusiness.ListSmartHomeAppliancesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListTags(ctx context.Context, params *alexaforbusiness.ListTagsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ListTagsOutput, error)
	PutConferencePreference(ctx context.Context, params *alexaforbusiness.PutConferencePreferenceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutInvitationConfiguration(ctx context.Context, params *alexaforbusiness.PutInvitationConfigurationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.PutInvitationConfigurationOutput, error)
	PutRoomSkillParameter(ctx context.Context, params *alexaforbusiness.PutRoomSkillParameterInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutSkillAuthorization(ctx context.Context, params *alexaforbusiness.PutSkillAuthorizationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	RegisterAVSDevice(ctx context.Context, params *alexaforbusiness.RegisterAVSDeviceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RejectSkill(ctx context.Context, params *alexaforbusiness.RejectSkillInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.RejectSkillOutput, error)
	ResolveRoom(ctx context.Context, params *alexaforbusiness.ResolveRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.ResolveRoomOutput, error)
	RevokeInvitation(ctx context.Context, params *alexaforbusiness.RevokeInvitationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.RevokeInvitationOutput, error)
	SearchAddressBooks(ctx context.Context, params *alexaforbusiness.SearchAddressBooksInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchContacts(ctx context.Context, params *alexaforbusiness.SearchContactsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchContactsOutput, error)
	SearchDevices(ctx context.Context, params *alexaforbusiness.SearchDevicesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchNetworkProfiles(ctx context.Context, params *alexaforbusiness.SearchNetworkProfilesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchNetworkProfilesOutput, error)
	SearchProfiles(ctx context.Context, params *alexaforbusiness.SearchProfilesInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchRooms(ctx context.Context, params *alexaforbusiness.SearchRoomsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchSkillGroups(ctx context.Context, params *alexaforbusiness.SearchSkillGroupsInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchUsers(ctx context.Context, params *alexaforbusiness.SearchUsersInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SearchUsersOutput, error)
	SendAnnouncement(ctx context.Context, params *alexaforbusiness.SendAnnouncementInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SendAnnouncementOutput, error)
	SendInvitation(ctx context.Context, params *alexaforbusiness.SendInvitationInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.SendInvitationOutput, error)
	StartDeviceSync(ctx context.Context, params *alexaforbusiness.StartDeviceSyncInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartSmartHomeApplianceDiscovery(ctx context.Context, params *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	TagResource(ctx context.Context, params *alexaforbusiness.TagResourceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *alexaforbusiness.UntagResourceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UntagResourceOutput, error)
	UpdateAddressBook(ctx context.Context, params *alexaforbusiness.UpdateAddressBookInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateBusinessReportSchedule(ctx context.Context, params *alexaforbusiness.UpdateBusinessReportScheduleInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateConferenceProvider(ctx context.Context, params *alexaforbusiness.UpdateConferenceProviderInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateContact(ctx context.Context, params *alexaforbusiness.UpdateContactInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateDevice(ctx context.Context, params *alexaforbusiness.UpdateDeviceInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateGateway(ctx context.Context, params *alexaforbusiness.UpdateGatewayInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateGatewayOutput, error)
	UpdateGatewayGroup(ctx context.Context, params *alexaforbusiness.UpdateGatewayGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateGatewayGroupOutput, error)
	UpdateNetworkProfile(ctx context.Context, params *alexaforbusiness.UpdateNetworkProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateNetworkProfileOutput, error)
	UpdateProfile(ctx context.Context, params *alexaforbusiness.UpdateProfileInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateRoom(ctx context.Context, params *alexaforbusiness.UpdateRoomInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateSkillGroup(ctx context.Context, params *alexaforbusiness.UpdateSkillGroupInput, optFns ...func(*alexaforbusiness.Options)) (*alexaforbusiness.UpdateSkillGroupOutput, error)
}
