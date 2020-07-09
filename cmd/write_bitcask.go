package cmd

import (
	"encoding/json"
	"log"
	"time"

	"github.com/prologic/bitcask"
	"github.com/spf13/cobra"
)

// bitcaskCmd represents the bitcask command
var bitcaskCmd = &cobra.Command{
	Use:   "bitcask <csv> <bitcask>",
	Short: "<csv> <bitcask> : reads a LifeFitness CSV and writes into a bitcask DB file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		csv := ReadCVS(args[0])
		db, _ := bitcask.Open("db")
		defer db.Close()
		ts := csv.Timestamp.Format(time.RFC3339)
		json, err := json.Marshal(csv)
		if err != nil {
			log.Fatalf("WriteJson marshal: %v", err)
		}
		db.Put([]byte(ts), []byte(json))
	},
}

func init() {
	writeCmd.AddCommand(bitcaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bitcaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bitcaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
