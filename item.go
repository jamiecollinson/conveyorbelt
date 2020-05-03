package conveyorbelt

import (
	"math/rand"
	"time"
)

// Item models the components & products on the ConveyorBelt
type Item int

const (
	Empty Item = iota
	ComponentA
	ComponentB
	FinishedProduct
)

// canPickup returns whether a Worker can pick up the Item from the line
func (i Item) canPickup() bool {
	return i == ComponentA || i == ComponentB
}

// canReplace returns whether a Worker can place an Item when this Item is present
func (i Item) canReplace() bool {
	return i == Empty
}

// prerequisites returns a list of necessary prerequisites to assemble the Item
func (i Item) prerequisites() []Item {
	if i == FinishedProduct {
		return []Item{ComponentA, ComponentB}
	}
	return []Item{}
}

// canAssemble returns whether assembly can start based on a list of prerequisites
func (i Item) canAssemble(items []Item) bool {
	if len(i.prerequisites()) == 0 {
		// we can't assemble things which don't have precursors
		return false
	}
	for _, item := range i.prerequisites() {
		found := false
		for _, testItem := range items {
			if item == testItem {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}
	return true
}

// String implements the Stringer interface
func (i Item) String() string {
	return []string{"Empty", "Component A", "Component B", "Finished Product"}[i]
}

// ItemGenerators model the arrival of new Items onto the ConveyorBelt
type ItemGenerator interface {
	Generate() Item
}

// EqualProbabilityItemGenerator creates Empty, ComponentA, ComponentB with equal chance
type EqualProbabilityItemGenerator struct{}

// Generate implements the ItemGenerator interface
func (i *EqualProbabilityItemGenerator) Generate() Item {
	rand.Seed(time.Now().UnixNano())
	return Item(rand.Intn(3))
}
