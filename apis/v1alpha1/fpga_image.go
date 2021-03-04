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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FPGAImageSpec defines the desired state of FPGAImage
type FPGAImageSpec struct {
	ClientToken *string `json:"clientToken,omitempty"`
	Description *string `json:"description,omitempty"`
	DryRun      *bool   `json:"dryRun,omitempty"`
	// +kubebuilder:validation:Required
	InputStorageLocation *StorageLocation    `json:"inputStorageLocation"`
	LogsStorageLocation  *StorageLocation    `json:"logsStorageLocation,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	TagSpecifications    []*TagSpecification `json:"tagSpecifications,omitempty"`
}

// FPGAImageStatus defines the observed state of FPGAImage
type FPGAImageStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions        []*ackv1alpha1.Condition `json:"conditions"`
	FPGAImageGlobalID *string                  `json:"fpgaImageGlobalID,omitempty"`
	FPGAImageID       *string                  `json:"fpgaImageID,omitempty"`
}

// FPGAImage is the Schema for the FPGAImages API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type FPGAImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FPGAImageSpec   `json:"spec,omitempty"`
	Status            FPGAImageStatus `json:"status,omitempty"`
}

// FPGAImageList contains a list of FPGAImage
// +kubebuilder:object:root=true
type FPGAImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FPGAImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FPGAImage{}, &FPGAImageList{})
}