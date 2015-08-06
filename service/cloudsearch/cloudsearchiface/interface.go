package cloudsearchiface

import (
	"github.com/datacratic/aws-sdk-go/service/cloudsearch"
)

type CloudSearchAPI interface {
	BuildSuggesters(*cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)

	CreateDomain(*cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)

	DefineAnalysisScheme(*cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)

	DefineExpression(*cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)

	DefineIndexField(*cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)

	DefineSuggester(*cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)

	DeleteAnalysisScheme(*cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)

	DeleteDomain(*cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)

	DeleteExpression(*cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)

	DeleteIndexField(*cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)

	DeleteSuggester(*cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)

	DescribeAnalysisSchemes(*cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)

	DescribeAvailabilityOptions(*cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)

	DescribeDomains(*cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)

	DescribeExpressions(*cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)

	DescribeIndexFields(*cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)

	DescribeScalingParameters(*cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)

	DescribeServiceAccessPolicies(*cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)

	DescribeSuggesters(*cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)

	IndexDocuments(*cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)

	ListDomainNames(*cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)

	UpdateAvailabilityOptions(*cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)

	UpdateScalingParameters(*cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)

	UpdateServiceAccessPolicies(*cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
}