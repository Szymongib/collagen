package collagen

import (
	"os"
	"path/filepath"

	"github.com/Szymongib/collagen/internal/generator"
)

type Config struct {
	StructName string
	PluralName string
	Path       string
	// OutDir is not yet supported
	OutputDir string
	Functions []string
	Pointer   bool
}

//go:generate mockery -name Parser
type Parser interface {
	FindStruct(structName, path string) (generator.Struct, error)
}

func Generate(config Config, parser Parser) error {
	structure, err := parser.FindStruct(config.StructName, config.Path)
	if err != nil {
		return err
	}

	structure.PluralType = config.PluralName

	content, err := generator.GenerateCollectionMethods(structure, config.Functions, config.Pointer)
	if err != nil {
		return err
	}

	if _, err := os.Stat(config.OutputDir); os.IsNotExist(err) {
		err = os.Mkdir(config.OutputDir, 0774)
		if err != nil {
			return err
		}
	}

	fileName := filepath.Join(config.OutputDir, "generated_"+structure.PluralType+".go")

	outFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0774)
	if err != nil {
		return err
	}

	_, err = outFile.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
