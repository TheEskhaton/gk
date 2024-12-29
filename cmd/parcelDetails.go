package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

var parcelDetailsCmd = &cobra.Command{
	Use:   "parcelDetails",
	Short: "Fetches details on parcel with specific parcel ID",
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
				if len(parcelDetail.PossessionSheet.Possessors) > 0 {
					for _, possessor := range parcelDetail.PossessionSheet.Possessors {
						possessors = possessors + "\t" + possessor.Name + " (" + possessor.Ownership + ")" + "\n"
					}
				} else {
					for _, share := range parcelDetail.LrUnit.OwnershipSheetB.LrUnitShares {
						for _, owner := range share.LrOwners {
							possessors = possessors + "\t" + owner.Name + " (" + share.Share + ")" + "\n"
						}
					}
				}
				fmt.Printf("Parcel Number: %s\nArea: %sm2\nMunicipality: %s\nPossesors:\n%s",
					parcelDetail.ParcelNumber,
					parcelDetail.Area,
					parcelDetail.CADMunicipalityName,
					possessors[:len(possessors)-1])

			}
		}
	},
	Example: `//Returns details on parcel with specified parcel ID
gk parcelDetails -p 21384620
//Returns details in JSON format on parcel with specified parcel ID
gk parcelDetails -p 21384620 --json
`,
}

func init() {
	rootCmd.AddCommand(parcelDetailsCmd)

	parcelDetailsCmd.Flags().IntP("parcelId", "p", 0, "Specify the parcel ID")
	parcelDetailsCmd.Flags().Bool("json", false, "Show full JSON")
	parcelDetailsCmd.MarkFlagRequired("parcelId")
}
