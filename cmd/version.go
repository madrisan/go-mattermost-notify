/*
  Copyright 2021-2024 Davide Madrisan <d.madrisan@proton.me>

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
	"os"
	"path/filepath"
	"runtime"

	"github.com/madrisan/go-mattermost-notify/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the program and name version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		processExecutable, err := os.Executable()
		checkErr(err)

		verInfo := version.GetVersion()
		version := verInfo.FullVersionNumber(true)
		fmt.Printf("%s v%s (%s/%s %s)\n",
			filepath.Base(processExecutable),
			version,
			runtime.GOOS, runtime.GOARCH,
			runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
