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

package transit_gateway_vpc_attachment

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
	_ = &svcapitypes.TransitGatewayVPCAttachment{}
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

	resp, respErr := rm.sdkapi.DescribeTransitGatewayVpcAttachmentsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeTransitGatewayVpcAttachments", respErr)
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
	for _, elem := range resp.TransitGatewayVpcAttachments {
		if elem.CreationTime != nil {
			ko.Status.CreationTime = &metav1.Time{*elem.CreationTime}
		}
		if elem.Options != nil {
			f1 := &svcapitypes.CreateTransitGatewayVPCAttachmentRequestOptions{}
			if elem.Options.ApplianceModeSupport != nil {
				f1.ApplianceModeSupport = elem.Options.ApplianceModeSupport
			}
			if elem.Options.DnsSupport != nil {
				f1.DNSSupport = elem.Options.DnsSupport
			}
			if elem.Options.Ipv6Support != nil {
				f1.IPv6Support = elem.Options.Ipv6Support
			}
			ko.Spec.Options = f1
		}
		if elem.State != nil {
			ko.Status.State = elem.State
		}
		if elem.SubnetIds != nil {
			f3 := []*string{}
			for _, f3iter := range elem.SubnetIds {
				var f3elem string
				f3elem = *f3iter
				f3 = append(f3, &f3elem)
			}
			ko.Spec.SubnetIDs = f3
		}
		if elem.Tags != nil {
			f4 := []*svcapitypes.Tag{}
			for _, f4iter := range elem.Tags {
				f4elem := &svcapitypes.Tag{}
				if f4iter.Key != nil {
					f4elem.Key = f4iter.Key
				}
				if f4iter.Value != nil {
					f4elem.Value = f4iter.Value
				}
				f4 = append(f4, f4elem)
			}
			ko.Status.Tags = f4
		}
		if elem.TransitGatewayAttachmentId != nil {
			ko.Status.TransitGatewayAttachmentID = elem.TransitGatewayAttachmentId
		}
		if elem.TransitGatewayId != nil {
			ko.Spec.TransitGatewayID = elem.TransitGatewayId
		}
		if elem.VpcId != nil {
			ko.Spec.VPCID = elem.VpcId
		}
		if elem.VpcOwnerId != nil {
			ko.Status.VPCOwnerID = elem.VpcOwnerId
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
) (*svcsdk.DescribeTransitGatewayVpcAttachmentsInput, error) {
	res := &svcsdk.DescribeTransitGatewayVpcAttachmentsInput{}

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

	resp, respErr := rm.sdkapi.CreateTransitGatewayVpcAttachmentWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTransitGatewayVpcAttachment", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.TransitGatewayVpcAttachment.CreationTime != nil {
		ko.Status.CreationTime = &metav1.Time{*resp.TransitGatewayVpcAttachment.CreationTime}
	}
	if resp.TransitGatewayVpcAttachment.State != nil {
		ko.Status.State = resp.TransitGatewayVpcAttachment.State
	}
	if resp.TransitGatewayVpcAttachment.Tags != nil {
		f4 := []*svcapitypes.Tag{}
		for _, f4iter := range resp.TransitGatewayVpcAttachment.Tags {
			f4elem := &svcapitypes.Tag{}
			if f4iter.Key != nil {
				f4elem.Key = f4iter.Key
			}
			if f4iter.Value != nil {
				f4elem.Value = f4iter.Value
			}
			f4 = append(f4, f4elem)
		}
		ko.Status.Tags = f4
	}
	if resp.TransitGatewayVpcAttachment.TransitGatewayAttachmentId != nil {
		ko.Status.TransitGatewayAttachmentID = resp.TransitGatewayVpcAttachment.TransitGatewayAttachmentId
	}
	if resp.TransitGatewayVpcAttachment.VpcOwnerId != nil {
		ko.Status.VPCOwnerID = resp.TransitGatewayVpcAttachment.VpcOwnerId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTransitGatewayVpcAttachmentInput, error) {
	res := &svcsdk.CreateTransitGatewayVpcAttachmentInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.Options != nil {
		f1 := &svcsdk.CreateTransitGatewayVpcAttachmentRequestOptions{}
		if r.ko.Spec.Options.ApplianceModeSupport != nil {
			f1.SetApplianceModeSupport(*r.ko.Spec.Options.ApplianceModeSupport)
		}
		if r.ko.Spec.Options.DNSSupport != nil {
			f1.SetDnsSupport(*r.ko.Spec.Options.DNSSupport)
		}
		if r.ko.Spec.Options.IPv6Support != nil {
			f1.SetIpv6Support(*r.ko.Spec.Options.IPv6Support)
		}
		res.SetOptions(f1)
	}
	if r.ko.Spec.SubnetIDs != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.SubnetIDs {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetSubnetIds(f2)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f3 := []*svcsdk.TagSpecification{}
		for _, f3iter := range r.ko.Spec.TagSpecifications {
			f3elem := &svcsdk.TagSpecification{}
			if f3iter.ResourceType != nil {
				f3elem.SetResourceType(*f3iter.ResourceType)
			}
			if f3iter.Tags != nil {
				f3elemf1 := []*svcsdk.Tag{}
				for _, f3elemf1iter := range f3iter.Tags {
					f3elemf1elem := &svcsdk.Tag{}
					if f3elemf1iter.Key != nil {
						f3elemf1elem.SetKey(*f3elemf1iter.Key)
					}
					if f3elemf1iter.Value != nil {
						f3elemf1elem.SetValue(*f3elemf1iter.Value)
					}
					f3elemf1 = append(f3elemf1, f3elemf1elem)
				}
				f3elem.SetTags(f3elemf1)
			}
			f3 = append(f3, f3elem)
		}
		res.SetTagSpecifications(f3)
	}
	if r.ko.Spec.TransitGatewayID != nil {
		res.SetTransitGatewayId(*r.ko.Spec.TransitGatewayID)
	}
	if r.ko.Spec.VPCID != nil {
		res.SetVpcId(*r.ko.Spec.VPCID)
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

	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.ModifyTransitGatewayVpcAttachmentWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyTransitGatewayVpcAttachment", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.TransitGatewayVpcAttachment.CreationTime != nil {
		ko.Status.CreationTime = &metav1.Time{*resp.TransitGatewayVpcAttachment.CreationTime}
	}
	if resp.TransitGatewayVpcAttachment.State != nil {
		ko.Status.State = resp.TransitGatewayVpcAttachment.State
	}
	if resp.TransitGatewayVpcAttachment.Tags != nil {
		f4 := []*svcapitypes.Tag{}
		for _, f4iter := range resp.TransitGatewayVpcAttachment.Tags {
			f4elem := &svcapitypes.Tag{}
			if f4iter.Key != nil {
				f4elem.Key = f4iter.Key
			}
			if f4iter.Value != nil {
				f4elem.Value = f4iter.Value
			}
			f4 = append(f4, f4elem)
		}
		ko.Status.Tags = f4
	}
	if resp.TransitGatewayVpcAttachment.TransitGatewayAttachmentId != nil {
		ko.Status.TransitGatewayAttachmentID = resp.TransitGatewayVpcAttachment.TransitGatewayAttachmentId
	}
	if resp.TransitGatewayVpcAttachment.VpcOwnerId != nil {
		ko.Status.VPCOwnerID = resp.TransitGatewayVpcAttachment.VpcOwnerId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ModifyTransitGatewayVpcAttachmentInput, error) {
	res := &svcsdk.ModifyTransitGatewayVpcAttachmentInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.Options != nil {
		f2 := &svcsdk.ModifyTransitGatewayVpcAttachmentRequestOptions{}
		if r.ko.Spec.Options.ApplianceModeSupport != nil {
			f2.SetApplianceModeSupport(*r.ko.Spec.Options.ApplianceModeSupport)
		}
		if r.ko.Spec.Options.DNSSupport != nil {
			f2.SetDnsSupport(*r.ko.Spec.Options.DNSSupport)
		}
		if r.ko.Spec.Options.IPv6Support != nil {
			f2.SetIpv6Support(*r.ko.Spec.Options.IPv6Support)
		}
		res.SetOptions(f2)
	}
	if r.ko.Status.TransitGatewayAttachmentID != nil {
		res.SetTransitGatewayAttachmentId(*r.ko.Status.TransitGatewayAttachmentID)
	}

	return res, nil
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
	_, respErr := rm.sdkapi.DeleteTransitGatewayVpcAttachmentWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteTransitGatewayVpcAttachment", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteTransitGatewayVpcAttachmentInput, error) {
	res := &svcsdk.DeleteTransitGatewayVpcAttachmentInput{}

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
	ko *svcapitypes.TransitGatewayVPCAttachment,
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
