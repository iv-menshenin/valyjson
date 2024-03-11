package test_extr

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"fill/test_any"
	"fill/test_string"
)

func TestExternalFill(t *testing.T) {
	t.Parallel()
	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var test External
		var def = "default"
		err := test.UnmarshalJSON([]byte(`{"test1":{"comment":"test","level":444},"test2":{"field":"foo","fieldRef":"bar"}}`))
		require.NoError(t, err)
		var expected = External{
			Test01: test_any.TestAllOfSecond{
				Comment: "test",
				Level:   444,
			},
			Test02: test_string.TestStr01{
				Field:    "foo",
				FieldRef: nil,
				DefRef:   &def,
			},
		}
		expected.Test02.FieldRef = new(string)
		*expected.Test02.FieldRef = "bar"
		require.Equal(t, expected, test)
	})
	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		const expected = `{"test1":{"comment":"test_2","level":456},"test2":{"field":"bar","fieldRef":"foo","defRef":null}}`
		var rF = "foo"
		var obj = External{
			Test01: test_any.TestAllOfSecond{
				Comment: "test_2",
				Level:   456,
			},
			Test02: test_string.TestStr01{
				Field:    "bar",
				FieldRef: &rF,
			},
		}
		data, err := obj.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("marshal_null", func(t *testing.T) {
		t.Parallel()
		var obj *External
		data, err := obj.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
}

func Test_IsZero(t *testing.T) {
	t.Parallel()
	t.Run("just_created", func(t *testing.T) {
		t.Parallel()
		var v External
		require.True(t, v.IsZero())
	})
	t.Run("not_zero_1", func(t *testing.T) {
		t.Parallel()
		var v External
		v.Test01.Comment = "0"
		require.False(t, v.IsZero())
	})
	t.Run("not_zero_2", func(t *testing.T) {
		t.Parallel()
		var v External
		var refStr = ""
		v.Test02.FieldRef = &refStr
		require.False(t, v.IsZero())
	})
}

func TestExternalNested(t *testing.T) {
	t.Parallel()
	t.Run("MarshalJSON", func(t *testing.T) {
		t.Parallel()
		var refA, refB = "A", "B"
		var v = ExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "someComment",
				Level:   12,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "CMD_test",
				Range:   22,
			},
			TestStr01: test_string.TestStr01{
				Field:    "Fld0001A",
				FieldRef: &refA,
				DefRef:   &refB,
			},
		}
		const expected = "{\"command\":\"CMD_test\", \"comment\":\"someComment\", \"defRef\":\"B\", \"field\":\"Fld0001A\", \"fieldRef\":\"A\", \"level\":12, \"range\":22}"
		data, err := v.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Parallel()
		const data = "{\"command\":\"CMD_test_2\", \"comment\":\"someComment_2\", \"defRef\":\"A\", \"field\":\"Fld0001B\", \"fieldRef\":\"B\", \"level\":21, \"range\":33}"
		var refA, refB = "A", "B"
		var expected = ExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "someComment_2",
				Level:   21,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "CMD_test_2",
				Range:   33,
			},
			TestStr01: test_string.TestStr01{
				Field:    "Fld0001B",
				FieldRef: &refB,
				DefRef:   &refA,
			},
		}
		var actual ExternalNested
		require.NoError(t, actual.UnmarshalJSON([]byte(data)))
		require.EqualValues(t, expected, actual)
	})
	t.Run("IsZero", func(t *testing.T) {
		t.Parallel()
		var empty ExternalNested
		require.True(t, empty.IsZero())

		var s string
		var nonEmpty = ExternalNested{
			TestStr01: test_string.TestStr01{
				FieldRef: &s,
			},
		}
		require.False(t, nonEmpty.IsZero())
	})
}

func TestReset(t *testing.T) {
	t.Parallel()
	t.Run("External", func(t *testing.T) {
		t.Parallel()
		var (
			s1, s2 = "foo1", "bar1"
		)
		var val = External{
			Test01: test_any.TestAllOfSecond{
				Comment: "foo",
				Level:   112,
			},
			Test02: test_string.TestStr01{
				Field:    "bar",
				FieldRef: &s1,
				DefRef:   &s2,
			},
		}
		val.Reset()
		require.Empty(t, val)
		require.Empty(t, val.Test01.Comment)
		require.Empty(t, val.Test01.Level)
		require.Empty(t, val.Test02.Field)
		require.Nil(t, val.Test02.FieldRef)
		require.Nil(t, val.Test02.DefRef)
	})
}

var unmarshalReuseRaceTestPool = sync.Pool{New: func() any { return &ExternalNested{} }}

func Test_Unmarshal_Reuse_Race(t *testing.T) {
	t.Parallel()
	const data = `{"defRef":"B"%s%s%s, "field":"Fld0001A", "fieldRef":"A", "range":22}`
	var comment = []string{
		"",
		`cmd1`,
		`cmd1`,
	}
	var command = []string{
		"",
		`983cf94b-0816-a412-21db-c01bef0e1d6a`,
		`22bfb495-52c3-46af-b2d0-972bc4c1d29d`,
		`42c61daa-782a-4244-94d1-0cfb950b202e`,
	}

	var wg sync.WaitGroup
	for n := 0; n < 100000; n++ {
		wg.Add(1)
		go func(comment, command string, level int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			defer wg.Done()
			var xcomment, xcommand, xlevel string
			if comment != "" {
				xcomment = fmt.Sprintf(`,"comment":"%s"`, comment)
			}
			if command != "" {
				xcommand = fmt.Sprintf(`,"command":"%s"`, command)
			}
			if level != 0 {
				xlevel = fmt.Sprintf(`,"level":%d`, level)
			}
			var dataJson = fmt.Sprintf(data, xcomment, xcommand, xlevel)
			var s = unmarshalReuseRaceTestPool.Get().(*ExternalNested)
			require.NoError(t, s.UnmarshalJSON([]byte(dataJson)))
			require.EqualValues(t, command, s.Command)
			require.EqualValues(t, comment, s.Comment)
			require.EqualValues(t, int64(level), s.Level)
			s.Reset()
			unmarshalReuseRaceTestPool.Put(s)
		}(comment[n%len(comment)], command[n%len(command)], n%9)
	}
	wg.Wait()
}

