/*
  Copyright 2021 Davide Madrisan <davide.madrisan@gmail.com>

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

// Package cmd implements the CLI interface using the Cobra and Viper libraries
package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var mattermostURL, mattermostAccessToken string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-mattermost-notify",
	Short: "Mattermost client in go",
	Long:  `Post a message to a Mattermost channel using its REST APIv4 interface.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError("%v", err)
	}
}

// handleError prints an error message to os.Stderr and termines the execution with an error code 1
func handleError(format string, a ...interface{}) {
	pFormat := fmt.Sprintf("Error: %v", format)
	message := fmt.Sprintf(pFormat, a...)
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is $HOME/.go-mattermost-notify.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mattermostURL,
		"url", "u", "",
		"Mattermost URL. The command-line value has precedence over the MATTERMOST_URL environment variable.")
	rootCmd.PersistentFlags().StringVarP(&mattermostAccessToken,
		"access-token", "a", "",
		"Mattermost Access Token. The command-line value has precedence over the MATTERMOST_ACCESS_TOKEN environment variable.")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet mode")

	err := viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	if err != nil {
		handleError("unable to bind 'quiet' flag: %v", err)
	}
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
			handleError("%v", err)
		}

		// Search config in home directory with name ".go-mattermost-notify" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-mattermost-notify")
	}

	var envVars = [...]string{
		"url",
		"access-token",
	}

	viper.SetEnvPrefix("mattermost")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	for _, envVar := range envVars {
		if err := viper.BindEnv(envVar); err != nil {
			handleError("%v", err)
		}
		if err := viper.BindPFlag(envVar, rootCmd.Flags().Lookup(envVar)); err != nil {
			handleError("%v", err)
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
