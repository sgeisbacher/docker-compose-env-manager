package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "compose-env-manager",
	Short: "compose-env-manager manages different environment-scenarios",
	Long: `compose-env-manager manages different environment-scenarios for docker-compose. 
it generates the 'docker-compose.yml'-file based on specified scenario. 
	
for missing-services in the specied scenario compared to 'all'-scenario, it adds 'extra-hosts'-entries to all other services. so your services running in docker will try to connect the missing services on your host-machine`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./compose-env-manager.yml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("compose-env-manager")
		viper.AddConfigPath(".")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("could not read config-file, aborting ...", err)
		os.Exit(1)
	}
}
