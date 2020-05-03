package conveyorbelt

import (
	"testing"
)

func TestWorkerAddRemove(t *testing.T) {
	worker := NewWorker(FinishedProduct, 1)

	if worker.hasItem(ComponentA) {
		t.Error("Worker should begin with no ComponentA")
	}
	if worker.hasItem(ComponentB) {
		t.Error("Worker should begin with no ComponentB")
	}
	if worker.hasItem(FinishedProduct) {
		t.Error("Worker should begin with no FinishedProduct")
	}

	worker.addItem(ComponentA)
	if !worker.hasItem(ComponentA) {
		t.Error("Worker should have ComponentA")
	}
	worker.removeItem(ComponentA)
	if worker.hasItem(ComponentA) {
		t.Error("Worker should no longer have ComponentA")
	}
}

func TestWorkerRun(t *testing.T) {
	worker := NewWorker(FinishedProduct, 1)

	item, didInteract := worker.Run(ComponentA, false)
	if item != Empty || didInteract {
		t.Error("Worker can not interact unless allowed")
	}

	item, didInteract = worker.Run(ComponentA, true)
	if item != Empty || !didInteract || !worker.hasItem(ComponentA) {
		t.Error("Worker should have picked up ComponentA")
	}
	item, didInteract = worker.Run(ComponentA, true)
	if item != Empty || didInteract || !worker.hasItem(ComponentA) {
		t.Error("Worker should not have picked up second ComponentA")
	}
	item, didInteract = worker.Run(ComponentB, true)
	if item != Empty || !didInteract || !worker.hasItem(FinishedProduct) || worker.hasCompleteProduct() {
		t.Error("Worker should have picked up ComponentB and begun making finished product")
	}
	item, didInteract = worker.Run(Empty, true)
	if item != Empty || didInteract {
		t.Error("Worker should be assembling")
	}
	item, didInteract = worker.Run(Empty, true)
	if item != FinishedProduct {
		t.Error("Worker should have completed assembly and deposited")
	}
}
