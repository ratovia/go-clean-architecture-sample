package main

import "ratovia/go-clean-architecture-sample/app/src/server"

func main() {
	env := "development"
	server.RunServer(env)
}
