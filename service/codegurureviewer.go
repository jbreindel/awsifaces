// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
)

// CodegurureviewerClient ...
type CodegurureviewerClient interface {
	Options() codegurureviewer.Options
	AssociateRepository(ctx context.Context, params *codegurureviewer.AssociateRepositoryInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.AssociateRepositoryOutput, error)
	CreateCodeReview(ctx context.Context, params *codegurureviewer.CreateCodeReviewInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.CreateCodeReviewOutput, error)
	DescribeCodeReview(ctx context.Context, params *codegurureviewer.DescribeCodeReviewInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.DescribeCodeReviewOutput, error)
	DescribeRecommendationFeedback(ctx context.Context, params *codegurureviewer.DescribeRecommendationFeedbackInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
	DescribeRepositoryAssociation(ctx context.Context, params *codegurureviewer.DescribeRepositoryAssociationInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DisassociateRepository(ctx context.Context, params *codegurureviewer.DisassociateRepositoryInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.DisassociateRepositoryOutput, error)
	ListCodeReviews(ctx context.Context, params *codegurureviewer.ListCodeReviewsInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.ListCodeReviewsOutput, error)
	ListRecommendationFeedback(ctx context.Context, params *codegurureviewer.ListRecommendationFeedbackInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
	ListRecommendations(ctx context.Context, params *codegurureviewer.ListRecommendationsInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.ListRecommendationsOutput, error)
	ListRepositoryAssociations(ctx context.Context, params *codegurureviewer.ListRepositoryAssociationsInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListTagsForResource(ctx context.Context, params *codegurureviewer.ListTagsForResourceInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.ListTagsForResourceOutput, error)
	PutRecommendationFeedback(ctx context.Context, params *codegurureviewer.PutRecommendationFeedbackInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
	TagResource(ctx context.Context, params *codegurureviewer.TagResourceInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *codegurureviewer.UntagResourceInput, optFns ...func(*codegurureviewer.Options)) (*codegurureviewer.UntagResourceOutput, error)
}
