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

package main

import (
	"flag"
	"fmt"
	"http"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"./util"
)

var deliciousUrl *string = flag.String(
	"delicious-url",
	// TODO(nictuku): Change to this right URL when Go supports HTTPS.
	// => "http://del.icio.us/api/posts/all",
	"http://api.del.icio.us/v1/posts/all",
	"URL of XML dump file from Delicious.com. HTTPS is currently NOT supported.")
var deliciousUser *string = flag.String("user", "", "Username")
// TODO(nictuku): read from user input instead.
var deliciousPassword *string = flag.String("password", "", "Password")
var deliciousDestFile *string = flag.String("dest-file",
	util.GetHomedir()+"/cloudbackups/delicious.xml",
	"Destination file path.")

func readData(url string) (b []byte, err os.Error) {
	r, _, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, os.ErrorString(r.Status)
	}
	b, _ = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return
}

func saveBackup(data []byte) os.Error {
	// Sub-directory must exist already.
	return ioutil.WriteFile(*deliciousDestFile, data, 0700)
}

func main() {
	flag.Parse()
	if len(*deliciousUser) == 0 || len(*deliciousPassword) == 0 {
		log.Stderrf("Missing user or password.\n\n")
		flag.Usage()
		os.Exit(2)
	}
	if strings.HasPrefix(*deliciousUrl, "http://") != true {
		log.Stderrf("URL missing 'http://' prefix. HTTPS is currently NOT supported.\n\n")
		flag.Usage()
		os.Exit(2)
	}
	split := strings.Split(*deliciousUrl, "http://", 0)[1:]
	userInfo := strings.Join(split, "")
	url := fmt.Sprintf("http://%s:%s@%s", *deliciousUser, *deliciousPassword, userInfo)

	b, err := readData(url)
	if err != nil {
		log.Stderrf("Error downloading %s: %s", url, err)
		os.Exit(1)
	}
	err = saveBackup(b)
	if err != nil {
		log.Stderrf("Error saving file to disk %s", err)
		os.Exit(1)
	}
}
