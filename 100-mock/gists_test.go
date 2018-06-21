package _00_mock

import (
	"io"
	"strings"
	"testing"
)

func TestGetGists(t *testing.T) {
	doGistsRequest = func(user string) (io.Reader, error) {
		return strings.NewReader(`
[
	{"html_url": "https:gist.github.com/example1"},
	{"html_url": "https:gist.github.com/example2"}
]
		`), nil
	}

	urls, err := ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
