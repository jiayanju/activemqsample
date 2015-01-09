package main

import (
	"fmt"
	"gopkg.in/stomp.v1"
)

func main() {
	conn, _ := stomp.Dial("tcp", "localhost:61613", stomp.Options{})
	fmt.Println("Consumer")
	sub, _ := conn.Subscribe("/queue/test", stomp.AckAuto)
	msg := <-sub.C

	fmt.Println(string(msg.Body))

	conn.Disconnect()
}
