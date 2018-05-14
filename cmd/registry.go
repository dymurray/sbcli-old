package cmd

import (
	"fmt"

	"github.com/automationbroker/bundle-lib/registries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registryConfig registries.Config
var whitelist string

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Configure registry adapters",
	Long:  `List, Add, or Delete registry adapters from configuration`,
}

var registryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new registry adapter",
	Long:  `Add a new registry adapter to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		addRegistry()
	},
}

var registryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the configured registry adapters",
	Long:  `List all registry adapters in the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		listRegistries()
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
	registryAddCmd.Flags().StringVar(&registryConfig.Type, "type", "dockerhub", "Type of registry adapter to add")
	registryAddCmd.Flags().StringVar(&registryConfig.Org, "org", "ansibleplaybookbundle", "Type of registry adapter to add")
	registryAddCmd.Flags().StringVar(&registryConfig.URL, "url", "docker.io", "URL of registry adapter to add")
	registryAddCmd.Flags().StringVar(&registryConfig.Name, "name", "docker", "Name of registry adapter to add")
	registryAddCmd.Flags().StringVar(&whitelist, "whitelist", ".*-apb$", "Whitelist for configuration of registry adapter")
	registryConfig.WhiteList = append(registryConfig.WhiteList, whitelist)

	registryCmd.AddCommand(registryAddCmd)
	registryCmd.AddCommand(registryListCmd)
}

func updateCachedRegistries(registries []registries.Config) error {
	viper.Set("Registries", registries)
	viper.WriteConfig()
	return nil
}

func addRegistry() {
	//	reg, err := registries.NewRegistry(registryConfig, "ansible-service-broker")
	//	if err != nil {
	//		fmt.Printf("Error creating new registry adapter: %v", err)
	//		return
	//	}
	var regList []registries.Config
	err := viper.UnmarshalKey("Registries", &regList)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return
	}

	regList = append(regList, registryConfig)
	updateCachedRegistries(regList)
	return
}

func listRegistries() {
	var regList []registries.Config
	err := viper.UnmarshalKey("Registries", &regList)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return
	}
	if len(regList) > 0 {
		fmt.Println("Found registries already in config:")
		for _, r := range regList {
			fmt.Printf("name: %v - type: %v - URL: %v\n", r.Name, r.Type, r.URL)
		}
	} else {
		fmt.Println("Found no registries in configuration. Try `sbcli registry add`.")
	}
	return

}
