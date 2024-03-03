package test_slice

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestSlice01_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("test-all-empty", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[], "ints":null}`))
		require.NoError(t, err)
		require.NotNil(t, test1.Field)
		require.Len(t, test1.Field, 0)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("test-all-nulled", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":null}`))
		require.NoError(t, err)
		require.Nil(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("filled-slice", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":["test1", "test2"], "ints":null}`))
		require.NoError(t, err)
		require.Len(t, test1.Field, 2)
		require.ElementsMatch(t, []string{"test1", "test2"}, test1.Field)
	})
	t.Run("filled-slice-ref", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 3)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
	t.Run("filled-slice-ref-with-null", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, null, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 4)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
			if value == 3 {
				expected = append(expected, nil)
			}
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
}

func Test_TestSlice_Unmarshal_Bad(t *testing.T) {
	t.Parallel()
	t.Run("empty-json", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON(nil))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON(nil))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON(nil))
	})
	t.Run("bad-json", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON([]byte{0, 1, 2, 3}))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON([]byte{0, 1, 2, 3}))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON([]byte{0, 1, 2, 3}))
	})
	t.Run("invalid-object", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON([]byte(`{"strs":null,"strs":null}`)))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON([]byte(`{}`)))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON([]byte(`[]`)))
	})
	t.Run("invalid-object-field-strs", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":{},"ints":null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain array")
	})
	t.Run("invalid-object-vals-strs", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[{},{}],"ints":null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("invalid-object-field-ints", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[],"ints":2}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain array")
	})
	t.Run("invalid-object-vals-ints", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[],"ints":[{},{}]}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain number")
	})
}

func Test_TestSlice01_Marshal(t *testing.T) {
	t.Parallel()
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		var s TestSlice01
		b, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":null,"ints":null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("comma_after_filled", func(t *testing.T) {
		t.Parallel()
		var s = TestSlice01{Field: []string{"1"}}
		b, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1"],"ints":null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("filled_both", func(t *testing.T) {
		t.Parallel()
		var a, b, c = 3, 2, 1
		var s = TestSlice01{
			Field:    []string{"1"},
			FieldRef: []*int{&a, &b, &c},
		}
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1"],"ints":[3,2,1]}`
		require.JSONEq(t, expected, string(data))
	})
	t.Run("with_null_elem", func(t *testing.T) {
		t.Parallel()
		var a, b, c = 3, 2, 1
		var s = TestSlice01{
			Field:    []string{"1", ""},
			FieldRef: []*int{&a, nil, &b, &c},
		}
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1",""],"ints":[3,null,2,1]}`
		require.JSONEq(t, expected, string(data))
	})
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var s *TestSlice01
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `null`
		require.JSONEq(t, expected, string(data))
	})
}

func TestTestSlice01_validate(t *testing.T) {
	t.Parallel()
	t.Run("ints_appears_twice", func(t *testing.T) {
		t.Parallel()
		const jData = `{"ints":[],"ints":[1,2]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "ints")
	})
	t.Run("not_object", func(t *testing.T) {
		t.Parallel()
		const jData = `122333`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "value doesn't contain object")
		require.ErrorContains(t, err, "it contains number")
	})
	t.Run("strs_appears_twice", func(t *testing.T) {
		t.Parallel()
		const jData = `{"strs":[],"strs":["1","2"]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "strs")
	})
	t.Run("data_appears_twice", func(t *testing.T) {
		t.Parallel()
		const jData = `{"data":{},"data":{},"extended":{}}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice03
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "data")
	})
	t.Run("both_appears_twice", func(t *testing.T) {
		t.Parallel()
		const jData = `{"strs":[],"ints":[],"strs":["1","2"],"ints":[1,2]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
	})
}

func TestTestSlice_IsZero(t *testing.T) {
	t.Parallel()
	t.Run("zero", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice01
		require.True(t, test1.IsZero())
		var test2 TestSlice02
		require.True(t, test2.IsZero())
		var test3 TestSlice03
		require.True(t, test3.IsZero())
	})
	t.Run("nonzero", func(t *testing.T) {
		t.Parallel()
		var test1 = TestSlice01{Field: []string{}}
		require.False(t, test1.IsZero())
		var test1_2 = TestSlice01{FieldRef: []*int{}}
		require.False(t, test1_2.IsZero())
		var test2 = TestSlice02{TestSlice03{Data: 3}}
		require.False(t, test2.IsZero())
		var test3 = TestSlice03{Data: 44}
		require.False(t, test3.IsZero())
	})
}

