package main

import (
	"bufio"
	"os"

	"Focus/service"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	// var out struct{}
	tasks := []string{}
	for input.Scan() {
		line := input.Text()
	tasks = append(tasks, line)
		if line==""{
			break
		}
}
service.Input(tasks)
	// fmt.Println(out)
}
