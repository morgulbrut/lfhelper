package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a file",
}

// csvCmd represents the csv command
var lfCSVCmd = &cobra.Command{
	Use:   "lf",
	Short: "Reads a life fitness csv",
	Run: func(cmd *cobra.Command, args []string) {
		csv := ReadCVS(args[0])
		spew.Dump(csv)
	},
}

func init() {
	readCmd.AddCommand(lfCSVCmd)
	rootCmd.AddCommand(readCmd)
}
