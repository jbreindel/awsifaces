// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

// Cloudhsmv2Client ...
type Cloudhsmv2Client interface {
	Options() cloudhsmv2.Options
	CopyBackupToRegion(ctx context.Context, params *cloudhsmv2.CopyBackupToRegionInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.CopyBackupToRegionOutput, error)
	CreateCluster(ctx context.Context, params *cloudhsmv2.CreateClusterInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.CreateClusterOutput, error)
	CreateHsm(ctx context.Context, params *cloudhsmv2.CreateHsmInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.CreateHsmOutput, error)
	DeleteBackup(ctx context.Context, params *cloudhsmv2.DeleteBackupInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DeleteBackupOutput, error)
	DeleteCluster(ctx context.Context, params *cloudhsmv2.DeleteClusterInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DeleteClusterOutput, error)
	DeleteHsm(ctx context.Context, params *cloudhsmv2.DeleteHsmInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DeleteHsmOutput, error)
	DescribeBackups(ctx context.Context, params *cloudhsmv2.DescribeBackupsInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeClusters(ctx context.Context, params *cloudhsmv2.DescribeClustersInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error)
	InitializeCluster(ctx context.Context, params *cloudhsmv2.InitializeClusterInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.InitializeClusterOutput, error)
	ListTags(ctx context.Context, params *cloudhsmv2.ListTagsInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ListTagsOutput, error)
	ModifyBackupAttributes(ctx context.Context, params *cloudhsmv2.ModifyBackupAttributesInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ModifyBackupAttributesOutput, error)
	ModifyCluster(ctx context.Context, params *cloudhsmv2.ModifyClusterInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ModifyClusterOutput, error)
	RestoreBackup(ctx context.Context, params *cloudhsmv2.RestoreBackupInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.RestoreBackupOutput, error)
	TagResource(ctx context.Context, params *cloudhsmv2.TagResourceInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *cloudhsmv2.UntagResourceInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.UntagResourceOutput, error)
}
