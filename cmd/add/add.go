package add

import (
	"github.com/gnarl/pomodoro/data"
	"github.com/gnarl/pomodoro/utils"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a favorite timer",
		Run:   runAddCmd,
	}

	data.SetTimerCmdFlags(addCmd)
	addCmd.Flags().StringP("name", "n", "", "The name of the favorite")
	addCmd.MarkFlagRequired("name")

	return addCmd
}

func runAddCmd(cmd *cobra.Command, args []string) {

	favoriteTimer := data.GetFavoriteTimerFromFlags(cmd)
	log := utils.GetLogger()
	log.Debug("Adding favorite timer: ", "favoriteTimer", favoriteTimer)

	currentFavorites := data.ReadFavorites()
	currentFavorites = append(currentFavorites, *favoriteTimer)
	data.WriteFavorites(currentFavorites)
}
