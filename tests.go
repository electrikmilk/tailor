/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if _, err := os.Stat("src/tailor"); errors.Is(err, os.ErrNotExist) {
		panic("Please build tailor binary!\n\nRun:\ncd src\ngo build")
	}
	var tests, dirErr = os.ReadDir("tests")
	if dirErr != nil {
		panic(dirErr)
	}
	for _, test := range tests {
		if !strings.Contains(test.Name(), ".min") {
			fmt.Printf("Running test %s...\n", test.Name())
			var testCmd = exec.Command("./src/tailor", "tests/"+test.Name())
			var output, testErr = testCmd.Output()
			if testErr != nil {
				var stderr bytes.Buffer
				testCmd.Stderr = &stderr
				if testErr != nil {
					fmt.Println(stderr.String())
				}
			}
			fmt.Println(string(output))
		}
	}
}
