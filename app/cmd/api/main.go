package main

import (
	"flag"
	"os"

	"ktolstikhin/damn/internal/logger"
	"ktolstikhin/damn/internal/server"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8000", "server address")
	flag.Parse()

	log := logger.New(os.Stdout, logger.LevelAll, true)
	srv := server.New(addr, log)
	log.Info("starting server on %s", addr)

	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("server stopped")
}
