// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTagDetails Properties used in tag create operations.
type CreateTagDetails struct {

	// The name of the tag in the case of a free form tag.
	// When linking to a glossary term, this field is not specified.
	Name *string `mandatory:"false" json:"name"`

	// Unique key of the related term or null in the case of a free form tag.
	TermKey *string `mandatory:"false" json:"termKey"`
}

func (m CreateTagDetails) String() string {
	return common.PointerString(m)
}
