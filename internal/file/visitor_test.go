package file

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Szymongib/go-collagen/internal/generator"
)

func TestNodeVisitor_Visit(t *testing.T) {

	for _, testCase := range []struct {
		description string
		packageName string
		source      string
		structs     []generator.Struct
	}{
		{
			description: "find single struct in source",
			packageName: "foo",
			source: `package foo

	type Thing struct {
    	Field1 string
    	Field2 []int
    	Field3 map[byte]float64
    	Field4 *Frog
  	}`,
			structs: []generator.Struct{
				{
					Type:    "Thing",
					Package: "foo",
				},
			},
		},
		{
			description: "find multiple structs in source",
			packageName: "foo",
			source: `package foo

	type Thing struct {
    	Field1 string
    	Field2 []int
    	Field3 map[byte]float64
    	Field4 *Frog
  	}

	type DifferentThing struct {
		Thing
    	Field map[Thing][]Thing
  	}`,
			structs: []generator.Struct{
				{
					Type:    "Thing",
					Package: "foo",
				},
				{
					Type:    "DifferentThing",
					Package: "foo",
				},
			},
		},
		{
			description: "not fail when no struct in source",
			packageName: "foo",
			source: `package foo

	func doSomething(frog *Frog, array []int) (bool, error) {
		return true, nil
	}

	type Test interface {
		Test()
	}	
	
			`,
			structs: []generator.Struct{},
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// given
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", testCase.source, 0)
			require.NoError(t, err)

			visitor := NewNodeVisitor(fset, "foo")

			// when
			ast.Walk(visitor, f)

			// then
			structs := visitor.Structs()

			assert.ElementsMatch(t, testCase.structs, structs)
		})
	}

}
