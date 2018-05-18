package cmd

import (
	"fmt"
	"os"

	"github.com/lnquy/bc/config"
	"github.com/lnquy/bc/storage"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "Basic blockchain implementation",
	Long:  ``,
}

var (
	cfg    *config.Config
	ledger storage.Ledger
)

func Execute(cfgArg *config.Config, ledgerArg storage.Ledger) {
	cfg = cfgArg
	ledger = ledgerArg

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".blockchain")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}