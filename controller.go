package main

import (
	"fmt"
	"strings"
)

type Controller struct {
	*connection
}

func (c *Controller) parse(msg string) {
	switch {
	case strings.HasPrefix(msg, "PING"):
		c.pingpong(msg)
	}
}

func (c *Controller) pingpong(ping string) {
	split := strings.Split(ping, ":")
	msg := fmt.Sprintf("PONG :%s\r\n", split[1])
	fmt.Println(msg)
	c.connection.write(msg)
}
