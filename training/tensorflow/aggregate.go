package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func Aggregate(updates []string) (aggregatedWeights string) {

	for i, update := range updates {
		path := prefix + "res/" + strconv.Itoa(i) + "_trainingWeights.in"
		err := ioutil.WriteFile(path, []byte(update), 0644)
		if err != nil {
			log.Fatalf("Can't write update %d to %s. Got error: %s", i, path, err)
		}
		log.Printf("Wrote update %d to %s", i, path)

		defer func() {
			err := os.Remove(path)

			if err != nil && !os.IsNotExist(err) {
				panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", path, err))
			}
		}()
	}

	cmd := exec.Command(pythonPath, prefix+"aggregate.py", prefix+"res/", outputPath)

	defer func() {
		err := os.Remove(outputPath)
		if err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("Tried deleting %s after aggregation but got %s", outputPath, err))
		}
	}()

	log.Print("Executing: ", cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	log.Println("Aggregating completed")

	content, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}

	aggregatedWeights = string(content)

	return
}
