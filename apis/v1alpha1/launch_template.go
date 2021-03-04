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

// LaunchTemplateSpec defines the desired state of LaunchTemplate
type LaunchTemplateSpec struct {
	ClientToken *string `json:"clientToken,omitempty"`
	DryRun      *bool   `json:"dryRun,omitempty"`
	// +kubebuilder:validation:Required
	LaunchTemplateData *RequestLaunchTemplateData `json:"launchTemplateData"`
	// +kubebuilder:validation:Required
	LaunchTemplateName *string             `json:"launchTemplateName"`
	TagSpecifications  []*TagSpecification `json:"tagSpecifications,omitempty"`
	VersionDescription *string             `json:"versionDescription,omitempty"`
}

// LaunchTemplateStatus defines the observed state of LaunchTemplate
type LaunchTemplateStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions     []*ackv1alpha1.Condition `json:"conditions"`
	LaunchTemplate *LaunchTemplate_SDK      `json:"launchTemplate,omitempty"`
	Warning        *ValidationWarning       `json:"warning,omitempty"`
}

// LaunchTemplate is the Schema for the LaunchTemplates API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec,omitempty"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

// LaunchTemplateList contains a list of LaunchTemplate
// +kubebuilder:object:root=true
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LaunchTemplate{}, &LaunchTemplateList{})
}