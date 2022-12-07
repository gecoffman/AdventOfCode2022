package lib

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetInputs(fileName string) (lines []string) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("File where art thou?: %s", err)
	}
	content := string(f)
	lines = strings.Split(content, "\r\n")
	return
}
