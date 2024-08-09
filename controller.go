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
	fmt.Println(msg)
}

func (c *Controller) pingpong(ping string) {
	split := strings.Split(ping, ":")
	c.connection.write(fmt.Sprintf("PONG :%s", split[1]))
}
