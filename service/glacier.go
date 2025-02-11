// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
)

// GlacierClient ...
type GlacierClient interface {
	Options() glacier.Options
	AbortMultipartUpload(ctx context.Context, params *glacier.AbortMultipartUploadInput, optFns ...func(*glacier.Options)) (*glacier.AbortMultipartUploadOutput, error)
	AbortVaultLock(ctx context.Context, params *glacier.AbortVaultLockInput, optFns ...func(*glacier.Options)) (*glacier.AbortVaultLockOutput, error)
	AddTagsToVault(ctx context.Context, params *glacier.AddTagsToVaultInput, optFns ...func(*glacier.Options)) (*glacier.AddTagsToVaultOutput, error)
	CompleteMultipartUpload(ctx context.Context, params *glacier.CompleteMultipartUploadInput, optFns ...func(*glacier.Options)) (*glacier.CompleteMultipartUploadOutput, error)
	CompleteVaultLock(ctx context.Context, params *glacier.CompleteVaultLockInput, optFns ...func(*glacier.Options)) (*glacier.CompleteVaultLockOutput, error)
	CreateVault(ctx context.Context, params *glacier.CreateVaultInput, optFns ...func(*glacier.Options)) (*glacier.CreateVaultOutput, error)
	DeleteArchive(ctx context.Context, params *glacier.DeleteArchiveInput, optFns ...func(*glacier.Options)) (*glacier.DeleteArchiveOutput, error)
	DeleteVault(ctx context.Context, params *glacier.DeleteVaultInput, optFns ...func(*glacier.Options)) (*glacier.DeleteVaultOutput, error)
	DeleteVaultAccessPolicy(ctx context.Context, params *glacier.DeleteVaultAccessPolicyInput, optFns ...func(*glacier.Options)) (*glacier.DeleteVaultAccessPolicyOutput, error)
	DeleteVaultNotifications(ctx context.Context, params *glacier.DeleteVaultNotificationsInput, optFns ...func(*glacier.Options)) (*glacier.DeleteVaultNotificationsOutput, error)
	DescribeJob(ctx context.Context, params *glacier.DescribeJobInput, optFns ...func(*glacier.Options)) (*glacier.DescribeJobOutput, error)
	DescribeVault(ctx context.Context, params *glacier.DescribeVaultInput, optFns ...func(*glacier.Options)) (*glacier.DescribeVaultOutput, error)
	GetDataRetrievalPolicy(ctx context.Context, params *glacier.GetDataRetrievalPolicyInput, optFns ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error)
	GetJobOutput(ctx context.Context, params *glacier.GetJobOutputInput, optFns ...func(*glacier.Options)) (*glacier.GetJobOutputOutput, error)
	GetVaultAccessPolicy(ctx context.Context, params *glacier.GetVaultAccessPolicyInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultLock(ctx context.Context, params *glacier.GetVaultLockInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error)
	GetVaultNotifications(ctx context.Context, params *glacier.GetVaultNotificationsInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error)
	InitiateJob(ctx context.Context, params *glacier.InitiateJobInput, optFns ...func(*glacier.Options)) (*glacier.InitiateJobOutput, error)
	InitiateMultipartUpload(ctx context.Context, params *glacier.InitiateMultipartUploadInput, optFns ...func(*glacier.Options)) (*glacier.InitiateMultipartUploadOutput, error)
	InitiateVaultLock(ctx context.Context, params *glacier.InitiateVaultLockInput, optFns ...func(*glacier.Options)) (*glacier.InitiateVaultLockOutput, error)
	ListJobs(ctx context.Context, params *glacier.ListJobsInput, optFns ...func(*glacier.Options)) (*glacier.ListJobsOutput, error)
	ListMultipartUploads(ctx context.Context, params *glacier.ListMultipartUploadsInput, optFns ...func(*glacier.Options)) (*glacier.ListMultipartUploadsOutput, error)
	ListParts(ctx context.Context, params *glacier.ListPartsInput, optFns ...func(*glacier.Options)) (*glacier.ListPartsOutput, error)
	ListProvisionedCapacity(ctx context.Context, params *glacier.ListProvisionedCapacityInput, optFns ...func(*glacier.Options)) (*glacier.ListProvisionedCapacityOutput, error)
	ListTagsForVault(ctx context.Context, params *glacier.ListTagsForVaultInput, optFns ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error)
	ListVaults(ctx context.Context, params *glacier.ListVaultsInput, optFns ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error)
	PurchaseProvisionedCapacity(ctx context.Context, params *glacier.PurchaseProvisionedCapacityInput, optFns ...func(*glacier.Options)) (*glacier.PurchaseProvisionedCapacityOutput, error)
	RemoveTagsFromVault(ctx context.Context, params *glacier.RemoveTagsFromVaultInput, optFns ...func(*glacier.Options)) (*glacier.RemoveTagsFromVaultOutput, error)
	SetDataRetrievalPolicy(ctx context.Context, params *glacier.SetDataRetrievalPolicyInput, optFns ...func(*glacier.Options)) (*glacier.SetDataRetrievalPolicyOutput, error)
	SetVaultAccessPolicy(ctx context.Context, params *glacier.SetVaultAccessPolicyInput, optFns ...func(*glacier.Options)) (*glacier.SetVaultAccessPolicyOutput, error)
	SetVaultNotifications(ctx context.Context, params *glacier.SetVaultNotificationsInput, optFns ...func(*glacier.Options)) (*glacier.SetVaultNotificationsOutput, error)
	UploadArchive(ctx context.Context, params *glacier.UploadArchiveInput, optFns ...func(*glacier.Options)) (*glacier.UploadArchiveOutput, error)
	UploadMultipartPart(ctx context.Context, params *glacier.UploadMultipartPartInput, optFns ...func(*glacier.Options)) (*glacier.UploadMultipartPartOutput, error)
}
