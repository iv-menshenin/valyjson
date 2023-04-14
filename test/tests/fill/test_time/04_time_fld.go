package test_time

import "time"

// TestTime01 tests time.Time fields
//
//json:strict
type TestTime01 struct {
	DateBegin  time.Time  `json:"date_begin"`
	DateCustom time.Time  `json:"date_custom" layout:"2006.01.02"`
	DateEnd    *time.Time `json:"date_end,omitempty"`
}

// TestTime2 test used-defined
//
//json:json
type TestTime2 time.Time
