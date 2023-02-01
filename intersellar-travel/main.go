package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Printf("You have %d gallons of fuel left!", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Printf("Welcome to %s", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	//when writing the below I forgot about the -= at the time which caused errors.
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Test your functions!
	//greetPlanet("Venus")
	//fuelGauge(100)
	//fmt.Println(flyToPlanet("Mercury", 600000))
	// Create `planetChoice` and `fuel`

	fuel := 1000000
	planetChoice := "Venus"

	// saves the result of flyToPlanet back to fuel
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

	// fmt.Println(calculateFuel("Venus"))

	// And then liftoff!

}
