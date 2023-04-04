package antena

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ping" {
		http.NotFound(w, r)
		return
	}
	var response Response
	response.setPong()
	jr, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
	}
	_, err = w.Write(jr)
	if err != nil {
		http.Error(w, "error sending response", http.StatusBadRequest)
	}
}

func StartServer(port int) error {
	http.HandleFunc("/ping", PingHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
