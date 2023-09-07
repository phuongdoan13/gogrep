package pkg

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func Grep(pattern string, fileName string) {
	matchWholeText(pattern, fileName)
}

func matchWholeText(pattern string, fileName string) {
	file, err := os.Open(fileName)
	CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Println(line)
		}
	}
}