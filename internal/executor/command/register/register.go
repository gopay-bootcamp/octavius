package register

import (
	"errors"
	"fmt"
	"octavius/internal/executor/daemon"
	"octavius/internal/executor/logger"

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
				logger.Error(fmt.Errorf("executor configuration failed, error details: %v", err.Error()), "")
			}
			logger.Info("executor configured successfully")
			registered, err := executorDaemon.RegisterClient()
			if err != nil {
				logger.Error(fmt.Errorf("registration failed, error details: %v", err.Error()), "")
				return
			}
			if registered {
				logger.Info("successfully registered executor in controller.")
				return
			}
			logger.Error(errors.New(fmt.Sprintln("registration blocked by control plane.")), "")
		},
	}
}
