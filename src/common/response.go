package common

import (
	"encoding/json"
	"net/http"
)

func CreateJSONResponse(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	result := map[string]any{}
	result["data"] = data
	_, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(500)
		result["data"] = nil
		error := map[string]any{}
		error["message"] = "something wrong when create response"
		result["error"] = error
	}

	response, _ := json.Marshal(result)

	w.Write(response)
}
