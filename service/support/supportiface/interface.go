package supportiface

import (
	"github.com/datacratic/aws-sdk-go/service/support"
)

type SupportAPI interface {
	AddAttachmentsToSet(*support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)

	AddCommunicationToCase(*support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)

	CreateCase(*support.CreateCaseInput) (*support.CreateCaseOutput, error)

	DescribeAttachment(*support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)

	DescribeCases(*support.DescribeCasesInput) (*support.DescribeCasesOutput, error)

	DescribeCommunications(*support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)

	DescribeServices(*support.DescribeServicesInput) (*support.DescribeServicesOutput, error)

	DescribeSeverityLevels(*support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)

	DescribeTrustedAdvisorCheckRefreshStatuses(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)

	DescribeTrustedAdvisorCheckResult(*support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)

	DescribeTrustedAdvisorCheckSummaries(*support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)

	DescribeTrustedAdvisorChecks(*support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)

	RefreshTrustedAdvisorCheck(*support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)

	ResolveCase(*support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
}