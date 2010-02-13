package util

import (
	"os"
	"strings"
)


func GetHomedir() string {
	env := os.Environ()
	for _, e := range env {
		if strings.HasPrefix(e, "HOME=") {
			return strings.Split(e, "=", 0)[1]
		}
	}
	return "/tmp"
}
