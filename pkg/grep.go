package pkg

import (
	"bufio"
	"os"
	"strings"
	"regexp"
	"github.com/phuongdoan13/gogrep/config"
	"github.com/spf13/viper"
)

func Grep(pattern string, fileName string) []PairLineNumberAndLine{
	if viper.GetBool(config.InvertMatchFlag) {
		return invertMatch(pattern, fileName)
	} else {
		return matchWholeText(pattern, fileName)
	}
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

func invertMatch(pattern string, fileName string) []PairLineNumberAndLine {
	file, err := os.Open(fileName)
	CheckError(err)
	defer file.Close()
	
	var result []PairLineNumberAndLine

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber += 1
		if !match(pattern, line) {
			result = append(result, PairLineNumberAndLine{lineNumber, line})
		}
	}

	return result
}

func match(pattern string, line string) bool {
	searchLine := line
	searchPattern := pattern

	if viper.GetBool(config.IgnoreCaseFlag) {
		searchLine = strings.ToLower(line)
		searchPattern = strings.ToLower(pattern)
	}

	patternRegex := makeRegex(searchPattern)

	return patternRegex.MatchString(searchLine)
}

func makeRegex(pattern string) *regexp.Regexp {
	if viper.GetBool(config.ExactMatchFlag) {
		return regexp.MustCompile(`\b` + regexp.QuoteMeta(pattern) + `\b`)
	} else {
		return regexp.MustCompile(pattern)
	}
}

type PairLineNumberAndLine struct {
	LineNumber int
	Line string
}
