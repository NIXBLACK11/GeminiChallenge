package routes

import (
	"GeminiChallenge/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func JobsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		vars := mux.Vars(r)
		
		var request models.Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err!=nil{
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}


	}
}