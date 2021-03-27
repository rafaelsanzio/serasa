package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"serasa/models"
)

func ListNegativationsByAPI() ([]models.Negativation, error) {
  NEGATIVATION_API_URL := os.Getenv("NEGATIVATION_API_URL")

  response := []models.Negativation{}
	resp, err := http.Get(NEGATIVATION_API_URL + "/negativations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	ResponseWithJSON(w, code, map[string]string{"error": msg})
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
