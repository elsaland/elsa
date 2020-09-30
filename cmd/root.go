package cmd

import (
	"fmt"
	"os"

	"github.com/elsaland/elsa/packager"
	"github.com/spf13/cobra"
)

type Perms struct {
	Fs bool
}

type Elsa struct {
	Run    func(file string, bundle string, flags Perms, args []string)
	Dev    func(file string, bundle string, args []string)
	Bundle func(file string) string
}

func Execute(elsa Elsa) {
	var fsFlag bool

	var rootCmd = &cobra.Command{
		Use:   "elsa [file]",
		Short: "Elsa is a simple Javascript and Typescript runtime written in Go",
	}

	var runCmd = &cobra.Command{
		Use:   "run [file]",
		Short: "Run a Javascript and Typescript source file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				bundle := elsa.Bundle(args[0])
				elsa.Run(args[0], bundle, Perms{fsFlag}, args[1:])
			}
		},
	}

	runCmd.Flags().BoolVar(&fsFlag, "fs", false, "Allow file system access")

	var devCmd = &cobra.Command{
		Use:   "dev [file]",
		Short: "Run a script in development mode.",
		Long:  `Run a script in development mode. It enables type-checking using the inbuilt typescript compiler.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				bundle := elsa.Bundle(args[0])
				elsa.Dev(args[0], bundle, args[1:])
			}
		},
	}

	var bundleCmd = &cobra.Command{
		Use:   "bundle [file]",
		Short: "Bundle your script to a single javascript file",
		Long:  `Bundle your script to a single javascript file. It utilises esbuild for super fast bundling.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				out := elsa.Bundle(args[0])
				fmt.Println(out)
			}
		},
	}

	var pkgCmd = &cobra.Command{
		Use:   "pkg [file]",
		Short: "Package your script to a standalone executable.",
		Long:  `Package your script to a standalone executable.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				packager.PkgSource(args[0])
			}
		},
	}

	rootCmd.AddCommand(bundleCmd, runCmd, pkgCmd, devCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
