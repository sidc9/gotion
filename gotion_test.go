package gotion_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
)

func setup(t *testing.T, saveRespFilename string, req *http.Request) *gotion.Client {
	is := is.New(t)

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	if saveResponse {
		t.Logf("    * saving to: %s\n", saveRespFilename)
		gotion.Init(apiKey, "")
		gotion.SaveResponse(saveRespFilename)
	} else {
		h := func(w http.ResponseWriter, r *http.Request) {
			// instead of asserting the request fields here,
			// save the request and assert in the main test
			req = r

			b, err := ioutil.ReadFile(saveRespFilename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(b)
		}

		srv := httptest.NewServer(http.HandlerFunc(h))
		t.Cleanup(srv.Close)
		gotion.Init(apiKey, srv.URL)
	}

	return gotion.GetClient()
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
