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

package route_table

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
	_ = &svcapitypes.RouteTable{}
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

	resp, respErr := rm.sdkapi.DescribeRouteTablesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeRouteTables", respErr)
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
	for _, elem := range resp.RouteTables {
		if elem.Associations != nil {
			f0 := []*svcapitypes.RouteTableAssociation{}
			for _, f0iter := range elem.Associations {
				f0elem := &svcapitypes.RouteTableAssociation{}
				if f0iter.AssociationState != nil {
					f0elemf0 := &svcapitypes.RouteTableAssociationState{}
					if f0iter.AssociationState.State != nil {
						f0elemf0.State = f0iter.AssociationState.State
					}
					if f0iter.AssociationState.StatusMessage != nil {
						f0elemf0.StatusMessage = f0iter.AssociationState.StatusMessage
					}
					f0elem.AssociationState = f0elemf0
				}
				if f0iter.GatewayId != nil {
					f0elem.GatewayID = f0iter.GatewayId
				}
				if f0iter.Main != nil {
					f0elem.Main = f0iter.Main
				}
				if f0iter.RouteTableAssociationId != nil {
					f0elem.RouteTableAssociationID = f0iter.RouteTableAssociationId
				}
				if f0iter.RouteTableId != nil {
					f0elem.RouteTableID = f0iter.RouteTableId
				}
				if f0iter.SubnetId != nil {
					f0elem.SubnetID = f0iter.SubnetId
				}
				f0 = append(f0, f0elem)
			}
			ko.Status.Associations = f0
		}
		if elem.OwnerId != nil {
			ko.Status.OwnerID = elem.OwnerId
		}
		if elem.PropagatingVgws != nil {
			f2 := []*svcapitypes.PropagatingVGW{}
			for _, f2iter := range elem.PropagatingVgws {
				f2elem := &svcapitypes.PropagatingVGW{}
				if f2iter.GatewayId != nil {
					f2elem.GatewayID = f2iter.GatewayId
				}
				f2 = append(f2, f2elem)
			}
			ko.Status.PropagatingVGWs = f2
		}
		if elem.RouteTableId != nil {
			ko.Status.RouteTableID = elem.RouteTableId
		}
		if elem.Routes != nil {
			f4 := []*svcapitypes.Route_SDK{}
			for _, f4iter := range elem.Routes {
				f4elem := &svcapitypes.Route_SDK{}
				if f4iter.CarrierGatewayId != nil {
					f4elem.CarrierGatewayID = f4iter.CarrierGatewayId
				}
				if f4iter.DestinationCidrBlock != nil {
					f4elem.DestinationCIDRBlock = f4iter.DestinationCidrBlock
				}
				if f4iter.DestinationIpv6CidrBlock != nil {
					f4elem.DestinationIPv6CIDRBlock = f4iter.DestinationIpv6CidrBlock
				}
				if f4iter.DestinationPrefixListId != nil {
					f4elem.DestinationPrefixListID = f4iter.DestinationPrefixListId
				}
				if f4iter.EgressOnlyInternetGatewayId != nil {
					f4elem.EgressOnlyInternetGatewayID = f4iter.EgressOnlyInternetGatewayId
				}
				if f4iter.GatewayId != nil {
					f4elem.GatewayID = f4iter.GatewayId
				}
				if f4iter.InstanceId != nil {
					f4elem.InstanceID = f4iter.InstanceId
				}
				if f4iter.InstanceOwnerId != nil {
					f4elem.InstanceOwnerID = f4iter.InstanceOwnerId
				}
				if f4iter.LocalGatewayId != nil {
					f4elem.LocalGatewayID = f4iter.LocalGatewayId
				}
				if f4iter.NatGatewayId != nil {
					f4elem.NatGatewayID = f4iter.NatGatewayId
				}
				if f4iter.NetworkInterfaceId != nil {
					f4elem.NetworkInterfaceID = f4iter.NetworkInterfaceId
				}
				if f4iter.Origin != nil {
					f4elem.Origin = f4iter.Origin
				}
				if f4iter.State != nil {
					f4elem.State = f4iter.State
				}
				if f4iter.TransitGatewayId != nil {
					f4elem.TransitGatewayID = f4iter.TransitGatewayId
				}
				if f4iter.VpcPeeringConnectionId != nil {
					f4elem.VPCPeeringConnectionID = f4iter.VpcPeeringConnectionId
				}
				f4 = append(f4, f4elem)
			}
			ko.Status.Routes = f4
		}
		if elem.Tags != nil {
			f5 := []*svcapitypes.Tag{}
			for _, f5iter := range elem.Tags {
				f5elem := &svcapitypes.Tag{}
				if f5iter.Key != nil {
					f5elem.Key = f5iter.Key
				}
				if f5iter.Value != nil {
					f5elem.Value = f5iter.Value
				}
				f5 = append(f5, f5elem)
			}
			ko.Status.Tags = f5
		}
		if elem.VpcId != nil {
			ko.Spec.VPCID = elem.VpcId
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
) (*svcsdk.DescribeRouteTablesInput, error) {
	res := &svcsdk.DescribeRouteTablesInput{}

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

	resp, respErr := rm.sdkapi.CreateRouteTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateRouteTable", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.RouteTable.Associations != nil {
		f0 := []*svcapitypes.RouteTableAssociation{}
		for _, f0iter := range resp.RouteTable.Associations {
			f0elem := &svcapitypes.RouteTableAssociation{}
			if f0iter.AssociationState != nil {
				f0elemf0 := &svcapitypes.RouteTableAssociationState{}
				if f0iter.AssociationState.State != nil {
					f0elemf0.State = f0iter.AssociationState.State
				}
				if f0iter.AssociationState.StatusMessage != nil {
					f0elemf0.StatusMessage = f0iter.AssociationState.StatusMessage
				}
				f0elem.AssociationState = f0elemf0
			}
			if f0iter.GatewayId != nil {
				f0elem.GatewayID = f0iter.GatewayId
			}
			if f0iter.Main != nil {
				f0elem.Main = f0iter.Main
			}
			if f0iter.RouteTableAssociationId != nil {
				f0elem.RouteTableAssociationID = f0iter.RouteTableAssociationId
			}
			if f0iter.RouteTableId != nil {
				f0elem.RouteTableID = f0iter.RouteTableId
			}
			if f0iter.SubnetId != nil {
				f0elem.SubnetID = f0iter.SubnetId
			}
			f0 = append(f0, f0elem)
		}
		ko.Status.Associations = f0
	}
	if resp.RouteTable.OwnerId != nil {
		ko.Status.OwnerID = resp.RouteTable.OwnerId
	}
	if resp.RouteTable.PropagatingVgws != nil {
		f2 := []*svcapitypes.PropagatingVGW{}
		for _, f2iter := range resp.RouteTable.PropagatingVgws {
			f2elem := &svcapitypes.PropagatingVGW{}
			if f2iter.GatewayId != nil {
				f2elem.GatewayID = f2iter.GatewayId
			}
			f2 = append(f2, f2elem)
		}
		ko.Status.PropagatingVGWs = f2
	}
	if resp.RouteTable.RouteTableId != nil {
		ko.Status.RouteTableID = resp.RouteTable.RouteTableId
	}
	if resp.RouteTable.Routes != nil {
		f4 := []*svcapitypes.Route_SDK{}
		for _, f4iter := range resp.RouteTable.Routes {
			f4elem := &svcapitypes.Route_SDK{}
			if f4iter.CarrierGatewayId != nil {
				f4elem.CarrierGatewayID = f4iter.CarrierGatewayId
			}
			if f4iter.DestinationCidrBlock != nil {
				f4elem.DestinationCIDRBlock = f4iter.DestinationCidrBlock
			}
			if f4iter.DestinationIpv6CidrBlock != nil {
				f4elem.DestinationIPv6CIDRBlock = f4iter.DestinationIpv6CidrBlock
			}
			if f4iter.DestinationPrefixListId != nil {
				f4elem.DestinationPrefixListID = f4iter.DestinationPrefixListId
			}
			if f4iter.EgressOnlyInternetGatewayId != nil {
				f4elem.EgressOnlyInternetGatewayID = f4iter.EgressOnlyInternetGatewayId
			}
			if f4iter.GatewayId != nil {
				f4elem.GatewayID = f4iter.GatewayId
			}
			if f4iter.InstanceId != nil {
				f4elem.InstanceID = f4iter.InstanceId
			}
			if f4iter.InstanceOwnerId != nil {
				f4elem.InstanceOwnerID = f4iter.InstanceOwnerId
			}
			if f4iter.LocalGatewayId != nil {
				f4elem.LocalGatewayID = f4iter.LocalGatewayId
			}
			if f4iter.NatGatewayId != nil {
				f4elem.NatGatewayID = f4iter.NatGatewayId
			}
			if f4iter.NetworkInterfaceId != nil {
				f4elem.NetworkInterfaceID = f4iter.NetworkInterfaceId
			}
			if f4iter.Origin != nil {
				f4elem.Origin = f4iter.Origin
			}
			if f4iter.State != nil {
				f4elem.State = f4iter.State
			}
			if f4iter.TransitGatewayId != nil {
				f4elem.TransitGatewayID = f4iter.TransitGatewayId
			}
			if f4iter.VpcPeeringConnectionId != nil {
				f4elem.VPCPeeringConnectionID = f4iter.VpcPeeringConnectionId
			}
			f4 = append(f4, f4elem)
		}
		ko.Status.Routes = f4
	}
	if resp.RouteTable.Tags != nil {
		f5 := []*svcapitypes.Tag{}
		for _, f5iter := range resp.RouteTable.Tags {
			f5elem := &svcapitypes.Tag{}
			if f5iter.Key != nil {
				f5elem.Key = f5iter.Key
			}
			if f5iter.Value != nil {
				f5elem.Value = f5iter.Value
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.Tags = f5
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateRouteTableInput, error) {
	res := &svcsdk.CreateRouteTableInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f1 := []*svcsdk.TagSpecification{}
		for _, f1iter := range r.ko.Spec.TagSpecifications {
			f1elem := &svcsdk.TagSpecification{}
			if f1iter.ResourceType != nil {
				f1elem.SetResourceType(*f1iter.ResourceType)
			}
			if f1iter.Tags != nil {
				f1elemf1 := []*svcsdk.Tag{}
				for _, f1elemf1iter := range f1iter.Tags {
					f1elemf1elem := &svcsdk.Tag{}
					if f1elemf1iter.Key != nil {
						f1elemf1elem.SetKey(*f1elemf1iter.Key)
					}
					if f1elemf1iter.Value != nil {
						f1elemf1elem.SetValue(*f1elemf1iter.Value)
					}
					f1elemf1 = append(f1elemf1, f1elemf1elem)
				}
				f1elem.SetTags(f1elemf1)
			}
			f1 = append(f1, f1elem)
		}
		res.SetTagSpecifications(f1)
	}
	if r.ko.Spec.VPCID != nil {
		res.SetVpcId(*r.ko.Spec.VPCID)
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
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
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
	_, respErr := rm.sdkapi.DeleteRouteTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteRouteTable", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRouteTableInput, error) {
	res := &svcsdk.DeleteRouteTableInput{}

	if r.ko.Spec.DryRun != nil {
		res.SetDryRun(*r.ko.Spec.DryRun)
	}
	if r.ko.Status.RouteTableID != nil {
		res.SetRouteTableId(*r.ko.Status.RouteTableID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.RouteTable,
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