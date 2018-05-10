package cmd

import (
	"fmt"

	"github.com/automationbroker/bundle-lib/registries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Configure registry adapters",
	Long:  `List, Add, or Delete registry adapters from configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		listRegistries()
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
}

func updateCachedRegistries(registries []registries.Config) error {
	viper.Set("Registries", registries)
	viper.WriteConfig()
	return nil
}

func addRegistry() {
}

func listRegistries() {
	var registries []*registries.Config = nil
	err := viper.UnmarshalKey("Registries", &registries)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return
	}
	if len(registries) > 0 {
		fmt.Println("Found registries already in config")
		for _, r := range registries {
			fmt.Printf("%v - %v\n", r.Name, r.URL)
		}
		return
	}
}
