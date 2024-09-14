package rm

import (
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/spf13/cobra"
)

var RmCmd = &cobra.Command{
	Use:   "rm <space>",
	Short: "remove existing workspaces",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := config.AppConfig.RemoveSpace(name)
		cobra.CheckErr(err)

		fmt.Println(name)
	},
}
