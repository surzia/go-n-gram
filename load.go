package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// loadSoHuData - load sohu.dat into memory and extract content to string list
func loadSoHuData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var ret []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if strings.HasPrefix(line, "<content>") {
			line = strings.Replace(line, "\n", "", -1)
			line = strings.Replace(line, "<content>", "", -1)
			line = strings.Replace(line, "</content>", "", -1)
			if len(line) > 1 {
				ret = append(ret, line)
			}
		}
	}

	return ret
}
