package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/LiddleChild/space/internal/cmd/create"
	"github.com/LiddleChild/space/internal/cmd/list"
	"github.com/LiddleChild/space/internal/cmd/open"
	"github.com/LiddleChild/space/internal/cmd/rm"
	"github.com/LiddleChild/space/internal/cmd/version"
	"github.com/LiddleChild/space/internal/config"
	"github.com/LiddleChild/space/internal/utils"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "space",
	Short: "(work)space is a workspace manager",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.EnsureStartUpScript()
		cobra.CheckErr(err)

		spaces := config.AppConfig.GetSpaces()
		whitespaces := utils.AlignString(config.AppConfig.GetSpaceNames(), 1)

		idx, err := fuzzyfinder.Find(spaces, func(i int) string {
			return fmt.Sprintf("%s%s%s", spaces[i].Name, strings.Repeat(" ", whitespaces[i]), spaces[i].Path)
		})

		if errors.Is(err, fuzzyfinder.ErrAbort) {
			os.Exit(0)
		} else if err != nil {
			cobra.CheckErr(err)
		}

		space := spaces[idx]
		space.LastOpened = time.Now()
		err = config.AppConfig.Save()
		cobra.CheckErr(err)

		if utils.IsInSession() {
			err = utils.Goto(space.Path)
			cobra.CheckErr(err)
		} else {
			err = utils.NewSession(space.Path)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd, create.CreateCmd, list.ListCmd, rm.RmCmd, open.OpenCmd, version.VersionCmd)
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
