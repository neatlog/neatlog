package main

import (
	"fmt"
	"github.com/neatlog/neatlog/server"
)

func main() {
	server.Start()
	fmt.Println("server started")
	if err := server.ListenAndServer(); err != nil {
		panic(err)
	}
}
