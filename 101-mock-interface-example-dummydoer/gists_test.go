package _01_mock_interface_example_dummydoer

import (
	"io"

	"strings"
	"testing"
)

// Doerのインターフェイスを満たすダミーのstruct
type dummyDoer struct{}

// doGistsRequestのダミー実装 HTTPRequestを送らないでダミーのデータを返します
func (d *dummyDoer) doGistRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
[
	{"html_url": "https:gist.github.com/example1"},
	{"html_url": "https:gist.github.com/example2"}
]
		`), nil
}

func TestGetGist2(t *testing.T) {
	// dummyDoerはDoerの実装なので、Clientに渡すことができます。
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}

}
