package internal

import (
	"testing"
)

func TestLinkParserValidLinks(t *testing.T) {
	var testsMap = []struct {
		link     string
		expected string
		start    string
	}{
		{"https://www.youtube.com/watch?v=abcdefghijk", "abcdefghijk", ""},
		{"https://www.youtube.com/watch?v=ABCDEFGHIJK&t=123s", "ABCDEFGHIJK", "123s"},
		{"https://www.youtube.com/watch?t=123&v=zyxwvutsrqp", "zyxwvutsrqp", "123"},
		{"https://youtu.be/ZYXWVUTSRQP?si=0000000000000000", "ZYXWVUTSRQP", ""},
		{"https://youtu.be/12345678901?si=0000000000000000&t=123", "12345678901", "123"},
		{"https://www.youtube.com/watch?v=a1b2c3d3e4f&list=WL&index=2&pp=gAQBiAQB", "a1b2c3d3e4f", ""},
		{"https://www.youtube.com/embed/A1B2C3D3E4F?si=ilpCXG_BnFBzj8q9", "A1B2C3D3E4F", ""},
		{"https://www.youtube.com/embed/0-0_0-0_0-0?si=ilpCXG_BnFBzj8q9&amp;start=6", "0-0_0-0_0-0", "6"},
	}

	for _, test := range testsMap {
		t.Run("", func(t *testing.T) {
			id, start, err := LinkParser(test.link)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			if id != test.expected {
				t.Errorf("got id %s when parsing %s", id, test.link)
			}
			if start != test.start {
				t.Errorf("got start %s when parsing %s", start, test.link)
			}
		})
	}
}
