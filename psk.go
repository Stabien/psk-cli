package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var languages = []string{
	"nodejs",
}

func nodejsCreation(args []string) {
	fmt.Println("Creating NodeJS project...")
	var command *exec.Cmd

	if len(args) > 1 {
		command = exec.Command("bash", "/mnt/d/Projets/Stabien cli/nodejs.sh", args[1])
	} else {
		command = exec.Command("bash", "/mnt/d/Projets/Stabien cli/nodejs.sh")
	}

	command.Dir = "."
	output, err := command.Output()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", output)
}

func isInArray(array []string, element string) bool {
	for _, item := range array {
		if item == element {
			return true
		}
	}
	return false
}

func handleCmd(args []string) {
	if isInArray(languages, args[0]) == false {
		fmt.Println("Unknown language")
	} else {
		if args[0] == "nodejs" {
			nodejsCreation(args)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "stringer",
	Short: "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding)

One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		handleCmd(args)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
