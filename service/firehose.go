// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

// FirehoseClient ...
type FirehoseClient interface {
	Options() firehose.Options
	CreateDeliveryStream(ctx context.Context, params *firehose.CreateDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.CreateDeliveryStreamOutput, error)
	DeleteDeliveryStream(ctx context.Context, params *firehose.DeleteDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.DeleteDeliveryStreamOutput, error)
	DescribeDeliveryStream(ctx context.Context, params *firehose.DescribeDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.DescribeDeliveryStreamOutput, error)
	ListDeliveryStreams(ctx context.Context, params *firehose.ListDeliveryStreamsInput, optFns ...func(*firehose.Options)) (*firehose.ListDeliveryStreamsOutput, error)
	ListTagsForDeliveryStream(ctx context.Context, params *firehose.ListTagsForDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.ListTagsForDeliveryStreamOutput, error)
	PutRecord(ctx context.Context, params *firehose.PutRecordInput, optFns ...func(*firehose.Options)) (*firehose.PutRecordOutput, error)
	PutRecordBatch(ctx context.Context, params *firehose.PutRecordBatchInput, optFns ...func(*firehose.Options)) (*firehose.PutRecordBatchOutput, error)
	StartDeliveryStreamEncryption(ctx context.Context, params *firehose.StartDeliveryStreamEncryptionInput, optFns ...func(*firehose.Options)) (*firehose.StartDeliveryStreamEncryptionOutput, error)
	StopDeliveryStreamEncryption(ctx context.Context, params *firehose.StopDeliveryStreamEncryptionInput, optFns ...func(*firehose.Options)) (*firehose.StopDeliveryStreamEncryptionOutput, error)
	TagDeliveryStream(ctx context.Context, params *firehose.TagDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.TagDeliveryStreamOutput, error)
	UntagDeliveryStream(ctx context.Context, params *firehose.UntagDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.UntagDeliveryStreamOutput, error)
	UpdateDestination(ctx context.Context, params *firehose.UpdateDestinationInput, optFns ...func(*firehose.Options)) (*firehose.UpdateDestinationOutput, error)
}
