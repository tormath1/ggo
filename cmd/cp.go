package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cpCmd)
}

var cpCmd = &cobra.Command{
	Use:   "cp <src> <destination>>",
	Short: "cp <src> to <destination>",
	Long:  `Do I need to write a description to copy a file`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cp(args[0], args[1])
	},
}

func cp(src string, dst string) {
	fmt.Println(src, dst)
}
