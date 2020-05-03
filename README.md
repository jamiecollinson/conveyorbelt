# Conveyor Belt

A program to simulate the action of a system consisting of a conveyor belt and number of workers, who produce products from input components.

The program calculates the output number of products, and unused components.

## Problem Description

There is a factory production line around a single a conveyor belt.

Components (of type A and B) come onto the start of the belt at random intervals; workers must take one component of each type from the belt as they come past, and combine them to make a finished product.

The belt is divided into fixed-size slots; each slot can hold only one component or one finished product. There are a number of worker stations on either side of the belt, spaced to match the size of the slots on the belt, like this (fixed-width font ASCII pic):

```
       v   v   v   v   v          workers
     ---------------------
  -> | A |   | B | A | P | ->     conveyor belt
     ---------------------
       ^   ^   ^   ^   ^          workers
```

In each unit of time, the belt moves forwards one position, and there is time for a worker on one side of each slot to EITHER take an item from the slot or replace an item onto the belt. The worker opposite them can't touch the same belt slot while they do this. (So you can't have one worker picking something from a slot while their counterpart puts something down in the same place).

Once a worker has collected one of both types of component, they can begin assembling the finished product. This takes an amount of time, so they will only be ready to place the assembled product back on the belt on the fourth subsequent slot. While they are assembling the product, they can't touch the conveyor belt. Workers can only hold two items (component or product) at a time: one in each hand.

Create a simulation of this, with three pairs of workers. At each time interval, the slot at the start of the conveyor belt should have an equal (1/3) chance of containing nothing, a component A or a component B.
Run the simulation for 100 steps, and compute how many finished products come off the production line, and how many components of each type go through the production line without being picked up by any workers.

A few pointers:
 - You should expect to spend no more than two or three hours on this challenge.
 - The code does not have to be 'production quality', but we will be looking for evidence that it's written to be somewhat flexible, and that a third party would be able to read and maintain it.
 - Be sure to state (or comment) your assumptions.
 - How would you alter your answer if the length of the conveyor belt were 3x the original?

## Assumptions

- On each slot the top worker acts (if possible) before the bottom worker. There is no coordination between them.
- Workers act greedily and without looking beyond their current slot. If they can place a product they always will, and similarly if they can pick up a needed component they always will.
- Conveyor belt length is relatively likely to be variable, as is the probabilities of inputs arriving. Comparatively the `A + B -> P` working logic, and the assumptions on number of items held, number of workers per slot etc. are more likely to be fixed.
- Given the input chances, we expect 33 finished products as an approximate upper limit.

## Sample output

```
go run cmd/main.go

Single conveyor belt run (100 steps):
Finished Product: 23, Component A: 9, Component B: 1

Averages over 1000 trials (100 steps per trial):
Finished Product: mean 25.36, standard deviation 2.87
Unused Component A: mean 2.89, standard deviation 3.58
Unused Component B: mean 2.98, standard deviation 3.68
```

The files `Finished Products Histogram.png`, `Unused Component A Histogram.png` and `Unused Component B Histogram.png` are also created in `img`:

![Finished Products](./img/Finished%20Products%20Histogram.png?raw=true)

![Unused Component A](./img/Unused%20Component%20A%20Histogram.png?raw=true)

![Unused Component B](./img/Unused%20Component%20B%20Histogram.png?raw=true)

## Use

The only dependency is `Go` (1.11 or greater recommended to support easier dependency management).

Run via `make run` or directly with `go run cmd/main.go`.

`make build` will compile for Windows, MacOS and Linux. For convenience compiled binaries are already located in `bin`.

## External API

Can be shown with `go doc -all`:

