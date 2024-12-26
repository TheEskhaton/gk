/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

// parcelDetailsCmd represents the parcelDetails command
var parcelDetailsCmd = &cobra.Command{
	Use:   "parcelDetails",
	Short: "Dohvaća detaljne podatke o čestici",
	Run: func(cmd *cobra.Command, args []string) {
		parcelId, err := cmd.Flags().GetInt("parcelId")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		isJson, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		parcelDetails := api.FetchParcelDetails(parcelId)
		if isJson {
			stringDetails, err := json.MarshalIndent(parcelDetails, "", " ")
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fmt.Print(string(stringDetails))
		} else {
			for _, parcelDetail := range parcelDetails {
				var possessors string
				for _, possessor := range parcelDetail.PossessionSheet.Possessors {
					possessors = possessors + "\t" + possessor.Name + "\n"
				}
				fmt.Printf("Parcel Number: %s\nArea: %sm2\nMunicipality: %s\nPossesors:\n%s",
					parcelDetail.ParcelNumber,
					parcelDetail.Area,
					parcelDetail.CADMunicipalityName,
					possessors[:len(possessors)-1])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parcelDetailsCmd)

	parcelDetailsCmd.Flags().IntP("parcelId", "p", 0, "Specify the parcel ID")
	parcelDetailsCmd.Flags().Bool("json", false, "Show full JSON")
}
