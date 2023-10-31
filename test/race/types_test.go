package race

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UnmarshalTestStruct(t *testing.T) {
	const data = `{"site_id":"5e8a266c308e7","user_id":"983cf94b-0816-a412-21db-c01bef0e1d6a","referer":"","url":"%s","ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.962 YaBrowser/23.9.1.962 Yowser/2.5 Safari/537.36","has_codes":true}`
	var urls = []string{
		"https://somedomain.ru/path/kak-est-goroh-lojkoy/",
		"https://somedomain.ru/path/kak-goroh/",
		"https://somedomain.ru/path/",
		"https://somedomain.ru/path/lorem-ipsum/",
	}

	var wg sync.WaitGroup
	for n := 0; n < 50000; n++ {
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
