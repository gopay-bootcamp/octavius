// Package start is used to turn on the controller
package start

import (
	"octavius/internal/controller/server"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  `Start server at AppPort`,
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Start()
			if err != nil {
				log.Error(err, "execution error in start server")
			}
		},
	}
}
