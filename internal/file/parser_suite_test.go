package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	testDataDir string
)

func TestMain(m *testing.M) {
	// setup
	var err error

	testDataDir, err = ioutil.TempDir(".", "testdata")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func() {
		err := os.RemoveAll(testDataDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	err = setupValidTestData(testDataDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = setupInvalidTestData(testDataDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// run tests
	m.Run()

}

var (
	fileContentWithStruct = []byte(`package valid

type Dog struct {
	Name   string
	Paws   []Paw
	Labels map[string][]string
}

type Paw struct {
	Size int
}
`)

	fileWithTestsContent = []byte(`package valid_test

func TestDog(t *testing.T){

}
`)

	fileContentWithoutStruct = []byte(`package valid

func Bark() string {
	return "woof"
}
`)

	fileContentWithoutPackage = []byte(`
func doStuff() {

}
`)
)

func setupValidTestData(testDataDir string) error {
	validDataDir := filepath.Join(testDataDir, "valid")

	err := os.Mkdir(validDataDir, os.ModePerm)
	if err != nil {
		return err
	}

	modelFile := filepath.Join(validDataDir, "model.go")
	err = ioutil.WriteFile(modelFile, fileContentWithStruct, os.ModePerm)
	if err != nil {
		return err
	}

	fileWithoutStruct := filepath.Join(validDataDir, "file_without_struct.go")
	err = ioutil.WriteFile(fileWithoutStruct, fileContentWithoutStruct, os.ModePerm)
	if err != nil {
		return err
	}

	testFile := filepath.Join(validDataDir, "file_without_struct_test.go")
	err = ioutil.WriteFile(testFile, fileWithTestsContent, os.ModePerm)
	if err != nil {
		return err
	}

	noGoFile := filepath.Join(validDataDir, "no_go_file")
	err = ioutil.WriteFile(noGoFile, fileContentWithStruct, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func setupInvalidTestData(testDataDir string) error {
	invalidDataDir := filepath.Join(testDataDir, "invalid")

	err := os.Mkdir(invalidDataDir, os.ModePerm)
	if err != nil {
		return err
	}

	modelFile := filepath.Join(invalidDataDir, "model.go")
	err = ioutil.WriteFile(modelFile, fileContentWithStruct, os.ModePerm)
	if err != nil {
		return err
	}

	fileWithoutStruct := filepath.Join(invalidDataDir, "no_pkg_file.go")
	err = ioutil.WriteFile(fileWithoutStruct, fileContentWithoutPackage, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
