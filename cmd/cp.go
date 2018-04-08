package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var recursive bool

func init() {
	var cpCmd = &cobra.Command{
		Use:   "cp <src> <destination>>",
		Short: "cp <src> to <destination>",
		Long:  `Do I need to write a description to copy a file`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			main(args[0], args[1])
		},
	}

	cpCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "copy a directory in a recursive way")
	rootCmd.AddCommand(cpCmd)
}

func main(src string, dst string) error {
	if !recursive {
		return cp(src, dst)
	}
	return cpRecursive(src, dst)
}

func cp(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return dstFile.Close()
}

func cpRecursive(src string, dst string) error {
	log.Println("recursive")
	return nil
}
