package loopia

import (
    "github.com/kolo/xmlrpc"
    "fmt"
    "errors"
    "os"
    "log"
)

type Record struct {
    Id int64 `xmlrpc:"record_id"`
    Ttl int `xmlrpc:"ttl"`
    Type string `xmlrpc:"type"`
    Value string `xmlrpc:"rdata"`
    Priority int `xmlrpc:"priority"`
}

func NewClient() (*xmlrpc.Client, error) {
  return xmlrpc.NewClient(API_URL, nil)
}

func GetZoneRecords(domain string, subdomain string) []Record {
    client, _ := NewClient()

    result := []Record{}
    args := []interface{}{
        os.Getenv("LOOPIA_USERNAME"),
        os.Getenv("LOOPIA_PASSWORD"),
        "",
        domain,
        subdomain,
    }

    if err := client.Call("getZoneRecords", args, &result); err != nil {
      fmt.Println("Error:", err)
    }

    return result
}

func GetZoneRecord(domain string, subdomain string, id int64) (Record, error) {
    for _, element := range GetZoneRecords(domain, subdomain) {
        if id == element.Id {
            return element, nil
        }
    }
    return Record{}, errors.New("ID Not found")
}
