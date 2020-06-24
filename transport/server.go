package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"whistlestop.com/attire"
)

// Serve starts the web server GET v1/weather?location=XYZ
func Serve(attire attire.Service) {
	http.HandleFunc("/v1/weather", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request")
		area := r.URL.Query().Get("location")
		if area == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		attire, err := attire.Recommend(area)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(attire)
	})

	http.ListenAndServe(":8080", nil)
}
