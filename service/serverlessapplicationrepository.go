// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
)

// ServerlessapplicationrepositoryClient ...
type ServerlessapplicationrepositoryClient interface {
	Options() serverlessapplicationrepository.Options
	CreateApplication(ctx context.Context, params *serverlessapplicationrepository.CreateApplicationInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.CreateApplicationOutput, error)
	CreateApplicationVersion(ctx context.Context, params *serverlessapplicationrepository.CreateApplicationVersionInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
	CreateCloudFormationChangeSet(ctx context.Context, params *serverlessapplicationrepository.CreateCloudFormationChangeSetInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
	CreateCloudFormationTemplate(ctx context.Context, params *serverlessapplicationrepository.CreateCloudFormationTemplateInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
	DeleteApplication(ctx context.Context, params *serverlessapplicationrepository.DeleteApplicationInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
	GetApplication(ctx context.Context, params *serverlessapplicationrepository.GetApplicationInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.GetApplicationOutput, error)
	GetApplicationPolicy(ctx context.Context, params *serverlessapplicationrepository.GetApplicationPolicyInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
	GetCloudFormationTemplate(ctx context.Context, params *serverlessapplicationrepository.GetCloudFormationTemplateInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
	ListApplicationDependencies(ctx context.Context, params *serverlessapplicationrepository.ListApplicationDependenciesInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
	ListApplicationVersions(ctx context.Context, params *serverlessapplicationrepository.ListApplicationVersionsInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
	ListApplications(ctx context.Context, params *serverlessapplicationrepository.ListApplicationsInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.ListApplicationsOutput, error)
	PutApplicationPolicy(ctx context.Context, params *serverlessapplicationrepository.PutApplicationPolicyInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
	UnshareApplication(ctx context.Context, params *serverlessapplicationrepository.UnshareApplicationInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.UnshareApplicationOutput, error)
	UpdateApplication(ctx context.Context, params *serverlessapplicationrepository.UpdateApplicationInput, optFns ...func(*serverlessapplicationrepository.Options)) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
}
