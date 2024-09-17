package open

import (
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/LiddleChild/space/internal/utils"
	"github.com/spf13/cobra"
)

// valid args function
var OpenCmd = &cobra.Command{
	Use:   "open <name>",
	Short: "change working directory to specific space in a new shell session",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return config.AppConfig.GetSpaceNames(), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		space, err := config.AppConfig.GetSpace(name)
		cobra.CheckErr(err)

		utils.Shell(fmt.Sprintf("SPACE_WD=%s", space.Path))
		cobra.CheckErr(err)
	},
}
