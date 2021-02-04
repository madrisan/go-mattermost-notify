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
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestHandleError(t *testing.T) {
	if os.Getenv("BE_HANDLE_ERROR") == "1" {
		HandleError("%v", "this is an error message")
		return
	}

	cases := []struct {
		format   string
		message  string
		shouldBe string
	}{
		{
			"%v",
			"this is an error message",
			"Error: this is an error message\n",
		},
	}

	for _, tc := range cases {
		var outbuf, errbuf bytes.Buffer

		cs := []string{"-test.run=TestHandleError", "--"}
		cs = append(cs, tc.format, tc.message)

		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = append(os.Environ(), "BE_HANDLE_ERROR=1")
		cmd.Stdout = &outbuf
		cmd.Stderr = &errbuf

		err := cmd.Run()

		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			v := errbuf.String()
			if v != tc.shouldBe {
				t.Fatalf("For \"%s\" expected \"%s\" got \"%s\"",
					tc.message, tc.shouldBe, v,
				)
			}
			return
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}
