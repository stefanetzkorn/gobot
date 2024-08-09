package main

import (
	"log"
)

const SERVER = "tngnet.nl.quakenet.org:6667"
const PORT = 6667

func main() {
	conn := connection{server: SERVER, port: PORT}
	controller := Controller{&conn}
	conn.connect()

	for {
		msg, err := conn.read()
		if err != nil {
			log.Fatal(err)
			break
		}
		controller.parse(msg)
	}

	conn.disconnect()
}
