/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

// departmentCmd represents the department command
var departmentCmd = &cobra.Command{
	Use:   "department",
	Short: "Fetches cadastral departments",
	Run: func(cmd *cobra.Command, args []string) {
		departments := api.FetchDepartments()
		for _, department := range departments {
			fmt.Printf("%d: %s, Office ID: %d\n", department.ID, department.Name, department.OfficeID)
		}
	},
	Example: "gk department",
}

func init() {
	rootCmd.AddCommand(departmentCmd)
}
