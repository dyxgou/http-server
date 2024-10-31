package main

import "log"

func main() {
	server := CreateServer(Config{
		ListenAddr: ":4221",
	})

	server.Router.Add("/")
	server.Router.Add("/user")
	server.Router.Add("/product")

	log.Fatal(server.Start())
}
