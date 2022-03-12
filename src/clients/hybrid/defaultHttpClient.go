package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
)

func NewDefaultHttpClient(
	config config.HybridConfigClient) DefaultHttpClient {

	return DefaultHttpClient{
		config: config,
	}
}

type DefaultHttpClient struct {
	clients.BaseHttpClient
	config config.HybridConfigClient
}
