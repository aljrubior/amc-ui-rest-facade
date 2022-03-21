package formatters

import (
	"net/http"
	"strings"
)

func (t DefaultFormatter) isRead(resource string) bool {

	permissions := (*t.response)[http.MethodGet]

	for _, v := range *permissions {
		if strings.Contains(v, resource) {
			return true
		}
	}

	return false
}
