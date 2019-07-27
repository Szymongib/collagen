package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Szymongib/go-collagen/internal/collagen"

	"github.com/stretchr/testify/assert"
)

func Test_validateConfig(t *testing.T) {

	for _, testCase := range []struct {
		description   string
		flagConfig    flagConfig
		shouldSucceed bool
	}{
		{
			description:   "should succeed if proper config provided",
			flagConfig:    flagConfig{name: "Dog"},
			shouldSucceed: true,
		},
		{
			description:   "should fail if name not provided",
			flagConfig:    flagConfig{name: ""},
			shouldSucceed: false,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			// when
			err := validateFlagConfig(testCase.flagConfig)

			// then
			succeeded := err == nil
			assert.Equal(t, testCase.shouldSucceed, succeeded)
		})
	}

}

func Test_prepareConfig(t *testing.T) {

	executionDir, err := os.Getwd()
	require.NoError(t, err)

	for _, testCase := range []struct {
		description    string
		flagConfig     flagConfig
		expectedConfig collagen.Config
	}{
		{
			description: "only name provided",
			flagConfig: flagConfig{
				name: "Dog",
			},
			expectedConfig: collagen.Config{
				StructName: "Dog",
				PluralName: "Dogs",
				Path:       executionDir,
				OutputDir:  executionDir,
				Functions:  nil,
			},
		},
		{
			description: "all values provided",
			flagConfig: flagConfig{
				name:      "Dog",
				plural:    "MyDogs",
				dir:       "/my-project",
				functions: "Filter,Contains",
			},
			expectedConfig: collagen.Config{
				StructName: "Dog",
				PluralName: "MyDogs",
				Path:       "/my-project",
				OutputDir:  "/my-project",
				Functions:  []string{"Filter", "Contains"},
			},
		},
	} {
		t.Run("should prepare config when "+testCase.description, func(t *testing.T) {
			// when
			config, err := prepareConfig(testCase.flagConfig)

			// then
			require.NoError(t, err)
			assert.Equal(t, testCase.expectedConfig, config)
		})
	}

}
