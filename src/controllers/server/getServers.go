package server

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/controllers/headers"
	"log"
	"net/http"
)

func (t DefaultController) GetServers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(headers.AuthorizationHttpHeader)
	orgId := r.Header.Get(headers.OrganizationIdHttpHeader)
	envId := r.Header.Get(headers.EnvironmentIdHttpHeader)

	targets, err := t.serverService.GetServers(token, orgId, envId)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(targets)
}
