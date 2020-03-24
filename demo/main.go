package main

import (
	"fmt"
	"log"

	"github.com/FMorsbach/DecFL/control"
)

// TODO: make CL argument
const CLIENTS = 3

func init() {

	// TODO: Check if chain reachable

	// TODO: Check if storage reachable

}

func main() {

	_, err := control.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	status, err := control.Status()
	fmt.Println(status)

	for i := 0; i < CLIENTS; i++ {
		err = control.Iterate()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = control.Aggregate()
	if err != nil {
		log.Fatal(err)
	}

	status, err = control.Status()
	fmt.Println(status)
}
