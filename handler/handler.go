package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// Echo receives whatever was sent and posts it back
func Echo(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"payload":     request,
		"processedAt": time.Now().Format(time.RFC1123),
		"status":      "received",
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
