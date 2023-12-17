package main

// IBurger is the interface that all products must implement
type iBurger interface {
	prepare()
}

type cheeseBurger struct{}

type chickenBurger struct{}

type veggieBurger struct{}

func (c *cheeseBurger) prepare() {
	println("Preparing Cheese Burger")
}

func (c *chickenBurger) prepare() {
	println("Preparing Chicken Burger")
}

func (c *veggieBurger) prepare() {
	println("Preparing Veggie Burger")
}

// BurgerFactory is the factory that creates products
type burgerFactory struct{}

func (b *burgerFactory) createBurger(burgerType string) iBurger {
	switch burgerType {
	case "cheese":
		return &cheeseBurger{}
	case "chicken":
		return &chickenBurger{}
	case "veggie":
		return &veggieBurger{}
	default:
		return nil
	}
}

func main() {
	burgerFactory := &burgerFactory{}
	burgerFactory.createBurger("cheese").prepare()
	burgerFactory.createBurger("chicken").prepare()
	burgerFactory.createBurger("veggie").prepare()
}
