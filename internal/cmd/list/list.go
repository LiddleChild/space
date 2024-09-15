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
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		longestName := 0
		spaces := config.AppConfig.GetSpaces()
		for _, space := range spaces {
			longestName = max(longestName, len(space.Name))
		}

		for _, space := range spaces {
			fmt.Printf("%s%s%s\n", space.Name, strings.Repeat(" ", longestName-len(space.Name)+4), space.Path)
		}
	},
}
