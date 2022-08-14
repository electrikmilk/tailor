/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var filename string
var css string

func main() {
	registerArg("c", "checker")
	handleArgs()
	if len(os.Args) > 1 {
		filename = os.Args[1]
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			errorf("File %s does not exist.", filename)
		}
		if strings.Split(filename, ".")[1] != "css" {
			errorf("Non-CSS file: %s", filename)
		}
		parser()
		if arg("checker") {
			check()
		}
		progressf("Minifying", "Minifying %s...", filename)
		writeToFile()
	} else {
		usage()
		return
	}
}

func errorf(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println("\n" + style(" FATAL ", BG_RED, BLACK, BOLD) + " " + style(message, RED, BOLD))
	os.Exit(1)
}

func issuef(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println("\n" + style(" ISSUE ", BG_RED, BLACK, BOLD) + " " + style(message, RED, BOLD))
}

func warningf(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println("\n" + style(" WARNING ", BG_YELLOW, BLACK, BOLD) + " " + style(message, YELLOW))
}

func progressf(label string, message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	label = strings.ToUpper(fmt.Sprintf(" %s ", label))
	fmt.Println("\n" + style(label, BG_BLUE, BLACK, BOLD) + " " + style(message, BLUE))
}

func success(label string, message string) {
	label = strings.ToUpper(fmt.Sprintf(" %s ", label))
	fmt.Println("\n" + style(label, BG_GREEN, BLACK, BOLD) + " " + style(message, GREEN))
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// func loadJSON(filename string, ref any) {
// 	file, err := os.ReadFile(filename)
// 	handle(err)
// 	err = json.Unmarshal(file, ref)
// 	handle(err)
// }
