package conveyorbelt

import "testing"

func TestItem(t *testing.T) {
	if FinishedProduct.String() != "Finished Product" {
		t.Error("Finished Product string incorrect")
	}

	if Empty.canPickup() != false || FinishedProduct.canPickup() != false {
		t.Error("Can only pick up components")
	}

	if ComponentA.canPickup() != true || ComponentB.canPickup() != true {
		t.Error("Components can be picked up")
	}

	if len(FinishedProduct.prerequisites()) != 2 {
		t.Error("FinishedProduct has two prerequisites")
	}

	if len(Empty.prerequisites()) > 0 || len(ComponentA.prerequisites()) > 0 || len(ComponentB.prerequisites()) > 0 {
		t.Error("Only FinishedProduct has prerequisites")
	}

	if FinishedProduct.canAssemble([]Item{}) || FinishedProduct.canAssemble([]Item{ComponentA}) || FinishedProduct.canAssemble([]Item{ComponentB}) {
		t.Error("FinishedProduct can only be created with all prerequisites met")
	}

	if !FinishedProduct.canAssemble([]Item{ComponentA, ComponentB}) {
		t.Error("FinishedProduct can be assembled with prerequisites")
	}

	if Empty.canAssemble([]Item{}) || ComponentA.canAssemble([]Item{}) || ComponentB.canAssemble([]Item{}) {
		t.Error("Only FinishedProduct can be assembled")
	}
}
