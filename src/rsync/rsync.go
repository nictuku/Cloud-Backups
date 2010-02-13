//   Copyright 2010 Yves Junqueira
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// The rsync package is a very simple interface to the command line rsync tool.
package rsync

import (
	"os"
	"io/ioutil"
	"exec"
	"log"
)

// Rsync sends a file via rsync to a remote host.
// Password input is not supported. You must use an SSH key in a standard location.
func Rsync(source string, user string, host string, dest string) (err os.Error) {
	r_path, err := exec.LookPath("rsync")
	if err != nil {
		log.Stderrf("rsync command not found (%s)\n", err)
		return
	}
	// MergeWithStdout makes error messages disappear.
	cmd, err := exec.Run(r_path, []string{r_path, "-az", source, user + "@" + host + ":" + dest},
		os.Environ(), exec.DevNull, exec.DevNull, exec.Pipe)
	// I love this in Go...
	defer cmd.Close()
	// .. but, man, these error checks look ugly.
	if err != nil {
		log.Stderrf("rsync run error (%s)\n", err)
		return
	}
	waitmsg, err := cmd.Wait(0)
	if err != nil {
		log.Stderrf("rsync wait error (%s)\n", err)
		return
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		log.Stderrf("Error reading from stderr (%s)\n", err)
	}
	log.Stdout(string(buf))
	if waitmsg.ExitStatus() != 0 {
		log.Stderrf("rsync returned with an error status (%s)\n", waitmsg)
		return
	}
	return
}
