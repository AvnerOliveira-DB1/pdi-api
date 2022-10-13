package main

import "github.com/avner.oliveira/pdi-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
