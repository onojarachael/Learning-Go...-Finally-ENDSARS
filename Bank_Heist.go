package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
rand.Seed(time.Now().UnixNano())
isHeistOn := true
eludedGuards := rand.Intn(100)

//Try to elude some guards at the bank,
if eludedGuards >= 50 {
fmt.Println("Looks like we've managed to make it past the guards. Awesome, now let's set up the Resistance!")
} else {
  isHeistOn := false
fmt.Println("Well, we live to fight another day.")

//This is what we came for, let's try to see if we can get as much money as we can from the bank.
openedVault := rand.Intn(100)
if isHeistOn && openedVault >= 70 {
  fmt.Println("Grab and GO!")
} else if isHeistOn {
fmt.Println("The vault can't be opened Boss, we gotta go!")
}
fmt.Println("eludedGuards has a value of", eludedGuards) 
fmt.Println("isheistOn has a value of", isHeistOn)
}

leftSafely := rand.Intn(5)
if isHeistOn { 
  switch leftSafely {
    case 0:
    isHeistOn = false
    fmt.Println("Shit, you tripped an alarm... run!") 
    case 1:
    isHeistOn = false
    fmt.Println("Turns out the vault doors don't open from the inside...") 
     case 2:
    isHeistOn = false
    fmt.Println("We didn't turn off this cam feed, we're screwed!...") 
     case 3:
    isHeistOn = false
    fmt.Println("One of the hostages escaped.")
    default:
    fmt.Println("Una mattina mi sono alzato O bella ciao, bella ciao, bella ciao, ciao, ciao...") 
}
if isHeistOn == true {
  amtStolen := 10000 + rand.Intn(1000000)
  fmt.Println("We got away with", amtStolen)
}

}
}
