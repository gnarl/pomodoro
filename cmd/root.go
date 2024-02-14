package cmd

import (
	"fmt"
	"os"

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
	fmt.Println("root Execute")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root init")

	// Persistent flags are displayed as global flags in help output.
	rootCmd.PersistentFlags().StringP("persist", "p", "PersistentDefault", "A persistent flag")

	// Local flag definitions
	rootCmd.Flags().BoolP("log", "l", false, "A boolean flag with shortcut 'l'")
	rootCmd.Flags().Bool("mog", false, "A boolean flag without a shortcut")

	// Add subcommands to rootCmd
	rootCmd.AddCommand(startCmd)
}
