package main

import "fmt"

func main() {
	calculateSteps(4)
	fmt.Println(paths[4])
}

var paths map[int][]string = make(map[int][]string)

func calculateSteps(stepNum int) {

	if stepNum == 0 {
		paths[0] = []string{""}
		return
	}

	if stepNum >= 2 {
		if _, ok := paths[stepNum-2]; !ok {
			calculateSteps(stepNum - 2)
		}

		for _, previousPath := range paths[stepNum-2] {
			paths[stepNum] = append(paths[stepNum], "2 "+previousPath)
		}
	}

	if stepNum >= 1 {
		if _, ok := paths[stepNum-1]; !ok {
			calculateSteps(stepNum - 1)
		}

		for _, previousPath := range paths[stepNum-1] {
			paths[stepNum] = append(paths[stepNum], "1 "+previousPath)
		}
	}

}
