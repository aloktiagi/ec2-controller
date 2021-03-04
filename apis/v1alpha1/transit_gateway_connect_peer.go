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

// TransitGatewayConnectPeerSpec defines the desired state of TransitGatewayConnectPeer
type TransitGatewayConnectPeerSpec struct {
	BGPOptions *TransitGatewayConnectRequestBGPOptions `json:"bgpOptions,omitempty"`
	DryRun     *bool                                   `json:"dryRun,omitempty"`
	// +kubebuilder:validation:Required
	InsideCIDRBlocks []*string `json:"insideCIDRBlocks"`
	// +kubebuilder:validation:Required
	PeerAddress           *string             `json:"peerAddress"`
	TagSpecifications     []*TagSpecification `json:"tagSpecifications,omitempty"`
	TransitGatewayAddress *string             `json:"transitGatewayAddress,omitempty"`
	// +kubebuilder:validation:Required
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentID"`
}

// TransitGatewayConnectPeerStatus defines the observed state of TransitGatewayConnectPeer
type TransitGatewayConnectPeerStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions                  []*ackv1alpha1.Condition                `json:"conditions"`
	ConnectPeerConfiguration    *TransitGatewayConnectPeerConfiguration `json:"connectPeerConfiguration,omitempty"`
	CreationTime                *metav1.Time                            `json:"creationTime,omitempty"`
	State                       *string                                 `json:"state,omitempty"`
	Tags                        []*Tag                                  `json:"tags,omitempty"`
	TransitGatewayConnectPeerID *string                                 `json:"transitGatewayConnectPeerID,omitempty"`
}

// TransitGatewayConnectPeer is the Schema for the TransitGatewayConnectPeers API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type TransitGatewayConnectPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayConnectPeerSpec   `json:"spec,omitempty"`
	Status            TransitGatewayConnectPeerStatus `json:"status,omitempty"`
}

// TransitGatewayConnectPeerList contains a list of TransitGatewayConnectPeer
// +kubebuilder:object:root=true
type TransitGatewayConnectPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayConnectPeer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TransitGatewayConnectPeer{}, &TransitGatewayConnectPeerList{})
}
