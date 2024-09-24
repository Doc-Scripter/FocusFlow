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
	count:=0
	for input.Scan() {
		line := input.Text()
		tasks = append(tasks, line)
		if count==2{
			break
		}
		count++
	}
	service.Input(tasks)
	// fmt.Println(out)~
}
