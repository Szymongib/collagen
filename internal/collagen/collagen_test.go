package collagen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"

	"github.com/Szymongib/collagen/internal/collagen/mocks"

	"github.com/Szymongib/collagen/internal/generator"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {

	tempDir, err := ioutil.TempDir(".", "tests")
	require.NoError(t, err)
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	for _, testCase := range []struct {
		description string
		config      Config
		structure   generator.Struct
		parseError  error
		shouldFail  bool
	}{
		{
			description: "generate file with collection methods",
			config: Config{
				StructName: "Dog",
				PluralName: "Dogs",
				Path:       tempDir,
				OutputDir:  tempDir,
				Functions:  []string{},
			},
			structure: generator.Struct{
				Type:       "Dog",
				PluralType: "Dogs",
				Package:    "animals",
			},
			parseError: nil,
			shouldFail: false,
		},
		{
			description: "create dir and generate file",
			config: Config{
				StructName: "Dog",
				PluralName: "Dogs",
				Path:       filepath.Join(tempDir, "dir_to_create"),
				OutputDir:  filepath.Join(tempDir, "dir_to_create"),
				Functions:  []string{},
			},
			structure: generator.Struct{
				Type:       "Dog",
				PluralType: "Dogs",
				Package:    "animals",
			},
			parseError: nil,
			shouldFail: false,
		},
		{
			description: "fail if invalid path provided",
			config: Config{
				StructName: "Dog",
				PluralName: "Dogs",
				Path:       "./invalid/path/for/struct",
				OutputDir:  "./invalid/path/for/struct",
				Functions:  []string{},
			},
			structure: generator.Struct{
				Type:       "Dog",
				PluralType: "Dogs",
				Package:    "animals",
			},
			parseError: nil,
			shouldFail: true,
		},
		{
			description: "return error when fail to parse files",
			config: Config{
				StructName: "Cat",
				PluralName: "Cats",
				Path:       tempDir,
				OutputDir:  tempDir,
				Functions:  []string{},
			},
			structure: generator.Struct{
				Type:       "Cat",
				PluralType: "Cats",
				Package:    "animals",
			},
			parseError: errors.New("error"),
			shouldFail: true,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// given
			defer func() {
				if r := recover(); r != nil {
					require.FailNow(t, "test panicked, ", r)
				}
			}()

			parser := &mocks.Parser{}
			parser.On("FindStruct", testCase.config.StructName, testCase.config.Path).
				Return(testCase.structure, testCase.parseError)

			// when
			err := Generate(testCase.config, parser)

			// then
			assert.Equal(t, testCase.shouldFail, err != nil)

			if !testCase.shouldFail {
				generatedFileName := filepath.Join(testCase.config.Path, "generated_Dogs.go")
				assert.FileExists(t, generatedFileName)

				content, err := ioutil.ReadFile(generatedFileName)
				require.NoError(t, err)

				assert.NotEmpty(t, content)
			}
		})
	}

}
