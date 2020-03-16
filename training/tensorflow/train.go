package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

const prefix string = "training/tensorflow/"
const pythonPath string = prefix + "venv/bin/python"
const scriptPath string = prefix + "train.py"
const configPath string = prefix + "res/configuration.txt"
const weightsPath string = prefix + "res/weights.txt"
const outputPath string = prefix + "res/output.txt"

func Run(configuration string, weights string) (updatedWeights string) {

	updatedWeights = trainByFile(configuration, weights)
	return
}

func trainByFile(configuration string, weights string) (updatedWeights string) {

	err := writeModelToDisk(configuration, weights)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Wrote model to disk")

	cmd := exec.Command(pythonPath, scriptPath, configPath, weightsPath, outputPath)
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	log.Println("Training completed")

	updatedWeights, err = readModelUpdatesFromDisk()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read model back from disk")

	return
}

func readModelUpdatesFromDisk() (weights string, err error) {

	var content []byte

	content, err = ioutil.ReadFile(outputPath)
	if err != nil {
		return
	}

	log.Printf("Read %d bytes from %s", len(content), outputPath)
	weights = string(content)

	err = nil
	return
}

func writeModelToDisk(configuration string, weights string) (err error) {

	err = ioutil.WriteFile(configPath, []byte(configuration), 0644)
	if err != nil {
		return
	}
	log.Printf("Wrote %d bytes as %s to disk.", len([]byte(configuration)), configPath)

	ioutil.WriteFile(weightsPath, []byte(weights), 0644)
	if err != nil {
		return
	}
	log.Printf("Wrote %d bytes as %s to disk.", len([]byte(weights)), weightsPath)

	return nil
}
