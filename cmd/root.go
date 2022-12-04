/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"time"

	"github.com/jatsrt/aoc2022/solution"
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
			solution.Solutions(cmd.Context())
		},
	}
)

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
