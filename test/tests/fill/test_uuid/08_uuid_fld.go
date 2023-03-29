package test_uuid

import "github.com/google/uuid"

// TestUUID tests uuid.UUID
//json:json
type TestUUID struct {
	UUID uuid.UUID `json:"uuid"`
}

// InheritUUID2 tests uuid.UUID inheritance
//json:json
type InheritUUID2 InheritUUID

// InheritUUID tests uuid.UUID inheritance
//json:json
type InheritUUID uuid.UUID
