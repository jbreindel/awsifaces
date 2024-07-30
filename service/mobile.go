// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mobile"
)

// MobileClient ...
type MobileClient interface {
	Options() mobile.Options
	CreateProject(ctx context.Context, params *mobile.CreateProjectInput, optFns ...func(*mobile.Options)) (*mobile.CreateProjectOutput, error)
	DeleteProject(ctx context.Context, params *mobile.DeleteProjectInput, optFns ...func(*mobile.Options)) (*mobile.DeleteProjectOutput, error)
	DescribeBundle(ctx context.Context, params *mobile.DescribeBundleInput, optFns ...func(*mobile.Options)) (*mobile.DescribeBundleOutput, error)
	DescribeProject(ctx context.Context, params *mobile.DescribeProjectInput, optFns ...func(*mobile.Options)) (*mobile.DescribeProjectOutput, error)
	ExportBundle(ctx context.Context, params *mobile.ExportBundleInput, optFns ...func(*mobile.Options)) (*mobile.ExportBundleOutput, error)
	ExportProject(ctx context.Context, params *mobile.ExportProjectInput, optFns ...func(*mobile.Options)) (*mobile.ExportProjectOutput, error)
	ListBundles(ctx context.Context, params *mobile.ListBundlesInput, optFns ...func(*mobile.Options)) (*mobile.ListBundlesOutput, error)
	ListProjects(ctx context.Context, params *mobile.ListProjectsInput, optFns ...func(*mobile.Options)) (*mobile.ListProjectsOutput, error)
	UpdateProject(ctx context.Context, params *mobile.UpdateProjectInput, optFns ...func(*mobile.Options)) (*mobile.UpdateProjectOutput, error)
}
