package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("there were errors reading, exiting program.")
	}

	fmt.Printf("your name is %s", input)

	switch input {
	case "ok1\n":
		fmt.Println("welcome ok1")
	case "ok2\n":
		fmt.Println("welcome ok2")
	case "ok3\n":
		fmt.Println("welcome ok3")
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}

	// version2
	switch input {
	case "ok1\n":
		fallthrough
	case "ok2\n":
		fallthrough
	case "ok3\n":
		fmt.Println("welcome", input)
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}

	// version3
	switch input {
	case "ok1\n", "ok2\n", "ok3\n":
		fmt.Printf("welcome %s", input)
	default:
		fmt.Println("you are not welcome here! goodbye!")
	}
}
