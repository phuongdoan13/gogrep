package pkg

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/phuongdoan13/gogrep/config"
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

		if match(pattern, line) {
			fmt.Println(line)
		}
	}
}

func match(pattern string, line string) bool {
	searchLine := line
	searchPattern := pattern
	if config.IsIgnoreCase {
		searchLine = strings.ToLower(line)
		searchPattern = strings.ToLower(pattern)
	}
	return strings.Contains(searchLine, searchPattern)
}