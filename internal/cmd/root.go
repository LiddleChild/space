package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
