package test_nested

import "time"

type (
	//json:json
	Root struct {
		Meta Meta    `json:"meta"`
		Data Middles `json:"data"`
	}
	//json:json
	Middles []Middle
	//json:json
	Meta struct {
		Count int `json:"count"`
	}
	//json:json
	Middle struct {
		Personal   `json:",inline"`
		DateOfBorn time.Time `json:"dateOfBorn"`
		Tags       Tags      `json:"tags"`
	}
	//json:json
	Personal struct {
		Name    UserName     `json:"name"`
		Surname UserSurname  `json:"surname"`
		Patname *UserPatname `json:"patname"`
	}
	//json:json
	Tags        map[TagName]TagValue
	TagName     string
	TagValue    string
	UserName    string
	UserSurname string
	UserPatname string

	//json:marshal
	CustomEvent struct {
		WRRetry `json:",inline"`
	}
	// Field with the same name
	//json:marshal
	WRRetry struct {
		WRRetry int `json:"WR-Retry,omitempty"`
	}
)
