package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello", phrase)
}
func slowGreet(phrase string, donechan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	donechan <- true
	// close the slowGreet function explict to avoid the error whikle running the for loop
	close(donechan)
}
func main() {
	//Using channels to get the slowGreet function value because of we using the goroutines.

	done := make(chan bool)

	//channels using slice method all commented lines are using the slice method

	//dones := make([]chan bool, 4)

	//dones[0] = make(chan bool)

	go func() {
		greet("Nice to meet You!")
	}()

	//dones[1] = make(chan bool)

	go func() {
		greet("how are you?")
	}()

	//dones[2] = make(chan bool)

	go slowGreet("HOW....ARE...YOU ?", done)

	//dones[3] = make(chan bool)

	go func() {
		greet("I hope you are liking the course")
	}()

	//for _, done := range dones {
	//	<-done
	//}

	for range done {

	}

}
