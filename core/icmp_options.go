// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// IcmpOptions. Optional object to specify a particular ICMP type and code. If you specify ICMP as the protocol
// but do not provide this object, then all ICMP types and codes are allowed. If you do provide
// this object, the type is required and the code is optional.
// See [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
// for allowed values. To enable MTU negotiation for ingress internet traffic, make sure to allow
// type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set").
// If you need to specify multiple codes for a single type, create a separate security list rule for each.
type IcmpOptions struct {

	// The ICMP type.
	Type_ *int `mandatory:"true" json:"type,omitempty"`

	// The ICMP code (optional).
	Code *int `mandatory:"false" json:"code,omitempty"`
}

func (model IcmpOptions) String() string {
	return common.PointerString(model)
}