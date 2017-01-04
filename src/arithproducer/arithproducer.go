package arithproducer

import (
	"fmt"
	"log"

	"github.com/devopsgig/publisher/src/publisher"
	"github.com/devopsgig/utilities/src/utilities"
)

// Run trigger the task generation
func Run() {
	taskCh := make(chan *Arith)
	respCh := make(chan string)
	go Assembly(taskCh)
	for {
		go publisher.Send(<-taskCh, respCh)
		select {
		case resp := <-respCh:
			fmt.Println(resp)
		}
	}
}

// Assembly an arithmetic operation
func Assembly(taskCh chan *Arith) {
	maxValue := 1000
	for {
		// Set random value for a
		a, err := utilities.RandomInt(maxValue)
		if err != nil {
			log.Fatal(err)
		}
		// Set random value for b
		b, err := utilities.RandomInt(maxValue)
		if err != nil {
			log.Fatal(err)
		}
		// Initialize Arith struct
		arith := new(Arith)
		arith.A, arith.B = a, b
		arith.SetTaskType()
		arith.Operator = arith.GetOperator()
		taskCh <- arith
	}
}
