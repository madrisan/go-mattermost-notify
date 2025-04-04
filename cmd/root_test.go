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
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCheckErr(t *testing.T) {
	testErrMsg := "this is an error message"
	shouldBe := fmt.Sprintln("Error:", testErrMsg)

	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		checkErr(testErrMsg)
	}

	var outbuf, errbuf bytes.Buffer

	cs := []string{"-test.run=TestCheckErr", "--"}

	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		v := errbuf.String()
		if v != shouldBe {
			t.Fatalf("For \"%s\" expected \"%s\" got \"%s\"",
				testErrMsg, shouldBe, v,
			)
		}
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
