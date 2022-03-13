package application

import (
	"encoding/json"
	"net/http"
)

func (t DefaultController) GetApplications(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}
