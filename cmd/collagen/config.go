package collagen

import (
	"flag"
	"os"
	"strings"

	"github.com/Szymongib/collagen/internal/collagen"

	"github.com/pkg/errors"
)

func parseFlagConfig() flagConfig {
	name := flag.String("name", "", "Name of the struct to generate collection methods for.")
	plural := flag.String("plural", "", "Name of the plural type alias.")
	dir := flag.String("dir", "", "Source directory of the struct location.")
	functions := flag.String("functions", "", "Coma separated list of functions to generate. If empty all functions will be generated.")
	pointer := flag.Bool("pointer", false, "Specifies if methods should be generated fo array of pointers")
	//outDir := flag.String("outDir", "", "Output directory.") // TODO: not yet supported.

	flag.Parse()

	return flagConfig{
		name:      *name,
		plural:    *plural,
		dir:       *dir,
		outDir:    "",
		functions: *functions,
		pointer:   *pointer,
	}
}

func validateFlagConfig(config flagConfig) error {
	if config.name == "" {
		return errors.New("Struct name not provided. Provide it with -name flag")
	}

	return nil
}

func prepareConfig(flagConfig flagConfig) (collagen.Config, error) {
	pluralName := flagConfig.plural

	// TODO: introduce better handling of generating plural type name
	if pluralName == "" {
		pluralName = flagConfig.name + "s"
	}

	executionDir, err := os.Getwd()
	if err != nil {
		return collagen.Config{}, errors.Wrap(err, "Failed to get current directory.")
	}

	path := flagConfig.dir

	if path == "" {
		path = executionDir
	}
	outDir := path // TODO: introduce support for out dir

	var functions []string

	if flagConfig.functions != "" {
		functions = strings.Split(flagConfig.functions, ",")
	}

	return collagen.Config{
		StructName: flagConfig.name,
		PluralName: pluralName,
		Path:       path,
		OutputDir:  outDir,
		Functions:  functions,
		Pointer:    flagConfig.pointer,
	}, nil
}
