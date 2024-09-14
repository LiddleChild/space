package create

import (
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "create workspace at working directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := config.AppConfig.CreateSpace(name)
		cobra.CheckErr(err)

		fmt.Println(name)
	},
}
