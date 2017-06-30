package main

import (
	"fmt"
	"github.com/wfxiang08/go-zookeeper/zk"
	"time"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second, 30) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
