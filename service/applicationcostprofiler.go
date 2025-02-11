// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
)

// ApplicationcostprofilerClient ...
type ApplicationcostprofilerClient interface {
	Options() applicationcostprofiler.Options
	DeleteReportDefinition(ctx context.Context, params *applicationcostprofiler.DeleteReportDefinitionInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.DeleteReportDefinitionOutput, error)
	GetReportDefinition(ctx context.Context, params *applicationcostprofiler.GetReportDefinitionInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.GetReportDefinitionOutput, error)
	ImportApplicationUsage(ctx context.Context, params *applicationcostprofiler.ImportApplicationUsageInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.ImportApplicationUsageOutput, error)
	ListReportDefinitions(ctx context.Context, params *applicationcostprofiler.ListReportDefinitionsInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.ListReportDefinitionsOutput, error)
	PutReportDefinition(ctx context.Context, params *applicationcostprofiler.PutReportDefinitionInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.PutReportDefinitionOutput, error)
	UpdateReportDefinition(ctx context.Context, params *applicationcostprofiler.UpdateReportDefinitionInput, optFns ...func(*applicationcostprofiler.Options)) (*applicationcostprofiler.UpdateReportDefinitionOutput, error)
}
