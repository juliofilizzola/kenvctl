/*
Copyright © 2026
*/
package main

import (
	"os"

	"github.com/juliofilizzola/kenvctl/internal/command"
	"github.com/juliofilizzola/kenvctl/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kenvctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	utils.TimeSleep()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(command.ListCmd)
	rootCmd.AddCommand(command.AddCmd)
	rootCmd.AddCommand(command.RemoveCmd)
	rootCmd.AddCommand(command.SetCmd)
	rootCmd.AddCommand(command.NewEnvCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
