package show

import (
	"encoding/json"
	"fmt"

	"github.com/gnarl/pomodoro/data"
	"github.com/gnarl/pomodoro/utils"
	"github.com/spf13/cobra"
)

func NewShowCmd() *cobra.Command {

	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show timers and favorites",
		Run:   runShowCmd,
	}

	// Local flag definitions
	showCmd.Flags().BoolP("favorites", "f", false, "Show favorite timers.")
	showCmd.Flags().BoolP("recent", "r", false, "Show recently run timers.")
	return showCmd
}

func runShowCmd(cmd *cobra.Command, args []string) {

	favorites, _ := cmd.Flags().GetBool("favorites")
	recent, _ := cmd.Flags().GetBool("recent")
	log := utils.GetLogger()
	log.Debug("runShowCmd: ", "favorites", favorites, " recent: ", recent)

	if !favorites {
		showTimers()
	} else {
		showFavorites()
		if recent {
			showTimers()
		}
	}

}

func showTimers() {
	// TODO: sort timers
	timers := data.ReadTimers()
	for _, timer := range timers {
		t, err := json.MarshalIndent(timer, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(t))
	}
}

func showFavorites() {
	// TODO: sort favorites
	favorites := data.ReadFavorites()
	for _, favorite := range favorites {
		f, err := json.MarshalIndent(favorite, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(f))
	}
}
