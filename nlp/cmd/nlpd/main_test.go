package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_healthHandler(t *testing.T) {
	// require := require.New(t)
	// require.Equal(a, b)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	healthHandler(w, r)
	/*
		h, _ := http.DefaultServeMux.Handler(r)
		h.ServeHTTP(w, r)
	*/

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "OK\n", string(body))
}
