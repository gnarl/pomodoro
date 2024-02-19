package start

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gnarl/pomodoro/data"
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
	startCmd.Flags().StringP("task", "t", "", "The name of the task")
	startCmd.Flags().StringP("message", "m", "Done!", "A message to display when the timer is done")

	return startCmd
}

func runStartCmd(cmd *cobra.Command, args []string) {
	timer := persistTimer(cmd)
	seconds := timer.Duration * 60

	fmt.Print("0 ")
	for i := 1; i < seconds; i++ {
		fmt.Print(".")
		if i%60 == 0 {
			fmt.Printf("\n%d ", i/60)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("")

	sendNotification(timer.Message, timer.Task)
}

func persistTimer(cmd *cobra.Command) *data.Timer {
	minutes, _ := cmd.Flags().GetInt("duration")
	message, _ := cmd.Flags().GetString("message")
	task, _ := cmd.Flags().GetString("task")
	timer := data.NewTimer(task, minutes, message)
	data.AppendTimer(timer)
	return &timer
}

func sendNotification(message, task string) {
	terminalNotifier := fmt.Sprintf("terminal-notifier -message \"%s\" -sound Glass -title \"%s\"", message, task)
	notifyCmd := exec.Command("sh", "-c", terminalNotifier)
	err := notifyCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
