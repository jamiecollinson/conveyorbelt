package conveyorbelt

import (
	"errors"
)

// Slot models a section of the ConveyorBelt
type Slot struct {
	item    Item
	workers [2]*Worker
}

// Run triggers the Run method on each Worker, offering the current item and notifying the Worker if they can interact or another has already done so
func (s *Slot) Run() {
	canInteract := true
	for _, worker := range s.workers {
		item, didInteract := worker.Run(s.item, canInteract)
		if didInteract {
			s.item = item
			canInteract = false
		}
	}
}

// Pop removes the current Item from the Slot, leaving it empty
func (s *Slot) pop() Item {
	item := s.item
	s.item = Empty
	return item
}

// Push adds an Item to the Slot
func (s *Slot) push(item Item) error {
	if s.item != Empty {
		return errors.New("slot already contains an item")
	}
	s.item = item
	return nil
}

// NewSlot creates a new Slot
func NewSlot(product Item) *Slot {
	return &Slot{
		item: Empty,
		workers: [2]*Worker{
			NewWorker(product, 3),
			NewWorker(product, 3),
		},
	}
}
