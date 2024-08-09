package main

import (
	"fmt"
	"log"
	"net"
)

type connection struct {
	server string
	port   int
	handle net.Conn
	buffer []byte
}

func (c *connection) connect() {
	var err error
	c.buffer = make([]byte, 1024)
	c.handle, err = net.Dial("tcp", c.server)
	if err != nil {
		fmt.Println(err)
	}

	c.handle.Write([]byte("PASS 34223424\r\n"))
	c.handle.Write([]byte("USER Gobottest * * : gobosttest4ever\r\n"))
	c.handle.Write([]byte("NICK gobottestnick\r\n"))
}

func (c *connection) read() (string, error) {
	//var msg []byte
	_, err := c.handle.Read(c.buffer)

	return string(c.buffer), err
}

func (c *connection) write(msg string) {
	_, err := c.handle.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *connection) disconnect() error {
	err := c.handle.Close()
	return err
}
