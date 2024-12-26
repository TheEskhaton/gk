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

// parcelCmd represents the parcel command
var parcelCmd = &cobra.Command{
	Use:   "parcel",
	Short: "Pretražuje katastarske čestice unutar općine",
	Run: func(cmd *cobra.Command, args []string) {
		municipalityRegNbr, err := cmd.Flags().GetInt("municipalityRegNbr")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		parcelNumber, err := cmd.Flags().GetString("parcelNumber")
		if err == nil {
			parcelIdentifiers := api.SearchParcelIdentifiers(parcelNumber, municipalityRegNbr)
			for _, parcelIdentifier := range parcelIdentifiers {
				fmt.Printf("ID: %d: %s\n", parcelIdentifier.ID, parcelIdentifier.Number)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parcelCmd)
	parcelCmd.Flags().IntP("municipalityRegNbr", "m", 0, "Specify the municipality registration number")
	parcelCmd.Flags().StringP("parcelNumber", "p", "", "Specify the search query for parcel numbers")
	parcelCmd.MarkFlagRequired("parcelNumber")
	parcelCmd.MarkFlagRequired("municipalityRegNbr")
}
