package solution

import (
	"context"
	"os"

	log "github.com/rs/zerolog/log"
)

func Solutions(ctx context.Context) {
	log.Ctx(ctx).Info().Msg("solving")

	bytes, err := os.ReadFile("./input/day01")
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("")
		os.Exit(1)
	}
	Day01Solution(ctx, bytes)

	bytes, err = os.ReadFile("./input/day02")
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("")
		os.Exit(1)
	}
	Day02Solution(ctx, bytes)

	bytes, err = os.ReadFile("./input/day03")
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("")
		os.Exit(1)
	}
	Day03Solution(ctx, bytes)
}
