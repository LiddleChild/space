package version

import (
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "config version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.AppConfig.Version)
	},
}
