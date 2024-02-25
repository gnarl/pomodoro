package add

import (
	"github.com/gnarl/pomodoro/cmd/common"
	"github.com/gnarl/pomodoro/internal/data"
	"github.com/spf13/cobra"

	log "github.com/gnarl/pomodoro/internal/utils"
)

func NewAddCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a favorite timer",
		Run:   runAddCmd,
	}

	common.SetTimerCmdFlags(addCmd)
	addCmd.Flags().StringP("name", "n", "", "The name of the favorite")
	addCmd.MarkFlagRequired("name")

	return addCmd
}

func runAddCmd(cmd *cobra.Command, args []string) {

	favoriteTimer := common.GetFavoriteTimerFromFlags(cmd)
	log.Logger.Debug("Adding favorite timer: ", "favoriteTimer", favoriteTimer)

	currentFavorites := data.ReadFavorites()
	currentFavorites = append(currentFavorites, *favoriteTimer)
	data.WriteFavorites(currentFavorites)
}
