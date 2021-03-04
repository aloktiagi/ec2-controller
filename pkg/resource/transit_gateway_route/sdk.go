// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package transit_gateway_route

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.EC2{}
	_ = &svcapitypes.TransitGatewayRoute{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// Believe it or not, there are API resources that can be created but there
	// is no read operation. Point in case: RDS' CreateDBInstanceReadReplica
	// has no corresponding read operation that I know of...
	return nil, ackerr.NotImplemented
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateTransitGatewayRouteWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTransitGatewayRoute", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Route.PrefixListId != nil {
		ko.Status.PrefixListID = resp.Route.PrefixListId
	}
	if resp.Route.State != nil {
		ko.Status.State = resp.Route.State
	}
	if resp.Route.TransitGatewayAttachments != nil {
		f3 := []*svcapitypes.TransitGatewayRouteAttachment{}
		for _, f3iter := range resp.Route.TransitGatewayAttachments {
			f3elem := &svcapitypes.TransitGatewayRouteAttachment{}
			if f3iter.ResourceId != nil {
				f3elem.ResourceID = f3iter.ResourceId
			}
			if f3iter.ResourceType != nil {
				f3elem.ResourceType = f3iter.ResourceType
			}
			if f3iter.TransitGatewayAttachmentId != nil {
				f3elem.TransitGatewayAttachmentID = f3iter.TransitGatewayAttachmentId
			}
			f3 = append(f3, f3elem)
		}
		ko.Status.TransitGatewayAttachments = f3
	}
	if resp.Route.Type != nil {
		ko.Status.Type = resp.Route.Type
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTransitGatewayRouteInput, error) {
	res := &svcsdk.CreateTransitGatewayRouteInput{}

	if r.ko.Spec.Blackhole != nil {
		res.SetBlackhole(*r.ko.Spec.Blackhole)
	}
	if r.ko.Spec.DestinationCIDRBlock != nil {
		res.SetDestinationCidrBlock(*r.ko.Spec.DestinationCIDRBlock)
	}
	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.TransitGatewayAttachmentID != nil {
		res.SetTransitGatewayAttachmentId(*r.ko.Spec.TransitGatewayAttachmentID)
	}
	if r.ko.Spec.TransitGatewayRouteTableID != nil {
		res.SetTransitGatewayRouteTableId(*r.ko.Spec.TransitGatewayRouteTableID)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {

	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteTransitGatewayRouteWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteTransitGatewayRoute", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteTransitGatewayRouteInput, error) {
	res := &svcsdk.DeleteTransitGatewayRouteInput{}

	if r.ko.Spec.DestinationCIDRBlock != nil {
		res.SetDestinationCidrBlock(*r.ko.Spec.DestinationCIDRBlock)
	}
	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.TransitGatewayRouteTableID != nil {
		res.SetTransitGatewayRouteTableId(*r.ko.Spec.TransitGatewayRouteTableID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TransitGatewayRoute,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
