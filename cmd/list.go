package cmd

import (
	"fmt"

	"github.com/automationbroker/bundle-lib/apb"
	"github.com/automationbroker/bundle-lib/registries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ServiceBundle images",
	Long:  `List ServiceBundles from a registry adapter`,
	Run: func(cmd *cobra.Command, args []string) {
		listImages()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func updateCachedList(specs []*apb.Spec) error {
	viper.Set("Specs", specs)
	viper.WriteConfig()
	return nil
}

func getImages() ([]*apb.Spec, error) {
	var regConfigList []registries.Config
	var regList []registries.Registry
	var specList []*apb.Spec
	err := viper.UnmarshalKey("Registries", &regConfigList)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return nil, err
	}

	authNamespace := ""
	for _, config := range regConfigList {
		registry, err := registries.NewRegistry(config, authNamespace)
		if err != nil {
			log.Error("Error from creating a NewRegistry")
			log.Error(err)
			return nil, err
		}
		regList = append(regList, registry)
	}
	for _, reg := range regList {
		specs, count, err := reg.LoadSpecs()
		if err != nil {
			log.Errorf("registry: %v was unable to complete bootstrap - %v",
				reg.RegistryName(), err)
			return nil, err
		}
		log.Infof("Registry %v has %d bundles available from %d images scanned", reg.RegistryName(), len(specs), count)
		specList = specs
	}
	fmt.Printf("Success! %v", specList)

	return specList, nil
}

func listImages() {
	var specs []*apb.Spec
	err := viper.UnmarshalKey("Specs", &specs)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return
	}
	if len(specs) > 0 {
		fmt.Println("Found specs already in config")
		for _, s := range specs {
			fmt.Printf("%v - %v - %v\n", s.FQName, s.Description, s.Image)
		}
		return
	}

	specs, err = getImages()
	if err != nil {
		fmt.Println("Error getting images")
		return
	}
	fmt.Printf("specs: %v", specs)
	err = updateCachedList(specs)
	if err != nil {
		fmt.Println("Error updating cache")
		return
	}

	for _, s := range specs {
		fmt.Printf("%v - %v - %v\n", s.FQName, s.Description, s.Image)
	}
}
