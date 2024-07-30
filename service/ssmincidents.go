// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssmincidents"
)

// SsmincidentsClient ...
type SsmincidentsClient interface {
	Options() ssmincidents.Options
	BatchGetIncidentFindings(ctx context.Context, params *ssmincidents.BatchGetIncidentFindingsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.BatchGetIncidentFindingsOutput, error)
	CreateReplicationSet(ctx context.Context, params *ssmincidents.CreateReplicationSetInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.CreateReplicationSetOutput, error)
	CreateResponsePlan(ctx context.Context, params *ssmincidents.CreateResponsePlanInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.CreateResponsePlanOutput, error)
	CreateTimelineEvent(ctx context.Context, params *ssmincidents.CreateTimelineEventInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.CreateTimelineEventOutput, error)
	DeleteIncidentRecord(ctx context.Context, params *ssmincidents.DeleteIncidentRecordInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.DeleteIncidentRecordOutput, error)
	DeleteReplicationSet(ctx context.Context, params *ssmincidents.DeleteReplicationSetInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.DeleteReplicationSetOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *ssmincidents.DeleteResourcePolicyInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.DeleteResourcePolicyOutput, error)
	DeleteResponsePlan(ctx context.Context, params *ssmincidents.DeleteResponsePlanInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.DeleteResponsePlanOutput, error)
	DeleteTimelineEvent(ctx context.Context, params *ssmincidents.DeleteTimelineEventInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.DeleteTimelineEventOutput, error)
	GetIncidentRecord(ctx context.Context, params *ssmincidents.GetIncidentRecordInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.GetIncidentRecordOutput, error)
	GetReplicationSet(ctx context.Context, params *ssmincidents.GetReplicationSetInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.GetReplicationSetOutput, error)
	GetResourcePolicies(ctx context.Context, params *ssmincidents.GetResourcePoliciesInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.GetResourcePoliciesOutput, error)
	GetResponsePlan(ctx context.Context, params *ssmincidents.GetResponsePlanInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.GetResponsePlanOutput, error)
	GetTimelineEvent(ctx context.Context, params *ssmincidents.GetTimelineEventInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.GetTimelineEventOutput, error)
	ListIncidentFindings(ctx context.Context, params *ssmincidents.ListIncidentFindingsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListIncidentFindingsOutput, error)
	ListIncidentRecords(ctx context.Context, params *ssmincidents.ListIncidentRecordsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListIncidentRecordsOutput, error)
	ListRelatedItems(ctx context.Context, params *ssmincidents.ListRelatedItemsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListRelatedItemsOutput, error)
	ListReplicationSets(ctx context.Context, params *ssmincidents.ListReplicationSetsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListReplicationSetsOutput, error)
	ListResponsePlans(ctx context.Context, params *ssmincidents.ListResponsePlansInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListResponsePlansOutput, error)
	ListTagsForResource(ctx context.Context, params *ssmincidents.ListTagsForResourceInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListTagsForResourceOutput, error)
	ListTimelineEvents(ctx context.Context, params *ssmincidents.ListTimelineEventsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.ListTimelineEventsOutput, error)
	PutResourcePolicy(ctx context.Context, params *ssmincidents.PutResourcePolicyInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.PutResourcePolicyOutput, error)
	StartIncident(ctx context.Context, params *ssmincidents.StartIncidentInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.StartIncidentOutput, error)
	TagResource(ctx context.Context, params *ssmincidents.TagResourceInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *ssmincidents.UntagResourceInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UntagResourceOutput, error)
	UpdateDeletionProtection(ctx context.Context, params *ssmincidents.UpdateDeletionProtectionInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateDeletionProtectionOutput, error)
	UpdateIncidentRecord(ctx context.Context, params *ssmincidents.UpdateIncidentRecordInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateIncidentRecordOutput, error)
	UpdateRelatedItems(ctx context.Context, params *ssmincidents.UpdateRelatedItemsInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateRelatedItemsOutput, error)
	UpdateReplicationSet(ctx context.Context, params *ssmincidents.UpdateReplicationSetInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateReplicationSetOutput, error)
	UpdateResponsePlan(ctx context.Context, params *ssmincidents.UpdateResponsePlanInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateResponsePlanOutput, error)
	UpdateTimelineEvent(ctx context.Context, params *ssmincidents.UpdateTimelineEventInput, optFns ...func(*ssmincidents.Options)) (*ssmincidents.UpdateTimelineEventOutput, error)
}
