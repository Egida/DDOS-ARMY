package antena

import (
	"container/list"
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

var orderList list.List

func Order(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/order" {
		http.NotFound(w, r)
		return
	}
	var response Response

	value := orderList.Front()
	if value == nil {
		response.setOrder(NOTHING)
	} else {
		response.setOrder(value.Value.(string))
	}

	jr, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error sending response", http.StatusBadRequest)
	}
	_, err = w.Write(jr)
	if err != nil {
		http.Error(w, "error sending response", http.StatusBadRequest)
	}
}
func StartServer(port int) error {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/order", Order)

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
