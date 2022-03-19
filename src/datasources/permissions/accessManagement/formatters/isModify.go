package formatters

import (
	"net/http"
	"strings"
)

func (t DefaultFormatter) isModify(resource string) bool {

	permissions := (*t.response)[http.MethodPut]

	for _, v := range *permissions {
		if strings.Contains(v, resource) {
			return true
		}
	}

	permissions = (*t.response)[http.MethodPatch]

	for _, v := range *permissions {
		if strings.Contains(v, resource) {
			return true
		}
	}

	return false
}
