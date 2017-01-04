package arithproducer

import (
	"reflect"
	"strings"
	"testing"
)

// TestAssemblyType test the task type.
// Expected: "*footypes.Arith"
func TestAssemblyType(t *testing.T) {
	testCh := make(chan *Arith)
	go Assembly(testCh)
	observed := reflect.TypeOf(<-testCh).String()
	expected := "*arithproducer.Arith"
	if observed != expected {
		t.Errorf("Expecting type %s. Observed %s", expected, observed)
	}
}

// TestAssemblyValueA test if the Arith.A value is in expected range.
// Expected: any int [0, 1000[
func TestAssemblyValueA(t *testing.T) {
	testCh := make(chan *Arith)
	go Assembly(testCh)
	task := <-testCh
	if task.A <= 0 || task.A > 1000 {
		t.Errorf("Expected an A in interval [0, 1000[. Observed %d", task.A)
	}
}

// TestAssemblyValueB test if the Arith.B value is in expected range.
// Expected: any int [0, 1000[
func TestAssemblyValueB(t *testing.T) {
	testCh := make(chan *Arith)
	go Assembly(testCh)
	task := <-testCh
	if task.B <= 0 || task.B > 1000 {
		t.Errorf("Expected an B in interval [0, 1000[. Observed %d", task.B)
	}
}

// TestTestAssemblyMetaTaskValue test Arith.Meta.Task value.
// Expected: "arithmetic"
func TestAssemblyMetaTaskValue(t *testing.T) {
	testCh := make(chan *Arith)
	go Assembly(testCh)
	task := <-testCh
	expected := "arithmetic"
	if task.Meta.TaskType != expected {
		t.Errorf("Expected meta task value to be %s. Observed %s", expected, task.Meta.TaskType)
	}
}

// TestAssemblyOperatorValue test Arith.Operator value.
// Expected: "+" or "-" or "*" or "/"
func TestAssemblyOperatorValue(t *testing.T) {
	operators := []string{"+", "-", "*", "/"}
	testCh := make(chan *Arith)
	go Assembly(testCh)
	task := <-testCh
	observed := task.Operator
	if strings.Index(strings.Join(operators[:], ""), observed) == -1 {
		t.Errorf("Expected one of the following operators %v. Observed '%s'", operators, task.Operator)
	}
}

// TestAssemblyOperatorValue test operators function.
// Expected: "+" or "-" or "*" or "/"
func TestGetOperator(t *testing.T) {
	op := []string{"+", "-", "*", "/"}
	arith := new(Arith)
	observed := arith.GetOperator()
	if strings.Index(strings.Join(op[:], ""), observed) == -1 {
		t.Errorf("Expected one of the following operators %v. Observed '%s'", op, observed)
	}
}

// TestSetUUID check if UUID is being correctly added to Arith struct
func TestSetUUID(t *testing.T) {
	UUID := "123e4567-e89b-12d3-a456-426655440000"
	arith := new(Arith)
	arith.SetUUID(UUID)
	observed := arith.Meta.UUID
	if observed != UUID {
		t.Errorf("For UUID: %s. The same value was exected. Instead, got: %s", UUID, observed)
	}
}

// TestSetTaskType check if method is adding the correct value to Arith struct
func TestSetTaskType(t *testing.T) {
	arith := new(Arith)
	arith.SetTaskType()
	expected := "arithmetic"
	observed := arith.Meta.TaskType
	if observed != expected {
		t.Errorf("Observed %s, expected: %s", observed, expected)
	}
}

// TestGetTaskType check if the correct value is being returned
func TestGetTaskType(t *testing.T) {
	arith := new(Arith)
	arith.Meta.TaskType = "foo"
	expected := "foo"
	observed := arith.GetTaskType()
	if observed != expected {
		t.Errorf("Observed %s, expected: %s", observed, expected)
	}
}
