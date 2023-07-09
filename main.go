package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ecscli",
	Short: "ECS CLI interacts with AWS ECS.",
	Long:  `A Fast and Flexible CLI for managing AWS ECS built with love by Ebi.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add default behavior here if no commands provided
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ECS task definitions",
	Long:  `List all the task definitions available in your ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
		fmt.Println("Listing task definitions...")
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an ECS task",
	Long:  `Deploy a task to ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
		fmt.Println("Deploying task...")
	},
}

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback an ECS task",
	Long:  `Rollback a deployed task in ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
		fmt.Println("Rolling back task...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(rollbackCmd)
}

func main() {
	Execute()
}
