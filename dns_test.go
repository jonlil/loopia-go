package loopia

import (
	"io/ioutil"
	"net/http"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetZoneRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		byteArray, _ := ioutil.ReadFile("fixtures/zone_records.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})
	record, _ := client.GetZoneRecord("example.com", "@", 14096733)
	assert.Equal(t, int64(14096733), record.ID)
}

func TestClient_AddSubdomain_OK(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		byteArray, _ := ioutil.ReadFile("fixtures/ok.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})
	result, _ := client.AddSubdomain("example.com", "test")
	assert.Equal(t, "success", result.Status)
}

func TestClient_AddSubdomain_AUTH_ERROR(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		byteArray, _ := ioutil.ReadFile("fixtures/auth_error.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})
	result, _ := client.AddSubdomain("example.com", "test")
	assert.Equal(t, "failed", result.Status)
	assert.Equal(t, "AUTH_ERROR", result.Cause)
}
