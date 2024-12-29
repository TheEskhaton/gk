package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

var parcelCmd = &cobra.Command{
	Use:   "parcel",
	Short: "Searches for parcel numbers starting with a specific string in a specific municipality",
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
	Example: `//Searches for parcels with a parcel number beginning with 32 in the municipality 
// with registration number 335576
gk parcel -m 335576 -p 32`,
}

func init() {
	rootCmd.AddCommand(parcelCmd)
	parcelCmd.Flags().IntP("municipalityRegNbr", "m", 0, "Specify the municipality registration number")
	parcelCmd.Flags().StringP("parcelNumber", "p", "", "Specify the search query for parcel numbers")
	parcelCmd.MarkFlagRequired("parcelNumber")
	parcelCmd.MarkFlagRequired("municipalityRegNbr")
}
