package solution

import (
	"context"
	"strings"

	log "github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

func Day02Solution(ctx context.Context, input []byte) {
	strategy := strings.Split(string(input), "\n")

	solution := funk.Reduce(strategy, func(acc int, x string) int {
		if x == "A X" {
			return 4 + acc
		} else if x == "A Y" {
			return 8 + acc
		} else if x == "A Z" {
			return 3 + acc
		} else if x == "B X" {
			return 1 + acc
		} else if x == "B Y" {
			return 5 + acc
		} else if x == "B Z" {
			return 9 + acc
		} else if x == "C X" {
			return 7 + acc
		} else if x == "C Y" {
			return 2 + acc
		} else if x == "C Z" {
			return 6 + acc
		}
		return acc
	}, 0).(int)
	log.Ctx(ctx).Info().Int("part", 1).Int("day", 2).Int("solution", solution).Msg("solved")

	solution = funk.Reduce(strategy, func(acc int, x string) int {
		if x == "A X" {
			return 3 + acc
		} else if x == "A Y" {
			return 4 + acc
		} else if x == "A Z" {
			return 8 + acc
		} else if x == "B X" {
			return 1 + acc
		} else if x == "B Y" {
			return 5 + acc
		} else if x == "B Z" {
			return 9 + acc
		} else if x == "C X" {
			return 2 + acc
		} else if x == "C Y" {
			return 6 + acc
		} else if x == "C Z" {
			return 7 + acc
		}
		return acc
	}, 0).(int)
	log.Ctx(ctx).Info().Int("part", 2).Int("day", 2).Int("solution", solution).Msg("solved")
}
