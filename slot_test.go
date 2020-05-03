package conveyorbelt

import (
	"testing"
)

func TestSlotPushPop(t *testing.T) {
	slot := NewSlot(FinishedProduct)
	if slot.pop() != Empty {
		t.Error("New slots should start empty")
	}
	err := slot.push(ComponentA)
	if err != nil {
		t.Error("Pushing into an empty slot should work")
	}
	if slot.pop() != ComponentA {
		t.Error("Pop should return last item pushed")
	}
	_ = slot.push(ComponentA)
	err = slot.push(ComponentB)
	if err == nil {
		t.Error("Pushing a full slot should throw an error")
	}
	if slot.pop() != ComponentA {
		t.Error("Pushing a full slot should not change the contents")
	}
}

func TestSlotRun(t *testing.T) {
	slot := NewSlot(FinishedProduct)
	slot.push(ComponentA)
	slot.Run()
	if slot.pop() != Empty {
		t.Error("ComponentA should have been picked up by worker")
	}
	slot.push(ComponentB)
	slot.Run()
	if slot.pop() != Empty {
		t.Error("ComponentB should have been picked up by worker")
	}
	slot.Run()
	if slot.pop() != Empty {
		t.Error("FinishedProduct should be under assembly")
	}
	slot.Run()
	if slot.pop() != Empty {
		t.Error("FinishedProduct should be under assembly")
	}
	slot.Run()
	if slot.pop() != Empty {
		t.Error("FinishedProduct should be under assembly")
	}
	slot.Run()
	if slot.pop() != FinishedProduct {
		t.Error("FinishedProduct should have been deposited by worker")
	}
}
