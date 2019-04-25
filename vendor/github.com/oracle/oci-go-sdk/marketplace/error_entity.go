// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ErrorEntity The model for the error entity.
type ErrorEntity struct {

	// A short error code that defines the error.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m ErrorEntity) String() string {
	return common.PointerString(m)
}
