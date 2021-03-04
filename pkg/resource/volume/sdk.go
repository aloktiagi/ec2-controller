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

package volume

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
	_ = &svcapitypes.Volume{}
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

	resp, respErr := rm.sdkapi.DescribeVolumesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeVolumes", respErr)
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
	for _, elem := range resp.Volumes {
		if elem.Attachments != nil {
			f0 := []*svcapitypes.VolumeAttachment{}
			for _, f0iter := range elem.Attachments {
				f0elem := &svcapitypes.VolumeAttachment{}
				if f0iter.AttachTime != nil {
					f0elem.AttachTime = &metav1.Time{*f0iter.AttachTime}
				}
				if f0iter.DeleteOnTermination != nil {
					f0elem.DeleteOnTermination = f0iter.DeleteOnTermination
				}
				if f0iter.Device != nil {
					f0elem.Device = f0iter.Device
				}
				if f0iter.InstanceId != nil {
					f0elem.InstanceID = f0iter.InstanceId
				}
				if f0iter.State != nil {
					f0elem.State = f0iter.State
				}
				if f0iter.VolumeId != nil {
					f0elem.VolumeID = f0iter.VolumeId
				}
				f0 = append(f0, f0elem)
			}
			ko.Status.Attachments = f0
		}
		if elem.AvailabilityZone != nil {
			ko.Spec.AvailabilityZone = elem.AvailabilityZone
		}
		if elem.CreateTime != nil {
			ko.Status.CreateTime = &metav1.Time{*elem.CreateTime}
		}
		if elem.Encrypted != nil {
			ko.Spec.Encrypted = elem.Encrypted
		}
		if elem.FastRestored != nil {
			ko.Status.FastRestored = elem.FastRestored
		}
		if elem.Iops != nil {
			ko.Spec.IOPS = elem.Iops
		}
		if elem.KmsKeyId != nil {
			ko.Spec.KMSKeyID = elem.KmsKeyId
		}
		if elem.MultiAttachEnabled != nil {
			ko.Spec.MultiAttachEnabled = elem.MultiAttachEnabled
		}
		if elem.OutpostArn != nil {
			ko.Spec.OutpostARN = elem.OutpostArn
		}
		if elem.Size != nil {
			ko.Spec.Size = elem.Size
		}
		if elem.SnapshotId != nil {
			ko.Spec.SnapshotID = elem.SnapshotId
		}
		if elem.State != nil {
			ko.Status.State = elem.State
		}
		if elem.Tags != nil {
			f12 := []*svcapitypes.Tag{}
			for _, f12iter := range elem.Tags {
				f12elem := &svcapitypes.Tag{}
				if f12iter.Key != nil {
					f12elem.Key = f12iter.Key
				}
				if f12iter.Value != nil {
					f12elem.Value = f12iter.Value
				}
				f12 = append(f12, f12elem)
			}
			ko.Status.Tags = f12
		}
		if elem.Throughput != nil {
			ko.Spec.Throughput = elem.Throughput
		}
		if elem.VolumeId != nil {
			ko.Status.VolumeID = elem.VolumeId
		}
		if elem.VolumeType != nil {
			ko.Spec.VolumeType = elem.VolumeType
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
) (*svcsdk.DescribeVolumesInput, error) {
	res := &svcsdk.DescribeVolumesInput{}

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

	resp, respErr := rm.sdkapi.CreateVolumeWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateVolume", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Attachments != nil {
		f0 := []*svcapitypes.VolumeAttachment{}
		for _, f0iter := range resp.Attachments {
			f0elem := &svcapitypes.VolumeAttachment{}
			if f0iter.AttachTime != nil {
				f0elem.AttachTime = &metav1.Time{*f0iter.AttachTime}
			}
			if f0iter.DeleteOnTermination != nil {
				f0elem.DeleteOnTermination = f0iter.DeleteOnTermination
			}
			if f0iter.Device != nil {
				f0elem.Device = f0iter.Device
			}
			if f0iter.InstanceId != nil {
				f0elem.InstanceID = f0iter.InstanceId
			}
			if f0iter.State != nil {
				f0elem.State = f0iter.State
			}
			if f0iter.VolumeId != nil {
				f0elem.VolumeID = f0iter.VolumeId
			}
			f0 = append(f0, f0elem)
		}
		ko.Status.Attachments = f0
	}
	if resp.CreateTime != nil {
		ko.Status.CreateTime = &metav1.Time{*resp.CreateTime}
	}
	if resp.FastRestored != nil {
		ko.Status.FastRestored = resp.FastRestored
	}
	if resp.State != nil {
		ko.Status.State = resp.State
	}
	if resp.Tags != nil {
		f12 := []*svcapitypes.Tag{}
		for _, f12iter := range resp.Tags {
			f12elem := &svcapitypes.Tag{}
			if f12iter.Key != nil {
				f12elem.Key = f12iter.Key
			}
			if f12iter.Value != nil {
				f12elem.Value = f12iter.Value
			}
			f12 = append(f12, f12elem)
		}
		ko.Status.Tags = f12
	}
	if resp.VolumeId != nil {
		ko.Status.VolumeID = resp.VolumeId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateVolumeInput, error) {
	res := &svcsdk.CreateVolumeInput{}

	if r.ko.Spec.AvailabilityZone != nil {
		res.SetAvailabilityZone(*r.ko.Spec.AvailabilityZone)
	}
	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.Encrypted != nil {
		res.SetEncrypted(*r.ko.Spec.Encrypted)
	}
	if r.ko.Spec.IOPS != nil {
		res.SetIops(*r.ko.Spec.IOPS)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.MultiAttachEnabled != nil {
		res.SetMultiAttachEnabled(*r.ko.Spec.MultiAttachEnabled)
	}
	if r.ko.Spec.OutpostARN != nil {
		res.SetOutpostArn(*r.ko.Spec.OutpostARN)
	}
	if r.ko.Spec.Size != nil {
		res.SetSize(*r.ko.Spec.Size)
	}
	if r.ko.Spec.SnapshotID != nil {
		res.SetSnapshotId(*r.ko.Spec.SnapshotID)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f9 := []*svcsdk.TagSpecification{}
		for _, f9iter := range r.ko.Spec.TagSpecifications {
			f9elem := &svcsdk.TagSpecification{}
			if f9iter.ResourceType != nil {
				f9elem.SetResourceType(*f9iter.ResourceType)
			}
			if f9iter.Tags != nil {
				f9elemf1 := []*svcsdk.Tag{}
				for _, f9elemf1iter := range f9iter.Tags {
					f9elemf1elem := &svcsdk.Tag{}
					if f9elemf1iter.Key != nil {
						f9elemf1elem.SetKey(*f9elemf1iter.Key)
					}
					if f9elemf1iter.Value != nil {
						f9elemf1elem.SetValue(*f9elemf1iter.Value)
					}
					f9elemf1 = append(f9elemf1, f9elemf1elem)
				}
				f9elem.SetTags(f9elemf1)
			}
			f9 = append(f9, f9elem)
		}
		res.SetTagSpecifications(f9)
	}
	if r.ko.Spec.Throughput != nil {
		res.SetThroughput(*r.ko.Spec.Throughput)
	}
	if r.ko.Spec.VolumeType != nil {
		res.SetVolumeType(*r.ko.Spec.VolumeType)
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

	resp, respErr := rm.sdkapi.ModifyVolumeWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyVolume", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.VolumeModification.VolumeId != nil {
		ko.Status.VolumeID = resp.VolumeModification.VolumeId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ModifyVolumeInput, error) {
	res := &svcsdk.ModifyVolumeInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.IOPS != nil {
		res.SetIops(*r.ko.Spec.IOPS)
	}
	if r.ko.Spec.MultiAttachEnabled != nil {
		res.SetMultiAttachEnabled(*r.ko.Spec.MultiAttachEnabled)
	}
	if r.ko.Spec.Size != nil {
		res.SetSize(*r.ko.Spec.Size)
	}
	if r.ko.Spec.Throughput != nil {
		res.SetThroughput(*r.ko.Spec.Throughput)
	}
	if r.ko.Status.VolumeID != nil {
		res.SetVolumeId(*r.ko.Status.VolumeID)
	}
	if r.ko.Spec.VolumeType != nil {
		res.SetVolumeType(*r.ko.Spec.VolumeType)
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
	_, respErr := rm.sdkapi.DeleteVolumeWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteVolume", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteVolumeInput, error) {
	res := &svcsdk.DeleteVolumeInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Status.VolumeID != nil {
		res.SetVolumeId(*r.ko.Status.VolumeID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Volume,
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