package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	// sand variable of type chan (Channel) expecting a Message
	send chan Message
}

// Method write receives a pointer to a Client
// This syntax is called 'receiver' o que faz isso um metodo do Client recebido
func (client *Client) write() {
	for msg := range client.send {
		// TODO: socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}
}

func (client *Client) subscribeChannels() {
	// Todo: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages() {
	// Todo: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}

func main() {
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}
