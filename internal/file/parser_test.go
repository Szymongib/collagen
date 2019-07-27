package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindStruct(t *testing.T) {

	t.Run("should parse tests files", func(t *testing.T) {
		// given
		dataPath := fmt.Sprintf("%s/valid", testDataDir)
		parser := &Parser{}

		// when
		dogStruct, err := parser.FindStruct("Dog", dataPath)

		// then
		require.NoError(t, err)
		assert.NotEmpty(t, dogStruct)
	})

}

func TestFindStruct_Error(t *testing.T) {

	t.Run("should return error when invalid path", func(t *testing.T) {
		// given
		parser := &Parser{}

		// when
		dogStruct, err := parser.FindStruct("Dog", "the path is invalid ___ \\//\\\\\\")

		// then
		require.Error(t, err)
		assert.Empty(t, dogStruct)
	})

	t.Run("should return error if file without package", func(t *testing.T) {
		// given
		dataPath := fmt.Sprintf("%s/invalid", testDataDir)
		parser := &Parser{}

		// when
		dogStruct, err := parser.FindStruct("Dog", dataPath)

		// then
		require.Error(t, err)
		assert.Empty(t, dogStruct)
	})

	t.Run("should return error if specified struct not found in path", func(t *testing.T) {
		// given
		dataPath := fmt.Sprintf("%s/valid", testDataDir)
		parser := &Parser{}

		// when
		dogStruct, err := parser.FindStruct("Cat", dataPath)

		// then
		require.Error(t, err)
		assert.Empty(t, dogStruct)
	})

}
