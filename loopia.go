package loopia

import (
  "github.com/kolo/xmlrpc"
)

const API_URL string = "https://api.loopia.se/RPCSERV"

type API struct {
    APIUsername string
    APIPassword string
    RCPEndpoint string
    xmlrpc *xmlrpc.Client
}

func (api *API) XmlRpcClient() *xmlrpc.Client {
  client, _ := xmlrpc.NewClient(api.RCPEndpoint, nil)
  return client
}

func New(username string, password string) (*API, error) {
  api := &API{
    RCPEndpoint: API_URL,
  }

  api.APIUsername = username
  api.APIPassword = password

  return api, nil
}
