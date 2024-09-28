package service

import (
	"bufio"
	"os"

)

func AddEvent() {
	input := bufio.NewScanner(os.Stdin)
	// var out struct{}
	tasks := []string{}
	count := 0
	for input.Scan() {
		line := input.Text()
		tasks = append(tasks, line)
		if count == 2 {
			break
		}
		count++
	}
	Input(tasks)
	// fmt.Println(out)~
}
