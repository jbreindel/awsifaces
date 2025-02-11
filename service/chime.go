// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/chime"
)

// ChimeClient ...
type ChimeClient interface {
	Options() chime.Options
	AssociatePhoneNumberWithUser(ctx context.Context, params *chime.AssociatePhoneNumberWithUserInput, optFns ...func(*chime.Options)) (*chime.AssociatePhoneNumberWithUserOutput, error)
	AssociatePhoneNumbersWithVoiceConnector(ctx context.Context, params *chime.AssociatePhoneNumbersWithVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, params *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error)
	AssociateSigninDelegateGroupsWithAccount(ctx context.Context, params *chime.AssociateSigninDelegateGroupsWithAccountInput, optFns ...func(*chime.Options)) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error)
	BatchCreateAttendee(ctx context.Context, params *chime.BatchCreateAttendeeInput, optFns ...func(*chime.Options)) (*chime.BatchCreateAttendeeOutput, error)
	BatchCreateChannelMembership(ctx context.Context, params *chime.BatchCreateChannelMembershipInput, optFns ...func(*chime.Options)) (*chime.BatchCreateChannelMembershipOutput, error)
	BatchCreateRoomMembership(ctx context.Context, params *chime.BatchCreateRoomMembershipInput, optFns ...func(*chime.Options)) (*chime.BatchCreateRoomMembershipOutput, error)
	BatchDeletePhoneNumber(ctx context.Context, params *chime.BatchDeletePhoneNumberInput, optFns ...func(*chime.Options)) (*chime.BatchDeletePhoneNumberOutput, error)
	BatchSuspendUser(ctx context.Context, params *chime.BatchSuspendUserInput, optFns ...func(*chime.Options)) (*chime.BatchSuspendUserOutput, error)
	BatchUnsuspendUser(ctx context.Context, params *chime.BatchUnsuspendUserInput, optFns ...func(*chime.Options)) (*chime.BatchUnsuspendUserOutput, error)
	BatchUpdatePhoneNumber(ctx context.Context, params *chime.BatchUpdatePhoneNumberInput, optFns ...func(*chime.Options)) (*chime.BatchUpdatePhoneNumberOutput, error)
	BatchUpdateUser(ctx context.Context, params *chime.BatchUpdateUserInput, optFns ...func(*chime.Options)) (*chime.BatchUpdateUserOutput, error)
	CreateAccount(ctx context.Context, params *chime.CreateAccountInput, optFns ...func(*chime.Options)) (*chime.CreateAccountOutput, error)
	CreateAppInstance(ctx context.Context, params *chime.CreateAppInstanceInput, optFns ...func(*chime.Options)) (*chime.CreateAppInstanceOutput, error)
	CreateAppInstanceAdmin(ctx context.Context, params *chime.CreateAppInstanceAdminInput, optFns ...func(*chime.Options)) (*chime.CreateAppInstanceAdminOutput, error)
	CreateAppInstanceUser(ctx context.Context, params *chime.CreateAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.CreateAppInstanceUserOutput, error)
	CreateAttendee(ctx context.Context, params *chime.CreateAttendeeInput, optFns ...func(*chime.Options)) (*chime.CreateAttendeeOutput, error)
	CreateBot(ctx context.Context, params *chime.CreateBotInput, optFns ...func(*chime.Options)) (*chime.CreateBotOutput, error)
	CreateChannel(ctx context.Context, params *chime.CreateChannelInput, optFns ...func(*chime.Options)) (*chime.CreateChannelOutput, error)
	CreateChannelBan(ctx context.Context, params *chime.CreateChannelBanInput, optFns ...func(*chime.Options)) (*chime.CreateChannelBanOutput, error)
	CreateChannelMembership(ctx context.Context, params *chime.CreateChannelMembershipInput, optFns ...func(*chime.Options)) (*chime.CreateChannelMembershipOutput, error)
	CreateChannelModerator(ctx context.Context, params *chime.CreateChannelModeratorInput, optFns ...func(*chime.Options)) (*chime.CreateChannelModeratorOutput, error)
	CreateMediaCapturePipeline(ctx context.Context, params *chime.CreateMediaCapturePipelineInput, optFns ...func(*chime.Options)) (*chime.CreateMediaCapturePipelineOutput, error)
	CreateMeeting(ctx context.Context, params *chime.CreateMeetingInput, optFns ...func(*chime.Options)) (*chime.CreateMeetingOutput, error)
	CreateMeetingDialOut(ctx context.Context, params *chime.CreateMeetingDialOutInput, optFns ...func(*chime.Options)) (*chime.CreateMeetingDialOutOutput, error)
	CreateMeetingWithAttendees(ctx context.Context, params *chime.CreateMeetingWithAttendeesInput, optFns ...func(*chime.Options)) (*chime.CreateMeetingWithAttendeesOutput, error)
	CreatePhoneNumberOrder(ctx context.Context, params *chime.CreatePhoneNumberOrderInput, optFns ...func(*chime.Options)) (*chime.CreatePhoneNumberOrderOutput, error)
	CreateProxySession(ctx context.Context, params *chime.CreateProxySessionInput, optFns ...func(*chime.Options)) (*chime.CreateProxySessionOutput, error)
	CreateRoom(ctx context.Context, params *chime.CreateRoomInput, optFns ...func(*chime.Options)) (*chime.CreateRoomOutput, error)
	CreateRoomMembership(ctx context.Context, params *chime.CreateRoomMembershipInput, optFns ...func(*chime.Options)) (*chime.CreateRoomMembershipOutput, error)
	CreateSipMediaApplication(ctx context.Context, params *chime.CreateSipMediaApplicationInput, optFns ...func(*chime.Options)) (*chime.CreateSipMediaApplicationOutput, error)
	CreateSipMediaApplicationCall(ctx context.Context, params *chime.CreateSipMediaApplicationCallInput, optFns ...func(*chime.Options)) (*chime.CreateSipMediaApplicationCallOutput, error)
	CreateSipRule(ctx context.Context, params *chime.CreateSipRuleInput, optFns ...func(*chime.Options)) (*chime.CreateSipRuleOutput, error)
	CreateUser(ctx context.Context, params *chime.CreateUserInput, optFns ...func(*chime.Options)) (*chime.CreateUserOutput, error)
	CreateVoiceConnector(ctx context.Context, params *chime.CreateVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.CreateVoiceConnectorOutput, error)
	CreateVoiceConnectorGroup(ctx context.Context, params *chime.CreateVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.CreateVoiceConnectorGroupOutput, error)
	DeleteAccount(ctx context.Context, params *chime.DeleteAccountInput, optFns ...func(*chime.Options)) (*chime.DeleteAccountOutput, error)
	DeleteAppInstance(ctx context.Context, params *chime.DeleteAppInstanceInput, optFns ...func(*chime.Options)) (*chime.DeleteAppInstanceOutput, error)
	DeleteAppInstanceAdmin(ctx context.Context, params *chime.DeleteAppInstanceAdminInput, optFns ...func(*chime.Options)) (*chime.DeleteAppInstanceAdminOutput, error)
	DeleteAppInstanceStreamingConfigurations(ctx context.Context, params *chime.DeleteAppInstanceStreamingConfigurationsInput, optFns ...func(*chime.Options)) (*chime.DeleteAppInstanceStreamingConfigurationsOutput, error)
	DeleteAppInstanceUser(ctx context.Context, params *chime.DeleteAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.DeleteAppInstanceUserOutput, error)
	DeleteAttendee(ctx context.Context, params *chime.DeleteAttendeeInput, optFns ...func(*chime.Options)) (*chime.DeleteAttendeeOutput, error)
	DeleteChannel(ctx context.Context, params *chime.DeleteChannelInput, optFns ...func(*chime.Options)) (*chime.DeleteChannelOutput, error)
	DeleteChannelBan(ctx context.Context, params *chime.DeleteChannelBanInput, optFns ...func(*chime.Options)) (*chime.DeleteChannelBanOutput, error)
	DeleteChannelMembership(ctx context.Context, params *chime.DeleteChannelMembershipInput, optFns ...func(*chime.Options)) (*chime.DeleteChannelMembershipOutput, error)
	DeleteChannelMessage(ctx context.Context, params *chime.DeleteChannelMessageInput, optFns ...func(*chime.Options)) (*chime.DeleteChannelMessageOutput, error)
	DeleteChannelModerator(ctx context.Context, params *chime.DeleteChannelModeratorInput, optFns ...func(*chime.Options)) (*chime.DeleteChannelModeratorOutput, error)
	DeleteEventsConfiguration(ctx context.Context, params *chime.DeleteEventsConfigurationInput, optFns ...func(*chime.Options)) (*chime.DeleteEventsConfigurationOutput, error)
	DeleteMediaCapturePipeline(ctx context.Context, params *chime.DeleteMediaCapturePipelineInput, optFns ...func(*chime.Options)) (*chime.DeleteMediaCapturePipelineOutput, error)
	DeleteMeeting(ctx context.Context, params *chime.DeleteMeetingInput, optFns ...func(*chime.Options)) (*chime.DeleteMeetingOutput, error)
	DeletePhoneNumber(ctx context.Context, params *chime.DeletePhoneNumberInput, optFns ...func(*chime.Options)) (*chime.DeletePhoneNumberOutput, error)
	DeleteProxySession(ctx context.Context, params *chime.DeleteProxySessionInput, optFns ...func(*chime.Options)) (*chime.DeleteProxySessionOutput, error)
	DeleteRoom(ctx context.Context, params *chime.DeleteRoomInput, optFns ...func(*chime.Options)) (*chime.DeleteRoomOutput, error)
	DeleteRoomMembership(ctx context.Context, params *chime.DeleteRoomMembershipInput, optFns ...func(*chime.Options)) (*chime.DeleteRoomMembershipOutput, error)
	DeleteSipMediaApplication(ctx context.Context, params *chime.DeleteSipMediaApplicationInput, optFns ...func(*chime.Options)) (*chime.DeleteSipMediaApplicationOutput, error)
	DeleteSipRule(ctx context.Context, params *chime.DeleteSipRuleInput, optFns ...func(*chime.Options)) (*chime.DeleteSipRuleOutput, error)
	DeleteVoiceConnector(ctx context.Context, params *chime.DeleteVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorOutput, error)
	DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error)
	DeleteVoiceConnectorGroup(ctx context.Context, params *chime.DeleteVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorGroupOutput, error)
	DeleteVoiceConnectorOrigination(ctx context.Context, params *chime.DeleteVoiceConnectorOriginationInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorOriginationOutput, error)
	DeleteVoiceConnectorProxy(ctx context.Context, params *chime.DeleteVoiceConnectorProxyInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorProxyOutput, error)
	DeleteVoiceConnectorStreamingConfiguration(ctx context.Context, params *chime.DeleteVoiceConnectorStreamingConfigurationInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error)
	DeleteVoiceConnectorTermination(ctx context.Context, params *chime.DeleteVoiceConnectorTerminationInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorTerminationOutput, error)
	DeleteVoiceConnectorTerminationCredentials(ctx context.Context, params *chime.DeleteVoiceConnectorTerminationCredentialsInput, optFns ...func(*chime.Options)) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
	DescribeAppInstance(ctx context.Context, params *chime.DescribeAppInstanceInput, optFns ...func(*chime.Options)) (*chime.DescribeAppInstanceOutput, error)
	DescribeAppInstanceAdmin(ctx context.Context, params *chime.DescribeAppInstanceAdminInput, optFns ...func(*chime.Options)) (*chime.DescribeAppInstanceAdminOutput, error)
	DescribeAppInstanceUser(ctx context.Context, params *chime.DescribeAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.DescribeAppInstanceUserOutput, error)
	DescribeChannel(ctx context.Context, params *chime.DescribeChannelInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelOutput, error)
	DescribeChannelBan(ctx context.Context, params *chime.DescribeChannelBanInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelBanOutput, error)
	DescribeChannelMembership(ctx context.Context, params *chime.DescribeChannelMembershipInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelMembershipOutput, error)
	DescribeChannelMembershipForAppInstanceUser(ctx context.Context, params *chime.DescribeChannelMembershipForAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelMembershipForAppInstanceUserOutput, error)
	DescribeChannelModeratedByAppInstanceUser(ctx context.Context, params *chime.DescribeChannelModeratedByAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelModeratedByAppInstanceUserOutput, error)
	DescribeChannelModerator(ctx context.Context, params *chime.DescribeChannelModeratorInput, optFns ...func(*chime.Options)) (*chime.DescribeChannelModeratorOutput, error)
	DisassociatePhoneNumberFromUser(ctx context.Context, params *chime.DisassociatePhoneNumberFromUserInput, optFns ...func(*chime.Options)) (*chime.DisassociatePhoneNumberFromUserOutput, error)
	DisassociatePhoneNumbersFromVoiceConnector(ctx context.Context, params *chime.DisassociatePhoneNumbersFromVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx context.Context, params *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error)
	DisassociateSigninDelegateGroupsFromAccount(ctx context.Context, params *chime.DisassociateSigninDelegateGroupsFromAccountInput, optFns ...func(*chime.Options)) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error)
	GetAccount(ctx context.Context, params *chime.GetAccountInput, optFns ...func(*chime.Options)) (*chime.GetAccountOutput, error)
	GetAccountSettings(ctx context.Context, params *chime.GetAccountSettingsInput, optFns ...func(*chime.Options)) (*chime.GetAccountSettingsOutput, error)
	GetAppInstanceRetentionSettings(ctx context.Context, params *chime.GetAppInstanceRetentionSettingsInput, optFns ...func(*chime.Options)) (*chime.GetAppInstanceRetentionSettingsOutput, error)
	GetAppInstanceStreamingConfigurations(ctx context.Context, params *chime.GetAppInstanceStreamingConfigurationsInput, optFns ...func(*chime.Options)) (*chime.GetAppInstanceStreamingConfigurationsOutput, error)
	GetAttendee(ctx context.Context, params *chime.GetAttendeeInput, optFns ...func(*chime.Options)) (*chime.GetAttendeeOutput, error)
	GetBot(ctx context.Context, params *chime.GetBotInput, optFns ...func(*chime.Options)) (*chime.GetBotOutput, error)
	GetChannelMessage(ctx context.Context, params *chime.GetChannelMessageInput, optFns ...func(*chime.Options)) (*chime.GetChannelMessageOutput, error)
	GetEventsConfiguration(ctx context.Context, params *chime.GetEventsConfigurationInput, optFns ...func(*chime.Options)) (*chime.GetEventsConfigurationOutput, error)
	GetGlobalSettings(ctx context.Context, params *chime.GetGlobalSettingsInput, optFns ...func(*chime.Options)) (*chime.GetGlobalSettingsOutput, error)
	GetMediaCapturePipeline(ctx context.Context, params *chime.GetMediaCapturePipelineInput, optFns ...func(*chime.Options)) (*chime.GetMediaCapturePipelineOutput, error)
	GetMeeting(ctx context.Context, params *chime.GetMeetingInput, optFns ...func(*chime.Options)) (*chime.GetMeetingOutput, error)
	GetMessagingSessionEndpoint(ctx context.Context, params *chime.GetMessagingSessionEndpointInput, optFns ...func(*chime.Options)) (*chime.GetMessagingSessionEndpointOutput, error)
	GetPhoneNumber(ctx context.Context, params *chime.GetPhoneNumberInput, optFns ...func(*chime.Options)) (*chime.GetPhoneNumberOutput, error)
	GetPhoneNumberOrder(ctx context.Context, params *chime.GetPhoneNumberOrderInput, optFns ...func(*chime.Options)) (*chime.GetPhoneNumberOrderOutput, error)
	GetPhoneNumberSettings(ctx context.Context, params *chime.GetPhoneNumberSettingsInput, optFns ...func(*chime.Options)) (*chime.GetPhoneNumberSettingsOutput, error)
	GetProxySession(ctx context.Context, params *chime.GetProxySessionInput, optFns ...func(*chime.Options)) (*chime.GetProxySessionOutput, error)
	GetRetentionSettings(ctx context.Context, params *chime.GetRetentionSettingsInput, optFns ...func(*chime.Options)) (*chime.GetRetentionSettingsOutput, error)
	GetRoom(ctx context.Context, params *chime.GetRoomInput, optFns ...func(*chime.Options)) (*chime.GetRoomOutput, error)
	GetSipMediaApplication(ctx context.Context, params *chime.GetSipMediaApplicationInput, optFns ...func(*chime.Options)) (*chime.GetSipMediaApplicationOutput, error)
	GetSipMediaApplicationLoggingConfiguration(ctx context.Context, params *chime.GetSipMediaApplicationLoggingConfigurationInput, optFns ...func(*chime.Options)) (*chime.GetSipMediaApplicationLoggingConfigurationOutput, error)
	GetSipRule(ctx context.Context, params *chime.GetSipRuleInput, optFns ...func(*chime.Options)) (*chime.GetSipRuleOutput, error)
	GetUser(ctx context.Context, params *chime.GetUserInput, optFns ...func(*chime.Options)) (*chime.GetUserOutput, error)
	GetUserSettings(ctx context.Context, params *chime.GetUserSettingsInput, optFns ...func(*chime.Options)) (*chime.GetUserSettingsOutput, error)
	GetVoiceConnector(ctx context.Context, params *chime.GetVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorOutput, error)
	GetVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *chime.GetVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error)
	GetVoiceConnectorGroup(ctx context.Context, params *chime.GetVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorGroupOutput, error)
	GetVoiceConnectorLoggingConfiguration(ctx context.Context, params *chime.GetVoiceConnectorLoggingConfigurationInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error)
	GetVoiceConnectorOrigination(ctx context.Context, params *chime.GetVoiceConnectorOriginationInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorOriginationOutput, error)
	GetVoiceConnectorProxy(ctx context.Context, params *chime.GetVoiceConnectorProxyInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorProxyOutput, error)
	GetVoiceConnectorStreamingConfiguration(ctx context.Context, params *chime.GetVoiceConnectorStreamingConfigurationInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error)
	GetVoiceConnectorTermination(ctx context.Context, params *chime.GetVoiceConnectorTerminationInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorTerminationOutput, error)
	GetVoiceConnectorTerminationHealth(ctx context.Context, params *chime.GetVoiceConnectorTerminationHealthInput, optFns ...func(*chime.Options)) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
	InviteUsers(ctx context.Context, params *chime.InviteUsersInput, optFns ...func(*chime.Options)) (*chime.InviteUsersOutput, error)
	ListAccounts(ctx context.Context, params *chime.ListAccountsInput, optFns ...func(*chime.Options)) (*chime.ListAccountsOutput, error)
	ListAppInstanceAdmins(ctx context.Context, params *chime.ListAppInstanceAdminsInput, optFns ...func(*chime.Options)) (*chime.ListAppInstanceAdminsOutput, error)
	ListAppInstanceUsers(ctx context.Context, params *chime.ListAppInstanceUsersInput, optFns ...func(*chime.Options)) (*chime.ListAppInstanceUsersOutput, error)
	ListAppInstances(ctx context.Context, params *chime.ListAppInstancesInput, optFns ...func(*chime.Options)) (*chime.ListAppInstancesOutput, error)
	ListAttendeeTags(ctx context.Context, params *chime.ListAttendeeTagsInput, optFns ...func(*chime.Options)) (*chime.ListAttendeeTagsOutput, error)
	ListAttendees(ctx context.Context, params *chime.ListAttendeesInput, optFns ...func(*chime.Options)) (*chime.ListAttendeesOutput, error)
	ListBots(ctx context.Context, params *chime.ListBotsInput, optFns ...func(*chime.Options)) (*chime.ListBotsOutput, error)
	ListChannelBans(ctx context.Context, params *chime.ListChannelBansInput, optFns ...func(*chime.Options)) (*chime.ListChannelBansOutput, error)
	ListChannelMemberships(ctx context.Context, params *chime.ListChannelMembershipsInput, optFns ...func(*chime.Options)) (*chime.ListChannelMembershipsOutput, error)
	ListChannelMembershipsForAppInstanceUser(ctx context.Context, params *chime.ListChannelMembershipsForAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.ListChannelMembershipsForAppInstanceUserOutput, error)
	ListChannelMessages(ctx context.Context, params *chime.ListChannelMessagesInput, optFns ...func(*chime.Options)) (*chime.ListChannelMessagesOutput, error)
	ListChannelModerators(ctx context.Context, params *chime.ListChannelModeratorsInput, optFns ...func(*chime.Options)) (*chime.ListChannelModeratorsOutput, error)
	ListChannels(ctx context.Context, params *chime.ListChannelsInput, optFns ...func(*chime.Options)) (*chime.ListChannelsOutput, error)
	ListChannelsModeratedByAppInstanceUser(ctx context.Context, params *chime.ListChannelsModeratedByAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.ListChannelsModeratedByAppInstanceUserOutput, error)
	ListMediaCapturePipelines(ctx context.Context, params *chime.ListMediaCapturePipelinesInput, optFns ...func(*chime.Options)) (*chime.ListMediaCapturePipelinesOutput, error)
	ListMeetingTags(ctx context.Context, params *chime.ListMeetingTagsInput, optFns ...func(*chime.Options)) (*chime.ListMeetingTagsOutput, error)
	ListMeetings(ctx context.Context, params *chime.ListMeetingsInput, optFns ...func(*chime.Options)) (*chime.ListMeetingsOutput, error)
	ListPhoneNumberOrders(ctx context.Context, params *chime.ListPhoneNumberOrdersInput, optFns ...func(*chime.Options)) (*chime.ListPhoneNumberOrdersOutput, error)
	ListPhoneNumbers(ctx context.Context, params *chime.ListPhoneNumbersInput, optFns ...func(*chime.Options)) (*chime.ListPhoneNumbersOutput, error)
	ListProxySessions(ctx context.Context, params *chime.ListProxySessionsInput, optFns ...func(*chime.Options)) (*chime.ListProxySessionsOutput, error)
	ListRoomMemberships(ctx context.Context, params *chime.ListRoomMembershipsInput, optFns ...func(*chime.Options)) (*chime.ListRoomMembershipsOutput, error)
	ListRooms(ctx context.Context, params *chime.ListRoomsInput, optFns ...func(*chime.Options)) (*chime.ListRoomsOutput, error)
	ListSipMediaApplications(ctx context.Context, params *chime.ListSipMediaApplicationsInput, optFns ...func(*chime.Options)) (*chime.ListSipMediaApplicationsOutput, error)
	ListSipRules(ctx context.Context, params *chime.ListSipRulesInput, optFns ...func(*chime.Options)) (*chime.ListSipRulesOutput, error)
	ListSupportedPhoneNumberCountries(ctx context.Context, params *chime.ListSupportedPhoneNumberCountriesInput, optFns ...func(*chime.Options)) (*chime.ListSupportedPhoneNumberCountriesOutput, error)
	ListTagsForResource(ctx context.Context, params *chime.ListTagsForResourceInput, optFns ...func(*chime.Options)) (*chime.ListTagsForResourceOutput, error)
	ListUsers(ctx context.Context, params *chime.ListUsersInput, optFns ...func(*chime.Options)) (*chime.ListUsersOutput, error)
	ListVoiceConnectorGroups(ctx context.Context, params *chime.ListVoiceConnectorGroupsInput, optFns ...func(*chime.Options)) (*chime.ListVoiceConnectorGroupsOutput, error)
	ListVoiceConnectorTerminationCredentials(ctx context.Context, params *chime.ListVoiceConnectorTerminationCredentialsInput, optFns ...func(*chime.Options)) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
	ListVoiceConnectors(ctx context.Context, params *chime.ListVoiceConnectorsInput, optFns ...func(*chime.Options)) (*chime.ListVoiceConnectorsOutput, error)
	LogoutUser(ctx context.Context, params *chime.LogoutUserInput, optFns ...func(*chime.Options)) (*chime.LogoutUserOutput, error)
	PutAppInstanceRetentionSettings(ctx context.Context, params *chime.PutAppInstanceRetentionSettingsInput, optFns ...func(*chime.Options)) (*chime.PutAppInstanceRetentionSettingsOutput, error)
	PutAppInstanceStreamingConfigurations(ctx context.Context, params *chime.PutAppInstanceStreamingConfigurationsInput, optFns ...func(*chime.Options)) (*chime.PutAppInstanceStreamingConfigurationsOutput, error)
	PutEventsConfiguration(ctx context.Context, params *chime.PutEventsConfigurationInput, optFns ...func(*chime.Options)) (*chime.PutEventsConfigurationOutput, error)
	PutRetentionSettings(ctx context.Context, params *chime.PutRetentionSettingsInput, optFns ...func(*chime.Options)) (*chime.PutRetentionSettingsOutput, error)
	PutSipMediaApplicationLoggingConfiguration(ctx context.Context, params *chime.PutSipMediaApplicationLoggingConfigurationInput, optFns ...func(*chime.Options)) (*chime.PutSipMediaApplicationLoggingConfigurationOutput, error)
	PutVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *chime.PutVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error)
	PutVoiceConnectorLoggingConfiguration(ctx context.Context, params *chime.PutVoiceConnectorLoggingConfigurationInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error)
	PutVoiceConnectorOrigination(ctx context.Context, params *chime.PutVoiceConnectorOriginationInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorOriginationOutput, error)
	PutVoiceConnectorProxy(ctx context.Context, params *chime.PutVoiceConnectorProxyInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorProxyOutput, error)
	PutVoiceConnectorStreamingConfiguration(ctx context.Context, params *chime.PutVoiceConnectorStreamingConfigurationInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error)
	PutVoiceConnectorTermination(ctx context.Context, params *chime.PutVoiceConnectorTerminationInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorTerminationOutput, error)
	PutVoiceConnectorTerminationCredentials(ctx context.Context, params *chime.PutVoiceConnectorTerminationCredentialsInput, optFns ...func(*chime.Options)) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
	RedactChannelMessage(ctx context.Context, params *chime.RedactChannelMessageInput, optFns ...func(*chime.Options)) (*chime.RedactChannelMessageOutput, error)
	RedactConversationMessage(ctx context.Context, params *chime.RedactConversationMessageInput, optFns ...func(*chime.Options)) (*chime.RedactConversationMessageOutput, error)
	RedactRoomMessage(ctx context.Context, params *chime.RedactRoomMessageInput, optFns ...func(*chime.Options)) (*chime.RedactRoomMessageOutput, error)
	RegenerateSecurityToken(ctx context.Context, params *chime.RegenerateSecurityTokenInput, optFns ...func(*chime.Options)) (*chime.RegenerateSecurityTokenOutput, error)
	ResetPersonalPIN(ctx context.Context, params *chime.ResetPersonalPINInput, optFns ...func(*chime.Options)) (*chime.ResetPersonalPINOutput, error)
	RestorePhoneNumber(ctx context.Context, params *chime.RestorePhoneNumberInput, optFns ...func(*chime.Options)) (*chime.RestorePhoneNumberOutput, error)
	SearchAvailablePhoneNumbers(ctx context.Context, params *chime.SearchAvailablePhoneNumbersInput, optFns ...func(*chime.Options)) (*chime.SearchAvailablePhoneNumbersOutput, error)
	SendChannelMessage(ctx context.Context, params *chime.SendChannelMessageInput, optFns ...func(*chime.Options)) (*chime.SendChannelMessageOutput, error)
	StartMeetingTranscription(ctx context.Context, params *chime.StartMeetingTranscriptionInput, optFns ...func(*chime.Options)) (*chime.StartMeetingTranscriptionOutput, error)
	StopMeetingTranscription(ctx context.Context, params *chime.StopMeetingTranscriptionInput, optFns ...func(*chime.Options)) (*chime.StopMeetingTranscriptionOutput, error)
	TagAttendee(ctx context.Context, params *chime.TagAttendeeInput, optFns ...func(*chime.Options)) (*chime.TagAttendeeOutput, error)
	TagMeeting(ctx context.Context, params *chime.TagMeetingInput, optFns ...func(*chime.Options)) (*chime.TagMeetingOutput, error)
	TagResource(ctx context.Context, params *chime.TagResourceInput, optFns ...func(*chime.Options)) (*chime.TagResourceOutput, error)
	UntagAttendee(ctx context.Context, params *chime.UntagAttendeeInput, optFns ...func(*chime.Options)) (*chime.UntagAttendeeOutput, error)
	UntagMeeting(ctx context.Context, params *chime.UntagMeetingInput, optFns ...func(*chime.Options)) (*chime.UntagMeetingOutput, error)
	UntagResource(ctx context.Context, params *chime.UntagResourceInput, optFns ...func(*chime.Options)) (*chime.UntagResourceOutput, error)
	UpdateAccount(ctx context.Context, params *chime.UpdateAccountInput, optFns ...func(*chime.Options)) (*chime.UpdateAccountOutput, error)
	UpdateAccountSettings(ctx context.Context, params *chime.UpdateAccountSettingsInput, optFns ...func(*chime.Options)) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAppInstance(ctx context.Context, params *chime.UpdateAppInstanceInput, optFns ...func(*chime.Options)) (*chime.UpdateAppInstanceOutput, error)
	UpdateAppInstanceUser(ctx context.Context, params *chime.UpdateAppInstanceUserInput, optFns ...func(*chime.Options)) (*chime.UpdateAppInstanceUserOutput, error)
	UpdateBot(ctx context.Context, params *chime.UpdateBotInput, optFns ...func(*chime.Options)) (*chime.UpdateBotOutput, error)
	UpdateChannel(ctx context.Context, params *chime.UpdateChannelInput, optFns ...func(*chime.Options)) (*chime.UpdateChannelOutput, error)
	UpdateChannelMessage(ctx context.Context, params *chime.UpdateChannelMessageInput, optFns ...func(*chime.Options)) (*chime.UpdateChannelMessageOutput, error)
	UpdateChannelReadMarker(ctx context.Context, params *chime.UpdateChannelReadMarkerInput, optFns ...func(*chime.Options)) (*chime.UpdateChannelReadMarkerOutput, error)
	UpdateGlobalSettings(ctx context.Context, params *chime.UpdateGlobalSettingsInput, optFns ...func(*chime.Options)) (*chime.UpdateGlobalSettingsOutput, error)
	UpdatePhoneNumber(ctx context.Context, params *chime.UpdatePhoneNumberInput, optFns ...func(*chime.Options)) (*chime.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberSettings(ctx context.Context, params *chime.UpdatePhoneNumberSettingsInput, optFns ...func(*chime.Options)) (*chime.UpdatePhoneNumberSettingsOutput, error)
	UpdateProxySession(ctx context.Context, params *chime.UpdateProxySessionInput, optFns ...func(*chime.Options)) (*chime.UpdateProxySessionOutput, error)
	UpdateRoom(ctx context.Context, params *chime.UpdateRoomInput, optFns ...func(*chime.Options)) (*chime.UpdateRoomOutput, error)
	UpdateRoomMembership(ctx context.Context, params *chime.UpdateRoomMembershipInput, optFns ...func(*chime.Options)) (*chime.UpdateRoomMembershipOutput, error)
	UpdateSipMediaApplication(ctx context.Context, params *chime.UpdateSipMediaApplicationInput, optFns ...func(*chime.Options)) (*chime.UpdateSipMediaApplicationOutput, error)
	UpdateSipMediaApplicationCall(ctx context.Context, params *chime.UpdateSipMediaApplicationCallInput, optFns ...func(*chime.Options)) (*chime.UpdateSipMediaApplicationCallOutput, error)
	UpdateSipRule(ctx context.Context, params *chime.UpdateSipRuleInput, optFns ...func(*chime.Options)) (*chime.UpdateSipRuleOutput, error)
	UpdateUser(ctx context.Context, params *chime.UpdateUserInput, optFns ...func(*chime.Options)) (*chime.UpdateUserOutput, error)
	UpdateUserSettings(ctx context.Context, params *chime.UpdateUserSettingsInput, optFns ...func(*chime.Options)) (*chime.UpdateUserSettingsOutput, error)
	UpdateVoiceConnector(ctx context.Context, params *chime.UpdateVoiceConnectorInput, optFns ...func(*chime.Options)) (*chime.UpdateVoiceConnectorOutput, error)
	UpdateVoiceConnectorGroup(ctx context.Context, params *chime.UpdateVoiceConnectorGroupInput, optFns ...func(*chime.Options)) (*chime.UpdateVoiceConnectorGroupOutput, error)
	ValidateE911Address(ctx context.Context, params *chime.ValidateE911AddressInput, optFns ...func(*chime.Options)) (*chime.ValidateE911AddressOutput, error)
}
