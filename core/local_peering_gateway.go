// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LocalPeeringGateway A local peering gateway (LPG) is an object on a VCN that lets that VCN peer
// with another VCN in the same region. *Peering* means that the two VCNs can
// communicate using private IP addresses, but without the traffic traversing the
// internet or routing through your on-premises network. For more information,
// see [VCN Peering]({{DOC_SERVER_URL}}/Content/Network/Tasks/VCNpeering.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type LocalPeeringGateway struct {

	// The OCID of the compartment containing the LPG.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The LPG's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// Whether the VCN at the other end of the peering is in a different tenancy.
	// Example: `false`
	IsCrossTenancyPeering *bool `mandatory:"true" json:"isCrossTenancyPeering,omitempty"`

	// The LPG's current lifecycle state.
	LifecycleState LocalPeeringGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been
	// peered. `PENDING` means the peering is being established. `REVOKED` means the
	// LPG at the other end of the peering has been deleted.
	PeeringStatus LocalPeeringGatewayPeeringStatusEnum `mandatory:"true" json:"peeringStatus,omitempty"`

	// The date and time the LPG was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VCN the LPG belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// The range of IP addresses available on the VCN at the other
	// end of the peering from this LPG. The value is `null` if the LPG is not peered.
	// You can use this as the destination CIDR for a route rule to route a subnet's
	// traffic to this LPG.
	// Example: `192.168.0.0/16`
	PeerAdvertisedCidr *string `mandatory:"false" json:"peerAdvertisedCidr,omitempty"`

	// Additional information regarding the peering status, if applicable.
	PeeringStatusDetails *string `mandatory:"false" json:"peeringStatusDetails,omitempty"`
}

func (model LocalPeeringGateway) String() string {
	return common.PointerString(model)
}

// LocalPeeringGatewayLifecycleStateEnum Enum with underlying type: string
type LocalPeeringGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayLifecycleState
const (
	LocalPeeringGatewayLifecycleStateProvisioning LocalPeeringGatewayLifecycleStateEnum = "PROVISIONING"
	LocalPeeringGatewayLifecycleStateAvailable    LocalPeeringGatewayLifecycleStateEnum = "AVAILABLE"
	LocalPeeringGatewayLifecycleStateTerminating  LocalPeeringGatewayLifecycleStateEnum = "TERMINATING"
	LocalPeeringGatewayLifecycleStateTerminated   LocalPeeringGatewayLifecycleStateEnum = "TERMINATED"
	LocalPeeringGatewayLifecycleStateUnknown      LocalPeeringGatewayLifecycleStateEnum = "UNKNOWN"
)

var mappingLocalPeeringGatewayLifecycleState = map[string]LocalPeeringGatewayLifecycleStateEnum{
	"PROVISIONING": LocalPeeringGatewayLifecycleStateProvisioning,
	"AVAILABLE":    LocalPeeringGatewayLifecycleStateAvailable,
	"TERMINATING":  LocalPeeringGatewayLifecycleStateTerminating,
	"TERMINATED":   LocalPeeringGatewayLifecycleStateTerminated,
	"UNKNOWN":      LocalPeeringGatewayLifecycleStateUnknown,
}

// GetLocalPeeringGatewayLifecycleStateEnumValues Enumerates the set of values for LocalPeeringGatewayLifecycleState
func GetLocalPeeringGatewayLifecycleStateEnumValues() []LocalPeeringGatewayLifecycleStateEnum {
	values := make([]LocalPeeringGatewayLifecycleStateEnum, 0)
	for _, v := range mappingLocalPeeringGatewayLifecycleState {
		if v != LocalPeeringGatewayLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}

// LocalPeeringGatewayPeeringStatusEnum Enum with underlying type: string
type LocalPeeringGatewayPeeringStatusEnum string

// Set of constants representing the allowable values for LocalPeeringGatewayPeeringStatus
const (
	LocalPeeringGatewayPeeringStatusInvalid LocalPeeringGatewayPeeringStatusEnum = "INVALID"
	LocalPeeringGatewayPeeringStatusNew     LocalPeeringGatewayPeeringStatusEnum = "NEW"
	LocalPeeringGatewayPeeringStatusPeered  LocalPeeringGatewayPeeringStatusEnum = "PEERED"
	LocalPeeringGatewayPeeringStatusPending LocalPeeringGatewayPeeringStatusEnum = "PENDING"
	LocalPeeringGatewayPeeringStatusRevoked LocalPeeringGatewayPeeringStatusEnum = "REVOKED"
	LocalPeeringGatewayPeeringStatusUnknown LocalPeeringGatewayPeeringStatusEnum = "UNKNOWN"
)

var mappingLocalPeeringGatewayPeeringStatus = map[string]LocalPeeringGatewayPeeringStatusEnum{
	"INVALID": LocalPeeringGatewayPeeringStatusInvalid,
	"NEW":     LocalPeeringGatewayPeeringStatusNew,
	"PEERED":  LocalPeeringGatewayPeeringStatusPeered,
	"PENDING": LocalPeeringGatewayPeeringStatusPending,
	"REVOKED": LocalPeeringGatewayPeeringStatusRevoked,
	"UNKNOWN": LocalPeeringGatewayPeeringStatusUnknown,
}

// GetLocalPeeringGatewayPeeringStatusEnumValues Enumerates the set of values for LocalPeeringGatewayPeeringStatus
func GetLocalPeeringGatewayPeeringStatusEnumValues() []LocalPeeringGatewayPeeringStatusEnum {
	values := make([]LocalPeeringGatewayPeeringStatusEnum, 0)
	for _, v := range mappingLocalPeeringGatewayPeeringStatus {
		if v != LocalPeeringGatewayPeeringStatusUnknown {
			values = append(values, v)
		}
	}
	return values
}