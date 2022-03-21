package atlas

import (
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
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
