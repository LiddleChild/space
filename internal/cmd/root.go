package cmd

import (
	"fmt"

	"github.com/LiddleChild/space/internal/cmd/create"
	"github.com/LiddleChild/space/internal/cmd/list"
	"github.com/LiddleChild/space/internal/cmd/open"
	"github.com/LiddleChild/space/internal/cmd/rm"
	"github.com/LiddleChild/space/internal/config"
	"github.com/LiddleChild/space/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
	Run: func(cmd *cobra.Command, args []string) {
		selection := promptui.Select{
			Label: "space",
			Items: config.AppConfig.GetSpaceNames(),
		}

		_, result, err := selection.Run()
		cobra.CheckErr(err)

		space, err := config.AppConfig.GetSpace(result)
		cobra.CheckErr(err)

		err = utils.Shell(fmt.Sprintf("SPACE_WD=%s", space.Path))
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd, create.CreateCmd, list.ListCmd, rm.RmCmd, open.OpenCmd)
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
