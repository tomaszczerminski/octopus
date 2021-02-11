package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"octopus/pkg/api"
	"os"
)

func main() {
	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("service_name", "domain").Str("version", os.Getenv("VERSION"))
	})
	log.Info().Msg("starting up...")
	r := api.Router()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Fatal().Err(err).Caller().Send()
	}
}
