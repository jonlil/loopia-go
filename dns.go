package loopia

import (
	"errors"
)

// Record descired Loopia record_obj
type Record struct {
	ID       int64  `xmlrpc:"record_id"`
	TTL      int    `xmlrpc:"ttl"`
	Type     string `xmlrpc:"type"`
	Value    string `xmlrpc:"rdata"`
	Priority int    `xmlrpc:"priority"`
}

// Status - operation status wrapper
type Status struct {
	Status string
	Cause string
}

// Subdomain ...
type Subdomain struct {
	Name string
}

// AddSubdomain - method for creating subdomain
func (api *API) AddSubdomain(domain string, subdomain string) (*Status, error) {
	var result string
	args := []interface{}{
		api.Username,
		api.Password,
		api.CustomerNumber,
		domain,
		subdomain,
	}
	if err := api.XMLRPCClient().Call("addSubdomain", args, &result); err != nil || result != "OK" {
		return &Status{
			Status: "failed",
			Cause: result,
		}, err
	}

	return &Status{
		Status: "success",
	}, nil
}

// GetSubdomains - Method for fetching all subdomains
func (api *API) GetSubdomains(domain string) ([]Subdomain, error) {
	result := []string{}
	args := []interface{}{
		api.Username,
		api.Password,
		api.CustomerNumber,
		domain,
	}

	if err := api.XMLRPCClient().Call("getSubdomains", args, &result); err != nil {
		return []Subdomain{}, err
	}

	subdomains := []Subdomain{}
	for _,value := range result {
		subdomains = append(subdomains, Subdomain{
			Name: value,
		})
	}

	return subdomains, nil
}

// GetZoneRecords - fetch subdomains records
func (api *API) GetZoneRecords(domain string, subdomain string) ([]Record, error) {
	result := []Record{}
	args := []interface{}{
		api.Username,
		api.Password,
		api.CustomerNumber,
		domain,
		subdomain,
	}

	if err := api.XMLRPCClient().Call("getZoneRecords", args, &result); err != nil {
		return []Record{}, err
	}

	return result, nil
}

// GetZoneRecord - fetch specific zone record
func (api *API) GetZoneRecord(domain string, subdomain string, id int64) (*Record, error) {
	results, err := api.GetZoneRecords(domain, subdomain)
	if err != nil {
		return &Record{}, err
	}

	for _, element := range results {
		if id == element.ID {
			return &element, nil
		}
	}
	return &Record{}, errors.New("ID Not found")
}
