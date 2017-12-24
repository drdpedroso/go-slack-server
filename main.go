package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Receive and http request and returns a boolean
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			outMessage = Message{"error", err}
			if err := socket.WriteJSON(outMessage); err != nil {
				fmt.Println(err)
				break
			}
		case "channel subscribe":
			go subscribeChannel(socket)
		}
	}
}

// Receives an empty interface and returns a Channel and a Error
func addChannel(data interface{}) error {
	var channel Channel
	// We tell go for a more specific type (type assertion). channelMap its a map with string keys and empty interfaces as values
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Println("add new channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add", Channel{"1", "Sofware Supp"}}
		fmt.Println("sent new channel")
		socket.WriteJSON(message)
	}
}
