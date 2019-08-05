package file

import (
	"go/ast"
	"go/token"

	"github.com/Szymongib/collagen/internal/generator"
)

type NodeVisitor struct {
	structs     []generator.Struct
	fileSet     *token.FileSet
	packageName string
}

func NewNodeVisitor(fileSet *token.FileSet, pckgName string) *NodeVisitor {
	return &NodeVisitor{
		structs:     []generator.Struct{},
		fileSet:     fileSet,
		packageName: pckgName,
	}
}

func (nv *NodeVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		if str, ok := n.Type.(*ast.StructType); ok {
			nv.structs = append(nv.structs, nv.newStructObject(n.Name.Name, str))
		}
	}
	return nv
}

func (nv *NodeVisitor) newStructObject(typeName string, str *ast.StructType) generator.Struct {
	structObject := generator.Struct{
		Type:    typeName,
		Package: nv.packageName,
	}

	return structObject
}

func (nv *NodeVisitor) Structs() []generator.Struct {
	return nv.structs
}
