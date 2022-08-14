package main

import (
	"os"
	"strings"
)

var args map[string]string
var registered map[string]string

func init() {
	registered = make(map[string]string)
	args = make(map[string]string)
}

func handleArgs() {
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i != 0 {
				arg = strings.TrimPrefix(arg, "--")
				arg = strings.TrimPrefix(arg, "-")
				if long, ok := registered[arg]; ok {
					arg = long
				}
				if strings.Contains(arg, "=") {
					var keyValue []string = strings.Split(arg, "=")
					if len(keyValue) > 1 {
						args[keyValue[0]] = keyValue[1]
					} else {
						args[keyValue[0]] = ""
					}
				} else {
					args[arg] = ""
				}
			}
		}
	}
}

func registerArg(short string, long string) {
	registered[short] = long
}

func arg(key string) bool {
	if len(args) > 0 {
		if _, ok := args[key]; ok {
			return true
		}
	}
	return false
}

func argValue(key string) (value string) {
	if len(args) > 0 {
		if val, ok := args[key]; ok {
			value = val
		}
	} else {
		value = ""
	}
	return
}
