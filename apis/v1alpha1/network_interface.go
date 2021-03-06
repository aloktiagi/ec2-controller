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

// NetworkInterfaceSpec defines the desired state of NetworkInterface
type NetworkInterfaceSpec struct {
	Description                    *string                          `json:"description,omitempty"`
	DryRun                         *bool                            `json:"dryRun,omitempty"`
	Groups                         []*string                        `json:"groups,omitempty"`
	InterfaceType                  *string                          `json:"interfaceType,omitempty"`
	IPv6AddressCount               *int64                           `json:"ipv6AddressCount,omitempty"`
	IPv6Addresses                  []*InstanceIPv6Address           `json:"ipv6Addresses,omitempty"`
	PrivateIPAddress               *string                          `json:"privateIPAddress,omitempty"`
	PrivateIPAddresses             []*PrivateIPAddressSpecification `json:"privateIPAddresses,omitempty"`
	SecondaryPrivateIPAddressCount *int64                           `json:"secondaryPrivateIPAddressCount,omitempty"`
	// +kubebuilder:validation:Required
	SubnetID          *string             `json:"subnetID"`
	TagSpecifications []*TagSpecification `json:"tagSpecifications,omitempty"`
}

// NetworkInterfaceStatus defines the observed state of NetworkInterface
type NetworkInterfaceStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions         []*ackv1alpha1.Condition     `json:"conditions"`
	Association        *NetworkInterfaceAssociation `json:"association,omitempty"`
	Attachment         *NetworkInterfaceAttachment  `json:"attachment,omitempty"`
	AvailabilityZone   *string                      `json:"availabilityZone,omitempty"`
	MacAddress         *string                      `json:"macAddress,omitempty"`
	NetworkInterfaceID *string                      `json:"networkInterfaceID,omitempty"`
	OutpostARN         *string                      `json:"outpostARN,omitempty"`
	OwnerID            *string                      `json:"ownerID,omitempty"`
	PrivateDNSName     *string                      `json:"privateDNSName,omitempty"`
	RequesterID        *string                      `json:"requesterID,omitempty"`
	RequesterManaged   *bool                        `json:"requesterManaged,omitempty"`
	SourceDestCheck    *bool                        `json:"sourceDestCheck,omitempty"`
	Status             *string                      `json:"status,omitempty"`
	TagSet             []*Tag                       `json:"tagSet,omitempty"`
	VPCID              *string                      `json:"vpcID,omitempty"`
}

// NetworkInterface is the Schema for the NetworkInterfaces API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

// NetworkInterfaceList contains a list of NetworkInterface
// +kubebuilder:object:root=true
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}
