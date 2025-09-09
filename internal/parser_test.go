package internal

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz-0987654321_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	size    = int64(len(letters))
)

func randomId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int64()%size]
	}
	return string(b)
}

func TestLinkParserValidLinks(t *testing.T) {
	var testsMap = []struct {
		link          string
		expectedStart string
	}{
		{"https://www.youtube.com/watch?v=%s", ""},
		{"https://www.youtube.com/watch?v=%s&t=123s", "123"},
		{"https://www.youtube.com/watch?t=123&v=%s", "123"},
		{"https://youtu.be/%s?si=0000000000000000", ""},
		{"https://youtu.be/%s?si=0000000000000000&t=123", "123"},
		{"https://www.youtube.com/watch?v=%s&list=WL&index=2&pp=gAQBiAQB", ""},
		{"https://www.youtube.com/embed/%s?si=ilpCXG_BnFBzj8q9", ""},
		{"https://www.youtube.com/embed/%s?si=ilpCXG_BnFBzj8q9&amp;start=6", "6"},
	}

	for _, test := range testsMap {
		t.Run("", func(t *testing.T) {
			// generate random id
			rid := randomId(11)
			link := fmt.Sprintf(test.link, rid)

			id, start, err := LinkParser(link)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			if id != rid {
				t.Errorf("got id %s when parsing %s", id, test.link)
			}
			if start != test.expectedStart {
				t.Errorf("got start %s when parsing %s", start, test.link)
			}
		})
	}
}

func TestLinkParserInvalidLinks(t *testing.T) {
	var testList = []string{
		"https://www.example.com",
		"https://www.youtube.com/watch?t=123s",
		"https://www.youtube.com/watch/bananas",
		"https://youtu.be/",
		"https://youtu.be",
		"https://www.youtube.com/embed/",
		"https://www.youtube.com/embed",
		"https://www.youtube.com/bananas",
	}

	for _, test := range testList {
		t.Run("", func(t *testing.T) {
			_, _, err := LinkParser(test)
			if err == nil {
				t.Errorf("no error when parsing %s", test)
			}
		})
	}
}

func TestLinkParserTimeParser(t *testing.T) {
	var testsMap = []struct {
		link          string
		expectedStart string
	}{
		{"https://www.youtube.com/watch?v=videoid&t=123", "123"},
		{"https://www.youtube.com/watch?v=videoid&t=123s", "123"},
		{"https://www.youtube.com/watch?v=videoid&t=12m", "720"},
		{"https://www.youtube.com/watch?v=videoid&t=1m23s", "83"},
		{"https://www.youtube.com/watch?v=videoid&t=1h4m", "3840"},
		{"https://www.youtube.com/watch?v=videoid&t=1h3s", "3603"},
		{"https://www.youtube.com/watch?v=videoid&t=1h5m3s", "3903"},
	}

	for _, test := range testsMap {
		t.Run("", func(t *testing.T) {
			// generate random id

			_, start, err := LinkParser(test.link)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			if start != test.expectedStart {
				t.Errorf("got start %s when parsing %s", start, test.link)
			}
		})
	}
}
