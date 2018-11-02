package loopia

import (

  "net/http"
  "testing"
  "io/ioutil"

  "github.com/stretchr/testify/assert"
  "fmt"
)

func TestClient_GetZoneRecord(t *testing.T) {
    setup()
    defer teardown()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
        byteArray, _ := ioutil.ReadFile("response.xml")
        fmt.Fprintf(w, string(byteArray[:]))
    })
    record, _ := client.GetZoneRecord("example.com", "@", 14096733)
    assert.Equal(t, int64(14096733), record.Id)
}
