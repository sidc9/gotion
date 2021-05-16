package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/kr/pretty"
	"github.com/matryer/is"
)

var saveResponse bool

func TestMain(m *testing.M) {
	flag.BoolVar(&saveResponse, "save", false, "call the real api and save the response")
	flag.Parse()

	if saveResponse {
		fmt.Println("-> Calling real api and saving response(s)")
	}

	code := m.Run()
	os.Exit(code)
}

func TestListDatabases(t *testing.T) {
	is := is.New(t)

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	c := NewClient(apiKey, "")
	resp, err := c.ListDatabases()

	is.NoErr(err)
	is.Equal(resp.Object, "list")
	is.Equal(resp.HasMore, false)
	is.Equal(resp.NextCursor, "")

	is.Equal(len(resp.Results), 1)

	item := resp.Results[0]
	is.Equal(item.Object, "database")
	pretty.Println(item.Properties)

}

func TestGetDatabase(t *testing.T) {
	is := is.New(t)

	c := &Client{}
	t.Run("returns error if id is empty", func(t *testing.T) {
		db, err := c.GetDatabase("")
		is.Equal(err, fmt.Errorf("id is required"))
		is.Equal(db, nil)
	})

	db, err := c.GetDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.NoErr(err)
	is.Equal(db.Object, "database")
	is.Equal(db.ID, "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")

	pretty.Println(db)
}

func TestQueryDatabase(t *testing.T) {
	is := is.New(t)

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	t.Run("simple query", func(t *testing.T) {
		c := NewClient(apiKey, "")
		pages, err := c.QueryDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f", nil)
		is.NoErr(err)
		pretty.Println(pages)
	})

	t.Run("with filter", func(t *testing.T) {
		var (
			c   *Client
			req *http.Request
		)

		saveRespFilename = filepath.Join("testdata", "query_db_with_filter.txt")

		if saveResponse {
			t.Logf("    * saving to: %s\n", saveRespFilename)
			c = NewClient(apiKey, "")
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
			c = NewClient(apiKey, srv.URL)
		}

		filt := NewNumberFilter("age").GreaterThan(24)
		q := NewDBQuery().WithFilter(filt)

		pages, err := c.QueryDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f", q)
		is.NoErr(err)
		is.Equal(len(pages.Results), 1)
		// pretty.Println(pages)

		if req != nil {
			is.Equal(req.Method, http.MethodPost)
			is.Equal(req.URL, "/databases/934c6132-4ea7-485e-9b0d-cf1a083e0f3f/query")
		}
	})
}
