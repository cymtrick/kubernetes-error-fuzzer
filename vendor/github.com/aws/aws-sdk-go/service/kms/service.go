// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Key Management Service (AWS KMS) is an encryption and key management
// web service. This guide describes the AWS KMS operations that you can call
// programmatically. For general information about AWS KMS, see the AWS Key
// Management Service Developer Guide (http://docs.aws.amazon.com/kms/latest/developerguide/).
//
// AWS provides SDKs that consist of libraries and sample code for various programming
// languages and platforms (Java, Ruby, .Net, iOS, Android, etc.). The SDKs
// provide a convenient way to create programmatic access to AWS KMS and other
// AWS services. For example, the SDKs take care of tasks such as signing requests
// (see below), managing errors, and retrying requests automatically. For more
// information about the AWS SDKs, including how to download and install them,
// see Tools for Amazon Web Services (http://aws.amazon.com/tools/).
//
// We recommend that you use the AWS SDKs to make programmatic API calls to
// AWS KMS.
//
// Clients must support TLS (Transport Layer Security) 1.0. We recommend TLS
// 1.2. Clients must also support cipher suites with Perfect Forward Secrecy
// (PFS) such as Ephemeral Diffie-Hellman (DHE) or Elliptic Curve Ephemeral
// Diffie-Hellman (ECDHE). Most modern systems such as Java 7 and later support
// these modes.
//
// Signing Requests
//
// Requests must be signed by using an access key ID and a secret access key.
// We strongly recommend that you do not use your AWS account (root) access
// key ID and secret key for everyday work with AWS KMS. Instead, use the access
// key ID and secret access key for an IAM user, or you can use the AWS Security
// Token Service to generate temporary security credentials that you can use
// to sign requests.
//
// All AWS KMS operations require Signature Version 4 (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// Logging API Requests
//
// AWS KMS supports AWS CloudTrail, a service that logs AWS API calls and related
// events for your AWS account and delivers them to an Amazon S3 bucket that
// you specify. By using the information collected by CloudTrail, you can determine
// what requests were made to AWS KMS, who made the request, when it was made,
// and so on. To learn more about CloudTrail, including how to turn it on and
// find your log files, see the AWS CloudTrail User Guide (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/).
//
// Additional Resources
//
// For more information about credentials and request signing, see the following:
//
//    * AWS Security Credentials (http://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
//    - This topic provides general information about the types of credentials
//    used for accessing AWS.
//
//    * Temporary Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html)
//    - This section of the IAM User Guide describes how to create and use temporary
//    security credentials.
//
//    * Signature Version 4 Signing Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
//    - This set of topics walks you through the process of signing a request
//    using an access key ID and a secret access key.
//
// Commonly Used APIs
//
// Of the APIs discussed in this guide, the following will prove the most useful
// for most applications. You will likely perform actions other than these,
// such as creating keys and assigning policies, by using the console.
//
//    * Encrypt
//
//    * Decrypt
//
//    * GenerateDataKey
//
//    * GenerateDataKeyWithoutPlaintext
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01
type KMS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "kms"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the KMS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a KMS client from just a session.
//     svc := kms.New(mySession)
//
//     // Create a KMS client with additional configuration
//     svc := kms.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *KMS {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *KMS {
	svc := &KMS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-11-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "TrentService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a KMS operation and runs any
// custom request initialization.
func (c *KMS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
