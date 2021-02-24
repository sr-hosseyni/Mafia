package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"time"
)

type WebSocketClient struct {
	incoming chan InMessage
	outgoing chan OutMessage
	ws       *websocket.Conn
}

func (client *WebSocketClient) read() {
	defer client.ws.Close()

	var user User
	err := client.ws.ReadJSON(&user)
	if err != nil {
		log.Printf("ERRRRROR: %v", err)
		return
	}

	if !users[usersNameIndex[user.Name]].checkPassword(user.Password) {
		fmt.Println("Invalid user [" + user.Name + "]'s password")
		return
	}

	//client.user = users[user.Name]
	getMafiaRoom().clients[users[usersNameIndex[user.Name]].Id] = client

	fmt.Println(user.Name + " joined")

	client.outbound(OutMessage{
		Status:        true,
		Message:       "You are in!",
		IsTrue:        rand.Int()%2 == 0,
		IsFinished:    false,
		IsNight:       false,
		AliveCitizens: getMafiaRoom().onlineUsers(),
		To:            0,
		CreatedAt:     time.Time{},
	})

	for {
		var msg InMessage
		// Read in a new message as JSON and map it to a InMessage object
		err := client.ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		msg.From = users[usersNameIndex[user.Name]].Id
		// Send the newly received message to the broadcast channel
		client.incoming <- msg
	}
}

func (client *WebSocketClient) write() {
	for data := range client.outgoing {
		err := client.ws.WriteJSON(data)
		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
		}
	}
}

func (client *WebSocketClient) listen() {
	go client.read()
	go client.write()
}

func (client *WebSocketClient) outbound(message OutMessage) {
	client.outgoing <- message
}

func (client *WebSocketClient) inbound() InMessage {
	return <-client.incoming
}

func newWebSocketClient(wsc websocket.Conn) *WebSocketClient {

	client := &WebSocketClient{
		incoming: make(chan InMessage),
		outgoing: make(chan OutMessage),
		ws:       &wsc,
	}

	client.listen()

	return client
}
