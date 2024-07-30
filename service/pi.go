// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/pi"
)

// PiClient ...
type PiClient interface {
	Options() pi.Options
	CreatePerformanceAnalysisReport(ctx context.Context, params *pi.CreatePerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.CreatePerformanceAnalysisReportOutput, error)
	DeletePerformanceAnalysisReport(ctx context.Context, params *pi.DeletePerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.DeletePerformanceAnalysisReportOutput, error)
	DescribeDimensionKeys(ctx context.Context, params *pi.DescribeDimensionKeysInput, optFns ...func(*pi.Options)) (*pi.DescribeDimensionKeysOutput, error)
	GetDimensionKeyDetails(ctx context.Context, params *pi.GetDimensionKeyDetailsInput, optFns ...func(*pi.Options)) (*pi.GetDimensionKeyDetailsOutput, error)
	GetPerformanceAnalysisReport(ctx context.Context, params *pi.GetPerformanceAnalysisReportInput, optFns ...func(*pi.Options)) (*pi.GetPerformanceAnalysisReportOutput, error)
	GetResourceMetadata(ctx context.Context, params *pi.GetResourceMetadataInput, optFns ...func(*pi.Options)) (*pi.GetResourceMetadataOutput, error)
	GetResourceMetrics(ctx context.Context, params *pi.GetResourceMetricsInput, optFns ...func(*pi.Options)) (*pi.GetResourceMetricsOutput, error)
	ListAvailableResourceDimensions(ctx context.Context, params *pi.ListAvailableResourceDimensionsInput, optFns ...func(*pi.Options)) (*pi.ListAvailableResourceDimensionsOutput, error)
	ListAvailableResourceMetrics(ctx context.Context, params *pi.ListAvailableResourceMetricsInput, optFns ...func(*pi.Options)) (*pi.ListAvailableResourceMetricsOutput, error)
	ListPerformanceAnalysisReports(ctx context.Context, params *pi.ListPerformanceAnalysisReportsInput, optFns ...func(*pi.Options)) (*pi.ListPerformanceAnalysisReportsOutput, error)
	ListTagsForResource(ctx context.Context, params *pi.ListTagsForResourceInput, optFns ...func(*pi.Options)) (*pi.ListTagsForResourceOutput, error)
	TagResource(ctx context.Context, params *pi.TagResourceInput, optFns ...func(*pi.Options)) (*pi.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *pi.UntagResourceInput, optFns ...func(*pi.Options)) (*pi.UntagResourceOutput, error)
}
