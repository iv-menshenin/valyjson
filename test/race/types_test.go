package race

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Unmarshal_TestStruct_URL_Race(t *testing.T) {
	t.Parallel()
	const data = `{"site_id":"5e8a266c308e7","user_id":"983cf94b-0816-a412-21db-c01bef0e1d6a","referer":"","url":"%s","ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.962 YaBrowser/23.9.1.962 Yowser/2.5 Safari/537.36","has_codes":true}`
	var urls = []string{
		"https://somedomain.ru/path/kak-est-goroh-lojkoy/",
		"https://somedomain.ru/path/kak-goroh/",
		"https://somedomain.ru/path/",
		"https://somedomain.ru/path/lorem-ipsum/",
	}

	var wg sync.WaitGroup
	for n := 0; n < 100000; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			var dataJson = fmt.Sprintf(data, urls[n%len(urls)])
			var s TestStruct
			require.NoError(t, s.UnmarshalJSON([]byte(dataJson)))
			if strings.Index(s.URL, "\"") > -1 {
				t.Errorf("race: %q", s.URL)
			}
		}(n)
	}
	wg.Wait()
}

func Test_Unmarshal_UserTyped_Race(t *testing.T) {
	t.Parallel()
	const data = `{"site_id":"5e8a266c308e7","user_id":"983cf94b-0816-a412-21db-c01bef0e1d6a","referer":"","url":"%s","ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.962 YaBrowser/23.9.1.962 Yowser/2.5 Safari/537.36","has_codes":true}`
	var urls = []string{
		"https://somedomain.ru/path/kak-est-goroh-lojkoy/",
		"https://somedomain.ru/path/kak-goroh/",
		"https://somedomain.ru/path/",
		"https://somedomain.ru/path/lorem-ipsum/",
	}

	var wg sync.WaitGroup
	for n := 0; n < 100000; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			var dataJson = fmt.Sprintf(data, urls[n%len(urls)])
			var s TestStructTyped
			require.NoError(t, s.UnmarshalJSON([]byte(dataJson)))
			if strings.Index(string(s.URL), "\"") > -1 {
				t.Errorf("race: %q", s.URL)
			}
		}(n)
	}
	wg.Wait()
}

var unmarshalReuseRaceTestPool = sync.Pool{New: func() any { return &TestStructTyped{} }}

func Test_Unmarshal_Reuse_Race(t *testing.T) {
	t.Parallel()
	const data = `{"site_id":"5e8a266c308e7"%s,"referer":""%s,"ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.962 YaBrowser/23.9.1.962 Yowser/2.5 Safari/537.36"%s}`
	var urls = []string{
		"",
		`https://somedomain.ru/path/kak-est-goroh-lojkoy/`,
		`https://somedomain.ru/path/kak-goroh/`,
		`https://somedomain.ru/path/`,
		`https://somedomain.ru/path/lorem-ipsum/`,
	}
	var users = []string{
		"",
		`983cf94b-0816-a412-21db-c01bef0e1d6a`,
		`22bfb495-52c3-46af-b2d0-972bc4c1d29d`,
		`42c61daa-782a-4244-94d1-0cfb950b202e`,
	}
	var codes = []string{
		"",
		`true`,
		`false`,
	}

	var wg sync.WaitGroup
	for n := 0; n < 100000; n++ {
		wg.Add(1)
		go func(url, user, code string) {
			defer wg.Done()
			var xurl, xuser, xcode string
			if url != "" {
				xurl = fmt.Sprintf(`,"url":"%s"`, url)
			}
			if user != "" {
				xuser = fmt.Sprintf(`,"user_id":"%s"`, user)
			}
			if code != "" {
				xcode = fmt.Sprintf(`,"has_codes":%s`, code)
			}
			var dataJson = fmt.Sprintf(data, xurl, xuser, xcode)
			var s = unmarshalReuseRaceTestPool.Get().(*TestStructTyped)
			require.NoError(t, s.UnmarshalJSON([]byte(dataJson)))
			require.EqualValues(t, url, s.URL)
			require.EqualValues(t, user, s.UserID)
			require.EqualValues(t, code == "true", s.HasCodes)
			s.Reset()
			unmarshalReuseRaceTestPool.Put(s)
		}(urls[n%len(urls)], users[n%len(users)], codes[n%len(codes)])
	}
	wg.Wait()
}
