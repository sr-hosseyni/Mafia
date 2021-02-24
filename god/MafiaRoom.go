package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var (
	SingletonMafiaRoom *MafiaRoom
	once               sync.Once
)

func getMafiaRoom() *MafiaRoom {
	once.Do(func() {
		SingletonMafiaRoom = newMafiaRoom()
	})

	return SingletonMafiaRoom
}

type MafiaRoom struct {
	clients  map[int]ClientInterface
	joins    chan websocket.Conn
	incoming chan InMessage
	outgoing chan InMessage
}

func (MafiaRoom *MafiaRoom) Broadcast(data OutMessage) {
	log.Println("broadcasting message", data, MafiaRoom.clients)
	for _, client := range MafiaRoom.clients {
		client.outbound(data)
	}
}

func (MafiaRoom *MafiaRoom) Send(data OutMessage) {
	log.Println("Sending message ", data)
	if MafiaRoom.clients[data.To] != nil {
		MafiaRoom.clients[data.To].outbound(data)
	} else {
		// @todo handle offline message
	}
}

func (MafiaRoom *MafiaRoom) webJoin(connection websocket.Conn) {
	fmt.Println("new web join")
	client := newWebSocketClient(connection)
	//MafiaRoom.clients = append(MafiaRoom.clients, client)
	go func() {
		for {
			MafiaRoom.incoming <- <-client.incoming
		}
	}()
}

func (MafiaRoom *MafiaRoom) listen() {
	go func() {
		fmt.Println("Mafia room is listening")
		for MafiaRoom.incoming != nil || MafiaRoom.joins != nil {
			select {
			case message, ok := <-MafiaRoom.incoming:
				if !ok {
					fmt.Println("Mafia room incoming channel is closed!!!")
					MafiaRoom.incoming = nil
					continue
				}
				log.Println(message)
				MafiaRoom.process(message)
			case conn, ok := <-MafiaRoom.joins:
				if !ok {
					fmt.Println("Mafia room ws join channel is closed!!!")
					MafiaRoom.joins = nil
					continue
				}
				MafiaRoom.webJoin(conn)
			}
		}
	}()
}

func (MafiaRoom *MafiaRoom) close() {
	close(MafiaRoom.incoming)
	close(MafiaRoom.joins)
	fmt.Println("Mafia room incoming channel is closing")
	fmt.Println("Mafia room tcp join channel is closing")
	fmt.Println("Mafia room ws join channel is closing")
}

func newMafiaRoom() *MafiaRoom {
	MafiaRoom := &MafiaRoom{
		clients:  make(map[int]ClientInterface, 0),
		joins:    make(chan websocket.Conn),
		incoming: make(chan InMessage),
		outgoing: make(chan InMessage),
	}

	MafiaRoom.listen()

	return MafiaRoom
}

func (MafiaRoom *MafiaRoom) process(message InMessage) {
	msg := &OutMessage{
		Status:        false,
		Message:       "",
		IsTrue:        false,
		IsNight:       false,
		AliveCitizens: UserList{},
		To:            0,
		CreatedAt:     time.Time{},
	}

	MafiaRoom.Send(*msg)
}

func (MafiaRoom *MafiaRoom) onlineUsers() UserList {
	onlineUsers := UserList{
		UserList: make([]IUser, len(MafiaRoom.clients)),
	}

	i := 0
	for userId := range MafiaRoom.clients {
		log.Println(userId, users[userId])
		onlineUsers.UserList[i] = users[userId].transform()
		i++
	}

	return onlineUsers
}
