package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Writes to a file or db",
}

var writeJSONCmd = &cobra.Command{
	Use:   "json <csv> <json>",
	Short: "<csv> <json> : reads a LifeFitness CSV and writes a json file",
	Run: func(cmd *cobra.Command, args []string) {
		csv := ReadCVS(args[0])
		json, err := json.Marshal(csv)
		if err != nil {
			log.Fatalf("WriteJson marshal: %v", err)
		}
		ioutil.WriteFile(args[1], json, 0644)
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.AddCommand(writeJSONCmd)
}
