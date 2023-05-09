package main

import (
	"flag"
	"log"

	"ktolstikhin/damn/internal/server"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8000", "server address")
	flag.Parse()

	srv := server.New(addr)
	log.Printf("starting server on %s", addr)

	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
