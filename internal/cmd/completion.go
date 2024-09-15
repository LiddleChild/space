package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(space completion bash)

# To load completions for each session, execute once:
Linux:
  $ space completion bash > /etc/bash_completion.d/space
MacOS:
  $ space completion bash > /usr/local/etc/bash_completion.d/space

Zsh:

$ source <(space completion zsh)

# To load completions for each session, execute once:
$ space completion zsh > "${fpath[1]}/_yourprogram"

Fish:

$ space completion fish | source

# To load completions for each session, execute once:
$ space completion fish > ~/.config/fish/completions/space.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
