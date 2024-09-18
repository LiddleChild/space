package rm

import (
	"errors"
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/manifoldco/promptui"
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

		result, err := confirmationPrompt(name)
		cobra.CheckErr(err)

		if result {
			err = config.AppConfig.RemoveSpace(name)
			cobra.CheckErr(err)

			fmt.Println(name)
		}
	},
}

func confirmationPrompt(name string) (bool, error) {
	input := promptui.Prompt{
		Label: fmt.Sprintf("type \"%s\" for confirmation", name),
		Validate: func(input string) error {
			if input != name {
				return errors.New("space name does not match")
			}

			return nil
		},
	}

	result, err := input.Run()
	if err != nil {
		return false, err
	}

	if result == name {
		return true, nil
	}

	return false, nil
}
