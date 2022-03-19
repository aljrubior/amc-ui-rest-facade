package formatters

import (
	"net/http"
	"strings"
)

func (t DefaultFormatter) isDelete(resource string) bool {

	permissions := (*t.response)[http.MethodDelete]

	for _, v := range *permissions {
		if strings.Contains(v, resource) {
			return true
		}
	}

	return false
}
