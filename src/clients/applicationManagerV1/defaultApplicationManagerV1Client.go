package applicationManagerV1

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
)

func NewDefaultApplicationManagerV1Client(
	config config.ApplicationManagerV1ConfigClient) DefaultApplicationManagerV1Client {

	return DefaultApplicationManagerV1Client{
		config: config,
	}
}

type DefaultApplicationManagerV1Client struct {
	clients.BaseHttpClient
	config config.ApplicationManagerV1ConfigClient
}
