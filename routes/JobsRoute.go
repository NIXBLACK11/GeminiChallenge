package routes

import (
	"GeminiChallenge/api"
	"GeminiChallenge/models"
	"GeminiChallenge/utils"
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

		searchValues := utils.GetSearchValues(text, request.Tags)

		for _, val := range searchValues {
			fmt.Println(val)
		}

		var results []api.ReturnData

		for _, searchValue := range searchValues {
			result, err := api.GetGoogleResponse(searchValue)

			if err!=nil {
				fmt.Println(err)
				http.Error(w, "Error in google api", http.StatusBadRequest)
			}

			results = append(results, result...)
		}

		w.WriteHeader(http.StatusOK)
		response := map[string][]api.ReturnData{"message": results}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}