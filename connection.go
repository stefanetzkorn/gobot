package main

import (
	"fmt"
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
}

func (c *connection) read() (string, error) {
	_, err := c.handle.Read(c.buffer)
	return string(c.buffer), err
}

func (c *connection) disconnect() error {
	err := c.handle.Close()
	return err
}
