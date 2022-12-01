/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jatsrt/aoc2022/solution"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	currentDay = time.Now().Day()
	rootCmd    = &cobra.Command{
		Use:   "aoc2022",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			input, _ := cmd.Flags().GetString("input")
			if input == "" {
				log.Ctx(cmd.Context()).Error().Msg("input required")
				os.Exit(1)
			}
			day, _ := cmd.Flags().GetInt("day")

			bytes, err := os.ReadFile(input)
			if err != nil {
				log.Ctx(cmd.Context()).Error().Err(err).Msg("")
				os.Exit(1)
			}

			solution.Solution(cmd.Context(), day, bytes)
		},
	}
)

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int("day", currentDay, "The day to run")
	rootCmd.PersistentFlags().String("input", fmt.Sprintf("./input/day%02d", currentDay), "The input data")
	rootCmd.MarkFlagRequired("day")
	rootCmd.MarkFlagRequired("input")
}
