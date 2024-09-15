package cmd

import (
	"github.com/LiddleChild/space/internal/cmd/create"
	"github.com/LiddleChild/space/internal/cmd/list"
	"github.com/LiddleChild/space/internal/cmd/open"
	"github.com/LiddleChild/space/internal/cmd/rm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
}

func init() {
	rootCmd.AddCommand(completionCmd, create.CreateCmd, list.ListCmd, rm.RmCmd, open.OpenCmd)
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
