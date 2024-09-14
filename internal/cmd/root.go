package cmd

import (
	"github.com/LiddleChild/space/internal/cmd/create"
	"github.com/LiddleChild/space/internal/cmd/list"
	"github.com/LiddleChild/space/internal/cmd/rm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
}

func init() {
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(rm.RmCmd)
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
