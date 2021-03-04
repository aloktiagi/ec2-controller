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

package reserved_instances_listing

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
	_ = &svcapitypes.ReservedInstancesListing{}
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

	resp, respErr := rm.sdkapi.DescribeReservedInstancesListingsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeReservedInstancesListings", respErr)
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
	for _, elem := range resp.ReservedInstancesListings {
		if elem.ClientToken != nil {
			ko.Spec.ClientToken = elem.ClientToken
		}
		if elem.PriceSchedules != nil {
			f3 := []*svcapitypes.PriceScheduleSpecification{}
			for _, f3iter := range elem.PriceSchedules {
				f3elem := &svcapitypes.PriceScheduleSpecification{}
				if f3iter.CurrencyCode != nil {
					f3elem.CurrencyCode = f3iter.CurrencyCode
				}
				if f3iter.Price != nil {
					f3elem.Price = f3iter.Price
				}
				if f3iter.Term != nil {
					f3elem.Term = f3iter.Term
				}
				f3 = append(f3, f3elem)
			}
			ko.Spec.PriceSchedules = f3
		}
		if elem.ReservedInstancesId != nil {
			ko.Spec.ReservedInstancesID = elem.ReservedInstancesId
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
) (*svcsdk.DescribeReservedInstancesListingsInput, error) {
	res := &svcsdk.DescribeReservedInstancesListingsInput{}

	if r.ko.Spec.ReservedInstancesID != nil {
		res.SetReservedInstancesId(*r.ko.Spec.ReservedInstancesID)
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

	resp, respErr := rm.sdkapi.CreateReservedInstancesListingWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateReservedInstancesListing", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ReservedInstancesListings != nil {
		f0 := []*svcapitypes.ReservedInstancesListing_SDK{}
		for _, f0iter := range resp.ReservedInstancesListings {
			f0elem := &svcapitypes.ReservedInstancesListing_SDK{}
			if f0iter.ClientToken != nil {
				f0elem.ClientToken = f0iter.ClientToken
			}
			if f0iter.CreateDate != nil {
				f0elem.CreateDate = &metav1.Time{*f0iter.CreateDate}
			}
			if f0iter.InstanceCounts != nil {
				f0elemf2 := []*svcapitypes.InstanceCount{}
				for _, f0elemf2iter := range f0iter.InstanceCounts {
					f0elemf2elem := &svcapitypes.InstanceCount{}
					if f0elemf2iter.InstanceCount != nil {
						f0elemf2elem.InstanceCount = f0elemf2iter.InstanceCount
					}
					if f0elemf2iter.State != nil {
						f0elemf2elem.State = f0elemf2iter.State
					}
					f0elemf2 = append(f0elemf2, f0elemf2elem)
				}
				f0elem.InstanceCounts = f0elemf2
			}
			if f0iter.PriceSchedules != nil {
				f0elemf3 := []*svcapitypes.PriceSchedule{}
				for _, f0elemf3iter := range f0iter.PriceSchedules {
					f0elemf3elem := &svcapitypes.PriceSchedule{}
					if f0elemf3iter.Active != nil {
						f0elemf3elem.Active = f0elemf3iter.Active
					}
					if f0elemf3iter.CurrencyCode != nil {
						f0elemf3elem.CurrencyCode = f0elemf3iter.CurrencyCode
					}
					if f0elemf3iter.Price != nil {
						f0elemf3elem.Price = f0elemf3iter.Price
					}
					if f0elemf3iter.Term != nil {
						f0elemf3elem.Term = f0elemf3iter.Term
					}
					f0elemf3 = append(f0elemf3, f0elemf3elem)
				}
				f0elem.PriceSchedules = f0elemf3
			}
			if f0iter.ReservedInstancesId != nil {
				f0elem.ReservedInstancesID = f0iter.ReservedInstancesId
			}
			if f0iter.ReservedInstancesListingId != nil {
				f0elem.ReservedInstancesListingID = f0iter.ReservedInstancesListingId
			}
			if f0iter.Status != nil {
				f0elem.Status = f0iter.Status
			}
			if f0iter.StatusMessage != nil {
				f0elem.StatusMessage = f0iter.StatusMessage
			}
			if f0iter.Tags != nil {
				f0elemf8 := []*svcapitypes.Tag{}
				for _, f0elemf8iter := range f0iter.Tags {
					f0elemf8elem := &svcapitypes.Tag{}
					if f0elemf8iter.Key != nil {
						f0elemf8elem.Key = f0elemf8iter.Key
					}
					if f0elemf8iter.Value != nil {
						f0elemf8elem.Value = f0elemf8iter.Value
					}
					f0elemf8 = append(f0elemf8, f0elemf8elem)
				}
				f0elem.Tags = f0elemf8
			}
			if f0iter.UpdateDate != nil {
				f0elem.UpdateDate = &metav1.Time{*f0iter.UpdateDate}
			}
			f0 = append(f0, f0elem)
		}
		ko.Status.ReservedInstancesListings = f0
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateReservedInstancesListingInput, error) {
	res := &svcsdk.CreateReservedInstancesListingInput{}

	if r.ko.Spec.ClientToken != nil {
		res.SetClientToken(*r.ko.Spec.ClientToken)
	}
	if r.ko.Spec.InstanceCount != nil {
		res.SetInstanceCount(*r.ko.Spec.InstanceCount)
	}
	if r.ko.Spec.PriceSchedules != nil {
		f2 := []*svcsdk.PriceScheduleSpecification{}
		for _, f2iter := range r.ko.Spec.PriceSchedules {
			f2elem := &svcsdk.PriceScheduleSpecification{}
			if f2iter.CurrencyCode != nil {
				f2elem.SetCurrencyCode(*f2iter.CurrencyCode)
			}
			if f2iter.Price != nil {
				f2elem.SetPrice(*f2iter.Price)
			}
			if f2iter.Term != nil {
				f2elem.SetTerm(*f2iter.Term)
			}
			f2 = append(f2, f2elem)
		}
		res.SetPriceSchedules(f2)
	}
	if r.ko.Spec.ReservedInstancesID != nil {
		res.SetReservedInstancesId(*r.ko.Spec.ReservedInstancesID)
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
	// TODO(jaypipes): Figure this out...
	return nil

}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ReservedInstancesListing,
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