/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var units = []string{"cm", "mm", "in", "px", "pt", "pc", "em", "ex", "ch", "rem", "vw", "vh", "vmin", "vmax", "%"}
var properties interface{}

func check() {
	downloadProperties()
	makeDeprecated()
	if len(ats) > 0 {
		for _, at := range ats {
			if at.property == "import" {
				warningf("@import \"%s\": Avoid using @import as it is render-blocking. Multiple <link> tags within HTML is more efficient and loads CSS files in parallel.", at.value)
				if _, err := os.Stat(at.value); errors.Is(err, os.ErrNotExist) {
					issuef("Unable to find import \"%s\".", at.value)
				} else {
					// maybe?
					// parser(at.value)
				}
			}
		}
	}
	if len(queries) > 0 {
		for _, q := range queries {
			for _, r := range q.rules {
				checkSelector(&r)
				for _, d := range r.declarations {
					checkDeclaration(&r, &d)
				}
			}
		}
	}
	if len(rules) > 0 {
		for _, r := range rules {
			checkSelector(&r)
			for _, d := range r.declarations {
				checkDeclaration(&r, &d)
			}
		}
	}
}

func checkSelector(rule *rule) {
	if !strings.ContainsAny(rule.selector, ".#,: [*") {
		var validHTMLTag bool = false
		for _, tag := range tags {
			if tag == strings.Split(rule.selector, ":")[0] {
				validHTMLTag = true
			}
		}
		if validHTMLTag == false {
			issuef("\"%s\": Not a valid HTML tag selector.", rule.selector)
		}
		for _, tag := range noStyle {
			if tag == strings.Split(rule.selector, ":")[0] {
				issuef("\"%s\": Not a stylable HTML tag.", rule.selector)
			}
		}
	}
	if dep, found := deprecated[rule.selector]; found {
		issuef("Deprecated HTML tag \"%s\": %s", rule.selector, dep)
	}
	// TODO: make sure values of attribute selectors need quotes
	// if !strings.Contains(rule.selector, " ") && strings.Contains(rule.selector, "[") {
	// 	var attr string = strings.Split(strings.Split(rule.selector, "[")[1], "]")[0]
	// 	if strings.Contains(attr, "\"") {
	// 		var attrValue string = strings.Trim(strings.Split(attr, "=")[1], "\"")
	// 		fmt.Println(attr)
	// 		fmt.Println(attrValue)
	// 	}
	// }
}

func checkDeclaration(rule *rule, declaration *declaration) {
	var propertyStripBrowser = strings.ReplaceAll(declaration.property, "-moz-", "")
	propertyStripBrowser = strings.ReplaceAll(propertyStripBrowser, "-webkit-", "")
	var valueStripBrowser = strings.ReplaceAll(declaration.value, "-moz-", "")
	valueStripBrowser = strings.ReplaceAll(valueStripBrowser, "-webkit-", "")
	var propMap = properties.(map[string]interface{})
	if def, found := propMap[propertyStripBrowser]; found {
		var definition = def.(map[string]interface{})
		if definition["initial"] != "see individual properties" {
			if definition["initial"] == declaration.value {
				warningf("\"%s\": Property \"%s\" set to it's initial value.", rule.selector, declaration.property)
			}
		}
		if declaration.property != "content" && !strings.ContainsAny(declaration.value, "()") {
			var validValue = false
			for _, val := range definition["values"].([]interface{}) {
				var definedValue string = fmt.Sprintf("%v", val)
				var rawValue = strings.Trim(strings.Replace(valueStripBrowser, "!important", "", 1), " ")
				if rawValue == definedValue {
					validValue = true
					break
				}
			}
			if validValue == false {
				var containsType = false
				for _, val := range definition["values"].([]interface{}) {
					var definedValue string = fmt.Sprintf("%v", val)
					if strings.ContainsAny(definedValue, "[]") {
						containsType = true
						break
					}
				}
				if containsType == false {
					issuef("\"%s\": Value of property \"%s\":\"%s\" does not match syntax: %v", rule.selector, declaration.property, declaration.value, definition["syntax"])
				} else {
					// TODO: check if value is valid type value
					// for _, val := range definition["values"].([]interface{}) {
					// 	var value string = fmt.Sprintf("%v", val)
					//
					// }
				}
			}
		}
	} else {
		warningf("\"%s\": Unknown property \"%s\".", rule.selector, declaration.property)
	}
	if strings.Contains(declaration.property, "-moz") || strings.Contains(declaration.property, "-webkit") {
		if def, found := propMap[propertyStripBrowser]; found {
			var definition = def.(map[string]interface{})
			if definition["moz"] == false && strings.Contains(declaration.property, "-moz") {
				warningf("\"%s\": Unknown -moz- property \"%s\".", rule.selector, declaration.property)
			}
			if definition["webkit"] == false && strings.Contains(declaration.property, "-webkit") {
				warningf("\"%s\": Unknown -webkit- property \"%s\".", rule.selector, declaration.property)
			}
		}
	}
	for _, unit := range units {
		if declaration.value == fmt.Sprintf("0%s", unit) {
			warningf("\"%s\": unit of measurement \"%s\" after 0 is redundant on \"%s\".", rule.selector, unit, declaration.property)
		}
	}
}

func downloadProperties() {
<<<<<<< Updated upstream
	fmt.Printf("Downloading... ")
=======
	fmt.Printf("Downloading property list... ")
>>>>>>> Stashed changes
	resp, downloadErr := http.Get("https://raw.githubusercontent.com/electrikmilk/tailor/main/properties.json")
	if downloadErr != nil {
		fmt.Printf("\n")
		panic(downloadErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		handle(err)
	}(resp.Body)
	fmt.Printf("done.\n")
	body, err := io.ReadAll(resp.Body)
	handle(err)
	err = json.Unmarshal(body, &properties)
	handle(err)
}
