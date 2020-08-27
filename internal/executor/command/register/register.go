package register

import (
	"errors"
	"fmt"
	"octavius/internal/executor/daemon"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(executorDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "register Executor",
		Long:  `Registers Executor in Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			err := executorDaemon.StartClient()
			if err != nil {
				log.Error(fmt.Errorf("executor configuration failed, error details: %v", err.Error()), "")
			}
			registered, err := executorDaemon.RegisterClient()
			if err != nil {
				log.Error(fmt.Errorf("registration failed, error details: %v", err.Error()), "")
				return
			}
			if registered {
				log.Info("successfully registered executor in controller.")
				return
			}
			log.Error(errors.New(fmt.Sprintln("registration blocked by control plane.")), "")
		},
	}
}
