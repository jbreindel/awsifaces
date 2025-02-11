// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
)

// LookoutequipmentClient ...
type LookoutequipmentClient interface {
	Options() lookoutequipment.Options
	CreateDataset(ctx context.Context, params *lookoutequipment.CreateDatasetInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateDatasetOutput, error)
	CreateInferenceScheduler(ctx context.Context, params *lookoutequipment.CreateInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateInferenceSchedulerOutput, error)
	CreateLabel(ctx context.Context, params *lookoutequipment.CreateLabelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateLabelOutput, error)
	CreateLabelGroup(ctx context.Context, params *lookoutequipment.CreateLabelGroupInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateLabelGroupOutput, error)
	CreateModel(ctx context.Context, params *lookoutequipment.CreateModelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateModelOutput, error)
	CreateRetrainingScheduler(ctx context.Context, params *lookoutequipment.CreateRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.CreateRetrainingSchedulerOutput, error)
	DeleteDataset(ctx context.Context, params *lookoutequipment.DeleteDatasetInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteDatasetOutput, error)
	DeleteInferenceScheduler(ctx context.Context, params *lookoutequipment.DeleteInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteInferenceSchedulerOutput, error)
	DeleteLabel(ctx context.Context, params *lookoutequipment.DeleteLabelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteLabelOutput, error)
	DeleteLabelGroup(ctx context.Context, params *lookoutequipment.DeleteLabelGroupInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteLabelGroupOutput, error)
	DeleteModel(ctx context.Context, params *lookoutequipment.DeleteModelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteModelOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *lookoutequipment.DeleteResourcePolicyInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteResourcePolicyOutput, error)
	DeleteRetrainingScheduler(ctx context.Context, params *lookoutequipment.DeleteRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DeleteRetrainingSchedulerOutput, error)
	DescribeDataIngestionJob(ctx context.Context, params *lookoutequipment.DescribeDataIngestionJobInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeDataIngestionJobOutput, error)
	DescribeDataset(ctx context.Context, params *lookoutequipment.DescribeDatasetInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeDatasetOutput, error)
	DescribeInferenceScheduler(ctx context.Context, params *lookoutequipment.DescribeInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeInferenceSchedulerOutput, error)
	DescribeLabel(ctx context.Context, params *lookoutequipment.DescribeLabelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeLabelOutput, error)
	DescribeLabelGroup(ctx context.Context, params *lookoutequipment.DescribeLabelGroupInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeLabelGroupOutput, error)
	DescribeModel(ctx context.Context, params *lookoutequipment.DescribeModelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeModelOutput, error)
	DescribeModelVersion(ctx context.Context, params *lookoutequipment.DescribeModelVersionInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeModelVersionOutput, error)
	DescribeResourcePolicy(ctx context.Context, params *lookoutequipment.DescribeResourcePolicyInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeResourcePolicyOutput, error)
	DescribeRetrainingScheduler(ctx context.Context, params *lookoutequipment.DescribeRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.DescribeRetrainingSchedulerOutput, error)
	ImportDataset(ctx context.Context, params *lookoutequipment.ImportDatasetInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ImportDatasetOutput, error)
	ImportModelVersion(ctx context.Context, params *lookoutequipment.ImportModelVersionInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ImportModelVersionOutput, error)
	ListDataIngestionJobs(ctx context.Context, params *lookoutequipment.ListDataIngestionJobsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListDataIngestionJobsOutput, error)
	ListDatasets(ctx context.Context, params *lookoutequipment.ListDatasetsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListDatasetsOutput, error)
	ListInferenceEvents(ctx context.Context, params *lookoutequipment.ListInferenceEventsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListInferenceEventsOutput, error)
	ListInferenceExecutions(ctx context.Context, params *lookoutequipment.ListInferenceExecutionsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListInferenceExecutionsOutput, error)
	ListInferenceSchedulers(ctx context.Context, params *lookoutequipment.ListInferenceSchedulersInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListInferenceSchedulersOutput, error)
	ListLabelGroups(ctx context.Context, params *lookoutequipment.ListLabelGroupsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListLabelGroupsOutput, error)
	ListLabels(ctx context.Context, params *lookoutequipment.ListLabelsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListLabelsOutput, error)
	ListModelVersions(ctx context.Context, params *lookoutequipment.ListModelVersionsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListModelVersionsOutput, error)
	ListModels(ctx context.Context, params *lookoutequipment.ListModelsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListModelsOutput, error)
	ListRetrainingSchedulers(ctx context.Context, params *lookoutequipment.ListRetrainingSchedulersInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListRetrainingSchedulersOutput, error)
	ListSensorStatistics(ctx context.Context, params *lookoutequipment.ListSensorStatisticsInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListSensorStatisticsOutput, error)
	ListTagsForResource(ctx context.Context, params *lookoutequipment.ListTagsForResourceInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.ListTagsForResourceOutput, error)
	PutResourcePolicy(ctx context.Context, params *lookoutequipment.PutResourcePolicyInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.PutResourcePolicyOutput, error)
	StartDataIngestionJob(ctx context.Context, params *lookoutequipment.StartDataIngestionJobInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.StartDataIngestionJobOutput, error)
	StartInferenceScheduler(ctx context.Context, params *lookoutequipment.StartInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.StartInferenceSchedulerOutput, error)
	StartRetrainingScheduler(ctx context.Context, params *lookoutequipment.StartRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.StartRetrainingSchedulerOutput, error)
	StopInferenceScheduler(ctx context.Context, params *lookoutequipment.StopInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.StopInferenceSchedulerOutput, error)
	StopRetrainingScheduler(ctx context.Context, params *lookoutequipment.StopRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.StopRetrainingSchedulerOutput, error)
	TagResource(ctx context.Context, params *lookoutequipment.TagResourceInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *lookoutequipment.UntagResourceInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UntagResourceOutput, error)
	UpdateActiveModelVersion(ctx context.Context, params *lookoutequipment.UpdateActiveModelVersionInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UpdateActiveModelVersionOutput, error)
	UpdateInferenceScheduler(ctx context.Context, params *lookoutequipment.UpdateInferenceSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UpdateInferenceSchedulerOutput, error)
	UpdateLabelGroup(ctx context.Context, params *lookoutequipment.UpdateLabelGroupInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UpdateLabelGroupOutput, error)
	UpdateModel(ctx context.Context, params *lookoutequipment.UpdateModelInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UpdateModelOutput, error)
	UpdateRetrainingScheduler(ctx context.Context, params *lookoutequipment.UpdateRetrainingSchedulerInput, optFns ...func(*lookoutequipment.Options)) (*lookoutequipment.UpdateRetrainingSchedulerOutput, error)
}
