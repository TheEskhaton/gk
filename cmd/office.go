package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theeskhaton/gk/pkg/api"
)

var officeCmd = &cobra.Command{
	Use:   "office",
	Short: "Fetchse cadastral offices",
	Run: func(cmd *cobra.Command, args []string) {
		offices := api.FetchOffices()
		for _, office := range offices {
			fmt.Printf("%d: %s\n", office.ID, office.Name)
		}
	},
	Example: "gk office",
}

func init() {
	rootCmd.AddCommand(officeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// officeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
