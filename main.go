package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hideout",
		Short: "Hideout is a simple CLI application",
	}

	cmdHello := &cobra.Command{
		Use:   "hello [name]",
		Short: "Prints Hello and an optional name",
		Args:  cobra.MaximumNArgs(1), // Accepts at most one argument
		Run: func(cmd *cobra.Command, args []string) {
			name := "World"
			if len(args) > 0 && args[0] != "" {
				name = args[0]
			}
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	rootCmd.AddCommand(cmdHello)

	// Reading from standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Hideout! Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // In case of Ctrl+D or any error in scanning
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		if line != "" {
			// Execute the command using Cobra
			args := strings.Fields(line)
			rootCmd.SetArgs(args)
			rootCmd.Execute()
		}
	}

	fmt.Println("Goodbye!")
}
