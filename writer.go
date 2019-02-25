package main

import (
	"fmt"
)

// FormatTasks will output a CLI friendly UI for the given list of tasks
func FormatTasks(tasks *[]Task) {
	fmt.Println("Task List:")

	for i, task := range *tasks {
		fmt.Printf("%d. %s\n", i+1, task.Task)
	}
}

// OutputTasks gets the tasks for today, and uses OutputTasks to show them
func OutputTasks() {
	// TODO ReadFile will create a new file if it doesn't exist, this is
	// undesirable
	tasks, err := ReadFile()

	if err != nil || len(*tasks) == 0 {
		fmt.Println("You have no tasks recorded for today.")
	} else {
		FormatTasks(tasks)
	}
}