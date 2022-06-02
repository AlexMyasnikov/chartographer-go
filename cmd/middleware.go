package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func setupDir() {
	path := os.Args[1]
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create dir")
	}
	err = os.Chdir(path)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to change dir")
	}
	log.Print("Dir path: ", path)
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Str("error", fmt.Sprintf("%+v", err)).Msg("PANIC")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
