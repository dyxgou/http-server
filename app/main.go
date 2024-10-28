package main

import "log"

func main() {
	server := CreateServer(Config{
		ListenAddr: ":4221",
	})

	log.Fatal(server.Start())
}
