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

// InternetGatewaySpec defines the desired state of InternetGateway
type InternetGatewaySpec struct {
	DryRun            *bool               `json:"dryRun,omitempty"`
	TagSpecifications []*TagSpecification `json:"tagSpecifications,omitempty"`
}

// InternetGatewayStatus defines the observed state of InternetGateway
type InternetGatewayStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions        []*ackv1alpha1.Condition     `json:"conditions"`
	Attachments       []*InternetGatewayAttachment `json:"attachments,omitempty"`
	InternetGatewayID *string                      `json:"internetGatewayID,omitempty"`
	OwnerID           *string                      `json:"ownerID,omitempty"`
	Tags              []*Tag                       `json:"tags,omitempty"`
}

// InternetGateway is the Schema for the InternetGateways API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type InternetGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InternetGatewaySpec   `json:"spec,omitempty"`
	Status            InternetGatewayStatus `json:"status,omitempty"`
}

// InternetGatewayList contains a list of InternetGateway
// +kubebuilder:object:root=true
type InternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InternetGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InternetGateway{}, &InternetGatewayList{})
}