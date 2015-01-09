package main

import (
	"gopkg.in/stomp.v1"
)

func main() {
	// conn, err := stomp.Dial("tcp", "localhost:61613", stomp.Options{})
	// if err != nil {

	// }

	// err = conn.Send("/queue/test", "", []byte("Test Message 1"), nil)
	// if err != nil {

	// }

	// conn.Disconnect()

	conn, _ := stomp.Dial("tcp", "localhost:61613", stomp.Options{})

	conn.Send("/queue/test", "", []byte("Simples Assim"), nil)

	conn.Disconnect()
}
