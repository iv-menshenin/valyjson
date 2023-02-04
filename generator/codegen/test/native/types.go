package njson

import (
	"time"
)

type (
	//json:strict
	Person struct {
		OriginName `json:",inline"`
		Middle     *string            `json:"middle,omitempty"`
		DOB        *time.Time         `json:"dob,omitempty"`
		Passport   *Passport          `json:"passport"`
		Tables     map[string]TableOf `json:"tables"`
	}
	OriginName struct {
		OriginNameName    `json:",inline"`
		OriginNameSurname `json:",inline"`
	}
	OriginNameName struct {
		Name string `json:"name"`
	}
	OriginNameSurname struct {
		Surname string `json:"surname"`
	}
	//json:strict
	Passport struct {
		Number  string    `json:"number"`
		DateDoc time.Time `json:"dateDoc"`
	}
	//json:strict
	TableOf struct {
		TableName string   `json:"tableName"`
		Tables    []*Table `json:"tables,omitempty"`
	}
	//json:strict
	Table struct {
		Counter     int       `json:"counter"`
		Assessments []int     `json:"assessments,omitempty"`
		Time        time.Time `json:"time"`
		Avg         float64   `json:"avg"`
		Tags        []Tag     `json:"tags"`
	}
	//json:strict
	Tag struct {
		TagName  string `json:"tagName"`
		TagValue string `json:"tagValue"`
	}
)
