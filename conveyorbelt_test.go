package conveyorbelt

import (
	"testing"
)

func TestConveyorBelt(t *testing.T) {
	belt := NewConveyorBelt(FinishedProduct, 3)
	for i := 0; i < 100; i++ {
		belt.Run()
	}
	if len(belt.output) < 100 {
		t.Error("Output should equal steps")
	}
}

func TestConveyorBeltHelpers(t *testing.T) {
	belt := NewConveyorBelt(FinishedProduct, 3)
	belt.output = []Item{Empty, ComponentA, ComponentA, ComponentB, FinishedProduct}

	output := belt.OutputCount()
	if output[FinishedProduct] != 1 || output[ComponentA] != 2 || output[ComponentB] != 1 {
		t.Error("Output counts incorrect")
	}

	if belt.String() != "Finished Product: 1, Component A: 2, Component B: 1" {
		t.Error("String format incorrect")
	}
}
