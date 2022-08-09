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

var properties interface{}
var tags interface{}
var deprecated interface{}
var noStyle interface{}
var units = []string{"cm", "mm", "in", "px", "pt", "pc", "em", "ex", "ch", "rem", "vw", "vh", "vmin", "vmax", "%"}

func check() {
	loadJSON("css/properties.json", &properties)
	loadJSON("html/tags.json", &tags)
	loadJSON("html/deprecated.json", &deprecated)
	loadJSON("html/no_style.json", &noStyle)
	for _, at := range ats {
		if at.property == "import" {
			if _, err := os.Stat(at.value); errors.Is(err, os.ErrNotExist) {
				warning(fmt.Sprintf("Unable to find import \"%s\".", at.value))
			}
		}
	}
	for _, q := range queries {
		for _, r := range q.rules {
			checkSelector(&r.selector)
			for _, d := range r.declarations {
				checkDeclaration(&r.selector, &d.property, &d.value)
			}
		}
	}
	for _, r := range rules {
		checkSelector(&r.selector)
		for _, d := range r.declarations {
			checkDeclaration(&r.selector, &d.property, &d.value)
		}
	}
}

func checkSelector(selector *string) {
	if !strings.ContainsAny(*selector, ".#,: [*") {
		var validHTMLTag bool = false
		for _, tag := range tags.([]interface{}) {
			if tag == strings.Split(*selector, ":")[0] {
				validHTMLTag = true
			}
		}
		if validHTMLTag == false {
			problem(fmt.Sprintf("\"%s\": Not a valid HTML tag selector.", *selector))
		}
		for _, tag := range noStyle.([]interface{}) {
			if tag == strings.Split(*selector, ":")[0] {
				problem(fmt.Sprintf("\"%s\": Not a stylable HTML tag.", *selector))
			}
		}
	}
	var depMap = deprecated.(map[string]interface{})
	if dep, found := depMap[*selector]; found {
		problem(fmt.Sprintf("Deprecated HTML tag \"%s\": %s", *selector, dep))
	}
	// if !strings.Contains(*selector, " ") && strings.Contains(*selector, "[") {
	// 	var attr string = strings.Split(strings.Split(*selector, "[")[1], "]")[0]
	// 	if strings.Contains(attr, "\"") {
	// 		var attrValue string = strings.Trim(strings.Split(attr, "=")[1], "\"")
	// 		fmt.Println(attr)
	// 		fmt.Println(attrValue)
	// 	}
	// }
}

func checkDeclaration(selector *string, propertyName *string, value *string) {
	var stripBrowser = strings.ReplaceAll(*propertyName, "-moz-", "")
	stripBrowser = strings.ReplaceAll(stripBrowser, "-webkit-", "")
	var propMap = properties.(map[string]interface{})
	if def, found := propMap[stripBrowser]; found {
		var definition = def.(map[string]interface{})
		// fmt.Println("property", *propertyName)
		// fmt.Println("moz", definition["moz"])
		// fmt.Println("webkit", definition["webkit"])
		// fmt.Println("syntax", definition["syntax"])
		// fmt.Println("initial", definition["initial"])
		// fmt.Println("values", definition["values"])
		if definition["initial"] != "see individual properties" {
			if definition["initial"] == *value {
				warning(fmt.Sprintf("\"%s\": Property \"%s\" set to it's initial value.", *selector, *propertyName))
			}
		}
		if *propertyName != "content" {
			var validValue = false
			for _, val := range definition["values"].([]interface{}) {
				var definedValue string = fmt.Sprintf("%v", val)
				var rawValue = strings.Trim(strings.Replace(*value, "!important", "", 1), " ")
				if rawValue == definedValue {
					validValue = true
					break
				}
			}
			// fmt.Println(validValue, *propertyName, *value)
			if validValue == false {
				var containsType = false
				for _, val := range definition["values"].([]interface{}) {
					var definedValue string = fmt.Sprintf("%def", val)
					if strings.ContainsAny(definedValue, "[]") {
						containsType = true
						break
					}
				}
				if containsType == false {
					problem(fmt.Sprintf("\"%s\": Value of property \"%s\":\"%s\" does not match syntax: %v", *selector, *propertyName, *value, definition["syntax"]))
				} else {
					// check if value is valid type value
					// for _, val := range definition["values"].([]interface{}) {
					// 	var value string = fmt.Sprintf("%def", val)
					//
					// }
				}
			}
		}
	} else {
		warning(fmt.Sprintf("\"%s\": Unknown property \"%s\".", *selector, *propertyName))
	}
	if strings.Contains(*propertyName, "-moz") || strings.Contains(*propertyName, "-webkit") {
		if def, found := propMap[stripBrowser]; found {
			var definition = def.(map[string]interface{})
			if definition["moz"] == false && strings.Contains(*propertyName, "-moz") {
				warning(fmt.Sprintf("\"%s\": Unknown -moz- property \"%s\".", *selector, *propertyName))
			}
			if definition["webkit"] == false && strings.Contains(*propertyName, "-webkit") {
				warning(fmt.Sprintf("\"%s\": Unknown -webkit- property \"%s\".", *selector, *propertyName))
			}
		}
	}
	for _, unit := range units {
		if *value == fmt.Sprintf("0%s", unit) {
			warning(fmt.Sprintf("\"%s\": unit of measurement \"%s\" after 0 is redundant on \"%s\".", *selector, unit, *propertyName))
		}
	}
}
