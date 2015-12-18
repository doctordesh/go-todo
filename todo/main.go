package main

import (
	"github.com/doctordesh/todo"
)

func main() {
	controller := todo.NewController("/Users/emil/.todo.json")
	controller.Run()
}
