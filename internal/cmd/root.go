package cmd

import (
	"github.com/LiddleChild/space/internal/cmd/create"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
}

func init() {
	rootCmd.AddCommand(create.CreateCmd)
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
