// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
)

// CloudsearchClient ...
type CloudsearchClient interface {
	Options() cloudsearch.Options
	BuildSuggesters(ctx context.Context, params *cloudsearch.BuildSuggestersInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.BuildSuggestersOutput, error)
	CreateDomain(ctx context.Context, params *cloudsearch.CreateDomainInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.CreateDomainOutput, error)
	DefineAnalysisScheme(ctx context.Context, params *cloudsearch.DefineAnalysisSchemeInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DefineAnalysisSchemeOutput, error)
	DefineExpression(ctx context.Context, params *cloudsearch.DefineExpressionInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DefineExpressionOutput, error)
	DefineIndexField(ctx context.Context, params *cloudsearch.DefineIndexFieldInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DefineIndexFieldOutput, error)
	DefineSuggester(ctx context.Context, params *cloudsearch.DefineSuggesterInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DefineSuggesterOutput, error)
	DeleteAnalysisScheme(ctx context.Context, params *cloudsearch.DeleteAnalysisSchemeInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DeleteAnalysisSchemeOutput, error)
	DeleteDomain(ctx context.Context, params *cloudsearch.DeleteDomainInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DeleteDomainOutput, error)
	DeleteExpression(ctx context.Context, params *cloudsearch.DeleteExpressionInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DeleteExpressionOutput, error)
	DeleteIndexField(ctx context.Context, params *cloudsearch.DeleteIndexFieldInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DeleteIndexFieldOutput, error)
	DeleteSuggester(ctx context.Context, params *cloudsearch.DeleteSuggesterInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DeleteSuggesterOutput, error)
	DescribeAnalysisSchemes(ctx context.Context, params *cloudsearch.DescribeAnalysisSchemesInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeAnalysisSchemesOutput, error)
	DescribeAvailabilityOptions(ctx context.Context, params *cloudsearch.DescribeAvailabilityOptionsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)
	DescribeDomainEndpointOptions(ctx context.Context, params *cloudsearch.DescribeDomainEndpointOptionsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error)
	DescribeDomains(ctx context.Context, params *cloudsearch.DescribeDomainsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeDomainsOutput, error)
	DescribeExpressions(ctx context.Context, params *cloudsearch.DescribeExpressionsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeExpressionsOutput, error)
	DescribeIndexFields(ctx context.Context, params *cloudsearch.DescribeIndexFieldsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeIndexFieldsOutput, error)
	DescribeScalingParameters(ctx context.Context, params *cloudsearch.DescribeScalingParametersInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeScalingParametersOutput, error)
	DescribeServiceAccessPolicies(ctx context.Context, params *cloudsearch.DescribeServiceAccessPoliciesInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)
	DescribeSuggesters(ctx context.Context, params *cloudsearch.DescribeSuggestersInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.DescribeSuggestersOutput, error)
	IndexDocuments(ctx context.Context, params *cloudsearch.IndexDocumentsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.IndexDocumentsOutput, error)
	ListDomainNames(ctx context.Context, params *cloudsearch.ListDomainNamesInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.ListDomainNamesOutput, error)
	UpdateAvailabilityOptions(ctx context.Context, params *cloudsearch.UpdateAvailabilityOptionsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)
	UpdateDomainEndpointOptions(ctx context.Context, params *cloudsearch.UpdateDomainEndpointOptionsInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error)
	UpdateScalingParameters(ctx context.Context, params *cloudsearch.UpdateScalingParametersInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.UpdateScalingParametersOutput, error)
	UpdateServiceAccessPolicies(ctx context.Context, params *cloudsearch.UpdateServiceAccessPoliciesInput, optFns ...func(*cloudsearch.Options)) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
}
