package start

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gnarl/pomodoro/cmd/common"
	"github.com/gnarl/pomodoro/internal/data"
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start a new timer",
		Run:   runStartCmd,
	}

	common.SetTimerCmdFlags(startCmd)

	return startCmd
}

func runStartCmd(cmd *cobra.Command, args []string) {

	timer := common.GetTimerFromFlags(cmd)
	data.AppendTimer(timer)

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

func sendNotification(message, task string) {
	terminalNotifier := fmt.Sprintf("terminal-notifier -message \"%s\" -sound Glass -title \"%s\"", message, task)
	notifyCmd := exec.Command("sh", "-c", terminalNotifier)
	err := notifyCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
