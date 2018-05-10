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

var dockerhubConfig = registries.Config{
	Type:      "dockerhub",
	Name:      "dh",
	Org:       "ansibleplaybookbundle",
	URL:       "docker.io",
	Tag:       "latest",
	WhiteList: []string{".*-apb$"},
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
	authNamespace := ""
	reg, err := registries.NewRegistry(dockerhubConfig, authNamespace)
	if err != nil {
		log.Error("Error from creating a NewRegistry")
		log.Error(err)
		return nil, err
	}

	specs, count, err := reg.LoadSpecs()
	if err != nil {
		log.Errorf("registry: %v was unable to complete bootstrap - %v",
			reg.RegistryName(), err)
		return nil, err
	}

	log.Infof("Registry %v has %d bundles available from %d images scanned", reg.RegistryName(), len(specs), count)
	return specs, nil
}

func listImages() {
	var specs []*apb.Spec = nil
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Println("Error unmarshalling config: ", err)
		return
	}
	if len(conf.Specs) > 0 {
		fmt.Println("Found specs already in config")
		for _, s := range conf.Specs {
			fmt.Printf("%v - %v\n", s.FQName, s.Description)
		}
		return
	}

	specs, err = getImages()
	if err != nil {
		fmt.Println("Error getting images")
		return
	}
	err = updateCachedList(specs)
	if err != nil {
		fmt.Println("Error updating cache")
		return
	}

	for _, s := range specs {
		fmt.Printf("%v - %v\n", s.FQName, s.Description)
	}
}
