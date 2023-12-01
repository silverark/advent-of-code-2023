package file

import (
	"os"
	"strings"
)

func GetFile(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to load file: " + err.Error())
	}
	return strings.Split(string(dat), "\n")
}
