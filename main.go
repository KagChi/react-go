package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"react-go/cmd"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	if !fiber.IsChild() {
		log.Info().Msg("🚀 Starting server on :8000")
	}
	cmd.Init()
}
