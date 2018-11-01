package loopia

import (
    "testing"
    "io/ioutil"
    "github.com/kolo/xmlrpc"
    "fmt"
)

func mock_get_records(domain string, subdomain string, id int) []Record {
    b, err := ioutil.ReadFile("response.xml") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    resp := xmlrpc.NewResponse(b)

    result := []Record{}

    if err := resp.Unmarshal(&result); err != nil {
        fmt.Println("Error:", err)
    }

    return result
}

func TestGetZoneRecord(t *testing.T) {
    d := NewGetZoneRecord(mock_get_records)
    record, err = d.getZoneRecord("example.com", "@", 14096733)
    if record.Id != 14096733 {
        t.Fatal("Didn't get correct record from query")
    }
}
