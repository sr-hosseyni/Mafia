package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	mafiaRoom := getMafiaRoom()

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println(err)
		} else {
			mafiaRoom.joins <- *ws
		}
	})

	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		w := bufio.NewWriter(writer)

		onlineUsers := getMafiaRoom().onlineUsers()

		log.Println(onlineUsers)
		rsp, err := json.Marshal(onlineUsers)
		if err != nil {
			log.Println(err)
		}
		w.Write(rsp)
		w.Flush()
	})

	log.Println("http server started on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
