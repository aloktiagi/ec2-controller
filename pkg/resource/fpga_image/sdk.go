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

package fpga_image

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
	_ = &svcapitypes.FPGAImage{}
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

	resp, respErr := rm.sdkapi.DescribeFpgaImagesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeFpgaImages", respErr)
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
	for _, elem := range resp.FpgaImages {
		if elem.Description != nil {
			ko.Spec.Description = elem.Description
		}
		if elem.FpgaImageGlobalId != nil {
			ko.Status.FPGAImageGlobalID = elem.FpgaImageGlobalId
		}
		if elem.FpgaImageId != nil {
			ko.Status.FPGAImageID = elem.FpgaImageId
		}
		if elem.Name != nil {
			ko.Spec.Name = elem.Name
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
) (*svcsdk.DescribeFpgaImagesInput, error) {
	res := &svcsdk.DescribeFpgaImagesInput{}

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

	resp, respErr := rm.sdkapi.CreateFpgaImageWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateFpgaImage", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.FpgaImageGlobalId != nil {
		ko.Status.FPGAImageGlobalID = resp.FpgaImageGlobalId
	}
	if resp.FpgaImageId != nil {
		ko.Status.FPGAImageID = resp.FpgaImageId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateFpgaImageInput, error) {
	res := &svcsdk.CreateFpgaImageInput{}

	if r.ko.Spec.ClientToken != nil {
		res.SetClientToken(*r.ko.Spec.ClientToken)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.InputStorageLocation != nil {
		f3 := &svcsdk.StorageLocation{}
		if r.ko.Spec.InputStorageLocation.Bucket != nil {
			f3.SetBucket(*r.ko.Spec.InputStorageLocation.Bucket)
		}
		if r.ko.Spec.InputStorageLocation.Key != nil {
			f3.SetKey(*r.ko.Spec.InputStorageLocation.Key)
		}
		res.SetInputStorageLocation(f3)
	}
	if r.ko.Spec.LogsStorageLocation != nil {
		f4 := &svcsdk.StorageLocation{}
		if r.ko.Spec.LogsStorageLocation.Bucket != nil {
			f4.SetBucket(*r.ko.Spec.LogsStorageLocation.Bucket)
		}
		if r.ko.Spec.LogsStorageLocation.Key != nil {
			f4.SetKey(*r.ko.Spec.LogsStorageLocation.Key)
		}
		res.SetLogsStorageLocation(f4)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f6 := []*svcsdk.TagSpecification{}
		for _, f6iter := range r.ko.Spec.TagSpecifications {
			f6elem := &svcsdk.TagSpecification{}
			if f6iter.ResourceType != nil {
				f6elem.SetResourceType(*f6iter.ResourceType)
			}
			if f6iter.Tags != nil {
				f6elemf1 := []*svcsdk.Tag{}
				for _, f6elemf1iter := range f6iter.Tags {
					f6elemf1elem := &svcsdk.Tag{}
					if f6elemf1iter.Key != nil {
						f6elemf1elem.SetKey(*f6elemf1iter.Key)
					}
					if f6elemf1iter.Value != nil {
						f6elemf1elem.SetValue(*f6elemf1iter.Value)
					}
					f6elemf1 = append(f6elemf1, f6elemf1elem)
				}
				f6elem.SetTags(f6elemf1)
			}
			f6 = append(f6, f6elem)
		}
		res.SetTagSpecifications(f6)
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
	_, respErr := rm.sdkapi.DeleteFpgaImageWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFpgaImage", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFpgaImageInput, error) {
	res := &svcsdk.DeleteFpgaImageInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Status.FPGAImageID != nil {
		res.SetFpgaImageId(*r.ko.Status.FPGAImageID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.FPGAImage,
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