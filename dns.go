package loopia

import (
    "errors"
)

type Record struct {
    Id int64 `xmlrpc:"record_id"`
    Ttl int `xmlrpc:"ttl"`
    Type string `xmlrpc:"type"`
    Value string `xmlrpc:"rdata"`
    Priority int `xmlrpc:"priority"`
}

func (api *API) GetZoneRecords(domain string, subdomain string) ([]Record, error) {
    result := []Record{}
    args := []interface{}{
        api.APIUsername,
        api.APIPassword,
        "",
        domain,
        subdomain,
    }

    if err := api.XmlRpcClient().Call("getZoneRecords", args, &result); err != nil {
      return []Record{}, err
    }

    return result, nil
}

func (api *API) GetZoneRecord(domain string, subdomain string, id int64) (Record, error) {
    results, err := api.GetZoneRecords(domain, subdomain)
    if err != nil {
        return Record{}, err
    }

    for _, element := range results {
        if id == element.Id {
            return element, nil
        }
    }
    return Record{}, errors.New("ID Not found")
}
