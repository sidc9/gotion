package gotion_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

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

// func setResponse(t *testing.T, c *gotion.Client, expectedRequest *http.Request, responseFile string) {
func setResponse(t *testing.T, c *gotion.Client, responseFile string) {
	filename := filepath.Join("testdata", responseFile)
	if saveResponse {
		t.Logf("    * saving to: %s\n", filename)
		c.SaveResponse(filename)
	} else {
		f, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		c.WithHTTPClient(&http.Client{
			Transport: &mockRoundTripper{
				fn: func(r *http.Request) (*http.Response, error) {
					// if r != expectedRequest {
					//     t.Fatal("alu")
					// }
					return &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBuffer(f)),
					}, nil
				},
			},
		})
	}
}
