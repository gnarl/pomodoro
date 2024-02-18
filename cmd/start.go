package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start a new timer",
		Long:  `Start a new timer. You can specify the duration, name, and action.`,
		Run:   runStartCmd,
	}

	// Local flag definitions
	startCmd.Flags().IntP("duration", "d", 35, "The duration the timer runs in minutes")
	startCmd.Flags().StringP("name", "n", "", "The name of the timer")
	startCmd.Flags().StringP("message", "m", "Done!", "The action to perform when the timer ends")

	return startCmd
}

func runStartCmd(cmd *cobra.Command, args []string) {
	minutes, _ := cmd.Flags().GetInt("duration")
	seconds := minutes * 60
	fmt.Print("0 ")
	for i := 1; i < seconds; i++ {
		fmt.Print(".")
		if i%60 == 0 {
			fmt.Printf("\n%d ", i/60)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("")

	message, _ := cmd.Flags().GetString("message")
	name, _ := cmd.Flags().GetString("name")
	sendNotification(message, name)
}

func sendNotification(message, name string) {
	terminalNotifier := fmt.Sprintf("terminal-notifier -message \"%s\" -sound Glass -title \"%s\"", message, name)
	notifyCmd := exec.Command("sh", "-c", terminalNotifier)
	err := notifyCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
