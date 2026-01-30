package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const (
	todoFileName = ".todo.json"
)

func main() {

	//Parsing command line flags
	task := flag.String("task", "", "Task to be included to the ToDo List")
	list := flag.Bool("list", false, "List all todos")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// List current to do items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		//Complete the given item
		if err := l.Completed(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)

		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid command")
		os.Exit(1)
	}
}
