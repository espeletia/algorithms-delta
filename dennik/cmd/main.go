package main

import (
	"dennik/internal/diary"
	"fmt"
	"os"
)

var mainList *diary.Diary

func add() {
	var task string
	fmt.Printf("Enter the text: ")
	fmt.Scanf("%s", &task)
	mainList.Add(task)
}

func save() {
	mainList.ToJSON("diary.json")
}

var functionList = map[string]func(){
	"add":    add,
	"remove": delete,
	"next":   next,
	"prev":   previous,
	"save":   save,
	"exit":   exit,
}

func delete() {
	mainList.Delete()
}

func next() {
	mainList.Next()
}

func previous() {
	mainList.Previous()
}

func exit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}

func displayCommands() {
	for k := range functionList {
		fmt.Println("- " + k)
	}
}

func main() {
	mainList = &diary.Diary{}
	fmt.Printf("Welcome to Dennik! %+v\n", functionList)
	err := mainList.FromJSON("diary.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		mainList = diary.NewDiary()
	}
	if err == nil {
		fmt.Println("Successfully loaded diary.json")
		fmt.Printf("Number of entries: %v\n", len(mainList.Entries))
	}
	for {
		printStatus()
		mainList.DisplayCurrentEntry()
		fmt.Printf("Enter command: ")
		var command string
		fmt.Scanf("%s", &command)
		function, ok := functionList[command]
		if ok {
			function()
		} else {
			fmt.Println("Invalid command")
		}
	}
}

func printStatus() {
	fmt.Printf("Head: %v\n", mainList.Head)
	if mainList.Head != nil {
		fmt.Printf("Head.Next: %v\n", mainList.Head.Next)
		fmt.Printf("Head.Previous: %v\n", mainList.Head.Previous)
	}
}
