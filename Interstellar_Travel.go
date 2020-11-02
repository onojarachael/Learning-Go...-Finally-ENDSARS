package main

import "fmt"

// Create the function fuelGauge() here
func fuelGuage(fuel int) {
 fmt.Println("We have", fuel, "left, how do we get back to Earth?")
}
// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
}
if planet == "Venus" {
  fuel = 300000
} else if planet == "Mercury" {
  fuel = 500000
} else if planet == "Mars" {
  fuel = 700000
} else {
  fuel = 0
}
  fmt.Println(calculateFuel("Mars"))
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome to", planet, "Chikas!")
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining == fuel
 fuelCost = calculateFuel(planet)
 if fuelRemaining >= fuelCost {
   greetPlanet(planet)
   fuelRemaining -= fuelCost
 } else {
   cantFly()
 }
 }

func main() {
  // Test your functions!
  var fuel int
  fuel = 1000000
  // Create `planetChoice` and `fuel`
  planetChoice := "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGuage(fuel)
}
