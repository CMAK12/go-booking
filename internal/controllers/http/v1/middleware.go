package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func ResponseWrapper(handler func(w http.ResponseWriter, r *http.Request) (any, int, int64, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, status, count, err := handler(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(w, status, resp, count)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any, count int64) {
	response := map[string]any{
		"data": data,
	}
	if count > 0 {
		response["count"] = count
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
