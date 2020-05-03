package conveyorbelt

import "errors"

// Worker models a worker assigned to a Slot on the ConveyorBelt
type Worker struct {
	product          Item
	items            [2]Item
	productCountdown int
	assemblyTime     int
}

// Run implements the logic of interacting with the Slot, and the assembly of the product
func (w *Worker) Run(slotItem Item, canInteract bool) (returnItem Item, didInteract bool) {
	didInteract = false

	// if we're working on assembly, reduce countdown and continue
	if w.productCountdown > 0 {
		w.productCountdown -= 1
		return Empty, false
	}

	// check if we can act further
	if !canInteract {
		return Empty, false
	}

	// if product is finished and it is possible to place it on conveyor, do so
	if w.hasCompleteProduct() && slotItem.canReplace() {
		w.removeItem(FinishedProduct)
		return FinishedProduct, true
	}

	// if we need the offered item, pick it up
	if slotItem.canPickup() && w.needsItem(slotItem) {
		err := w.addItem(slotItem)
		if err == nil {
			didInteract = true
		}
	}

	// if we have all prerequisites, begin work on assembly
	if w.product.canAssemble(w.items[:]) {
		for _, item := range w.product.prerequisites() {
			w.removeItem(item)
		}
		w.addItem(FinishedProduct)
		w.productCountdown = w.assemblyTime
	}

	return Empty, didInteract
}

// hasItem returns whether a Worker has the Item
func (w *Worker) hasItem(newItem Item) bool {
	for _, item := range w.items {
		if item == newItem {
			return true
		}
	}
	return false
}

// needsItem returns whether a Worker needs the Item
func (w *Worker) needsItem(newItem Item) bool {
	for _, item := range w.product.prerequisites() {
		if newItem == item && !w.hasItem(newItem) {
			return true
		}
	}
	return false
}

// addItem adds the Item to the Worker's storage
func (w *Worker) addItem(newItem Item) error {
	for i, item := range w.items {
		if item == Empty {
			w.items[i] = newItem
			return nil
		}
	}
	return errors.New("No empty hand")
}

// removeItem removes the Item from the Worker's storage
func (w *Worker) removeItem(newItem Item) {
	for i, item := range w.items {
		if item == newItem {
			w.items[i] = Empty
			break
		}
	}
}

// hasCompleteProduct returns whether assembly has finished
func (w *Worker) hasCompleteProduct() bool {
	return w.productCountdown == 0 && w.hasItem(w.product)
}

// NewWorker creates a new Worker who produces an Item in a given assemblyTime
func NewWorker(product Item, assemblyTime int) *Worker {
	return &Worker{
		product:          product,
		items:            [2]Item{Empty, Empty},
		productCountdown: 0,
		assemblyTime:     assemblyTime,
	}
}
