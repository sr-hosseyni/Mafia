package main

import (
	"time"
)

type InMessage struct {
	User      string    `json:"user"`
	Selection int       `json:"selection"`
	From      int       `json:"from"`
	CreatedAt time.Time `json:"createdAt"`
}

type OutMessage struct {
	Status        bool      `json:"status"`
	Message       string    `json:"message"`
	IsTrue        bool      `json:"isTrue"`
	IsNight       bool      `json:"isNight"`
	IsFinished    bool      `json:"isFinished"`
	AliveCitizens UserList  `json:"aliveCitizens"`
	To            int       `json:"to"`
	CreatedAt     time.Time `json:"createdAt"`
}
