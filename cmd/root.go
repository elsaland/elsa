package cmd

import (
	"fmt"
	"os"

	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/elsa/module"
	"github.com/elsaland/elsa/util"
	"github.com/fatih/color"

	"github.com/elsaland/elsa/packager"
	"github.com/spf13/cobra"
)

type Elsa struct {
	Run    func(opt options.Options)
	Dev    func(opt options.Options)
	Bundle func(file string, minify bool, config *module.Config) string
}

func Execute(elsa Elsa) {
	config, err := module.ConfigLoad()
	util.Check(err)

	color.NoColor = config.Options.NoColor

	var fsFlag bool
	var minifyFlag bool

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
				bundle := elsa.Bundle(args[0], true, config)
				env := options.Environment{
					NoColor: config.Options.NoColor,
					Args:    args[1:],
				}
				opt := options.Options{
					SourceFile: args[0],
					Source:     bundle,
					Perms:      &options.Perms{fsFlag},
					Env:        env,
				}
				elsa.Run(opt)
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
				bundle := elsa.Bundle(args[0], true, config)
				env := options.Environment{
					NoColor: config.Options.NoColor,
					Args:    args[1:],
				}
				opt := options.Options{
					SourceFile: args[0],
					Source:     bundle,
					Perms:      &options.Perms{fsFlag},
					Env:        env,
				}
				elsa.Dev(opt)
			}
		},
	}

	var bundleCmd = &cobra.Command{
		Use:   "bundle [file]",
		Short: "Bundle your script to a single javascript file",
		Long:  `Bundle your script to a single javascript file. It utilizes esbuild for super fast bundling.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 0 {
				out := elsa.Bundle(args[0], minifyFlag, config)
				fmt.Println(out)
			}
		},
	}

	bundleCmd.Flags().BoolVarP(&minifyFlag, "minify", "m", false, "Minify the output bundle")

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
