package atlas

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
)

func NewDefaultHttpClient(
	config config.AtlasConfigClient) DefaultHttpClient {

	return DefaultHttpClient{
		config: config,
	}
}

type DefaultHttpClient struct {
	clients.BaseHttpClient
	config config.AtlasConfigClient
}
