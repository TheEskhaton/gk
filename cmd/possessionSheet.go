package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

var extractType api.ExtractType

const ExtractBaseUri = "https://oss.uredjenazemlja.hr/oss/public"

var possessionSheetCmd = &cobra.Command{
	Use:   "possessionSheet",
	Short: "Fetches unofficial posseession sheet extracts, cadastral plans and LR extracts",
	Run: func(cmd *cobra.Command, args []string) {
		parcelId, err := cmd.Flags().GetInt("parcelId")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		possessionSheetId, err := cmd.Flags().GetInt("possessionSheetId")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		switch extractType {
		case api.LR, api.CadastralPlan:
			if parcelId == 0 {
				log.Fatal("For LR and CadastralPlan extract types, a parcelId should be provided")
				os.Exit(1)
			}
		case api.PropertyDeed:
			if possessionSheetId == 0 {
				log.Fatal("For PropertyDeed, a possessionSheetId should be provided")
				os.Exit(1)
			}
		}
		propertySheet := api.FetchPossesionSheetExtract(parcelId, possessionSheetId, extractType)
		fmt.Println(ExtractBaseUri + propertySheet[0].FileUrl)
	},
	Example: `//The following commands fetch a file URL pointing to the specified document
gk possessionSheet --type CadastralPlan -p 123456
gk possessionSheet --type LR -p 123456
gk possessionSheet --type PropertyDeed -s 654321`,
}

func init() {
	extractType = api.PropertyDeed
	rootCmd.AddCommand(possessionSheetCmd)
	possessionSheetCmd.Flags().IntP("parcelId", "p", 0, "Parcel ID (user for LR and CadastralPlan extract types)")
	possessionSheetCmd.Flags().IntP("possessionSheetId", "s", 0, "Possession Sheet ID (Property Deed extract type)")
	possessionSheetCmd.Flags().Var(&extractType, "type", "Extract type. Allowed values: LR, PropertyDeed, CadastralPlan")
}
