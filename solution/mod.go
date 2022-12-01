package solution

import (
	"context"
	"os"

	log "github.com/rs/zerolog/log"
)

func Solution(ctx context.Context, day int, input []byte) {
	log.Ctx(ctx).Info().Int("day", day).Msg("solving")

	switch day {
	case 1:
		Day01Solution(ctx, day, input)
	default:
		log.Ctx(ctx).Error().Int("day", day).Msg("not yet solved")
		os.Exit(1)
	}
}
