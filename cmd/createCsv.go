package cmd

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// createCsvCmd for create hierachie board
var createCsvCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate new csv file",
	Long:  `Generate new csv file with the minimal config for generate the end render`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "hierachy.csv"
		} else {
			name += ".csv"
		}
		data, err := ioutil.ReadFile("csvHeader")
		checkError("csvHeader can be reach", err)
		err = ioutil.WriteFile(name, data, 0644)
		checkError("Cannot create file", err)
	},
}

func init() {
	rootCmd.AddCommand(createCsvCmd)
	createCsvCmd.Flags().StringP("name", "n", "", "Define your csv file name")
}
