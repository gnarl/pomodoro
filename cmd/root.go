package cmd

import (
	"os"

	"github.com/gnarl/pomodoro/cmd/add"
	"github.com/gnarl/pomodoro/cmd/show"
	"github.com/gnarl/pomodoro/cmd/start"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomodoro",
	Short: "A simple pomodoro timer for the command line.",
	Long: `A simple pomodoro timer for the command line.
You can start a timer, stop a timer, and list today's timers.
You can also add and remove a timer from the list.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Add subcommands to rootCmd
	rootCmd.AddCommand(start.NewStartCmd())
	rootCmd.AddCommand(show.NewShowCmd())
	rootCmd.AddCommand(add.NewAddCmd())
}
