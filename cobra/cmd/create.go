package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// go run main.go create
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
