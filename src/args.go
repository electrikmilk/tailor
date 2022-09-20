/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

type argument struct {
	name        string
	short       string
	description string
}

var args map[string]string
var registered []argument

var customUsage string

func init() {
	args = make(map[string]string)
	if len(os.Args) > 1 {
		for i, a := range os.Args {
			if i != 0 {
				a = strings.TrimPrefix(a, "--")
				a = strings.TrimPrefix(a, "-")
				if strings.Contains(a, "=") {
					var keyValue []string = strings.Split(a, "=")
					if len(keyValue) > 1 {
						args[a] = keyValue[1]
					} else {
						args[a] = ""
					}
				} else {
					args[a] = ""
				}
			}
		}
	}
}

// Prints a usage message based on the arguments and usage you have registered.
func usage() {
	fmt.Printf("USAGE: %s %s", os.Args[0], customUsage)
	fmt.Printf("\nOptions:\n")
	for _, arg := range registered {
		fmt.Printf("\t-%s --%s\t%s\n", arg.short, arg.name, arg.description)
	}
}

// Register an argument
func registerArg(name string, shorthand string, description string) {
	for _, r := range registered {
		if r.name == name {
			panic(fmt.Sprintf("Argument %s is already registered!", name))
		}
	}
	registered = append(registered, argument{
		name:        name,
		short:       shorthand,
		description: description,
	})
}

// Register a custom usage message.
// It will be printed after USAGE: executable.
func registerUsage(usage string) {
	customUsage = usage
}

// Returns a boolean indicating if argument name was passed to your executable.
func arg(name string) bool {
	if len(args) > 0 {
		if _, ok := args[name]; ok {
			return true
		}
		for _, r := range registered {
			if r.name == name {
				if _, ok := args[r.short]; ok {
					return true
				}
			}
		}
	}
	return false
}

// Returns a string of the value of argument name if passed to your executable.
func argValue(name string) (value string) {
	if len(args) > 0 {
		if val, ok := args[name]; ok {
			value = val
		}
	} else {
		value = ""
	}
	return
}
