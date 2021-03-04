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

package transit_gateway_peering_attachment

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
	_ = &svcapitypes.TransitGatewayPeeringAttachment{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeTransitGatewayPeeringAttachmentsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeTransitGatewayPeeringAttachments", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.TransitGatewayPeeringAttachments {
		if elem.AccepterTgwInfo != nil {
			f0 := &svcapitypes.PeeringTgwInfo{}
			if elem.AccepterTgwInfo.OwnerId != nil {
				f0.OwnerID = elem.AccepterTgwInfo.OwnerId
			}
			if elem.AccepterTgwInfo.Region != nil {
				f0.Region = elem.AccepterTgwInfo.Region
			}
			if elem.AccepterTgwInfo.TransitGatewayId != nil {
				f0.TransitGatewayID = elem.AccepterTgwInfo.TransitGatewayId
			}
			ko.Status.AccepterTgwInfo = f0
		}
		if elem.CreationTime != nil {
			ko.Status.CreationTime = &metav1.Time{*elem.CreationTime}
		}
		if elem.RequesterTgwInfo != nil {
			f2 := &svcapitypes.PeeringTgwInfo{}
			if elem.RequesterTgwInfo.OwnerId != nil {
				f2.OwnerID = elem.RequesterTgwInfo.OwnerId
			}
			if elem.RequesterTgwInfo.Region != nil {
				f2.Region = elem.RequesterTgwInfo.Region
			}
			if elem.RequesterTgwInfo.TransitGatewayId != nil {
				f2.TransitGatewayID = elem.RequesterTgwInfo.TransitGatewayId
			}
			ko.Status.RequesterTgwInfo = f2
		}
		if elem.State != nil {
			ko.Status.State = elem.State
		}
		if elem.Status != nil {
			f4 := &svcapitypes.PeeringAttachmentStatus{}
			if elem.Status.Code != nil {
				f4.Code = elem.Status.Code
			}
			if elem.Status.Message != nil {
				f4.Message = elem.Status.Message
			}
			ko.Status.Status = f4
		}
		if elem.Tags != nil {
			f5 := []*svcapitypes.Tag{}
			for _, f5iter := range elem.Tags {
				f5elem := &svcapitypes.Tag{}
				if f5iter.Key != nil {
					f5elem.Key = f5iter.Key
				}
				if f5iter.Value != nil {
					f5elem.Value = f5iter.Value
				}
				f5 = append(f5, f5elem)
			}
			ko.Status.Tags = f5
		}
		if elem.TransitGatewayAttachmentId != nil {
			ko.Status.TransitGatewayAttachmentID = elem.TransitGatewayAttachmentId
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeTransitGatewayPeeringAttachmentsInput, error) {
	res := &svcsdk.DescribeTransitGatewayPeeringAttachmentsInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}

	return res, nil
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

	resp, respErr := rm.sdkapi.CreateTransitGatewayPeeringAttachmentWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTransitGatewayPeeringAttachment", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.TransitGatewayPeeringAttachment.AccepterTgwInfo != nil {
		f0 := &svcapitypes.PeeringTgwInfo{}
		if resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.OwnerId != nil {
			f0.OwnerID = resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.OwnerId
		}
		if resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.Region != nil {
			f0.Region = resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.Region
		}
		if resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.TransitGatewayId != nil {
			f0.TransitGatewayID = resp.TransitGatewayPeeringAttachment.AccepterTgwInfo.TransitGatewayId
		}
		ko.Status.AccepterTgwInfo = f0
	}
	if resp.TransitGatewayPeeringAttachment.CreationTime != nil {
		ko.Status.CreationTime = &metav1.Time{*resp.TransitGatewayPeeringAttachment.CreationTime}
	}
	if resp.TransitGatewayPeeringAttachment.RequesterTgwInfo != nil {
		f2 := &svcapitypes.PeeringTgwInfo{}
		if resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.OwnerId != nil {
			f2.OwnerID = resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.OwnerId
		}
		if resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.Region != nil {
			f2.Region = resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.Region
		}
		if resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.TransitGatewayId != nil {
			f2.TransitGatewayID = resp.TransitGatewayPeeringAttachment.RequesterTgwInfo.TransitGatewayId
		}
		ko.Status.RequesterTgwInfo = f2
	}
	if resp.TransitGatewayPeeringAttachment.State != nil {
		ko.Status.State = resp.TransitGatewayPeeringAttachment.State
	}
	if resp.TransitGatewayPeeringAttachment.Status != nil {
		f4 := &svcapitypes.PeeringAttachmentStatus{}
		if resp.TransitGatewayPeeringAttachment.Status.Code != nil {
			f4.Code = resp.TransitGatewayPeeringAttachment.Status.Code
		}
		if resp.TransitGatewayPeeringAttachment.Status.Message != nil {
			f4.Message = resp.TransitGatewayPeeringAttachment.Status.Message
		}
		ko.Status.Status = f4
	}
	if resp.TransitGatewayPeeringAttachment.Tags != nil {
		f5 := []*svcapitypes.Tag{}
		for _, f5iter := range resp.TransitGatewayPeeringAttachment.Tags {
			f5elem := &svcapitypes.Tag{}
			if f5iter.Key != nil {
				f5elem.Key = f5iter.Key
			}
			if f5iter.Value != nil {
				f5elem.Value = f5iter.Value
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.Tags = f5
	}
	if resp.TransitGatewayPeeringAttachment.TransitGatewayAttachmentId != nil {
		ko.Status.TransitGatewayAttachmentID = resp.TransitGatewayPeeringAttachment.TransitGatewayAttachmentId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTransitGatewayPeeringAttachmentInput, error) {
	res := &svcsdk.CreateTransitGatewayPeeringAttachmentInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.PeerAccountID != nil {
		res.SetPeerAccountId(*r.ko.Spec.PeerAccountID)
	}
	if r.ko.Spec.PeerRegion != nil {
		res.SetPeerRegion(*r.ko.Spec.PeerRegion)
	}
	if r.ko.Spec.PeerTransitGatewayID != nil {
		res.SetPeerTransitGatewayId(*r.ko.Spec.PeerTransitGatewayID)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f4 := []*svcsdk.TagSpecification{}
		for _, f4iter := range r.ko.Spec.TagSpecifications {
			f4elem := &svcsdk.TagSpecification{}
			if f4iter.ResourceType != nil {
				f4elem.SetResourceType(*f4iter.ResourceType)
			}
			if f4iter.Tags != nil {
				f4elemf1 := []*svcsdk.Tag{}
				for _, f4elemf1iter := range f4iter.Tags {
					f4elemf1elem := &svcsdk.Tag{}
					if f4elemf1iter.Key != nil {
						f4elemf1elem.SetKey(*f4elemf1iter.Key)
					}
					if f4elemf1iter.Value != nil {
						f4elemf1elem.SetValue(*f4elemf1iter.Value)
					}
					f4elemf1 = append(f4elemf1, f4elemf1elem)
				}
				f4elem.SetTags(f4elemf1)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTagSpecifications(f4)
	}
	if r.ko.Spec.TransitGatewayID != nil {
		res.SetTransitGatewayId(*r.ko.Spec.TransitGatewayID)
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
	_, respErr := rm.sdkapi.DeleteTransitGatewayPeeringAttachmentWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteTransitGatewayPeeringAttachment", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteTransitGatewayPeeringAttachmentInput, error) {
	res := &svcsdk.DeleteTransitGatewayPeeringAttachmentInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Status.TransitGatewayAttachmentID != nil {
		res.SetTransitGatewayAttachmentId(*r.ko.Status.TransitGatewayAttachmentID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TransitGatewayPeeringAttachment,
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
