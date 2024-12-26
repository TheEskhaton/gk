/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

// officeCmd represents the office command
var officeCmd = &cobra.Command{
	Use:   "office",
	Short: "Dohvaća katastarske urede",
	Run: func(cmd *cobra.Command, args []string) {
		offices := api.FetchOffices()
		for _, office := range offices {
			fmt.Printf("%d: %s\n", office.ID, office.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(officeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// officeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
