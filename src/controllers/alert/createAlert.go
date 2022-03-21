package alert

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
	"github.com/aljrubior/go-facade/controllers/headers"
	"github.com/gorilla/mux"
	"net/http"
)

func (t DefaultController) CreateAlert(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(headers.AuthorizationHttpHeader)
	orgId := r.Header.Get(headers.OrganizationIdHttpHeader)
	envId := r.Header.Get(headers.EnvironmentIdHttpHeader)

	vars := mux.Vars(r)

	productId, ok := vars["productId"]

	if !ok {
		// TODO: Implement
		http.Error(w, "productId is missing in parameters", http.StatusBadRequest)
	}

	var request requests.AlertRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp, err := t.alertService.PostAlert(token, orgId, envId, productId, request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(resp)
}
