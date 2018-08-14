package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		fmt.Println("Someone tried to open the websocket with a regular request :(")
		return
	}

	go connection(conn)
}

func connection(conn *websocket.Conn) {
	// Annonce
	fmt.Println("Connection opened")
	defer fmt.Println("Connection closed")
	defer conn.Close()

	// Loop
	for {
		if err := conn.WriteJSON(time.Now()); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
