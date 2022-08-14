/*
 * Copyright (c) 2022 Brandon Jordan
 */

// TODO: parse variables

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var content string

var line int
var char int
var chars []string
var currentChar string

type rule struct {
	selector     string
	declarations []declaration
}

type declaration struct {
	property string
	value    string
}

type media struct {
	query string
	rules []rule
}

var rules []rule

var ats []declaration

var queries []media

var EOL = "\n"

func parser() {
	if runtime.GOOS == "windows" {
		EOL = "\r\n"
	}
	bytes, err := os.ReadFile(filename)
	handle(err)
	progressf("Parsing", "Parsing %s...", filename)
	content = string(bytes)
	chars = strings.Split(content, "")
	char = 0
	currentChar = chars[0]
	for currentChar != "" {
		if currentChar == " " || currentChar == "\t" || currentChar == EOL {
			advance()
		} else if currentChar == "/" && next(1) == "*" {
			waitForComment()
		} else if currentChar == "@" {
			collectQuery()
		} else {
			rules = append(rules, collectRule())
		}
	}
}

func collectQuery() (dec declaration) {
	var property string
	var value string
	advance()
	for currentChar != " " {
		property += currentChar
		advance()
	}
	if property == "media" || property == "supports" || property == "keyframes" {
		for currentChar != "{" {
			value += currentChar
			advance()
		}
		value = strings.Trim(value, " ")
		advance()
		var mediaRules []rule
		for currentChar != "" {
			if currentChar == " " || currentChar == "\t" || currentChar == EOL {
				advance()
			} else if currentChar == "/" && next(1) == "*" {
				waitForComment()
			} else if currentChar == "}" {
				advance()
				break
			} else {
				mediaRules = append(mediaRules, collectRule())
			}
		}
		queries = append(queries, media{
			query: value,
			rules: mediaRules,
		})
	} else {
		char++
		advance()
		if currentChar != "\"" {
			value = collectString()
		}
		char++
		advance()
		ats = append(ats, declaration{
			property: property,
			value:    value,
		})
	}
	return
}

func collectRule() (rul rule) {
	var selector string
	var declarations []declaration
	for currentChar != "" {
		if currentChar == "{" {
			advance()
			for next(1) != "}" {
				if next(1) == "/" {
					waitForComment()
					continue
				}
				declarations = append(declarations, collectDeclaration())
			}
			selector = strings.Trim(selector, " ")
			rul = rule{
				selector:     selector,
				declarations: declarations,
			}
			for currentChar != "}" {
				advance()
			}
			advance()
			break
		} else {
			if currentChar != "\t" && currentChar != EOL {
				selector += currentChar
			}
			advance()
		}
	}
	return
}

func collectDeclaration() (dec declaration) {
	var property string
	var value string
	for currentChar != ";" {
		if len(property) == 0 {
			if currentChar != " " && currentChar != "\t" && currentChar != EOL {
				for currentChar != ":" {
					if currentChar != " " && currentChar != "\t" && currentChar != EOL && currentChar != ":" {
						property += currentChar
					} else {
						parserError(fmt.Sprintf("invalid property: %s", property))
					}
					advance()
				}
			}
		} else {
			if len(value) == 0 && currentChar == " " {
				advance()
				continue
			}
			// TODO: collect variable property value (e.g. url(), attr())
			// if currentChar == "(" {
			// 	fmt.Println(currentChar)
			// for currentChar != ")" {
			// 	value += currentChar
			// 	advance()
			// }
			// }
			if currentChar == "\"" {
				advance()
				value = collectString()
				advance()
				continue
			}
			if currentChar == EOL {
				if len(value) == 0 {
					parserError(fmt.Sprintf("No value given to property: %s", property))
				}
				break
			}
			if currentChar != "\t" && currentChar != ":" {
				value += currentChar
			} else {
				parserError(fmt.Sprintf("Invalid property value: %s", value))
			}
		}
		advance()
	}
	if currentChar != EOL {
		advance()
	}
	dec = declaration{
		property: property,
		value:    value,
	}
	return
}

func collectString() (str string) {
	for currentChar != "\"" {
		str += currentChar
		advance()
	}
	return
}

func waitForComment() {
	for currentChar != "" {
		if currentChar == "*" && next(1) == "/" {
			char++
			advance()
			break
		}
		advance()
	}
}

func advance() {
	char++
	if len(chars) > char {
		currentChar = chars[char]
		if currentChar == EOL {
			line++
		}
	} else {
		currentChar = ""
	}
}

func getChar(c int) (char string) {
	if len(chars) > c {
		char = chars[c]
	} else {
		char = ""
	}
	return
}

func next(mov int) string {
	return seek(&mov, false)
}

// func prev(mov int) string {
// 	return seek(&mov, true)
// }

func seek(mov *int, reverse bool) (requestedChar string) {
	var nextChar int = char
	if reverse == true {
		nextChar -= *mov
	} else {
		nextChar += *mov
	}
	requestedChar = getChar(nextChar)
	for requestedChar == " " || requestedChar == "\t" || requestedChar == EOL {
		if reverse == true {
			nextChar -= 1
		} else {
			nextChar += 1
		}
		requestedChar = getChar(nextChar)
	}
	return
}

func parserError(error string) {
	fmt.Println("\n" + style(error, RED, BOLD) + "\n")
	var lines []string = strings.Split(content, EOL)
	var offsetLine = line + 1
	var prev int = offsetLine - 1
	var next int = offsetLine + 1
	if len(lines) > prev {
		fmt.Printf(style(fmt.Sprintf("%d |", prev), DIM))
		fmt.Printf("%s\n", lines[line-1])
	}
	fmt.Printf("%d ", offsetLine)
	fmt.Printf(style("|", DIM))
	fmt.Println(style(fmt.Sprintf("%s", lines[line]), RED, BOLD))
	if len(lines) >= next {
		fmt.Printf(style(fmt.Sprintf("%d |", next), DIM))
		fmt.Printf("%s\n", lines[line+1])
	}
	fmt.Printf("\n")
	// fmt.Printf("%d:%d | char: %s, next: %s, prev: %s\n", line, char, prev(1), next(1), currentChar)
	// panic(error)
	os.Exit(1)
}

// func printCurrentChar() {
// 	var char string
// 	switch currentChar {
// 	case "\t":
// 		char = "TAB"
// 		break
// 	case " ":
// 		char = "SPACE"
// 		break
// 	case EOL:
// 		char = "EOL"
// 		break
// 	default:
// 		char = currentChar
// 	}
// 	fmt.Println(char)
// }
