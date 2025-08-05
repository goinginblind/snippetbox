package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// newTestApplication returns a pointer to an application struct which has
// proper logger in it. It is needed for some of the middleware to work.
func newTestApplication() *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

// testServer is a wrapper for an http.Server
type testServer struct {
	*httptest.Server
}

// newTestServer returns a pointer to a new instance of testServer
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, string(body)
}
