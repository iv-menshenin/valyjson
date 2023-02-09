package test_uuid

import "github.com/google/uuid"

// TestUUID tests uuid.UUID
//json:optional
type TestUUID struct {
	UUID uuid.UUID `json:"uuid"`
}
