package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Perms struct {
	Fs bool
}

func Execute(cb func(file string, flags Perms)) {
	var fsFlag bool

	var rootCmd = &cobra.Command{
		Use:   "done [file]",
		Short: "Done is a simple Javascript and Typescript runtime written in Go",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				fmt.Printf("Running %v\n", args[0])
				cb(args[0], Perms{
					fsFlag,
				})
			}
		},
	}

	rootCmd.Flags().BoolVarP(&fsFlag, "fs", "f", false, "Allow file system access")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
