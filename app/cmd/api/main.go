package main

import (
	"flag"
	"os"

	"ktolstikhin/damn/internal/logger"
	"ktolstikhin/damn/internal/server"

	"github.com/rs/zerolog"
)

func main() {
	var (
		addr    string
		debug   bool
		jsonlog bool
	)
	flag.StringVar(&addr, "addr", ":8000", "server address")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&jsonlog, "jsonlog", false, "json logs")
	flag.Parse()

	level := zerolog.InfoLevel
	if debug {
		level = zerolog.DebugLevel
	}

	log := logger.New(os.Stdout, level, jsonlog)
	srv := server.New(addr, log)

	log.Info().Msgf("Starting server on %s", addr)

	err := srv.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}

	log.Info().Msg("Server stopped")
}
