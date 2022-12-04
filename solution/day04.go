package solution

import (
	"context"
	"strconv"
	"strings"

	log "github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

func Day04Solution(ctx context.Context, input []byte) {
	ranges := funk.Map(strings.Split(string(input), "\n"), func(pairs string) [][]int {
		return funk.Map(strings.Split(pairs, ","), func(pair string) []int {
			things := strings.Split(pair, "-")
			lo, _ := strconv.Atoi(things[0])
			hi, _ := strconv.Atoi(things[1])
			return create_range(lo, hi)
		}).([][]int)
	})

	intersections := funk.Map(ranges, func(xy [][]int) [2]int {
		interesection := funk.Join(xy[0], xy[1], funk.InnerJoin).([]int)
		return [2]int{len(interesection), min(len(xy[0]), len(xy[1]))}
	})

	solution := funk.Reduce(intersections, func(acc int, v [2]int) int {
		if v[0] >= v[1] {
			return acc + 1
		} else {
			return acc
		}
	}, 0).(int)
	log.Ctx(ctx).Info().Int("part", 1).Int("day", 4).Int("solution", solution).Msg("solved")

	solution = funk.Reduce(intersections, func(acc int, v [2]int) int {
		if v[0] > 0 {
			return acc + 1
		} else {
			return acc
		}
	}, 0).(int)
	log.Ctx(ctx).Info().Int("part", 2).Int("day", 4).Int("solution", solution).Msg("solved")
}

func create_range(lo int, hi int) []int {
	s := make([]int, hi-lo+1)
	for i := range s {
		s[i] = i + lo
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
