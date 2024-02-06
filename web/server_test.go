package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/x?a=1&b=1", nil)
	require.NoErrorf(t, err, "failed to create request: %v", err)

	rec := httptest.NewRecorder()

	// Testing the sumHandler.
	sumHandler(rec, req)

	resp := rec.Result()
	require.Equalf(t, http.StatusOK, resp.StatusCode, "response status code is not 200 OK: %v", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	require.NoErrorf(t, err, "failed to read body: %v", err)
	defer resp.Body.Close()

	s, err := strconv.Atoi(string(b))
	require.NoErrorf(t, err, "failed to convert body to integer: %v", err)

	require.Equalf(t, 2, s, "one plus one should be 2, got: %v", s)
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	resp, err := http.Get(fmt.Sprintf("%s/sum?a=1&b=1", srv.URL))
	require.NoErrorf(t, err, "failed to send request to /sum: %v", err)

	require.Equalf(t, http.StatusOK, resp.StatusCode, "response status code is not 200 OK: %v", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	require.NoErrorf(t, err, "failed to read body: %v", err)
	defer resp.Body.Close()

	s, err := strconv.Atoi(string(b))
	require.NoErrorf(t, err, "failed to convert body to integer: %v", err)

	require.Equalf(t, 2, s, "one plus one should be 2, got: %v", s)
}
