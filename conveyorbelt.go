package conveyorbelt

import (
	"fmt"
	"strings"
)

// ConveyorBelt models a conveyor belt
type ConveyorBelt struct {
	product Item
	slots   []*Slot
	input   ItemGenerator
	output  []Item
}

// Run triggers the Run method of all child Slots, brings new Items onto the belt,
// moves Items between the Slots and stores the output
func (c *ConveyorBelt) Run() {
	var next, last Item
	next = c.input.Generate()
	for _, slot := range c.slots {
		slot.Run()
		last = slot.pop()
		slot.push(next)
		next = last
	}
	c.output = append(c.output, next)
}

// OutputCount gives a summary of the current output of the ConveyorBelt
func (c *ConveyorBelt) OutputCount() map[Item]int {
	var result = make(map[Item]int)

	for _, item := range c.output {
		result[item] += 1
	}

	return result
}

// String implements the Stringer interface
func (c *ConveyorBelt) String() string {
	output := c.OutputCount()
	relevantCounts := []string{fmt.Sprintf("%s: %d", c.product, output[c.product])}
	for _, item := range c.product.prerequisites() {
		relevantCounts = append(relevantCounts, fmt.Sprintf("%s: %d", item, output[item]))
	}
	return strings.Join(relevantCounts, ", ")
}

// NewConveyorBelt creates a new ConveyorBelt producing an Item with a number of Slots
func NewConveyorBelt(product Item, nSlots int) *ConveyorBelt {
	var slots []*Slot
	for i := 0; i < nSlots; i++ {
		slots = append(slots, NewSlot(product))
	}
	return &ConveyorBelt{
		product: product,
		slots:   slots,
		input:   &EqualProbabilityItemGenerator{},
	}
}
