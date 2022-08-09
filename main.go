/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var css string

func main() {
	if len(os.Args) > 1 {
		parser(os.Args[1])
		check()
		writeToFile(os.Args[1])
	} else {
		fmt.Println("USAGE: tailor [FILE]")
		return
	}
}

func errorf(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println(style(" FATAL ", BG_RED, BLACK, BOLD) + " " + style(message, RED, BOLD) + "\n")
	os.Exit(1)
}

func issuef(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println(style(" ISSUE ", BG_RED, BLACK, BOLD) + " " + style(message, RED, BOLD) + "\n")
}

func warningf(message string, vars ...any) {
	message = fmt.Sprintf(message, vars...)
	fmt.Println(style(" WARNING ", BG_YELLOW, BLACK, BOLD) + " " + style(message, YELLOW) + "\n")
}

func success(label string, message string) {
	label = strings.ToUpper(fmt.Sprintf(" %s ", label))
	fmt.Println(style(label, BG_GREEN, BLACK, BOLD) + " " + style(message, GREEN) + "\n")
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func loadJSON(filename string, ref any) {
	file, err := os.ReadFile(filename)
	handle(err)
	err = json.Unmarshal(file, ref)
	handle(err)
}
