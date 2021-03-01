package main

import (
	"fmt"
	"github.com/bensema/library/log"
)

func main() {
	c := &log.LumberjackConfig{
		Dir: "./log",
	}
	t := log.NewLbj(c, "info")
	fmt.Println(t.Compress)
}
