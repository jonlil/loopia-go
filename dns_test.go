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
		byteArray, _ := ioutil.ReadFile("fixtures/response.xml")
		fmt.Fprintf(w, string(byteArray[:]))
	})
	record, _ := client.GetZoneRecord("example.com", "@", 14096733)
	assert.Equal(t, int64(14096733), record.ID)
}
