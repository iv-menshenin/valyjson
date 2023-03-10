package njson

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPerson_MarshalJSON(t *testing.T) {
	person := fulfilledPerson()
	data, err := json.Marshal(person)
	require.NoError(t, err)
	require.JSONEq(t, fulfilledPersonData, string(data))
}

func TestPerson_UnmarshalJSON(t *testing.T) {
	var person Person
	require.NoError(t, json.Unmarshal([]byte(fulfilledPersonData), &person))
	require.Equal(t, fulfilledPerson(), &person)
}

// BenchmarkMarshal-8     	  423813	      2741 ns/op	     872 B/op	       9 allocs/op
func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	person := fulfilledPerson()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(person)
	}
}

// BenchmarkUnmarshal-8   	  133050	      8904 ns/op	    2528 B/op	      40 allocs/op
func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var person Person
		_ = json.Unmarshal([]byte(fulfilledPersonData), &person)
	}
}

func fulfilledPerson() *Person {
	var (
		middle = "Vladimirovich"
		dob    = time.Date(1984, time.February, 14, 0, 0, 0, 0, time.UTC)
	)
	return &Person{
		OriginName: OriginName{
			OriginNameName: OriginNameName{
				Name: "Igor",
			},
			OriginNameSurname: OriginNameSurname{
				Surname: "Menshenin",
			},
		},
		Middle: &middle,
		DOB:    &dob,
		Passport: &Passport{
			Number:  "123-456 789000",
			DateDoc: time.Date(2010, 10, 10, 10, 10, 10, 10, time.UTC),
		},
		Tables: map[string]TableOf{
			"math": {
				TableName: "Mathematics",
				Tables: []*Table{
					{
						Counter:     3,
						Assessments: []int{3, 3, 4},
						Time:        time.Date(2023, 1, 9, 1, 1, 1, 1, time.UTC),
						Avg:         3.3333,
						Tags: []Tag{
							{
								TagName:  "test-name",
								TagValue: "Person-Struct",
							},
							{
								TagName:  "author",
								TagValue: "valyjson",
							},
						},
					},
				},
			},
		},
	}
}

const fulfilledPersonData = `
{
  "name": "Igor",
  "surname": "Menshenin",
  "middle": "Vladimirovich",
  "dob": "1984-02-14T00:00:00Z",
  "passport": {
    "number": "123-456 789000",
    "dateDoc": "2010-10-10T10:10:10.00000001Z"
  },
  "tables": {
    "math": {
      "tableName": "Mathematics",
      "tables": [
        {
          "counter": 3,
          "assessments": [
            3,
            3,
            4
          ],
          "time": "2023-01-09T01:01:01.000000001Z",
          "avg": 3.3333,
          "tags": [
            {
              "tagName": "test-name",
              "tagValue": "Person-Struct"
            },
            {
              "tagName": "author",
              "tagValue": "valyjson"
            }
          ]
        }
      ]
    }
  }
}`
