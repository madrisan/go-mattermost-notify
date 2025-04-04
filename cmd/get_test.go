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
	"bytes"
	"strings"
	"testing"
)

func TestCmdGet(t *testing.T) {
	check := "missing endpoint"
	args := []string{"get"}
	shouldBe := "Error: An endpoint must be specified in the command-line arguments"

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err == nil {
		t.Fatalf("The test \"%s`\" did not exit with error: %v", check, err)
	}

	if !strings.HasPrefix(buf.String(), shouldBe) {
		t.Fatalf("For \"%s\" expected \"%s\" got \"%s\"",
			check, shouldBe, buf.String(),
		)
	}
}
