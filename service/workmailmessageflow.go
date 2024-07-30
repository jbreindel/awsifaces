// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
)

// WorkmailmessageflowClient ...
type WorkmailmessageflowClient interface {
	Options() workmailmessageflow.Options
	GetRawMessageContent(ctx context.Context, params *workmailmessageflow.GetRawMessageContentInput, optFns ...func(*workmailmessageflow.Options)) (*workmailmessageflow.GetRawMessageContentOutput, error)
	PutRawMessageContent(ctx context.Context, params *workmailmessageflow.PutRawMessageContentInput, optFns ...func(*workmailmessageflow.Options)) (*workmailmessageflow.PutRawMessageContentOutput, error)
}
