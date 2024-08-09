package main

const SERVER = "tngnet.nl.quakenet.org:6667"
const PORT = 6667

func main() {
	conn := connection{server: SERVER, port: PORT}
	conn.connect()

	conn.read()

	conn.disconnect()
}
