package main

import (
	"os"
)
import "github.com/rs/zerolog"
import "github.com/rs/zerolog/log"

func main() {
	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("service_name", "domain").Str("version", os.Getenv("VERSION"))
	})
	log.Info().Msg("starting up...")
}
