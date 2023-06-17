package main

import "ratovia/go-clean-architecture-sample/app/src/server"

func main() {
	env := "development"
	port := "8081"
	server.RunServer(env, port)
}
