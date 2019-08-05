package collagen

import (
	"fmt"
	"os"

	"github.com/Szymongib/collagen/internal/collagen"
	"github.com/Szymongib/collagen/internal/file"
)

type flagConfig struct {
	name      string
	plural    string
	dir       string
	outDir    string
	functions string
}

func Collagen() {
	flagConfig := parseFlagConfig()

	if err := validateFlagConfig(flagConfig); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	collagenConfig, err := prepareConfig(flagConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := collagen.Generate(collagenConfig, &file.Parser{}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s generated with methods.", collagenConfig.PluralName)
}
