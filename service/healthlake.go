// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/healthlake"
)

// HealthlakeClient ...
type HealthlakeClient interface {
	Options() healthlake.Options
	CreateFHIRDatastore(ctx context.Context, params *healthlake.CreateFHIRDatastoreInput, optFns ...func(*healthlake.Options)) (*healthlake.CreateFHIRDatastoreOutput, error)
	DeleteFHIRDatastore(ctx context.Context, params *healthlake.DeleteFHIRDatastoreInput, optFns ...func(*healthlake.Options)) (*healthlake.DeleteFHIRDatastoreOutput, error)
	DescribeFHIRDatastore(ctx context.Context, params *healthlake.DescribeFHIRDatastoreInput, optFns ...func(*healthlake.Options)) (*healthlake.DescribeFHIRDatastoreOutput, error)
	DescribeFHIRExportJob(ctx context.Context, params *healthlake.DescribeFHIRExportJobInput, optFns ...func(*healthlake.Options)) (*healthlake.DescribeFHIRExportJobOutput, error)
	DescribeFHIRImportJob(ctx context.Context, params *healthlake.DescribeFHIRImportJobInput, optFns ...func(*healthlake.Options)) (*healthlake.DescribeFHIRImportJobOutput, error)
	ListFHIRDatastores(ctx context.Context, params *healthlake.ListFHIRDatastoresInput, optFns ...func(*healthlake.Options)) (*healthlake.ListFHIRDatastoresOutput, error)
	ListFHIRExportJobs(ctx context.Context, params *healthlake.ListFHIRExportJobsInput, optFns ...func(*healthlake.Options)) (*healthlake.ListFHIRExportJobsOutput, error)
	ListFHIRImportJobs(ctx context.Context, params *healthlake.ListFHIRImportJobsInput, optFns ...func(*healthlake.Options)) (*healthlake.ListFHIRImportJobsOutput, error)
	ListTagsForResource(ctx context.Context, params *healthlake.ListTagsForResourceInput, optFns ...func(*healthlake.Options)) (*healthlake.ListTagsForResourceOutput, error)
	StartFHIRExportJob(ctx context.Context, params *healthlake.StartFHIRExportJobInput, optFns ...func(*healthlake.Options)) (*healthlake.StartFHIRExportJobOutput, error)
	StartFHIRImportJob(ctx context.Context, params *healthlake.StartFHIRImportJobInput, optFns ...func(*healthlake.Options)) (*healthlake.StartFHIRImportJobOutput, error)
	TagResource(ctx context.Context, params *healthlake.TagResourceInput, optFns ...func(*healthlake.Options)) (*healthlake.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *healthlake.UntagResourceInput, optFns ...func(*healthlake.Options)) (*healthlake.UntagResourceOutput, error)
}
