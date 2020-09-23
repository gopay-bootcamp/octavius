package register

import (
	"errors"
	"fmt"
	"octavius/internal/executor/config"
	daemon "octavius/internal/executor/daemon/registration"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

//NewCmd returns register command
func NewCmd(executorDaemon daemon.RegistrationServicesClient, executorConfig config.OctaviusExecutorConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "register Executor",
		Long:  `Registers Executor in Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			registered, err := executorDaemon.RegisterClient(executorConfig)
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
