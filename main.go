/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"os"

	"github.com/jatsrt/aoc2022/cmd"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	ctx = log.With().Str("app", "aoc2022").Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout}).WithContext(ctx)

	cmd.Execute(ctx)
}
