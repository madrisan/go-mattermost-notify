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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var mattermostURL string
var mattermostAccessToken string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-mattermost-notify",
	Short: "Mattermost client in go",
	Long: `Post a message to a Mattermost channel using its REST APIv4 interface.

Usage: ADD-ME.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is $HOME/.go-mattermost-notify.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mattermostURL,
		"url", "u", "", "mattermost URL")
	rootCmd.PersistentFlags().StringVarP(&mattermostAccessToken,
		"access-token", "a", "", "mattermost Access Token")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".go-mattermost-notify" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-mattermost-notify")
	}

	viper.SetEnvPrefix("mattermost")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	viper.BindEnv("url")
	viper.BindPFlag("url", rootCmd.Flags().Lookup("url"))

	viper.BindEnv("access-token")
	viper.BindPFlag("access-token", rootCmd.Flags().Lookup("access-token"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
