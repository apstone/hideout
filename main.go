package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	db := initializeDB()
	defer db.Close()

	rootCmd := &cobra.Command{
		Use:   "hideout",
		Short: "Hideout is a simple password management CLI application",
	}

	cmdAdd := &cobra.Command{
		Use:   "add [name] [value]",
		Short: "Adds a new password",
		Args:  cobra.ExactArgs(2), // Requires exactly two arguments
		Run: func(cmd *cobra.Command, args []string) {
			passwordName := args[0]
			passwordValue := args[1]
			insertPassword(db, passwordName, passwordValue)
			fmt.Println("Password added successfully.")
		},
	}

	cmdList := &cobra.Command{
		Use:   "list",
		Short: "Lists all stored passwords",
		Run: func(cmd *cobra.Command, args []string) {
			listPasswords(db)
		},
	}

	rootCmd.AddCommand(cmdAdd, cmdList)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
