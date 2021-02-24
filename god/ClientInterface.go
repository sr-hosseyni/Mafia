package main

type ClientInterface interface {
	read()
	write()
	listen()
	outbound(message OutMessage)
	inbound() InMessage
}