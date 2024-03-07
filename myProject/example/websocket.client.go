package example

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
)

func Client() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	log.SetFlags(0)

	u := "ws://" + *addr + "/echo"
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 发送消息
	message := []byte("hello, my baby")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// 接收响应
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("received: %s", message)
}
