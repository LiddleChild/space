package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/LiddleChild/space/internal/config"
	"github.com/LiddleChild/space/internal/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list workspaces",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		spaces := config.AppConfig.GetSpaces()
		if len(spaces) == 0 {
			fmt.Println("no space created")
			os.Exit(0)
		}

		whitspaces := utils.AlignString(config.AppConfig.GetSpaceNames(), 4)

		for i, space := range spaces {
			fmt.Printf("%s%s%s\n", space.Name, strings.Repeat(" ", whitspaces[i]), space.Path)
		}
	},
}
