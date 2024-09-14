package list

import (
	"fmt"
	"strings"

	"github.com/LiddleChild/space/internal/config"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		longestName := 0
		for _, space := range config.AppConfig.GetSpaces() {
			longestName = max(longestName, len(space.Name))
		}

		for _, space := range config.AppConfig.GetSpaces() {
			fmt.Printf("%s%s%s\n", space.Name, strings.Repeat(" ", longestName-len(space.Name)+4), space.Path)
		}
	},
}