```
package conveyorbelt // import "conveyorbelt"


TYPES

type ConveyorBelt struct {
	// Has unexported fields.
}
    ConveyorBelt models a conveyor belt

func NewConveyorBelt(product Item, nSlots int) *ConveyorBelt
    NewConveyorBelt creates a new ConveyorBelt producing an Item with a number
    of Slots

func (c *ConveyorBelt) OutputCount() map[Item]int
    OutputCount gives a summary of the current output of the ConveyorBelt

func (c *ConveyorBelt) Run()
    Run triggers the Run method of all child Slots, brings new Items onto the
    belt, moves Items between the Slots and stores the output

func (c *ConveyorBelt) String() string
    String implements the Stringer interface

type EqualProbabilityItemGenerator struct{}
    EqualProbabilityItemGenerator creates Empty, ComponentA, ComponentB with
    equal chance

func (i *EqualProbabilityItemGenerator) Generate() Item
    Generate implements the ItemGenerator interface

type Item int
    Item models the components & products on the ConveyorBelt

const (
	Empty Item = iota
	ComponentA
	ComponentB
	FinishedProduct
)
func (i Item) String() string
    String implements the Stringer interface

type ItemGenerator interface {
	Generate() Item
}
    ItemGenerators model the arrival of new Items onto the ConveyorBelt

type Slot struct {
	// Has unexported fields.
}
    Slot models a section of the ConveyorBelt

func NewSlot(product Item) *Slot
    NewSlot creates a new Slot

func (s *Slot) Run()
    Run triggers the Run method on each Worker, offering the current item and
    notifying the Worker if they can interact or another has already done so

type Worker struct {
	// Has unexported fields.
}
    Worker models a worker assigned to a Slot on the ConveyorBelt

func NewWorker(product Item, assemblyTime int) *Worker
    NewWorker creates a new Worker who produces an Item in a given assemblyTime

func (w *Worker) Run(slotItem Item, canInteract bool) (returnItem Item, didInteract bool)
    Run implements the logic of interacting with the Slot, and the assembly of
    the product
```

## Testing

Run unit tests with `make test`, or check coverage with `make cover`:

```
go test -coverprofile=coverage.out && go tool cover -func=coverage.out
PASS
coverage: 100.0% of statements
ok  	conveyorbelt	0.367s
conveyorbelt/conveyorbelt.go:18:	Run			100.0%
conveyorbelt/conveyorbelt.go:31:	OutputCount		100.0%
conveyorbelt/conveyorbelt.go:42:	String			100.0%
conveyorbelt/conveyorbelt.go:52:	NewConveyorBelt		100.0%
conveyorbelt/item.go:19:		canPickup		100.0%
conveyorbelt/item.go:24:		canReplace		100.0%
conveyorbelt/item.go:29:		prerequisites		100.0%
conveyorbelt/item.go:37:		canAssemble		100.0%
conveyorbelt/item.go:58:		String			100.0%
conveyorbelt/item.go:71:		Generate		100.0%
conveyorbelt/slot.go:14:		Run			100.0%
conveyorbelt/slot.go:26:		pop			100.0%
conveyorbelt/slot.go:33:		push			100.0%
conveyorbelt/slot.go:42:		NewSlot			100.0%
conveyorbelt/worker.go:14:		Run			100.0%
conveyorbelt/worker.go:55:		hasItem			100.0%
conveyorbelt/worker.go:65:		needsItem		100.0%
conveyorbelt/worker.go:75:		addItem			100.0%
conveyorbelt/worker.go:86:		removeItem		100.0%
conveyorbelt/worker.go:96:		hasCompleteProduct	100.0%
conveyorbelt/worker.go:101:		NewWorker		100.0%
total:					(statements)		100.0%
```

## Todo / notes

- To make more extendable consider the use of interfaces over concrete structs for `Item`, `Slot`, `Worker`. An example is `ItemGenerator` which allows easy replacement with different `Item` generation rules.
- The logic within the `Worker.Run` method could be better expressed as a state machine.
- The code was developed using test driven development, but the tests are largely covering happy paths and would benefit from more edge case coverage and better segregation.
- The multi-trial run could trivially be made faster on multi-core machines using goroutines & channels.
- The driver program `cmd/main.go` is a quick script and could be generalised
