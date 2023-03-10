package test_uuid

import "github.com/google/uuid"

// TestUUID tests uuid.UUID
//json:optional,decode
type TestUUID struct {
	UUID uuid.UUID `json:"uuid"`
}
