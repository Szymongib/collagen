package generator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

const (
	testTemplate = `global:
  domainName: {{ .DomainName }}
  testValue: {{ .TestValue }}`
)

func Test_createFromTemplate(t *testing.T) {

	t.Run("should parse overrides", func(t *testing.T) {
		// given
		expectedOutput := `global:
  domainName: my.domain
  testValue: 100`

		data := struct {
			DomainName string
			TestValue  int
		}{
			DomainName: "my.domain",
			TestValue:  100,
		}

		// when
		output, err := createFromTemplate(data, testTemplate)

		// then
		assert.NoError(t, err)
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("should return error if invalid data provided", func(t *testing.T) {
		// given
		data := struct {
			WrongData string
		}{
			WrongData: "wrongData",
		}

		// when
		output, err := createFromTemplate(data, testTemplate)

		// then
		assert.Error(t, err)
		assert.Equal(t, "", output)
	})

	t.Run("should return error if template is invalid", func(t *testing.T) {
		// given
		data := struct{}{}

		// when
		output, err := createFromTemplate(data, `{{ .invalid template }{{  }}`)

		// then
		assert.Error(t, err)
		assert.Equal(t, "", output)
	})
}

const (
	renderedType = `// Code generated by collagen. DO NOT EDIT
package tests

type Tests []Test

func (collection Tests) ToSlice() []Test {
	return []Test(collection)
}
`
	renderedContainsFunc = `
// Contains determines if the slice contains the provided element
func (collection Tests) Contains(element Test) bool {
	for _, e := range collection {
		if e == element {
			return true
		}
	}
	return false
}
`

	renderedDropFunc = `
// Drop returns a slice without first n elements
func (collection Tests) Drop(n int) Tests {
	length := len(collection)

	if n > length {
		return Tests{}
	}
	return collection[:length-n]
}
`

	renderedExistsFunc = `
// Exist returns if the collection contains the element that satisfy passed function
func (collection Tests) Exist(f func(element Test) bool) bool {
	for _, e := range collection {
		if f(e) {
			return true
		}
	}
	return false
}
`

	renderedFilterFunc = `
// Filter returns collection of elements that satisfied the check
func (collection Tests) Filter(f func(item Test) bool) Tests {
	filtered := Tests{}
	for _, e := range collection {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
`
	renderedMapFunc = `
// Map maps all collection elements with use of provided function
func (collection Tests) Map(f func(item Test) interface{}) []interface{} {
	result := make([]interface{}, len(collection))
	for i, e := range collection {
		result[i] = f(e)
	}
	return result
}
`

	renderedTakeFunc = `
// Take returns first n elements of the slice
func (collection Tests) Take(n int) Tests {
	if n > len(collection) {
		return collection
	}
	return collection[:n]
}
`
)

func TestGenerateCollectionMethods(t *testing.T) {

	testStruct := Struct{
		Type:       "Test",
		PluralType: "Tests",
		Package:    "tests",
	}

	for _, testCase := range []struct {
		description    string
		structure      Struct
		functions      []string
		expectedOutput string
	}{
		{
			description:    "generate all functions if non provided",
			structure:      testStruct,
			functions:      []string{},
			expectedOutput: renderedType + renderedContainsFunc + renderedDropFunc + renderedExistsFunc + renderedFilterFunc + renderedMapFunc + renderedTakeFunc,
		},
		{
			description:    "generate only specified functions",
			structure:      testStruct,
			functions:      []string{"Contains", "Filter", "Take"},
			expectedOutput: renderedType + renderedContainsFunc + renderedFilterFunc + renderedTakeFunc,
		},
		{
			description:    "generate only one function",
			structure:      testStruct,
			functions:      []string{"Filter"},
			expectedOutput: renderedType + renderedFilterFunc,
		},
		{
			description:    "ignore unrecognized functions",
			structure:      testStruct,
			functions:      []string{"NonFunc", "AnotherFunc"},
			expectedOutput: renderedType,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			output, err := GenerateCollectionMethods(testCase.structure, testCase.functions)

			// then
			require.NoError(t, err)
			assert.Equal(t, testCase.expectedOutput, output)
		})
	}

}
