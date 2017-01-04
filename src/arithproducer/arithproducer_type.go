package arithproducer

import (
	"math/rand"

	"github.com/devopsgig/publisher/src/publisher"
)

// Arith wraps arithmetic task type
type Arith struct {
	Meta     publisher.Meta
	A        int    `json:"a"`
	B        int    `json:"b"`
	Operator string `json:"operator"`
	Err      error  `json:"error"`
}

// SetUUID add a UUID to the task
func (a *Arith) SetUUID(UUID string) {
	a.Meta.UUID = UUID
}

// SetTaskType so the publish can decide where to send it
func (a *Arith) SetTaskType() {
	a.Meta.TaskType = "arithmetic"
}

// GetTaskType return the type of the task
func (a *Arith) GetTaskType() string {
	return a.Meta.TaskType
}

// GetOperator return a random arithmetic operation
func (a *Arith) GetOperator() string {
	operators := []string{"+", "-", "*", "/"}
	return operators[rand.Intn(len(operators))]
}
