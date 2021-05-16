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

	respOut := filepath.Join("testdata", "list_db.txt")
	var req *http.Request
	c := setup(t, respOut, req)
	resp, err := c.ListDatabases()

	is.NoErr(err)
	is.Equal(resp.Object, "list")
	is.Equal(resp.HasMore, false)
	is.Equal(resp.NextCursor, "")

	is.Equal(len(resp.Results), 1)

	item := resp.Results[0]
	is.Equal(item.Object, "database")
	// pretty.Println(item.Properties)
}

func TestGetDatabase(t *testing.T) {
	is := is.New(t)

	t.Run("returns error if id is empty", func(t *testing.T) {
		c := &Client{}
		db, err := c.GetDatabase("")
		is.Equal(err, fmt.Errorf("id is required"))
		is.Equal(db, nil)
	})

	respOut := filepath.Join("testdata", "get_db.txt")
	var req *http.Request
	c := setup(t, respOut, req)

	db, err := c.GetDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.NoErr(err)
	is.Equal(db.Object, "database")
	is.Equal(db.ID, "934c6132-4ea7-485e-9b0d-cf1a083e0f3f")
	is.Equal(db.Properties["age"].Number.Format, "number")

	// pretty.Println(db)
}

func TestQueryDatabase(t *testing.T) {
	t.Run("simple query", func(t *testing.T) {
		is := is.NewRelaxed(t)

		respOut := filepath.Join("testdata", "query_db.txt")
		var req *http.Request
		c := setup(t, respOut, req)

		pages, err := c.QueryDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f", nil)
		is.NoErr(err)
		is.Equal(len(pages.Results), 2)
		// pretty.Println(pages)
		if req != nil {
			is.Equal(req.Method, http.MethodPost)
			is.Equal(req.URL.String(), "/databases/934c6132-4ea7-485e-9b0d-cf1a083e0f3f/query")
		}
	})

	t.Run("with filter", func(t *testing.T) {
		is := is.NewRelaxed(t)

		respOut := filepath.Join("testdata", "query_db_with_filter.txt")
		var req *http.Request
		c := setup(t, respOut, req)

		filt := NewNumberFilter("age").GreaterThan(24)
		q := NewDBQuery().WithFilter(filt)

		pages, err := c.QueryDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f", q)
		is.NoErr(err)
		is.Equal(len(pages.Results), 1)
		is.Equal(pages.Results[0].Properties["age"].Number, 25)
		// pretty.Println(pages)

		if req != nil {
			is.Equal(req.Method, http.MethodPost)
			is.Equal(req.URL.String(), "/databases/934c6132-4ea7-485e-9b0d-cf1a083e0f3f/query")
		}
	})

	t.Run("with sort", func(t *testing.T) {
		is := is.NewRelaxed(t)

		respOut := filepath.Join("testdata", "query_db_with_sort.txt")
		var req *http.Request
		c := setup(t, respOut, req)

		sort := NewPropertySort("age", SortAscending)
		q := NewDBQuery().WithSorts([]*Sort{sort})

		pages, err := c.QueryDatabase("934c6132-4ea7-485e-9b0d-cf1a083e0f3f", q)
		is.NoErr(err)
		is.Equal(len(pages.Results), 2)

		is.Equal(pages.Results[0].Properties["age"].Number, 23)
		is.Equal(pages.Results[1].Properties["age"].Number, 25)

		if req != nil {
			is.Equal(req.Method, http.MethodPost)
			is.Equal(req.URL.String(), "/databases/934c6132-4ea7-485e-9b0d-cf1a083e0f3f/query")
		}
	})
}

func setup(t *testing.T, saveRespFilename string, req *http.Request) *Client {
	is := is.New(t)

	var c *Client

	apiKey, err := loadAPIKey()
	is.NoErr(err)

	if saveResponse {
		t.Logf("    * saving to: %s\n", saveRespFilename)
		c = NewClient(apiKey, "")
		c.SaveResponse(saveRespFilename)
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
		// TODO
		// t.Cleanup(srv.Close)
		c = NewClient(apiKey, srv.URL)
	}

	return c
}
