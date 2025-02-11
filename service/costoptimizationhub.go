// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
)

// CostoptimizationhubClient ...
type CostoptimizationhubClient interface {
	Options() costoptimizationhub.Options
	GetPreferences(ctx context.Context, params *costoptimizationhub.GetPreferencesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetPreferencesOutput, error)
	GetRecommendation(ctx context.Context, params *costoptimizationhub.GetRecommendationInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.GetRecommendationOutput, error)
	ListEnrollmentStatuses(ctx context.Context, params *costoptimizationhub.ListEnrollmentStatusesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListEnrollmentStatusesOutput, error)
	ListRecommendationSummaries(ctx context.Context, params *costoptimizationhub.ListRecommendationSummariesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationSummariesOutput, error)
	ListRecommendations(ctx context.Context, params *costoptimizationhub.ListRecommendationsInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.ListRecommendationsOutput, error)
	UpdateEnrollmentStatus(ctx context.Context, params *costoptimizationhub.UpdateEnrollmentStatusInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdateEnrollmentStatusOutput, error)
	UpdatePreferences(ctx context.Context, params *costoptimizationhub.UpdatePreferencesInput, optFns ...func(*costoptimizationhub.Options)) (*costoptimizationhub.UpdatePreferencesOutput, error)
}
