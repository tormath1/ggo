package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gg",
	Long:  `All software has versions. This is gg's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("GG GNU Golang v0.1 -- HEAD")
	},
}