func TestExternalStringSlice(t *testing.T) {
	t.Parallel()
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		var ts = ExternalStringSlice{"foo", "bar"}
		data, err := ts.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `["foo", "bar"]`, string(data))
	})
	t.Run("MarshalEmpty", func(t *testing.T) {
		t.Parallel()
		var ts = ExternalStringSlice{}
		data, err := ts.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `[]`, string(data))
		require.True(t, ts.IsZero())
	})
	t.Run("MarshalNil", func(t *testing.T) {
		t.Parallel()
		var ts ExternalStringSlice
		data, err := ts.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `null`, string(data))
		require.True(t, ts.IsZero())
	})
	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()
		var test ExternalStringSlice
		require.NoError(t, test.UnmarshalJSON([]byte(`["foo", "bar", ""]`)))
		require.Equal(t, ExternalStringSlice{"foo", "bar", ""}, test)
		require.False(t, test.IsZero())
	})
}

func TestExternalStructSlice(t *testing.T) {
	t.Parallel()
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		var a, b = "a", "b"
		var test = ExternalStructSlice{
			{
				Field:    "foo",
				FieldRef: &a,
			},
			{
				Field:  "bar",
				DefRef: &b,
			},
		}
		const expected = `[{"field":"foo","fieldRef":"a","defRef":null},{"field":"bar","fieldRef":null,"defRef":"b"}]`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		var a, b = "foo", "bar"
		var expected = ExternalStructSlice{
			{
				Field:    "1",
				FieldRef: &a,
			},
			{
				Field:  "2",
				DefRef: &b,
			},
		}
		const data = `[{"field":"1","fieldRef":"foo","defRef":null},{"field":"2","fieldRef":null,"defRef":"bar"}]`
		var test ExternalStructSlice
		err := test.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.EqualValues(t, expected, test)
	})
	t.Run("Reset", func(t *testing.T) {
		t.Parallel()
		var a, b = "foo", "bar"
		var test = ExternalStructSlice{
			{
				Field:    "1",
				FieldRef: &a,
			},
			{
				Field:  "2",
				DefRef: &b,
			},
		}
		require.False(t, test.IsZero())
		test.Reset()
		require.True(t, test.IsZero())
		require.Equal(t, ExternalStructSlice{}, test)
		require.Equal(t, 2, cap(test))
	})
}

var unmarshalExternalStructSlicePool = sync.Pool{New: func() any { return &ExternalStructSlice{} }}

func Test_ExternalStructSlice_Reuse_Race(t *testing.T) {
	t.Parallel()
	const data = `{"n":"v"%s%s%s}`
	var comment = []string{
		"",
		`cmd1`,
		`cmd1-stop`,
	}
	var command = []string{
		"",
		`983cf94b-a412-c01bef0e1d6a`,
		`22bfb495-46af-b2d0-972bc4c1d29d`,
		`42c61daa-782a-4244-94d1-0cfb950b202e`,
	}

	var wg sync.WaitGroup
	for n := 0; n < 100000; n++ {
		wg.Add(1)
		var cmm, cmd []string
		var lln = n % 4
		for i := 0; i < lln; i++ {
			cmm = append(cmm, comment[(i+n)%len(comment)])
			cmd = append(cmd, command[(i+n)%len(command)])
		}
		go func(comments, commands []string, level, ln int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			defer wg.Done()
			var dataJson []string
			for i := 0; i < len(comments); i++ {
				comment1 := comments[i]
				command1 := commands[i]
				var xcomment, xcommand, xlevel string
				if comment1 != "" {
					xcomment = fmt.Sprintf(`,"field":"%s"`, comment1)
				}
				if command1 != "" {
					xcommand = fmt.Sprintf(`,"fieldRef":"%s"`, command1)
				}
				if level != 0 {
					xlevel = fmt.Sprintf(`,"defRef":"%d"`, level)
				}
				dataJson = append(dataJson, fmt.Sprintf(data, xcomment, xcommand, xlevel))
			}
			var s = unmarshalExternalStructSlicePool.Get().(*ExternalStructSlice)
			require.NoError(t, s.UnmarshalJSON([]byte(fmt.Sprintf("[%s]", strings.Join(dataJson, ",")))))
			require.Len(t, []test_string.TestStr01(*s), ln)
			for i, v := range *s {
				require.EqualValues(t, comments[i], v.Field)
				if commands[i] == "" {
					require.Nil(t, v.FieldRef)
				} else {
					require.EqualValues(t, commands[i], *v.FieldRef)
				}
				if level == 0 {
					require.EqualValues(t, "default", *v.DefRef)
				} else {
					require.EqualValues(t, strconv.Itoa(level), *v.DefRef)
				}
			}
			s.Reset()
			unmarshalExternalStructSlicePool.Put(s)
		}(cmm, cmd, n%19, lln)
	}
	wg.Wait()
}
