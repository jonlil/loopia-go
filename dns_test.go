package loopia

import (
	"io/ioutil"
	"net/http"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
)

func zoneRecordsHandler(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error("Unexpected error when reading response Body")
		}

		strBody := string(body[:])
		if strings.Contains(strBody, "getZoneRecords") {
			byteArray, _ := ioutil.ReadFile("fixtures/zone_records.xml")
			fmt.Fprintf(w, string(byteArray[:]))
		} else if strings.Contains(strBody, "addZoneRecord") {
			byteArray, _ := ioutil.ReadFile("fixtures/ok.xml")
			fmt.Fprintf(w, string(byteArray[:]))
		}
	}
}

func TestClient_GetZoneRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", zoneRecordsHandler(t))
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

func TestClient_GetSubdomains(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		byteArray, _ := ioutil.ReadFile("fixtures/subdomains.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})

	result, _ := client.GetSubdomains("example.com")
	assert.Equal(t, 8, len(result), "Expected result to have 8 items")
}

func TestClient_GetSubdomain(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		byteArray, _ := ioutil.ReadFile("fixtures/subdomains.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})

	result, _ := client.GetSubdomain("example.com", "www")
	assert.Equal(t, "www", result.Name, "Expected result equal www")
}

func TestClient_AddZoneRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", zoneRecordsHandler(t))

	record := Record{
		TTL:      300,
		Type:     "A",
		Value:    "1.1.1.1",
		Priority: 0,
	}

	err := client.AddZoneRecord("example.com", "api", &record)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, int64(14096733), record.ID, "AddZoneRecord expects to find ID")
}

func TestClient_UpdateZoneRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error("Unexpected error when reading response Body")
		}

		assert.Contains(t,
			string(body[:]),
			"<member><name>record_id</name><value><int>14096733</int></value></member>",
			"ID should be converted to record_id")

		byteArray, _ := ioutil.ReadFile("fixtures/ok.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})

	result, err := client.UpdateZoneRecord("example.com", "api", Record{
		ID:       14096733,
		TTL:      300,
		Type:     "A",
		Value:    "1.1.1.1",
		Priority: 0,
	})

	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "success", result.Status)
}
