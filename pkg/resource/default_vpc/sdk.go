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

package default_vpc

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
	_ = &svcapitypes.DefaultVPC{}
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

	resp, respErr := rm.sdkapi.CreateDefaultVpcWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateDefaultVpc", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Vpc.CidrBlock != nil {
		ko.Status.CIDRBlock = resp.Vpc.CidrBlock
	}
	if resp.Vpc.CidrBlockAssociationSet != nil {
		f1 := []*svcapitypes.VPCCIDRBlockAssociation{}
		for _, f1iter := range resp.Vpc.CidrBlockAssociationSet {
			f1elem := &svcapitypes.VPCCIDRBlockAssociation{}
			if f1iter.AssociationId != nil {
				f1elem.AssociationID = f1iter.AssociationId
			}
			if f1iter.CidrBlock != nil {
				f1elem.CIDRBlock = f1iter.CidrBlock
			}
			if f1iter.CidrBlockState != nil {
				f1elemf2 := &svcapitypes.VPCCIDRBlockState{}
				if f1iter.CidrBlockState.State != nil {
					f1elemf2.State = f1iter.CidrBlockState.State
				}
				if f1iter.CidrBlockState.StatusMessage != nil {
					f1elemf2.StatusMessage = f1iter.CidrBlockState.StatusMessage
				}
				f1elem.CIDRBlockState = f1elemf2
			}
			f1 = append(f1, f1elem)
		}
		ko.Status.CIDRBlockAssociationSet = f1
	}
	if resp.Vpc.DhcpOptionsId != nil {
		ko.Status.DHCPOptionsID = resp.Vpc.DhcpOptionsId
	}
	if resp.Vpc.InstanceTenancy != nil {
		ko.Status.InstanceTenancy = resp.Vpc.InstanceTenancy
	}
	if resp.Vpc.Ipv6CidrBlockAssociationSet != nil {
		f4 := []*svcapitypes.VPCIPv6CIDRBlockAssociation{}
		for _, f4iter := range resp.Vpc.Ipv6CidrBlockAssociationSet {
			f4elem := &svcapitypes.VPCIPv6CIDRBlockAssociation{}
			if f4iter.AssociationId != nil {
				f4elem.AssociationID = f4iter.AssociationId
			}
			if f4iter.Ipv6CidrBlock != nil {
				f4elem.IPv6CIDRBlock = f4iter.Ipv6CidrBlock
			}
			if f4iter.Ipv6CidrBlockState != nil {
				f4elemf2 := &svcapitypes.VPCCIDRBlockState{}
				if f4iter.Ipv6CidrBlockState.State != nil {
					f4elemf2.State = f4iter.Ipv6CidrBlockState.State
				}
				if f4iter.Ipv6CidrBlockState.StatusMessage != nil {
					f4elemf2.StatusMessage = f4iter.Ipv6CidrBlockState.StatusMessage
				}
				f4elem.IPv6CIDRBlockState = f4elemf2
			}
			if f4iter.Ipv6Pool != nil {
				f4elem.IPv6Pool = f4iter.Ipv6Pool
			}
			if f4iter.NetworkBorderGroup != nil {
				f4elem.NetworkBorderGroup = f4iter.NetworkBorderGroup
			}
			f4 = append(f4, f4elem)
		}
		ko.Status.IPv6CIDRBlockAssociationSet = f4
	}
	if resp.Vpc.IsDefault != nil {
		ko.Status.IsDefault = resp.Vpc.IsDefault
	}
	if resp.Vpc.OwnerId != nil {
		ko.Status.OwnerID = resp.Vpc.OwnerId
	}
	if resp.Vpc.State != nil {
		ko.Status.State = resp.Vpc.State
	}
	if resp.Vpc.Tags != nil {
		f8 := []*svcapitypes.Tag{}
		for _, f8iter := range resp.Vpc.Tags {
			f8elem := &svcapitypes.Tag{}
			if f8iter.Key != nil {
				f8elem.Key = f8iter.Key
			}
			if f8iter.Value != nil {
				f8elem.Value = f8iter.Value
			}
			f8 = append(f8, f8elem)
		}
		ko.Status.Tags = f8
	}
	if resp.Vpc.VpcId != nil {
		ko.Status.VPCID = resp.Vpc.VpcId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateDefaultVpcInput, error) {
	res := &svcsdk.CreateDefaultVpcInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
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
	ko *svcapitypes.DefaultVPC,
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