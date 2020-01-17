// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

const (

	// ErrCodeAlreadyExistsFault for service response error code
	// "AlreadyExists".
	//
	// You already have an Auto Scaling group or launch configuration with this
	// name.
	ErrCodeAlreadyExistsFault = "AlreadyExists"

	// ErrCodeInvalidNextToken for service response error code
	// "InvalidNextToken".
	//
	// The NextToken value is not valid.
	ErrCodeInvalidNextToken = "InvalidNextToken"

	// ErrCodeLimitExceededFault for service response error code
	// "LimitExceeded".
	//
	// You have already reached a limit for your Amazon EC2 Auto Scaling resources
	// (for example, Auto Scaling groups, launch configurations, or lifecycle hooks).
	// For more information, see DescribeAccountLimits.
	ErrCodeLimitExceededFault = "LimitExceeded"

	// ErrCodeResourceContentionFault for service response error code
	// "ResourceContention".
	//
	// You already have a pending update to an Amazon EC2 Auto Scaling resource
	// (for example, an Auto Scaling group, instance, or load balancer).
	ErrCodeResourceContentionFault = "ResourceContention"

	// ErrCodeResourceInUseFault for service response error code
	// "ResourceInUse".
	//
	// The operation can't be performed because the resource is in use.
	ErrCodeResourceInUseFault = "ResourceInUse"

	// ErrCodeScalingActivityInProgressFault for service response error code
	// "ScalingActivityInProgress".
	//
	// The operation can't be performed because there are scaling activities in
	// progress.
	ErrCodeScalingActivityInProgressFault = "ScalingActivityInProgress"

	// ErrCodeServiceLinkedRoleFailure for service response error code
	// "ServiceLinkedRoleFailure".
	//
	// The service-linked role is not yet ready for use.
	ErrCodeServiceLinkedRoleFailure = "ServiceLinkedRoleFailure"
)
