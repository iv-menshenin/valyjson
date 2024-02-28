package vjson

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMapTable_UnmarshalJSON(t *testing.T) {
	var table MapTable
	require.NoError(t, table.UnmarshalJSON([]byte(`{
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
  }`)))
	require.Equal(t, fulfilledPerson().Tables, table)
}

func TestPerson_MarshalJSON(t *testing.T) {
	person := fulfilledPerson()
	data, err := json.Marshal(person)
	require.NoError(t, err)
	require.JSONEq(t, fulfilledPersonData, string(data))
}

func TestPerson_UnmarshalJSON(t *testing.T) {
	var person Person
	require.NoError(t, person.UnmarshalJSON([]byte(fulfilledPersonData)))
	require.Equal(t, fulfilledPerson(), &person)
}

// BenchmarkMarshal-8     	  425655	      2666 ns/op	     872 B/op	       9 allocs/op
// BenchmarkMarshal-8   	  676693	      1866 ns/op	     768 B/op	       2 allocs/op
// EASY JSON Benchmark     	  685834	      1847 ns/op	    1192 B/op	       8 allocs/op
// BenchmarkMarshal-8     	 1000000	      5538 ns/op	    1256 B/op	       9 allocs/op
func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	person := fulfilledPerson()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = person.MarshalJSON()
	}
}

// BenchmarkUnmarshal-8   	  383025	      3094 ns/op	    1056 B/op	      25 allocs/op
// BenchmarkUnmarshal-8   	  400539	      3015 ns/op	     988 B/op	      18 allocs/op
// EASY JSON Benchmark   	  403358	      3021 ns/op	    1872 B/op	      23 allocs/op
// BenchmarkUnmarshal-8   	  493149	     12653 ns/op	    1000 B/op	      21 allocs/op
func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var person Person
		_ = person.UnmarshalJSON([]byte(fulfilledPersonData))
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
