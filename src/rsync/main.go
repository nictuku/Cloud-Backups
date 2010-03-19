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

// Silly test tool.
package main

import (
	"flag"
	"log"
	"./rsync"
)

var src string
var destUser string
var destHost string
var destFile string

func init() {
	flag.StringVar(&src, "source", "", "Source file or directory")
	flag.StringVar(&destUser, "dest-user", "", "Destination user")
	flag.StringVar(&destHost, "dest-host", "", "Destination host")
	flag.StringVar(&destFile, "dest-file", "", "Destination file or directory")
}

func main() {
	flag.Parse()
	if destUser == "" || destHost == "" || destFile == "" || src == "" {
		flag.Usage()
		return
	}
	if err := rsync.Rsync(src, destUser, destHost, destFile); err != nil {
		log.Stderr("Rsync() err %v\n", err)
	}
}
