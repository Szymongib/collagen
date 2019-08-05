package file

import (
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/Szymongib/collagen/internal/generator"

	"golang.org/x/tools/go/packages"
)

type fileEntry struct {
	fileName    string
	packageName string
	fset        *token.FileSet
	syntax      *ast.File
}

type Parser struct{}

// FindStruct finds the struct with specified name in the specified path
func (p *Parser) FindStruct(structName, path string) (generator.Struct, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return generator.Struct{}, errors.Wrap(err, "Failed to find absolute path")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return generator.Struct{}, errors.WithMessagef(err, "Failed to read %s directory", path)
	}

	entries, err := parseFilesForSyntax(path, files)
	if err != nil {
		return generator.Struct{}, errors.Wrap(err, "Failed to parse files for syntax")
	}

	var structs []generator.Struct

	wg := &sync.WaitGroup{}

	for _, entry := range entries {
		wg.Add(1)
		go func(entry fileEntry) {
			defer wg.Done()

			nodeVisitor := NewNodeVisitor(entry.fset, entry.packageName)

			ast.Walk(nodeVisitor, entry.syntax)

			structs = append(structs, nodeVisitor.Structs()...)
		}(entry)
	}

	wg.Wait()

	structToLookFor, found := findStruct(structName, structs)
	if !found {
		return generator.Struct{}, errors.Errorf("Struct %s not found in the %s path", structName, path)
	}

	return structToLookFor, nil
}

// parseFilesForSyntax parses all go, non test files and extracts the structs with names
func parseFilesForSyntax(path string, files []os.FileInfo) ([]fileEntry, error) {
	config := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax | packages.NeedFiles | packages.NeedName,
	}

	entries := map[string]fileEntry{}

	for _, file := range files {
		fileName := file.Name()

		if !isGoNonTestFile(fileName) {
			continue
		}

		filePath := filepath.Join(path, fileName)

		pkgs, err := packages.Load(config, fmt.Sprintf("file=%s", filePath))
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to load packages from %s file", filePath)
		}

		if len(pkgs) != 1 {
			return nil, errors.Errorf("File %s should contain single package", filePath)
		}

		pkg := pkgs[0]

		for i, f := range pkg.GoFiles {
			if _, ok := entries[f]; ok {
				continue
			}

			entry := fileEntry{
				fileName:    f,
				packageName: pkg.Name,
				fset:        pkg.Fset,
				syntax:      pkg.Syntax[i],
			}

			entries[f] = entry
		}
	}

	var entryValues = make([]fileEntry, len(entries))

	i := 0
	for _, entry := range entries {
		entryValues[i] = entry
		i++
	}

	return entryValues, nil
}

func isGoNonTestFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".go") && !strings.HasSuffix(fileName, "_test.go")
}

func findStruct(structName string, structs []generator.Struct) (generator.Struct, bool) {
	for _, s := range structs {
		if s.Type == structName {
			return s, true
		}
	}

	return generator.Struct{}, false
}
