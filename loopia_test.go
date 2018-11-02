package loopia

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "io/ioutil"

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
    client.RCPEndpoint = server.URL
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
            fmt.Println(err)
            return
        }
        assert.Contains(t, string(body[:]), "<value><string>loopia@loopiaapi</string></value>", "Expected username inside XML body")
        assert.Contains(t, string(body[:]), "<value><string>verysecret</string></value>", "Expected password inside XML body")
    })
    client.GetZoneRecord("example.com", "www", 12345)
    teardown()
}
