package loopia 

import (
    "github.com/kolo/xmlrpc"
    "fmt"
    "errors"
)

type Record struct {
    Id int `xmlrpc:"record_id"`
    Ttt int `xmlrpc:"ttl"`
    Type string `xmlrpc:"type"`
    Value string `xmlrpc:"rdata"`
    Priority int `xmlrpc:"priority"`
}

func GetZoneRecords(b []byte) []Record {
    resp := xmlrpc.NewResponse(b)

    result := []Record{}

    if err := resp.Unmarshal(&result); err != nil {
        fmt.Println("Error:", err)
    }

    return result
}

type RecordGetter func(domain string, subdomain string, id int) []Record

type GetZoneRecord struct {
    get_records RecordGetter
}

func NewGetZoneRecord (rg RecordGetter) *GetZoneRecord {
    return &GetZoneRecord{get_records: rg}
}

func (r *GetZoneRecord) getZoneRecord(domain string, subdomain string, id int) (Record, error) {
    for _, element := range r.get_records(domain, subdomain, id) {
        if id == element.Id {
            return element, nil
        }
    }
    return Record{}, errors.New("ID Not found") 
}
