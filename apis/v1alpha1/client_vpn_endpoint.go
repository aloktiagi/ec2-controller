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

// ClientVPNEndpointSpec defines the desired state of ClientVPNEndpoint
type ClientVPNEndpointSpec struct {
	// +kubebuilder:validation:Required
	AuthenticationOptions []*ClientVPNAuthenticationRequest `json:"authenticationOptions"`
	// +kubebuilder:validation:Required
	ClientCIDRBlock *string `json:"clientCIDRBlock"`
	ClientToken     *string `json:"clientToken,omitempty"`
	// +kubebuilder:validation:Required
	ConnectionLogOptions *ConnectionLogOptions `json:"connectionLogOptions"`
	Description          *string               `json:"description,omitempty"`
	DNSServers           []*string             `json:"dnsServers,omitempty"`
	DryRun               *bool                 `json:"dryRun,omitempty"`
	SecurityGroupIDs     []*string             `json:"securityGroupIDs,omitempty"`
	// +kubebuilder:validation:Required
	ServerCertificateARN *string             `json:"serverCertificateARN"`
	SplitTunnel          *bool               `json:"splitTunnel,omitempty"`
	TagSpecifications    []*TagSpecification `json:"tagSpecifications,omitempty"`
	TransportProtocol    *string             `json:"transportProtocol,omitempty"`
	VPCID                *string             `json:"vpcID,omitempty"`
	VPNPort              *int64              `json:"vpnPort,omitempty"`
}

// ClientVPNEndpointStatus defines the observed state of ClientVPNEndpoint
type ClientVPNEndpointStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions          []*ackv1alpha1.Condition     `json:"conditions"`
	ClientVPNEndpointID *string                      `json:"clientVPNEndpointID,omitempty"`
	DNSName             *string                      `json:"dnsName,omitempty"`
	Status              *ClientVPNEndpointStatus_SDK `json:"status,omitempty"`
}

// ClientVPNEndpoint is the Schema for the ClientVPNEndpoints API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ClientVPNEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientVPNEndpointSpec   `json:"spec,omitempty"`
	Status            ClientVPNEndpointStatus `json:"status,omitempty"`
}

// ClientVPNEndpointList contains a list of ClientVPNEndpoint
// +kubebuilder:object:root=true
type ClientVPNEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientVPNEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClientVPNEndpoint{}, &ClientVPNEndpointList{})
}
