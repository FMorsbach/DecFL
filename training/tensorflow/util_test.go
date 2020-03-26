package tensorflow

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/FMorsbach/dlog"
)

var testConfiguration string
var testWeights string

func init() {

	content, err := ioutil.ReadFile("testData/configuration.in")
	if err != nil {
		panic(err)
	}
	testConfiguration = string(content)

	content, err = ioutil.ReadFile("testData/initialWeights.in")
	if err != nil {
		panic(err)
	}
	testWeights = string(content)

	EnableDebug(false)
}

func TestCleanUpRessources(t *testing.T) {

	files := []string{configPath, weightsPath, outputPath}

	for _, path := range files {
		err := ioutil.WriteFile(path, []byte("RandomData"), 0644)
		if err != nil {
			dlog.Fatal(err)
		}
	}

	cleanUpRessources()

	for _, path := range files {
		_, err := os.Stat(path)
		if err == nil {
			t.Errorf("%s still exists", path)
		} else if os.IsNotExist(err) {

		} else {
			panic(err)
		}
	}
}

func TestWriteModelToDisk(t *testing.T) {

	err := writeModelToDisk(testConfiguration, testWeights)
	if err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if string(content) != testConfiguration {
		t.Errorf("Wrote %s as configuration but wanted %s", string(content), testConfiguration)
	}

	content, err = ioutil.ReadFile(weightsPath)
	if err != nil {
		panic(err)
	}

	if string(content) != testWeights {
		t.Errorf("Wrote %s as configuration but wanted %s", string(content), testWeights)
	}

	cleanUpRessources()
}

func TestReadModelUpdatesFromDisk(t *testing.T) {

	err := ioutil.WriteFile(outputPath, []byte(testWeights), 0644)
	if err != nil {
		panic(err)
	}

	content, err := readUpdatesFromDisk()
	if err != nil {
		t.Error(err)
	}

	if content != testWeights {
		t.Errorf("Read %s as weights but wanted %s", content, testWeights)
	}

	cleanUpRessources()
}