func Test_TestSlice02_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`null`))
		require.NoError(t, err)
		require.Nil(t, test1)
		require.Len(t, test1, 0)
	})
	t.Run("empty-array", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 0)
	})
	t.Run("filled-array", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},{"data": 54}]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 2)
		require.Equal(t, TestSlice02{{Data: 13}, {Data: 54}}, test1)
	})
	t.Run("null-instead-of-object", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},null]`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("filled-array-2", func(t *testing.T) {
		t.Parallel()
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},{},{"data": 54}]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 3)
		require.Equal(t, TestSlice02{{Data: 13}, {}, {Data: 54}}, test1)
	})
}

func Test_TestSlice02_Marshal(t *testing.T) {
	t.Parallel()
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var test TestSlice02
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		var test = TestSlice02{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "[]", string(data))
	})
	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		var test = TestSlice02{{Data: 190}, {Data: 191}}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `[{"data":190},{"data":191}]`, string(data))
	})
}

func Test_TestSlice03_Marshal(t *testing.T) {
	t.Parallel()
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var test *TestSlice03
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		var test = TestSlice03{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"data":0}`, string(data))
	})
}

func TestSliceSlice_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("zero", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":null,"ints":null}`
		var test TestSliceSlice
		require.NoError(t, test.UnmarshalJSON([]byte(data)))
		require.Empty(t, test)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":[],"ints":[]}`
		var test TestSliceSlice
		require.NoError(t, test.UnmarshalJSON([]byte(data)))
		require.Empty(t, test.FieldInt)
		require.Empty(t, test.FieldStr)
	})
	t.Run("well", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":[["a","b","c"],["1","2","3"]],"ints":[[1,2,3,4],[0,0,0,1]]}`
		var test TestSliceSlice
		var expected = TestSliceSlice{
			FieldStr: [][]InnerString{
				{"a", "b", "c"},
				{"1", "2", "3"},
			},
			FieldInt: [][]int{
				{1, 2, 3, 4},
				{0, 0, 0, 1},
			},
		}
		require.NoError(t, test.UnmarshalJSON([]byte(data)))
		require.EqualValues(t, expected, test)
	})
	t.Run("nulls", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":[["a","b","c"],null,["1","2","3"]],"ints":[[1,2,3,4],[],[-1,0,0,1]]}`
		var test TestSliceSlice
		var expected = TestSliceSlice{
			FieldStr: [][]InnerString{
				{"a", "b", "c"},
				nil,
				{"1", "2", "3"},
			},
			FieldInt: [][]int{
				{1, 2, 3, 4},
				{},
				{-1, 0, 0, 1},
			},
		}
		require.NoError(t, test.UnmarshalJSON([]byte(data)))
		require.EqualValues(t, expected, test)
	})
	t.Run("error_strings", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":[["a","b","c"],["1",2,"3"]],"ints":[[1,2,3,4],[0,0,0,1]]}`
		var test TestSliceSlice
		err := test.UnmarshalJSON([]byte(data))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("error_integers", func(t *testing.T) {
		t.Parallel()
		const data = `{"strs":[["a","b","c"],["1","2","3"]],"ints":[[1,2,3,4],[0,"0",0,1]]}`
		var test TestSliceSlice
		err := test.UnmarshalJSON([]byte(data))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain number")
	})
}

