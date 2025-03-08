package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlag struct {
	Add  string
	Del  int
	Edit string
	Done int
	List bool
}

func NewCmdFlag() *CmdFlag {
	cf := CmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify the title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title")
	flag.IntVar(&cf.Del, "delete", -1, "Delete a todo by specifying the index")
	flag.IntVar(&cf.Done, "done", -1, "Specify a index to mark as done")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf

}

func (cf *CmdFlag) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Done != -1:
		todos.done(cf.Done)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("No command specified")
	}
}
