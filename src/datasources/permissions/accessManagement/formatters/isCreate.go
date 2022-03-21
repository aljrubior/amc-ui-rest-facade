package formatters

import (
	"net/http"
	"strings"
)

func (t DefaultFormatter) isCreate(resource string) bool {

	permissions := (*t.response)[http.MethodPost]

	for _, v := range *permissions {
		if strings.Contains(v, resource) {
			return true
		}
	}

	return false
}
