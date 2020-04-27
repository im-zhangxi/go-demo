package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rabbitConsumeCmd = &cobra.Command{
	Use:                        "rabbit_consume",
	Short:                      "hi",
	Long:                       "This subcommand rabbit_consume",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	},
}

func init()  {
	rootCmd.AddCommand(rabbitConsumeCmd)
}
