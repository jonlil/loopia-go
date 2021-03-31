package loopia

import (
	"github.com/kolo/xmlrpc"
)

// APIURL - where API is hosted
const APIURL string = "https://api.loopia.se/RPCSERV"

// API Struct to store runtime info
type API struct {
	Username       string
	Password       string
	RPCEndpoint    string
	CustomerNumber string
}

// XMLRPCClient to interact with Loopia XMLRPC
func (api *API) XMLRPCClient() *xmlrpc.Client {
	client, _ := xmlrpc.NewClient(api.RPCEndpoint, nil)
	return client
}

func (api *API) getAuthenticationArgs() []interface{} {
	return []interface{}{
		api.Username,
		api.Password,
		api.CustomerNumber,
	}
}

func (api *API) Call(serviceMethod string, args []interface{}, reply interface{}) error {
	return api.XMLRPCClient().Call(
		serviceMethod,
		append(api.getAuthenticationArgs(), args...),
		reply,
	)
}

// New returns a loopia.API instance
func New(username string, password string) (*API, error) {
	return &API{
		RPCEndpoint: APIURL,
		Username: username,
		Password: password,
		CustomerNumber: "",
	}, nil
}
