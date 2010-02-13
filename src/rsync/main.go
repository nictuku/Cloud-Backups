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

var src *string = flag.String("source", "", "Source file or directory")
var destUser *string = flag.String("dest-user", "", "Destination user")
var destHost *string = flag.String("dest-host", "", "Destination host")
var destFile *string = flag.String("dest-file", "", "Destination file or directory")

func main() {
	flag.Parse()
	if err := rsync.Rsync(*src, *destUser, *destHost, *destFile); err != nil {
		log.Stderr("Rsync() err %v\n", err)
	}
}
