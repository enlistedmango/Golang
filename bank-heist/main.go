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

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened!")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm.... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out the vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("Looks like the camera got you.... hide!")
		case 3:
			isHeistOn = false
			fmt.Println("There's extra guards you missed!!")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 1000 + rand.Intn(1000000)
		fmt.Println("You managed to steal:", amtStolen)
	}

	fmt.Println("isHeistOn is currently:", isHeistOn)

}
