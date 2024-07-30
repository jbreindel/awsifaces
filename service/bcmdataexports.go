// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
)

// BcmdataexportsClient ...
type BcmdataexportsClient interface {
	Options() bcmdataexports.Options
	CreateExport(ctx context.Context, params *bcmdataexports.CreateExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.CreateExportOutput, error)
	DeleteExport(ctx context.Context, params *bcmdataexports.DeleteExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.DeleteExportOutput, error)
	GetExecution(ctx context.Context, params *bcmdataexports.GetExecutionInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExecutionOutput, error)
	GetExport(ctx context.Context, params *bcmdataexports.GetExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetExportOutput, error)
	GetTable(ctx context.Context, params *bcmdataexports.GetTableInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.GetTableOutput, error)
	ListExecutions(ctx context.Context, params *bcmdataexports.ListExecutionsInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExecutionsOutput, error)
	ListExports(ctx context.Context, params *bcmdataexports.ListExportsInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListExportsOutput, error)
	ListTables(ctx context.Context, params *bcmdataexports.ListTablesInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTablesOutput, error)
	ListTagsForResource(ctx context.Context, params *bcmdataexports.ListTagsForResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.ListTagsForResourceOutput, error)
	TagResource(ctx context.Context, params *bcmdataexports.TagResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *bcmdataexports.UntagResourceInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.UntagResourceOutput, error)
	UpdateExport(ctx context.Context, params *bcmdataexports.UpdateExportInput, optFns ...func(*bcmdataexports.Options)) (*bcmdataexports.UpdateExportOutput, error)
}
