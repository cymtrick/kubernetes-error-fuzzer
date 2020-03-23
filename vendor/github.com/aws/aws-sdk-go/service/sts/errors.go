// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

const (

	// ErrCodeExpiredTokenException for service response error code
	// "ExpiredTokenException".
	//
	// The web identity token that was passed is expired or is not valid. Get a
	// new identity token from the identity provider and then retry the request.
	ErrCodeExpiredTokenException = "ExpiredTokenException"

	// ErrCodeIDPCommunicationErrorException for service response error code
	// "IDPCommunicationError".
	//
	// The request could not be fulfilled because the identity provider (IDP) that
	// was asked to verify the incoming identity token could not be reached. This
	// is often a transient error caused by network conditions. Retry the request
	// a limited number of times so that you don't exceed the request rate. If the
	// error persists, the identity provider might be down or not responding.
	ErrCodeIDPCommunicationErrorException = "IDPCommunicationError"

	// ErrCodeIDPRejectedClaimException for service response error code
	// "IDPRejectedClaim".
	//
	// The identity provider (IdP) reported that authentication failed. This might
	// be because the claim is invalid.
	//
	// If this error is returned for the AssumeRoleWithWebIdentity operation, it
	// can also mean that the claim has expired or has been explicitly revoked.
	ErrCodeIDPRejectedClaimException = "IDPRejectedClaim"

	// ErrCodeInvalidAuthorizationMessageException for service response error code
	// "InvalidAuthorizationMessageException".
	//
	// The error returned if the message passed to DecodeAuthorizationMessage was
	// invalid. This can happen if the token contains invalid characters, such as
	// linebreaks.
	ErrCodeInvalidAuthorizationMessageException = "InvalidAuthorizationMessageException"

	// ErrCodeInvalidIdentityTokenException for service response error code
	// "InvalidIdentityToken".
	//
	// The web identity token that was passed could not be validated by AWS. Get
	// a new identity token from the identity provider and then retry the request.
	ErrCodeInvalidIdentityTokenException = "InvalidIdentityToken"

	// ErrCodeMalformedPolicyDocumentException for service response error code
	// "MalformedPolicyDocument".
	//
	// The request was rejected because the policy document was malformed. The error
	// message describes the specific error.
	ErrCodeMalformedPolicyDocumentException = "MalformedPolicyDocument"

	// ErrCodePackedPolicyTooLargeException for service response error code
	// "PackedPolicyTooLarge".
	//
	// The request was rejected because the total packed size of the session policies
	// and session tags combined was too large. An AWS conversion compresses the
	// session policy document, session policy ARNs, and session tags into a packed
	// binary format that has a separate limit. The error message indicates by percentage
	// how close the policies and tags are to the upper size limit. For more information,
	// see Passing Session Tags in STS (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html)
	// in the IAM User Guide.
	//
	// You could receive this error even though you meet other defined session policy
	// and session tag limits. For more information, see IAM and STS Entity Character
	// Limits (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
	// in the IAM User Guide.
	ErrCodePackedPolicyTooLargeException = "PackedPolicyTooLarge"

	// ErrCodeRegionDisabledException for service response error code
	// "RegionDisabledException".
	//
	// STS is not activated in the requested region for the account that is being
	// asked to generate credentials. The account administrator must use the IAM
	// console to activate STS in that region. For more information, see Activating
	// and Deactivating AWS STS in an AWS Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
	// in the IAM User Guide.
	ErrCodeRegionDisabledException = "RegionDisabledException"
)
