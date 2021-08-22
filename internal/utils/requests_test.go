package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRequest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("body"))
	}))

	defer testServer.Close()

	resp, err := SendRequest(testServer.URL, "GET", nil)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, string(resp), "body")
}
