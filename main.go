/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/phuongdoan13/gogrep/cmd"
// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// )
func main() {
	cmd.Execute()
	// fileName := "~/Users/phuongdoantuminh/Projects/gogrep/test/DummyFile.txt"         
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// pattern := "of"
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if strings.Contains(line, pattern) {
	// 		fmt.Println(line)
	// 	}
	// }
}
