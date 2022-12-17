package response

import (
	"encoding/json"
	"fmt"
	"golang-chi/config"
	"net/http"
	"time"
)

// SUCCESS
func RespondwithJSON(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(map[string]interface{}{"status": "SUCCESS", "data": payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}

// FAILED
func RespondWithError(w http.ResponseWriter, err error) {
	response, _ := json.Marshal(map[string]interface{}{"status": "FAILED", "error": err})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := config.Router
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}
