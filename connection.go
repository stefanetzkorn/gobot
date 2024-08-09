package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type connection struct {
	server  string
	port    int
	conn    net.Conn
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func (c *connection) connect() {
	var err error
	c.conn, err = net.Dial("tcp", c.server)
	c.scanner = bufio.NewScanner(c.conn)
	c.writer = bufio.NewWriter(c.conn)

	if err != nil {
		fmt.Println(err)
	}

	c.writer.Write([]byte("PASS 34223424\r\n"))
	c.writer.Write([]byte("USER Gobottest * * : gobosttest4ever\r\n"))
	c.writer.Write([]byte("NICK gobottestnick\r\n"))
	c.writer.Flush()
}

func (c *connection) read() {
	controller := Controller{c}
	for c.scanner.Scan() {
		controller.parse(c.scanner.Text())
	}
}

func (c *connection) write(msg string) {
	_, err := c.writer.Write([]byte(msg))
	c.writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *connection) disconnect() error {
	err := c.conn.Close()
	return err
}
