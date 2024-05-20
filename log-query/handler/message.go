package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

//func (h *Handler) QueryMessage(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	response := map[string]string{"message": "OK"}
//
//	// Get query parameters
//	queryValues := r.URL.Query()
//	logLevel := queryValues.Get("log_level")
//
//	if logLevel != "" {
//		response = h.Repo.
//	}
//
//
//	err := json.NewEncoder(w).Encode(response)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

func (h *Handler) RetrieveMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	re := regexp.MustCompile(`/messages/([a-f0-9\-]+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 2 {
		id := matches[1]

		message, err := h.Repo.RetrieveMessage(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("{\"error\":\"%s\"}", err.Error()), http.StatusInternalServerError)
			return
		}

		// Serialize the message to JSON
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "{\"error\":\"Failed to serialize response\"}", http.StatusInternalServerError)
			return
		}

		// Write the JSON response
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResponse)
		if err != nil {
			return
		}

	} else {
		http.NotFound(w, r)
	}
}
