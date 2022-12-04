package solution

import (
	"context"
	"strings"
	"unicode"

	"github.com/juliangruber/go-intersect"
	log "github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

func Day03Solution(ctx context.Context, input []byte) {
	sacks := strings.Split(string(input), "\n")

	compartments := funk.Map(sacks, func(sack string) [2][]rune { return [2][]rune{[]rune(sack[:len(sack)/2]), []rune(sack[len(sack)/2:])} }).([][2][]rune)
	common := funk.Map(compartments, func(x [2][]rune) int32 { return intersect.Hash(x[0], x[1])[0].(int32) }).([]int32)
	solution := funk.Reduce(common, func(acc int32, c int32) int32 { return char_val(c) + acc }, 0).(int32)

	log.Ctx(ctx).Info().Int("part", 1).Int("day", 3).Int32("solution", solution).Msg("solved")

	solution = funk.Reduce(funk.Chunk(sacks, 3), func(acc int32, group []string) int32 {
		return char_val(int32(intersect.Hash(group[0], intersect.Hash(group[1], group[2]))[0].(uint8))) + acc
	}, 0).(int32)

	log.Ctx(ctx).Info().Int("part", 2).Int("day", 3).Int32("solution", solution).Msg("solved")
}

func char_val(c int32) int32 {
	if unicode.IsUpper(c) {
		return c - 38
	} else {
		return c - 96
	}
}

// defmodule Aoc2022elixir.Day03 do
//   require Logger

//   def run(contents) do
//     sacks = contents |> String.split("\n")

//     priority_sum =
//       Enum.map(sacks, fn sack ->
//         String.codepoints(sack)
//         |> Enum.split(div(String.length(sack), 2))
//         |> Tuple.to_list()
//         |> Enum.map(&MapSet.new/1)
//         |> Enum.reduce(&MapSet.intersection/2)
//         |> Enum.join("")
//         |> alpha
//       end)
//       |> Enum.sum()

//     Logger.info("solved", solution: priority_sum, part: :one)

//     priority_sum =
//       Enum.map(sacks, &String.codepoints/1)
//       |> Enum.map(&MapSet.new/1)
//       |> Enum.chunk_every(3)
//       |> Enum.map(fn sackset ->
//         Enum.reduce(sackset, &MapSet.intersection/2) |> Enum.join("") |> alpha
//       end)
//       |> Enum.sum()

//     Logger.info("solved", solution: priority_sum, part: :two)
//   end

//   def alpha(<<c::8>>) when c in ?a..?z, do: c - 96
//   def alpha(<<c::8>>) when c in ?A..?Z, do: c - 38
// end
