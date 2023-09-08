package pkg

import (
	"bufio"
	"os"
	"strings"
	"github.com/phuongdoan13/gogrep/config"
)

func Grep(pattern string, fileName string) []PairLineNumberAndLine{
	return matchWholeText(pattern, fileName)
}

func matchWholeText(pattern string, fileName string) []PairLineNumberAndLine {
	file, err := os.Open(fileName)
	CheckError(err)
	defer file.Close()
	
	var result []PairLineNumberAndLine

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber += 1
		if match(pattern, line) {
			result = append(result, PairLineNumberAndLine{lineNumber, line})
		}
	}

	return result
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

type PairLineNumberAndLine struct {
	LineNumber int
	Line string
}
