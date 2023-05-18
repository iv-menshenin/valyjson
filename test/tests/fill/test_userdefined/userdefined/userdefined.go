package userdefined

import "time"

type (
	// DefinedFieldAsUserDefined tests inherited user-defined nested field
	//
	//json:json
	DefinedFieldAsUserDefined1 struct {
		Status DefinedFieldAsUserDefinedStatus `json:"status"`
	}
	DefinedFieldAsUserDefinedStatus string
	//json:json
	DefinedFieldAsUserDefined2 struct {
		Time time.Time `json:"time"`
	}
)

const (
	DefinedFieldAsUserDefinedStatusActive  DefinedFieldAsUserDefinedStatus = "active"
	DefinedFieldAsUserDefinedStatusPassive DefinedFieldAsUserDefinedStatus = "passive"
)
