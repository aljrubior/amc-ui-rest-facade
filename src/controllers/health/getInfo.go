package health

import (
	"encoding/json"
	"net/http"
)

func (t DefaultController) GetInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Info{
		Name:    "Facade",
		Code:    "facade",
		Version: "2.12.63",
	})
}
