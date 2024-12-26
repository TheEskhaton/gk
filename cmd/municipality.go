/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

// municipalityCmd represents the municipality command
var municipalityCmd = &cobra.Command{
	Use:   "municipality",
	Short: "Fetches municipalities. Note that you will need the registration number rather than the ID to fetch parcels later",
	Run: func(cmd *cobra.Command, args []string) {
		departmentId, err := cmd.Flags().GetInt("department")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		officeId, err := cmd.Flags().GetInt("office")
		if err == nil {
			municipalities := api.FetchMunicipalities(officeId, departmentId)
			for _, municipality := range municipalities {
				fmt.Printf("ID: %d: %s\n", municipality.ID, municipality.Name)
			}
		}
	},
	Example: `//Returns all municipalities in specifed department
gk municipality -d 34
//Returns all municipalities in specifed office
gk municipality -o 100`,
}

func init() {
	rootCmd.AddCommand(municipalityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// municipalityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	municipalityCmd.Flags().IntP("department", "d", 0, "Specify the department ID to filter for")
	municipalityCmd.Flags().IntP("office", "o", 0, "Specify the office ID to filter for")
}
