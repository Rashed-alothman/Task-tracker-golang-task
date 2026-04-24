package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// load reads all tasks from the file
func load() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{} // file doesn't exist yet, start empty
	}
	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

// save writes all tasks back to the file
func save(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("tasks.json", data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run Task_tracker.go [command]")
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <description>")
			return
		}
		tasks := load()
		newTask := Task{
			ID:          len(tasks) + 1,
			Description: os.Args[2],
			Status:      "todo",
			CreatedAt:   time.Now().Format(time.RFC3339),
		}
		tasks = append(tasks, newTask)
		save(tasks)
		fmt.Printf("Task added with ID %d\n", newTask.ID)

	case "list":
		tasks := load()
		for _, task := range tasks {
			fmt.Printf("[%d] %s — %s\n", task.ID, task.Description, task.Status)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: complete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		tasks := load()
		for index, task := range tasks {
			if task.ID == id {
				tasks[index].Status = "done"
				tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)
				save(tasks)
				fmt.Printf("Task %d marked as complete.\n", id)
				return
			}
		}
		fmt.Printf("Task %d not found.\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		tasks := load()
		for index, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:index], tasks[index+1:]...)
				save(tasks)
				fmt.Printf("Task %d deleted.\n", id)
				return
			}
		}
		fmt.Printf("Task %d not found.\n", id)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> <new description>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		tasks := load()
		for index,task := range tasks {
			if task.ID == id {
				tasks[index].Description = os.Args[3]
				tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)
				save(tasks)
				fmt.Printf("Task %d updated.\n", id)
				return
			}
		}
		fmt.Printf("Task %d not found.\n", id)
	case "help":
		fmt.Printf("this is a simple task tracker the command it is \n./task_tracker add <expalin your task>, \n ./task_tracker update <id>, \n./task_tracker list, \n./task_tracker delete <id>")
	}
}
