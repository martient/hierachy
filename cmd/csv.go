package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// csvCmd for create hierachie board
var csvCmd = &cobra.Command{
	Use:   "append",
	Short: "Add a new member of the hierachy board",
	Long:  `Append someone to the hierachy board, with this personal information (name, position, location, etc)`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		manager, _ := cmd.Flags().GetString("manager")
		if len(args) != 4 {
			println(len(args))
			log.Fatal("Please follow the method for append one member\nFor more information exec: hierachy append -h")
		}
		if filename == "" {
			log.Fatal("Please define your filename\nFor more information exec: hierachy append -h")
		}
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644) //Create("result.csv")
		checkError("Cannot open "+filename, err)
		defer file.Close()
		Reader := csv.NewReader(file)
		records, err := Reader.ReadAll()
		checkError("Reading issue on file "+filename, err)
		println(records)
		writer := csv.NewWriter(file)
		var id = records[len(records)-1][2]
		if id == "id" {
			id = "0"
		} else {
			i1, err := strconv.Atoi(id)
			checkError("Cannot write to file", err)
			id = strconv.Itoa(i1 + 1)
		}
		defer writer.Flush()
		var user = []string{args[0], args[1], id, args[2], manager, args[3], "#d5e8d4", "#82b366", "", "", "https://cdn3.iconfinder.com/data/icons/user-avatars-1/512/users-9-2-128.png"}
		err = writer.Write(user)
		checkError("Cannot write to file", err)
		log.Print(user)
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)
	csvCmd.Flags().StringP("filename", "f", "", "Csv filename, already create")
	csvCmd.Flags().StringP("manager", "m", "", "Manager id")
}
