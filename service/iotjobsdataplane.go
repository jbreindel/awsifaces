// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
)

// IotjobsdataplaneClient ...
type IotjobsdataplaneClient interface {
	Options() iotjobsdataplane.Options
	DescribeJobExecution(ctx context.Context, params *iotjobsdataplane.DescribeJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.DescribeJobExecutionOutput, error)
	GetPendingJobExecutions(ctx context.Context, params *iotjobsdataplane.GetPendingJobExecutionsInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error)
	StartNextPendingJobExecution(ctx context.Context, params *iotjobsdataplane.StartNextPendingJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error)
	UpdateJobExecution(ctx context.Context, params *iotjobsdataplane.UpdateJobExecutionInput, optFns ...func(*iotjobsdataplane.Options)) (*iotjobsdataplane.UpdateJobExecutionOutput, error)
}
