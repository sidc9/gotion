package gotion_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

type mockRoundTripper struct {
	fn func(*http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.fn(r)
}

func loadAPIKey() (string, error) {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(b), "\n"), nil
}

func getClient(t *testing.T) *gotion.Client {
	is := is.New(t)

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	return gotion.NewClient(apiKey, "")
}

func setResponse(t *testing.T, c *gotion.Client, response, method, path string) {
	is := is.New(t)

	// replace the http client with a mock one
	c.WithHTTPClient(&http.Client{
		Transport: &mockRoundTripper{
			fn: func(r *http.Request) (*http.Response, error) {
				is.Equal(method, r.Method)
				is.Equal(path, r.URL.Path)
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(response)),
				}, nil
			},
		},
	})
}

func setResponseFromFile(t *testing.T, c *gotion.Client, responseFile, method, path string) {
	t.Helper()
	filename := filepath.Join("testdata", responseFile)
	if saveResponse {
		// saveResponse means client will call the actual API and save the response
		t.Logf("    * saving to: %s\n", filename)
		c.SaveResponse(filename)
	} else {
		f, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		setResponse(t, c, string(f), method, path)
	}
}
