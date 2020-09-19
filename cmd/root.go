package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute(cb func(file string)) {

	var rootCmd = &cobra.Command{
		Use:   "done [file]",
		Short: "Done is a simple Javascript and Typescript runtime written in Go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Running %v\n", args[0])
			cb(args[0])
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
