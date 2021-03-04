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

// NetworkACLEntrySpec defines the desired state of NetworkACLEntry
type NetworkACLEntrySpec struct {
	CIDRBlock *string `json:"cidrBlock,omitempty"`
	DryRun    *bool   `json:"dryRun,omitempty"`
	// +kubebuilder:validation:Required
	Egress        *bool         `json:"egress"`
	ICMPTypeCode  *ICMPTypeCode `json:"icmpTypeCode,omitempty"`
	IPv6CIDRBlock *string       `json:"ipv6CIDRBlock,omitempty"`
	// +kubebuilder:validation:Required
	NetworkACLID *string    `json:"networkACLID"`
	PortRange    *PortRange `json:"portRange,omitempty"`
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol"`
	// +kubebuilder:validation:Required
	RuleAction *string `json:"ruleAction"`
	// +kubebuilder:validation:Required
	RuleNumber *int64 `json:"ruleNumber"`
}

// NetworkACLEntryStatus defines the observed state of NetworkACLEntry
type NetworkACLEntryStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// NetworkACLEntry is the Schema for the NetworkACLEntries API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type NetworkACLEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLEntrySpec   `json:"spec,omitempty"`
	Status            NetworkACLEntryStatus `json:"status,omitempty"`
}

// NetworkACLEntryList contains a list of NetworkACLEntry
// +kubebuilder:object:root=true
type NetworkACLEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACLEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkACLEntry{}, &NetworkACLEntryList{})
}