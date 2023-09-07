package pkg

import (
	"fmt"
	"github.com/phuongdoan13/gogrep/config"
)
func CheckError(e error) {
	if e != nil {
			panic(e)
	}
}

func fmtPrintLn(line string, lineNumber int) {
	if config.IsPrintLnWithNumLine {
		fmt.Println(lineNumber, line)
	} else {
		fmt.Println(line)
	}
}