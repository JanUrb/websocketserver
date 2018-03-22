package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type client struct {
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024} //research this

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error while upgrading http to ws: ", err)
			return
		}
		log.Printf("Registering client: address: %s\t", conn.RemoteAddr().String())
		register(conn)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func register(conn *websocket.Conn) {
	client := &client{
		Conn: conn,
	}
	go client.listen()
}

func (c *client) listen() {
	for {
		msgType, b, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading from connection: ", err)
			break
		}

		log.Println("Message type: ", msgType)

		log.Printf("%s:\t%s", c.Conn.RemoteAddr(), b)

	}
}
