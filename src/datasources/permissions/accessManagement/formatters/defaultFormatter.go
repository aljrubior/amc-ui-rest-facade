package formatters

import "github.com/aljrubior/go-facade/clients/accessManagement/responses"

func NewDefaultFormatter(response *map[string]*responses.AuthorizeResponse) DefaultFormatter {

	return DefaultFormatter{
		response: response,
	}
}

type DefaultFormatter struct {
	response *map[string]*responses.AuthorizeResponse
}
