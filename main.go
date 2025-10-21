package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func addTodo(filename, task string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err2 := file.WriteString("[ ] " + task + "\n")
	return err2
}

func listTodos(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No todos yet!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}
}

func markDone(filename string, index int) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(file), "\n")
	if index <= 0 || index > len(lines)-1 { // -1 because last line may be empty
		return fmt.Errorf("invalid task number")
	}

	line := lines[index-1]
	if strings.HasPrefix(line, "[ ]") {
		lines[index-1] = "[x]" + line[3:]
	} else {
		return fmt.Errorf("task already done")
	}

	return os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
}

func main() {
	filename := "todos.txt"

	addFlag := flag.String("add", "", "Add a new todo")
	doneFlag := flag.Int("done", 0, "Mark todo as done by number")
	listFlag := flag.Bool("list", false, "List all todos")
	flag.Parse()

	if *addFlag != "" {
		err := addTodo(filename, *addFlag)
		if err != nil {
			fmt.Println("Error adding todo:", err)
		} else {
			fmt.Println("Todo added!")
		}
	} else if *doneFlag != 0 {
		err := markDone(filename, *doneFlag)
		if err != nil {
			fmt.Println("Error marking todo:", err)
		} else {
			fmt.Println("Task marked done!")
		}
	} else if *listFlag {
		listTodos(filename)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  --add \"Task name\"    Add a todo")
		fmt.Println("  --done <number>      Mark task as done")
		fmt.Println("  --list               List all todos")
	}
}
