/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alwindoss/orca/internal/inventory"
	"github.com/alwindoss/orca/internal/playbook"
	"github.com/alwindoss/orca/internal/runner"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("run called")
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		inv, err := inventory.LoadInventory(inventoryFile)
		if err != nil {
			return fmt.Errorf("failed to load inventory: %v", err)
		}

		plays, err := playbook.LoadPlaybook(playbookFile)
		if err != nil {
			return fmt.Errorf("failed to load playbook: %v", err)
		}

		runner, err := runner.NewRunner(inv, plays, privateKeyPath)
		if err != nil {
			return fmt.Errorf("failed to create runner: %v", err)
		}

		return runner.Run()
	},
}

var inventoryFile, playbookFile, privateKeyPath string

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().StringVarP(&inventoryFile, "inventory", "i", "inventory.yml", "Path to inventory file")
	runCmd.Flags().StringVarP(&playbookFile, "playbook", "p", "playbook.yml", "Path to playbook file")
	runCmd.Flags().StringVarP(&privateKeyPath, "key", "k", "~/.ssh/id_rsa", "Path to SSH private key")
}
