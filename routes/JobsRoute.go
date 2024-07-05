package routes

import (
	"GeminiChallenge/api"
	"GeminiChallenge/models"
	"encoding/json"

	"fmt"
	"net/http"
)

func JobsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		var request models.RequestType

		err := json.NewDecoder(r.Body).Decode(&request)
		if err!=nil{
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		text, err := api.GetGeminiResponse(request.Resume)

		if err!=nil {
			fmt.Println(err)
			http.Error(w, "Server Error", http.StatusBadRequest)
		}
		fmt.Println("API response2")
		fmt.Println(text)

		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": "Request Parsed"}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}