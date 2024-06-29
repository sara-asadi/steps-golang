package main

import "fmt"

func main() {

	// steps := os.Args[1]
	calculateSteps(5, "")

}

func calculateSteps(stepNum int, steps string) {
	if stepNum == 0 {
		fmt.Println(steps)
		return
	}
	if stepNum >= 2 {
		calculateSteps(stepNum-2, steps + "2 ")
	}
	if stepNum >= 1 {
		calculateSteps(stepNum-1, steps + "1 ")
	}
}
