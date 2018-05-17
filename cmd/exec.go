package cmd

import (
	"fmt"

	"github.com/automationbroker/bundle-lib/apb"
	//	"github.com/automationbroker/bundle-lib/registries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	log "github.com/sirupsen/logrus"
)

var execName string

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Perform specified action on Service Bundles",
	Long:  `Perform actions (Provision, Deprovision, Bind, Unbind) on Service Bundles`,
}

var execProvisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision ServiceBundle images",
	Long:  `Provision ServiceBundles from a registry adapter`,
	Run: func(cmd *cobra.Command, args []string) {
		createExecutor()
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execProvisionCmd.Flags().StringVarP(&execName, "name", "n", "", "Name of spec to provision")
	execCmd.AddCommand(execProvisionCmd)
}

func createExecutor() {
	exec := apb.NewExecutor()
	si := apb.ServiceInstance{}
	specs := []*apb.Spec{}
	viper.UnmarshalKey("Specs", &specs)
	for _, s := range specs {
		if s.FQName == execName {
			si.Spec = s
		}
	}
	fmt.Printf("Spec: %v", si.Spec)
	exec.ExecuteApb("provision", &si, nil)
}
