package health

import (
	"encoding/json"
	"net/http"
)

func (t DefaultController) GetPing(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Ping{
		Status: "ok",
	})
}
