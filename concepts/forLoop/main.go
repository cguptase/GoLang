package main

import "fmt"

func main() {
	names := []string{"Chirag", "Sudha", "Uma"}
	for i, name := range names {
		fmt.Println(i, name)
	}
}
