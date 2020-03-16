package tensorflow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

func Aggregate(updates []string) (aggregatedWeights string) {

	for i, update := range updates {
		path := prefix + "testData/aggregation/weight" + strconv.Itoa(i)
		err := ioutil.WriteFile(path, []byte(update), 0644)
		if err != nil {
			log.Fatalf("Can't write update %d to %s. Got %s", i, path, err)
		}
	}

	cmd := exec.Command(pythonPath, prefix+"aggregate.py", prefix+"testData/aggregation/", prefix+"testData/output.txt")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}
	log.Println("Aggregating completed")

	content, err := ioutil.ReadFile(prefix + "testData/output.txt")
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}

	aggregatedWeights = string(content)
	return
}
