package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
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
	writeCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
