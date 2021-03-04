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

// VPNConnectionSpec defines the desired state of VPNConnection
type VPNConnectionSpec struct {
	// +kubebuilder:validation:Required
	CustomerGatewayID *string                            `json:"customerGatewayID"`
	DryRun            *bool                              `json:"dryRun,omitempty"`
	Options           *VPNConnectionOptionsSpecification `json:"options,omitempty"`
	TagSpecifications []*TagSpecification                `json:"tagSpecifications,omitempty"`
	TransitGatewayID  *string                            `json:"transitGatewayID,omitempty"`
	// +kubebuilder:validation:Required
	Type         *string `json:"type_"`
	VPNGatewayID *string `json:"vpnGatewayID,omitempty"`
}

// VPNConnectionStatus defines the observed state of VPNConnection
type VPNConnectionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions                   []*ackv1alpha1.Condition `json:"conditions"`
	Category                     *string                  `json:"category,omitempty"`
	CustomerGatewayConfiguration *string                  `json:"customerGatewayConfiguration,omitempty"`
	Routes                       []*VPNStaticRoute        `json:"routes,omitempty"`
	State                        *string                  `json:"state,omitempty"`
	Tags                         []*Tag                   `json:"tags,omitempty"`
	VGWTelemetry                 []*VGWTelemetry          `json:"vgwTelemetry,omitempty"`
	VPNConnectionID              *string                  `json:"vpnConnectionID,omitempty"`
}

// VPNConnection is the Schema for the VPNConnections API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VPNConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNConnectionSpec   `json:"spec,omitempty"`
	Status            VPNConnectionStatus `json:"status,omitempty"`
}

// VPNConnectionList contains a list of VPNConnection
// +kubebuilder:object:root=true
type VPNConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPNConnection{}, &VPNConnectionList{})
}