package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"time"
)

func main() {
	ListExample()
}

func ListExample() {

	c, _ := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	defer c.Close()
	// Create list pushing onto left side
	c.Cmd("LPUSH", "mylist", "a")
	c.Cmd("LPUSH", "mylist", "b")
	c.Cmd("LPUSH", "mylist", "c")

	// Print out list -> [c b a]
	list := c.Cmd("LRANGE", "mylist", "0", "2")
	fmt.Println(list)

	// DEL - clear list
	c.Cmd("DEL", "mylist")

	// Create list pushing onto right side
	c.Cmd("RPUSH", "mylist", "a")
	c.Cmd("RPUSH", "mylist", "b")
	c.Cmd("RPUSH", "mylist", "c")

	// Print out list -> [a b c]
	list = c.Cmd("LRANGE", "mylist", "0", "2")
	fmt.Println(list)
}
