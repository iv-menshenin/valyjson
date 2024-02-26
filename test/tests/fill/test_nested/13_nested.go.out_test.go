package test_nested

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("path_to_error_array", func(t *testing.T) {
		t.Parallel()
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_error_inh_array", func(t *testing.T) {
		t.Parallel()
		const data = `[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]`
		var v Middles
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_map", func(t *testing.T) {
		t.Parallel()
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{"surname":"Petrov","tags":{"foo":"bar","count": 7}}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.1.tags.count")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_inlined_map", func(t *testing.T) {
		t.Parallel()
		const data = `{"test":"test","foo":"bar","conf":null}`
		var v Tags
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "conf")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
}

func Test_InlinedNestedStructures(t *testing.T) {
	t.Parallel()
	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		var ce = CustomEvent{
			WRRetry{
				WRRetry: 3,
			},
		}
		data, err := ce.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"WR-Retry": 3}`, string(data))
	})
	t.Run("marshal-the-same-as-std", func(t *testing.T) {
		t.Parallel()
		var ce = CustomEvent{WRRetry{WRRetry: 312}}
		data1, err := ce.MarshalJSON()
		require.NoError(t, err)

		data2, err := json.Marshal(ce)
		require.NoError(t, err)

		require.JSONEq(t, string(data1), string(data2))
	})
	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		const jsonData = `{"WR-Retry": 304}`
		var expected = CustomEvent{
			WRRetry{
				WRRetry: 304,
			},
		}
		var got CustomEvent
		require.NoError(t, got.UnmarshalJSON([]byte(jsonData)))
		require.Equal(t, expected, got)
	})
	t.Run("unmarshal-the-same-as-std", func(t *testing.T) {
		t.Parallel()
		const jsonData = `{"WR-Retry": 5666}`
		var expected = CustomEvent{WRRetry{WRRetry: 5666}}
		var got1 CustomEvent

		require.NoError(t, got1.UnmarshalJSON([]byte(jsonData)))
		require.Equal(t, expected, got1)

		var got2 CustomEvent
		require.NoError(t, json.Unmarshal([]byte(jsonData), &got2))
		require.Equal(t, expected, got2)
	})
}

func TestReset(t *testing.T) {
	t.Parallel()
	t.Run("Root", func(t *testing.T) {
		t.Parallel()
		var (
			p1, p2 UserPatname = "Fedorovich", "Petrovich"
		)
		var val = Root{
			Meta: Meta{Count: 3},
			Data: Middles{
				{
					Personal: Personal{
						Name:    "John",
						Surname: "Doe",
						Patname: nil,
					},
					DateOfBorn: time.Date(2024, 1, 1, 12, 54, 31, 0, time.UTC),
					Tags: map[TagName]TagValue{
						"AgeClass": "young",
						"Place":    "America",
					},
				},
				{
					Personal: Personal{
						Name:    "Ivan",
						Surname: "Ivanov",
						Patname: &p1,
					},
					DateOfBorn: time.Date(2024, 1, 1, 12, 54, 33, 0, time.UTC),
					Tags:       nil,
				},
				{
					Personal: Personal{
						Name:    "Mike",
						Surname: "Potapov",
						Patname: &p2,
					},
					DateOfBorn: time.Date(2024, 1, 1, 12, 54, 32, 0, time.UTC),
					Tags: map[TagName]TagValue{
						"AgeClass": "old",
						"Place":    "Russia",
					},
				},
			},
		}
		val.Reset()
		require.Empty(t, val.Data)
		require.Zero(t, val.Meta.Count)
	})
}

var unmarshalReuseRaceTestPool = sync.Pool{New: func() any { return &Root{} }}

func Test_Unmarshal_Reuse_Race(t *testing.T) {
	t.Parallel()
	const data = `{"meta":{"count": %d}, "data":[%s]}`
	var names = []string{
		"",
		`Igor`,
		`Natasha`,
	}
	var surnames = []string{
		"",
		`Po`,
		`Popov`,
		`Sidorov`,
	}
	var patnames = []string{
		"",
		`Ogly`,
		`Popovich`,
		`Petrovich_0`,
		`Gorgorothovich`,
	}

	var wg sync.WaitGroup
	for n := 0; n < 10000; n++ {
		wg.Add(1)
		go func(name, surname, patname string, i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			defer wg.Done()
			var xname, xsurname, xpatname string
			if name != "" {
				xname = fmt.Sprintf(`,"name":"%s"`, name)
			}
			if surname != "" {
				xsurname = fmt.Sprintf(`,"surname":"%s"`, surname)
			}
			if patname != "" {
				xpatname = fmt.Sprintf(`,"patname":"%s"`, patname)
			}
			var person []string
			for d := i; d > 0; d-- {
				person = append(person, fmt.Sprintf(`{"tags": {"n":"%d"%s%s%s}%s%s%s}`, d, xname, xsurname, xpatname, xname, xsurname, xpatname))
			}
			var dataJson = fmt.Sprintf(data, i, strings.Join(person, ","))
			var s = unmarshalReuseRaceTestPool.Get().(*Root)
			require.NoError(t, s.UnmarshalJSON([]byte(dataJson)))
			require.Len(t, s.Data, i)
			require.Equal(t, s.Meta.Count, i)

			for d := 0; d < i; d++ {
				require.EqualValues(t, name, s.Data[d].Name)
				require.EqualValues(t, surname, s.Data[d].Surname)
				if patname == "" {
					require.Nil(t, s.Data[d].Patname)
				} else {
					require.NotNil(t, s.Data[d].Patname)
					require.EqualValues(t, patname, *s.Data[d].Patname)
				}
				require.NotEmpty(t, s.Data[d].Tags["n"])
				require.EqualValues(t, name, s.Data[d].Tags["name"])
				require.EqualValues(t, surname, s.Data[d].Tags["surname"])
				require.EqualValues(t, patname, s.Data[d].Tags["patname"])
			}

			s.Reset()
			unmarshalReuseRaceTestPool.Put(s)
		}(names[n%len(names)], surnames[n%len(surnames)], patnames[n%len(patnames)], n%6)
	}
	wg.Wait()
}

var unmarshalReuseAllocTestPool = sync.Pool{New: func() any { return &Root{} }}

func Benchmark_Unmarshal_Reuse_Alloc(b *testing.B) {
	b.ReportAllocs()
	const data = `{"meta":{"count": %d}, "data":[%s]}`
	var names = []string{
		"",
		`Igor`,
		`Natasha`,
	}
	var surnames = []string{
		"",
		`Po`,
		`Popov`,
		`Sidorov`,
	}
	var patnames = []string{
		"",
		`Ogly`,
		`Popovich`,
		`Petrovich_0`,
		`Gorgorothovich`,
	}
	var dataJson []string
	for n := 0; n < len(names)*len(surnames)*len(patnames); n++ {
		var name, surname, patname = names[n%len(names)], surnames[n%len(surnames)], patnames[n%len(patnames)]
		var xname, xsurname, xpatname string
		var i = n % 6
		if name != "" {
			xname = fmt.Sprintf(`,"name":"%s"`, name)
		}
		if surname != "" {
			xsurname = fmt.Sprintf(`,"surname":"%s"`, surname)
		}
		if patname != "" {
			xpatname = fmt.Sprintf(`,"patname":"%s"`, patname)
		}
		var person []string
		for d := i; d > 0; d-- {
			person = append(person, fmt.Sprintf(`{"tags": {"n":"%d"%s%s%s}%s%s%s}`, d, xname, xsurname, xpatname, xname, xsurname, xpatname))
		}
		dataJson = append(dataJson, fmt.Sprintf(data, i, strings.Join(person, ",")))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s = unmarshalReuseAllocTestPool.Get().(*Root)
		if err := s.UnmarshalJSON([]byte(dataJson[n%len(dataJson)])); err != nil {
			b.Error(err)
		}
		s.Reset()
		unmarshalReuseAllocTestPool.Put(s)
	}
}

func Test_Unmarshal_Reuse_Alloc(t *testing.T) {
	const data = `{"meta":{"count": %d}, "data":[%s]}`
	var names = []string{
		"",
		`Igor`,
		`Natasha`,
	}
	var surnames = []string{
		"",
		`Po`,
		`Popov`,
		`Sidorov`,
	}
	var patnames = []string{
		"",
		`Ogly`,
		`Popovich`,
		`Petrovich_0`,
		`Gorgorothovich`,
	}
	var dataJson []string
	for n := 0; n < len(names)*len(surnames)*len(patnames); n++ {
		var name, surname, patname = names[n%len(names)], surnames[n%len(surnames)], patnames[n%len(patnames)]
		var xname, xsurname, xpatname string
		var i = n % 6
		if name != "" {
			xname = fmt.Sprintf(`,"name":"%s"`, name)
		}
		if surname != "" {
			xsurname = fmt.Sprintf(`,"surname":"%s"`, surname)
		}
		if patname != "" {
			xpatname = fmt.Sprintf(`,"patname":"%s"`, patname)
		}
		var person []string
		for d := i; d > 0; d-- {
			person = append(person, fmt.Sprintf(`{"tags": {"n":"%d"%s%s%s}%s%s%s}`, d, xname, xsurname, xpatname, xname, xsurname, xpatname))
		}
		dataJson = append(dataJson, fmt.Sprintf(data, i, strings.Join(person, ",")))
	}

	var i int64
	allocs := testing.AllocsPerRun(10000, func() {
		var s = unmarshalReuseAllocTestPool.Get().(*Root)
		n := int(atomic.AddInt64(&i, 1))
		if err := s.UnmarshalJSON([]byte(dataJson[n%len(dataJson)])); err != nil {
			t.Error(err)
		}
		s.Reset()
		unmarshalReuseAllocTestPool.Put(s)
	})
	require.LessOrEqual(t, allocs, float64(20))
}
