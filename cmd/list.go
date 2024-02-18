package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gnarl/pomodoro/data"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all timers",
		Long:  "List all timers",
		Run:   runListCmd,
	}

	return listCmd
}

func runListCmd(cmd *cobra.Command, args []string) {
	timers := data.ReadTimers()
	for _, timer := range timers {
		t, err := json.MarshalIndent(timer, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(t))
	}
}
