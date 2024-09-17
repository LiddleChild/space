package create

import (
	"errors"
	"fmt"

	"github.com/LiddleChild/space/internal/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "create workspace at working directory",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var name string

		var err error
		if len(args) == 0 {
			name, err = inputPrompt()
			cobra.CheckErr(err)
		} else {
			name = args[0]
		}

		err = config.AppConfig.CreateSpace(name)
		cobra.CheckErr(err)

		fmt.Println(name)
	},
}

func inputPrompt() (string, error) {
	input := promptui.Prompt{
		Label: "space name",
		Validate: func(input string) error {
			_, err := config.AppConfig.GetSpace(input)
			if err != nil {
				return nil
			}
			return errors.New("space name already exists")
		},
	}

	result, err := input.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
