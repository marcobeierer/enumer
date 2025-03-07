package main

import (
	"fmt"
	"strings"
)

const valuesMethod = `func (i %[1]s) Values() []string {
	return []string{
		%s
	}
}
`

// TODO is runsThreshold necessary?
func (g *Generator) addValuesMethod(runs [][]Value, typeName string) {
	var valuesSlice []string

	for _, values := range runs {
		for _, value := range values {
			valuesSlice = append(valuesSlice, fmt.Sprintf("%s.String(),", value.name))
		}
	}

	g.Printf("\n")
	g.Printf(valuesMethod, typeName, strings.Join(valuesSlice, "\n"))
}
