package main

import "github.com/neatlog/neatlog/server"

func main() {
	server.Start()
	if err := server.ListenAndServer(); err != nil {
		panic(err)
	}
}
