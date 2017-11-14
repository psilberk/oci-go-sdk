// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateBackupDetails struct {

	// The OCID of the database.
	DatabaseID *string `mandatory:"true" json:"databaseId,omitempty"`

	// The user-friendly name for the backup. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`
}

func (model CreateBackupDetails) String() string {
	return common.PointerString(model)
}