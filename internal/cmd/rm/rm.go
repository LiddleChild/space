package rm

import (
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/spf13/cobra"
)

var RmCmd = &cobra.Command{
	Use:   "rm <space>",
	Short: "remove existing workspaces",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return config.AppConfig.GetSpaceNames(), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := config.AppConfig.RemoveSpace(name)
		cobra.CheckErr(err)

		fmt.Println(name)
	},
}
