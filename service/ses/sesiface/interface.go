package sesiface

import (
	"github.com/datacratic/aws-sdk-go/service/ses"
)

type SESAPI interface {
	DeleteIdentity(*ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)

	DeleteVerifiedEmailAddress(*ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)

	GetIdentityDKIMAttributes(*ses.GetIdentityDKIMAttributesInput) (*ses.GetIdentityDKIMAttributesOutput, error)

	GetIdentityNotificationAttributes(*ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)

	GetIdentityVerificationAttributes(*ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)

	GetSendQuota(*ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)

	GetSendStatistics(*ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)

	ListIdentities(*ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)

	ListVerifiedEmailAddresses(*ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)

	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)

	SendRawEmail(*ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)

	SetIdentityDKIMEnabled(*ses.SetIdentityDKIMEnabledInput) (*ses.SetIdentityDKIMEnabledOutput, error)

	SetIdentityFeedbackForwardingEnabled(*ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)

	SetIdentityNotificationTopic(*ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)

	VerifyDomainDKIM(*ses.VerifyDomainDKIMInput) (*ses.VerifyDomainDKIMOutput, error)

	VerifyDomainIdentity(*ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)

	VerifyEmailAddress(*ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)

	VerifyEmailIdentity(*ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
}