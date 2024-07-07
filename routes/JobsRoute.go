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
			http.Error(w, "Error in gemini api", http.StatusBadRequest)
		}

		result, err := api.GetGoogleResponse(`intext:"golang" intext:"internship"`)

		if err!=nil {
			fmt.Println(err)
			http.Error(w, "Error in google api", http.StatusBadRequest)
		}

		fmt.Println(result)

		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": text}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}