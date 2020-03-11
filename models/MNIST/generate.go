package MNIST

import (
	"io/ioutil"
	"log"
)

func GenerateInitialModel() (configuration string, weights string) {

	configuration = loadDataFromDisk("models/MNIST/configuration.txt")
	weights = loadDataFromDisk("models/MNIST/weights.txt")
	return
}

func loadDataFromDisk(file string) (data string) {
	var content []byte

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %d bytes from %s", len(content), file)

	data = string(content)
	return
}
