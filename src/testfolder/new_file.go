package main

import (
	"fmt"
	"workingmyassoff"
)

// dude!
/**
*
* Note: you need to create a folder inside the src folder, basically need to make a "package"
 Java analogy
*
Need to build other files so that you can use the functions from other modules or packages/folders under the src dir
*
*/

func main() {

	fmt.Println("Hello")
	workingmyassoff.Working("SId")

	const name string = "Sid"

	workingmyassoff.IsthisWorking(name == "Sid")

	workingmyassoff.IsthisWorking(name == "Owl")
}
