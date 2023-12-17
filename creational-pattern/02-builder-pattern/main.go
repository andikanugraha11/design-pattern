package main

import "fmt"

// burger struct, your basic canvas
type burger struct {
	bun     string
	patty   string
	salad   string
	sauce   string
	topping string
}

// burgerBuilder is the interface that all builders must implement
type burgerBuilder interface {
	setBun(bun string)
	setPatty(patty string)
	setSalad(salad string)
	setSauce(sauce string)
	setTopping(topping string)
	getBurger() *burger
}

// genericBurgerBuilder for all your burger needs
type genericBurgerBuilder struct {
	burger *burger
}

// setBun sets the bun of the burger
func (b *genericBurgerBuilder) setBun(bun string) {
	b.burger.bun = bun
}

// setPatty sets the patty of the burger
func (b *genericBurgerBuilder) setPatty(patty string) {
	b.burger.patty = patty
}

// setSalad sets the salad of the burger
func (b *genericBurgerBuilder) setSalad(salad string) {
	b.burger.salad = salad
}

// setSauce sets the sauce of the burger
func (b *genericBurgerBuilder) setSauce(sauce string) {
	b.burger.sauce = sauce
}

// setTopping sets the topping of the burger
func (b *genericBurgerBuilder) setTopping(topping string) {
	b.burger.topping = topping
}

// getBurger returns the built burger
func (b *genericBurgerBuilder) getBurger() *burger {
	return b.burger
}

// burgerDirector is the director that controls the building process
type burgerDirector struct {
	builder burgerBuilder
}

// setBuilder sets the builder for the director
func (d *burgerDirector) setBuilder(builder burgerBuilder) {
	d.builder = builder
}

// buildBurger builds the burger using the builder
func (d *burgerDirector) buildBurger() *burger {
	if d.builder == nil {
		return nil // handle the case where builder is not set
	}
	return d.builder.getBurger()
}

func main() {
	director := burgerDirector{}

	// Building a cheeseburger
	cheeseBurgerBuilder := genericBurgerBuilder{&burger{}}
	director.setBuilder(&cheeseBurgerBuilder)
	director.builder.setPatty("beef patty") // set beef patty for cheeseburger
	director.builder.setTopping("cheese")   // set cheese for cheeseburger
	director.builder.setSauce("mayo")       // set mayo for cheeseburger
	director.builder.setSalad("lettuce")    // set lettuce for cheeseburger
	director.builder.setBun("sesame")       // set sesame bun for cheeseburger
	cheeseburger := director.buildBurger()
	fmt.Printf("Cheeseburger: %+v\n", cheeseburger)

	// Building a chicken burger
	chickenBurgerBuilder := genericBurgerBuilder{&burger{}}
	director.setBuilder(&chickenBurgerBuilder)
	director.builder.setPatty("chicken patty") // set chicken patty for chicken burger
	director.builder.setTopping("cheese")      // set cheese for chicken burger
	director.builder.setSauce("mayo")          // set mayo for chicken burger
	director.builder.setSalad("lettuce")       // set lettuce for chicken burger
	director.builder.setBun("sesame")          // set sesame bun for chicken burger
	chickenburger := director.buildBurger()
	fmt.Printf("Chicken Burger: %+v\n", chickenburger)
}
