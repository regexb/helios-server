package helios

import (
	"fmt"
	"log"

	"github.com/googollee/go-socket.io"
)

func initSocket() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatalf("Error on socket.io server", err.Error())
		return nil
	}

	server.On("connection", func(so socketio.Socket) {
		fmt.Printf("New socket.io connection: %s", so.Id())
		so.Join("helios")
		so.On("disconnection", func() {
			// no op
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Fatalf("Error on socket.io server", err.Error())
	})

	return server
}

func (h *Engine) NewBroadcastChannel(room, message string) {
	go func() {
		for {
			msg := <-h.SocketChan
			fmt.Println("Got message to broadcast")
			h.Socket.BroadcastTo(room, message, msg)
		}
	}()
}