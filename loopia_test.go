package loopia

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"fmt"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *API

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = New("loopia@loopiaapi", "verysecret")
	client.RPCEndpoint = server.URL
}

func teardown() {
	server.Close()
}

func TestClient_Credentials(t *testing.T) {
	setup()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error("Unexpected error when reading response Body")
		}

		xmlParamPattern := "<param><value>%s</value></param>"
		search := []string {
			fmt.Sprintf(xmlParamPattern, "<string>loopia@loopiaapi</string>"),
			fmt.Sprintf(xmlParamPattern, "<string>verysecret</string>"),
			fmt.Sprintf(xmlParamPattern, "<string></string>"),
			fmt.Sprintf(xmlParamPattern, "<string>example.com</string>"),
			fmt.Sprintf(xmlParamPattern, "<string>www</string>"),
		}
		assert.Contains(t, string(body[:]), strings.Join(search, ""), "Expected username inside XML body")
	})
	client.GetZoneRecord("example.com", "www", 12345)
	teardown()
}
