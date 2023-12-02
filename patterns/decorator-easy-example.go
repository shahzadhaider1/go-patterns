package patterns

/*

Decorator Pattern:

The decorator pattern is a structural design pattern
that allows behavior to be added to an individual object,
either statically or dynamically, without affecting the
behavior of other objects from the same class.

Let's break down the key components and concepts of the decorator pattern:

Participants in the Decorator Pattern:

Component:
Defines the interface for objects that can have responsibilities added to them dynamically.

Concrete Component:
Implements the Component interface and provides the core functionality.

Decorator:
Maintains a reference to a Component object and defines an interface that conforms to the Component interface.
Can have additional responsibilities.

Concrete Decorator:
Adds new responsibilities to the Component and extends its behavior.

*/

// Coffee is Component Interface
type Coffee interface {
	Cost() int
}

// SimpleCoffee is Concrete Component
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() int {
	return 5
}

// CoffeeDecorator is a Decorator Interface
type CoffeeDecorator interface {
	Cost() int
}

// MilkDecorator - Concrete Decorators:
type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) Cost() int {
	return m.Coffee.Cost() + 2
}

type SugarDecorator struct {
	Coffee Coffee
}

func (s *SugarDecorator) Cost() int {
	return s.Coffee.Cost() + 1
}

//
// main code:
//
//coffee := &patterns.SimpleCoffee{}
//
// Wrap the simple coffee with decorators
//milkCoffee := &patterns.MilkDecorator{Coffee: coffee}
//sweetMilkCoffee := &patterns.SugarDecorator{Coffee: milkCoffee}
//
// Calculate the cost
//fmt.Println("Cost of simple coffee:", coffee.Cost())
//fmt.Println("Cost of milk coffee:", milkCoffee.Cost())
//fmt.Println("Cost of sweet milk coffee:", sweetMilkCoffee.Cost())
//
