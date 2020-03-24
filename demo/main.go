package main

import (
	"fmt"

	"github.com/FMorsbach/DecFL/control"
	"github.com/FMorsbach/dlog"
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
		dlog.Fatal(err)
	}

	status, err := control.Status()
	fmt.Println(status)

	for i := 0; i < CLIENTS; i++ {
		err = control.Iterate()
		if err != nil {
			dlog.Fatal(err)
		}
	}

	err = control.Aggregate()
	if err != nil {
		dlog.Fatal(err)
	}

	status, err = control.Status()
	fmt.Println(status)
}
