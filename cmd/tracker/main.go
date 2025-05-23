package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seektor/habits-tracker-go/internal/command"
	"github.com/seektor/habits-tracker-go/internal/habits"
	"github.com/seektor/habits-tracker-go/internal/utils"
)

func main() {
	fmt.Println(utils.FgColors.Yellow + utils.FgColors.Bold +
		"=== Habit Tracker ===" +
		utils.FgColors.Reset)

	habits := habits.NewHabits()

	if err := habits.Load(); err != nil {
		fmt.Println(utils.FgColors.Red + err.Error())
		os.Exit(1)
	}

	fmt.Println()
	isUpdated := habits.UpdateToPresent()
	fmt.Println()

	if isUpdated {
		habits.Save(utils.FileName)
		utils.PrintlnSuccess("Habits have been updated")
		fmt.Println()
	}

	if len(habits.Habits) > 0 {
		habits.PrintAll()
	} else {
		habits.PrintCommands()
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Print("Enter command: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		command := command.NewCommand(input)
		habits.Execute(command)
	}
}
