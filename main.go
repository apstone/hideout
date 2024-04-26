package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	db := initializeDB() // Initialize the database
	defer db.Close()

	rootCmd := &cobra.Command{
		Use:   "hideout",
		Short: "Hideout is a simple password management CLI application",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Print("Enter Master Password: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				masterPassword := scanner.Text()
				if !verifyMasterPassword(db, masterPassword) {
					fmt.Println("Incorrect Master Password!")
					os.Exit(1)
				}
			}
		},
	}

	cmdSetMaster := &cobra.Command{
		Use:   "setmaster",
		Short: "Sets the master password",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Set Master Password: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				masterPassword := scanner.Text()
				setMasterPassword(db, masterPassword)
				fmt.Println("Master password set successfully.")
			}
		},
	}

	cmdAdd := &cobra.Command{
		Use:   "add [name] [value]",
		Short: "Adds a new password",
		Args:  cobra.ExactArgs(2),
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

	cmdInteractive := &cobra.Command{
		Use:   "interactive",
		Short: "Enters interactive mode, allowing multiple commands without exiting",
		Run: func(cmd *cobra.Command, args []string) {
			interactiveMode(rootCmd)
		},
	}

	rootCmd.AddCommand(cmdSetMaster, cmdAdd, cmdList, cmdInteractive)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func interactiveMode(rootCmd *cobra.Command) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Entering interactive mode. Type 'exit' to quit.")
	for {
		fmt.Print("hideout> ")
		if !scanner.Scan() {
			break // Handle EOF (Ctrl+D)
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		if line != "" {
			args := strings.Fields(line)
			rootCmd.SetArgs(args)
			rootCmd.Execute()
		}
	}
	fmt.Println("Exiting interactive mode.")
}
