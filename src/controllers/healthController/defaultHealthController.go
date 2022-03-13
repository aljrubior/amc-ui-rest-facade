package healthController

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Version string `json:"version"`
}

type Ping struct {
	Status string `json:"status"`
}

func NewDefaultHealthController() DefaultHealthController {
	return DefaultHealthController{}
}

type DefaultHealthController struct {
}

func (t DefaultHealthController) GetInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Info{
		Name:    "Facade",
		Code:    "facade",
		Version: "2.12.63",
	})
}

func (t DefaultHealthController) GetPing(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Ping{
		Status: "ok",
	})
}
