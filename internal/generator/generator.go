package generator

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"

	"github.com/pkg/errors"
)

const (
	containsFunc = "Contains"
	dropFunc     = "Drop"
	existsFunc   = "Exists"
	filterFunc   = "Filter"
	mapFunc      = "Map"
	takeFunc     = "Take"
)

var supportedFunctions = []string{
	containsFunc,
	dropFunc,
	existsFunc,
	filterFunc,
	mapFunc,
	takeFunc,
}

var functionTemplates = map[string]string{
	containsFunc: containsFuncTemplate,
	dropFunc:     dropFuncTemplate,
	existsFunc:   existFuncTemplate,
	filterFunc:   filterFuncTemplate,
	mapFunc:      mapFuncTemplate,
	takeFunc:     takeFuncTemplate,
}

func GenerateCollectionMethods(struc Struct, functions []string, usePointers bool) (string, error) {
	if len(functions) == 0 {
		functions = supportedFunctions
	}

	if usePointers {
		struc.Type = fmt.Sprintf("*%s", struc.Type)
	}

	var generatedFunctions = map[string]string{}
	var requestedFunctions []string

	for _, fun := range functions {
		funTemplate, found := functionTemplates[fun]
		if !found {
			fmt.Printf("Function %s not recognized", fun)
			continue
		}

		generatedFun, err := createFromTemplate(struc, funTemplate)
		if err != nil {
			return "", errors.WithMessagef(err, "Failed to generate %s function", fun)
		}

		generatedFunctions[fun] = generatedFun
		requestedFunctions = append(requestedFunctions, fun)
	}

	sort.Strings(requestedFunctions)

	code := struct {
		PluralType         string
		Type               string
		Package            string
		RequestedFunctions []string
		GeneratedFunctions map[string]string
	}{
		PluralType:         struc.PluralType,
		Type:               struc.Type,
		Package:            struc.Package,
		RequestedFunctions: requestedFunctions,
		GeneratedFunctions: generatedFunctions,
	}

	return createFromTemplate(code, generatedCodeTemplate)
}

func createFromTemplate(data interface{}, rawTemplate string) (string, error) {
	tmpl, err := template.New("").Parse(rawTemplate)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
