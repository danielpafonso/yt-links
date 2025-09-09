package internal

import (
	"errors"
	"fmt"
	"html"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// LinkParser parses youtube links and extracts video id
func LinkParser(link string) (string, string, error) {
	urlLink, err := url.Parse(html.UnescapeString(link))
	if err != nil {
		return "", "", err
	}
	params := urlLink.Query()
	id := params["v"]
	t := params["t"]

	// try parsing t hms
	if len(t) > 0 {
		_, err := strconv.Atoi(t[0])
		if err != nil {
			pt, err := time.ParseDuration(t[0])
			if err != nil {
				t = nil
			}
			t[0] = fmt.Sprint(pt.Seconds())
		}
	}

	if strings.Contains(link, "youtube.com/watch") {
		if len(id) == 0 {
			return "", "", errors.New("url without video id")
		}
		if len(t) > 0 {
			return id[0], t[0], nil
		} else {
			return id[0], "", nil
		}

	} else if strings.Contains(link, "youtu.be/") {
		if urlLink.Path[1:] == "" {
			return "", "", errors.New("url without video id")
		}
		if len(t) > 0 {
			return urlLink.Path[1:], t[0], nil
		} else {
			return urlLink.Path[1:], "", nil
		}

	} else if strings.Contains(link, "youtube.com/embed/") {
		if urlLink.Path[7:] == "" {
			return "", "", errors.New("url without video id")
		}
		start := params["start"]
		if len(start) > 0 {
			return urlLink.Path[7:], start[0], nil
		} else {
			return urlLink.Path[7:], "", nil
		}

	} else {
		return "", "", errors.New("unknow url")
	}
}
