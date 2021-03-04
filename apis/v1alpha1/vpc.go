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

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	AmazonProvidedIPv6CIDRBlock *bool `json:"amazonProvidedIPv6CIDRBlock,omitempty"`
	// +kubebuilder:validation:Required
	CIDRBlock                       *string             `json:"cidrBlock"`
	DryRun                          *bool               `json:"dryRun,omitempty"`
	InstanceTenancy                 *string             `json:"instanceTenancy,omitempty"`
	IPv6CIDRBlock                   *string             `json:"ipv6CIDRBlock,omitempty"`
	IPv6CIDRBlockNetworkBorderGroup *string             `json:"ipv6CIDRBlockNetworkBorderGroup,omitempty"`
	IPv6Pool                        *string             `json:"ipv6Pool,omitempty"`
	TagSpecifications               []*TagSpecification `json:"tagSpecifications,omitempty"`
}

// VPCStatus defines the observed state of VPC
type VPCStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions                  []*ackv1alpha1.Condition       `json:"conditions"`
	CIDRBlockAssociationSet     []*VPCCIDRBlockAssociation     `json:"cidrBlockAssociationSet,omitempty"`
	DHCPOptionsID               *string                        `json:"dhcpOptionsID,omitempty"`
	IPv6CIDRBlockAssociationSet []*VPCIPv6CIDRBlockAssociation `json:"ipv6CIDRBlockAssociationSet,omitempty"`
	IsDefault                   *bool                          `json:"isDefault,omitempty"`
	OwnerID                     *string                        `json:"ownerID,omitempty"`
	State                       *string                        `json:"state,omitempty"`
	Tags                        []*Tag                         `json:"tags,omitempty"`
	VPCID                       *string                        `json:"vpcID,omitempty"`
}

// VPC is the Schema for the VPCS API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec,omitempty"`
	Status            VPCStatus `json:"status,omitempty"`
}

// VPCList contains a list of VPC
// +kubebuilder:object:root=true
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}