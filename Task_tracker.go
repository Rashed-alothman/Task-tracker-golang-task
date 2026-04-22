package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Task Tracker initialized.")
	// Check if an argument was provided to avoid panic
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		fmt.Println("Usage: go run Task_tracker.go [command]")
		return
	}

	switch args := os.Args[1:]; args[0] {
	case "add", "new", "create", "a":
		fmt.Println("Task added.")

	case "list", "ls", "show", "L":

		fmt.Println("Listing tasks.")
	case "complete", "done", "finish", "C":

		fmt.Println("Task completed.")
	case "delete", "remove", "del", "D":

		fmt.Println("Task deleted.")
	case "update", "edit", "modify", "change", "edit-task", "U":

		fmt.Println("Task updated.")
	case "help", "h", "?":

		fmt.Println("Available commands: add, list, complete, delete, update, help")
	default:

		fmt.Printf("Unknown command: %s\n", args[0])
	}

}
