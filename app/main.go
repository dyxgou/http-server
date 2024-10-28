package main

import "log"

func main() {
	server := CreateServer(Config{
		ListenAddr: ":3000",
	})

	log.Fatal(server.Start())
}
