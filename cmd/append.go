package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("append called")
	},
}

var appendJSONCmd = &cobra.Command{
	Use:   "json",
	Short: "append to an existing json",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("json called")
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
	appendCmd.AddCommand(appendJSONCmd)
}
