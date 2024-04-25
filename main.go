package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	db := initializeDB() // Initialize the database
	defer db.Close()

	rootCmd := &cobra.Command{Use: "hideout", Short: "Hideout is a simple CLI application"}

	cmdHello := &cobra.Command{
		Use:   "hello [name]",
		Short: "Prints Hello and an optional name",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := "World"
			if len(args) > 0 && args[0] != "" {
				name = args[0]
			}
			fmt.Printf("Hello, %s!\n", name)
			insertGreeting(db, name) // Insert greeting into the database
		},
	}

	cmdList := &cobra.Command{
		Use:   "list",
		Short: "Lists all greetings",
		Run: func(cmd *cobra.Command, args []string) {
			listGreetings(db) // List all greetings from the database
		},
	}

	rootCmd.AddCommand(cmdHello, cmdList)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
