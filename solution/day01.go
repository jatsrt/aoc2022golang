package solution

import (
	"context"
	"sort"
	"strconv"
	"strings"

	log "github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

func Day01Solution(ctx context.Context, day int, input []byte) {
	caloric_values := funk.Map(strings.Split(string(input), "\n\n"), func(x string) int {
		return funk.SumInt(funk.Map(strings.Split(x, "\n"), func(x string) int {
			v, _ := strconv.Atoi(x)
			return v
		}).([]int))
	}).([]int)
	sort.Ints(caloric_values)

	solution := int64(caloric_values[len(caloric_values)-1])
	log.Ctx(ctx).Info().Int("part", 1).Int64("solution", solution).Msg("solved")
	solution = int64(funk.SumInt(caloric_values[len(caloric_values)-3:]))
	log.Ctx(ctx).Info().Int("part", 2).Int64("solution", solution).Msg("solved")
}
