package workingmyassoff

import (
	"fmt"
)

func Working(myname string) {

	fmt.Println("My name is", myname)
}

func IsthisWorking(bollean bool) {

	if bollean {

		fmt.Printf("%t", bollean)
	} else {

		fmt.Printf("%s", "BuzzFizz")
	}
}
