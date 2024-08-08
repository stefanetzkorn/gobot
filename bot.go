package main

import (
	"fmt"
	"log"
)

const SERVER = "tngnet.nl.quakenet.org:6667"
const PORT = 6667

func main() {
	c := connection{server: SERVER, port: PORT}
	c.connect()

	msg, err := c.read()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	c.disconnect()
}
