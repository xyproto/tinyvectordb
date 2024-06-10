package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xyproto/ask"
)

const versionString = "Tiny Vector Database 0.0.1"

var numbers []int

func main() {
	fmt.Println(versionString)
	for {
		command := strings.ToLower(strings.TrimSpace(ask.Ask("> ")))
		switch command {
		case "hi":
			fmt.Println("Hello")
		case "store":
			for {
				numText := ask.Ask("One or more numbers: ")
				if numText == "" {
					break
				}
				numTextSlice := []string{numText}
				if strings.Contains(numText, ";") {
					numTextSlice = strings.Split(numText, ";")
				} else if strings.Contains(numText, ",") {
					numTextSlice = strings.Split(numText, ",")
				} else if strings.Contains(numText, " ") {
					numTextSlice = strings.Split(numText, " ")
				}
				if len(numTextSlice) == 0 {
					break
				}
				for _, num := range numTextSlice {
					n, err := strconv.Atoi(strings.TrimSpace(num))
					if err == nil {
						numbers = append(numbers, n)
					} else {
						fmt.Printf("warning, could not add %v: %v\n", n, err)
					}
				}
			}
		case "show":
			fmt.Println(numbers)
		case "bye", "end", "quit":
			fmt.Println("bye")
			return
		default:
			fmt.Println("available commands: hi, store, show, bye, end, quit")
		}
	}
}