func TestSliceSlice_Marshal(t *testing.T) {
	t.Parallel()
	t.Run("well", func(t *testing.T) {
		t.Parallel()
		const expected = `{"strs":[["bar","foo"],["test"]],"ints":[[123,321],[999]]}`
		var test = TestSliceSlice{
			FieldStr: [][]InnerString{
				{"bar", "foo"},
				{"test"},
			},
			FieldInt: [][]int{
				{123, 321},
				{999},
			},
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("nulls_empty_slices", func(t *testing.T) {
		t.Parallel()
		const expected = `{"strs":[["bar","foo"],[],["test"]],"ints":[[123,321],[],[999]]}`
		var test = TestSliceSlice{
			FieldStr: [][]InnerString{
				{"bar", "foo"},
				nil,
				{"test"},
			},
			FieldInt: [][]int{
				{123, 321},
				nil,
				{999},
			},
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("nulls", func(t *testing.T) {
		t.Parallel()
		const expected = `{"strs":null,"ints":null}`
		var test = TestSliceSlice{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		const expected = `{"strs":[],"ints":[]}`
		var test = TestSliceSlice{
			FieldStr: [][]InnerString{},
			FieldInt: [][]int{},
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("top_null", func(t *testing.T) {
		t.Parallel()
		const expected = `null`
		var test *TestSliceSlice
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func TestSliceSlice_IsZero(t *testing.T) {
	t.Parallel()
	t.Run("zero", func(t *testing.T) {
		var test TestSliceSlice
		require.True(t, test.IsZero())
	})
	t.Run("nonzero_ints", func(t *testing.T) {
		var test = TestSliceSlice{FieldInt: [][]int{{1}}}
		require.False(t, test.IsZero())
	})
	t.Run("nonzero_strs", func(t *testing.T) {
		var test = TestSliceSlice{FieldStr: [][]InnerString{{"a"}}}
		require.False(t, test.IsZero())
	})
}

func TestSliceSlice_Unmarshal_Reuse(t *testing.T) {
	t.Parallel()
	const testsCount = 100000
	var (
		ints = []int{0, 1000, 1001, 1002, 1123, 1140, 5300, 6712, 33432, 546466, 746564666}
		strs = []InnerString{``, `null`, `[]`, `[[abc]]`, `[[foo],[bar,test]]`, `[[foo_test,bar_test],[],[a,b,c]]`}
	)
	var pool = sync.Pool{New: func() any { return &TestSliceSlice{} }}
	var wg sync.WaitGroup
	for n := 0; n < testsCount; n++ {
		wg.Add(1)
		var i = make([]int, n%9)
		for n1 := range i {
			i[n1] = ints[(n+n1)%len(ints)]
		}
		var s = make([]InnerString, n%10)
		for n1 := range s {
			s[n1] = strs[(n+n1)%len(strs)]
		}
		go func(i []int, s []InnerString, n int) {
			defer wg.Done()
			var (
				ints [][]int
				strs [][]InnerString
			)
			jsonData := `{"strs":`
			if len(s) == 0 {
				jsonData += `null,`
			} else {
				jsonData += `[[`
				var x1 = (n % 3) % (len(s) + 1)
				v := make([]InnerString, 0, x1)
				for y1, x := range s[:x1] {
					v = append(v, x)
					jsonData += `"` + string(x) + `"`
					if x1-1 > y1 {
						jsonData += `,`
					}
				}
				jsonData += `],[`
				strs = append(strs, v)
				v = make([]InnerString, 0)
				for y1, x := range s[x1:] {
					v = append(v, x)
					jsonData += `"` + string(x) + `"`
					if len(s)-1 > y1+x1 {
						jsonData += `,`
					}
				}
				jsonData += `]],`
				strs = append(strs, v)
			}
			ii := i
			jsonData += `"ints":`
			if len(ii) == 0 {
				jsonData += `null`
			} else {
				jsonData += `[[`
				var ttx int
				for {
					var x2 = ((n + ttx) % 5) % (len(ii) + 1)
					ttx++
					v := make([]int, 0)
					for y2, x := range ii[:x2] {
						v = append(v, x)
						jsonData += strconv.Itoa(x)
						if x2-1 > y2 {
							jsonData += `,`
						}
					}
					ints = append(ints, v)
					if ii = ii[x2:]; len(ii) == 0 {
						break
					}
					jsonData += `],[`
				}
				jsonData += `]]`
			}
			jsonData += `}`
			var test = pool.Get().(*TestSliceSlice)
			require.NoError(t, test.UnmarshalJSON([]byte(jsonData)))
			var expected = TestSliceSlice{
				FieldStr: strs,
				FieldInt: ints,
			}
			if len(test.FieldInt) == 0 {
				test.FieldInt = nil
			}
			if len(test.FieldStr) == 0 {
				test.FieldStr = nil
			}
			require.Equal(t, expected, *test)
			test.Reset()
			pool.Put(test)
		}(i, s, n+1)
	}
	wg.Wait()
}

func TestCampaignSitesUnmarshal(t *testing.T) {
	t.Parallel()
	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()
		var cs CampaignSites
		const data = `{"excluded":["a","b","c"],"included":["1","2","3","4","5"]}`
		require.NoError(t, cs.UnmarshalJSON([]byte(data)))
		var expected = CampaignSites{
			Excluded: []FieldValueString{"a", "b", "c"},
			Included: [5]FieldValueString{"1", "2", "3", "4", "5"},
		}
		require.Equal(t, expected, cs)
	})
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		var cs = CampaignSites{
			Excluded: []FieldValueString{"foo", "bar"},
			Included: [5]FieldValueString{"a", "b", "c", "d", "e"},
		}
		const expected = `{"excluded":["foo","bar"],"included":["a","b","c","d","e"]}`
		data, err := cs.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("Reset", func(t *testing.T) {
		t.Parallel()
		var fvs = []FieldValueString{"foo", "bar"}
		var cs = CampaignSites{
			Excluded: fvs,
			Included: [5]FieldValueString{"", "", "4", "", ""},
		}
		cs.Reset()
		require.Len(t, cs.Excluded, 0)
		require.Equal(t, [5]FieldValueString{}, cs.Included)
		require.Equal(t, []FieldValueString{"", ""}, fvs)
	})
}
