/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func addRule(r rule) {
	css += fmt.Sprintf("%s{", r.selector)
	for i, d := range r.declarations {
		if (i + 1) < len(r.declarations) {
			addDeclaration(d, false)
		} else {
			addDeclaration(d, true)
		}
	}
	css += "}"
}

func addDeclaration(d declaration, last bool) {
	var value string = d.value
	value = strings.ReplaceAll(value, ", ", ",")
	css += fmt.Sprintf("%s:%s", d.property, value)
	if last == false {
		css += ";"
	}
}

func writeToFile() {
	for _, at := range ats {
		css += fmt.Sprintf("@%s \"%s\";", at.property, at.value)
	}
	for _, query := range queries {
		css += fmt.Sprintf("@media %s{", query.query)
		for _, r := range query.rules {
			addRule(r)
		}
		css += "}"
	}
	for _, r := range rules {
		addRule(r)
	}
	filename = strings.Replace(filename, ".css", "", 1)
	filename = fmt.Sprintf("%s.min.css", filename)
	err := os.WriteFile(filename, []byte(css), 0774)
	handle(err)
	success("Minified", filename)
}
