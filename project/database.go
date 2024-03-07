package example

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有CORS请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("Received: %s\n", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func Server() {
	http.HandleFunc("/echo", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
