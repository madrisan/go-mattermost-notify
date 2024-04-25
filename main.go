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

// Package main is the entrypoint of the go-mattermost-notify CLI.
//
// Source code and other details for the project are available at GitHub:
//
//	https://github.com/madrisan/go-mattermost-notify
package main

import "github.com/madrisan/go-mattermost-notify/cmd"

func main() {
	cmd.Execute()
}
