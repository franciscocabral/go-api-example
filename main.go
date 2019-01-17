package main

import "github.com/franciscocabral/go-api-example/server"

func main() {
	apiServer := server.Setup()
	apiServer.Run()
}
