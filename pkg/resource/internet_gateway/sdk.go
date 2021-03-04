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

package internet_gateway

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
	_ = &svcapitypes.InternetGateway{}
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

	resp, respErr := rm.sdkapi.DescribeInternetGatewaysWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeInternetGateways", respErr)
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
	for _, elem := range resp.InternetGateways {
		if elem.Attachments != nil {
			f0 := []*svcapitypes.InternetGatewayAttachment{}
			for _, f0iter := range elem.Attachments {
				f0elem := &svcapitypes.InternetGatewayAttachment{}
				if f0iter.State != nil {
					f0elem.State = f0iter.State
				}
				if f0iter.VpcId != nil {
					f0elem.VPCID = f0iter.VpcId
				}
				f0 = append(f0, f0elem)
			}
			ko.Status.Attachments = f0
		}
		if elem.InternetGatewayId != nil {
			ko.Status.InternetGatewayID = elem.InternetGatewayId
		}
		if elem.OwnerId != nil {
			ko.Status.OwnerID = elem.OwnerId
		}
		if elem.Tags != nil {
			f3 := []*svcapitypes.Tag{}
			for _, f3iter := range elem.Tags {
				f3elem := &svcapitypes.Tag{}
				if f3iter.Key != nil {
					f3elem.Key = f3iter.Key
				}
				if f3iter.Value != nil {
					f3elem.Value = f3iter.Value
				}
				f3 = append(f3, f3elem)
			}
			ko.Status.Tags = f3
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
) (*svcsdk.DescribeInternetGatewaysInput, error) {
	res := &svcsdk.DescribeInternetGatewaysInput{}

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

	resp, respErr := rm.sdkapi.CreateInternetGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateInternetGateway", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.InternetGateway.Attachments != nil {
		f0 := []*svcapitypes.InternetGatewayAttachment{}
		for _, f0iter := range resp.InternetGateway.Attachments {
			f0elem := &svcapitypes.InternetGatewayAttachment{}
			if f0iter.State != nil {
				f0elem.State = f0iter.State
			}
			if f0iter.VpcId != nil {
				f0elem.VPCID = f0iter.VpcId
			}
			f0 = append(f0, f0elem)
		}
		ko.Status.Attachments = f0
	}
	if resp.InternetGateway.InternetGatewayId != nil {
		ko.Status.InternetGatewayID = resp.InternetGateway.InternetGatewayId
	}
	if resp.InternetGateway.OwnerId != nil {
		ko.Status.OwnerID = resp.InternetGateway.OwnerId
	}
	if resp.InternetGateway.Tags != nil {
		f3 := []*svcapitypes.Tag{}
		for _, f3iter := range resp.InternetGateway.Tags {
			f3elem := &svcapitypes.Tag{}
			if f3iter.Key != nil {
				f3elem.Key = f3iter.Key
			}
			if f3iter.Value != nil {
				f3elem.Value = f3iter.Value
			}
			f3 = append(f3, f3elem)
		}
		ko.Status.Tags = f3
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateInternetGatewayInput, error) {
	res := &svcsdk.CreateInternetGatewayInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f1 := []*svcsdk.TagSpecification{}
		for _, f1iter := range r.ko.Spec.TagSpecifications {
			f1elem := &svcsdk.TagSpecification{}
			if f1iter.ResourceType != nil {
				f1elem.SetResourceType(*f1iter.ResourceType)
			}
			if f1iter.Tags != nil {
				f1elemf1 := []*svcsdk.Tag{}
				for _, f1elemf1iter := range f1iter.Tags {
					f1elemf1elem := &svcsdk.Tag{}
					if f1elemf1iter.Key != nil {
						f1elemf1elem.SetKey(*f1elemf1iter.Key)
					}
					if f1elemf1iter.Value != nil {
						f1elemf1elem.SetValue(*f1elemf1iter.Value)
					}
					f1elemf1 = append(f1elemf1, f1elemf1elem)
				}
				f1elem.SetTags(f1elemf1)
			}
			f1 = append(f1, f1elem)
		}
		res.SetTagSpecifications(f1)
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
	_, respErr := rm.sdkapi.DeleteInternetGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteInternetGateway", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteInternetGatewayInput, error) {
	res := &svcsdk.DeleteInternetGatewayInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Status.InternetGatewayID != nil {
		res.SetInternetGatewayId(*r.ko.Status.InternetGatewayID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.InternetGateway,
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