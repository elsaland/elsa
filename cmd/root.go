package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/elsa/module"
	"github.com/elsaland/elsa/util"
	"github.com/fatih/color"

	"github.com/elsaland/elsa/packager"
	"github.com/spf13/cobra"
)

// Elsa functions expected to be passed into cmd
type Elsa struct {
	Run    func(opt options.Options)
	Dev    func(og string, opt options.Options)
	Bundle func(file string, minify bool, config *module.Config) string
}

// Execute start the CLI
func Execute(elsa Elsa) {
	// TODO: need to come to a concrete conclusion
	// Load mod.toml (if exists)
	config, err := module.ConfigLoad()
	util.Check(err)

	color.NoColor = config.Options.NoColor

	var fsFlag bool
	var netFlag bool
	var minifyFlag bool

	// Rool command
	var rootCmd = &cobra.Command{
		Use:   "elsa [file]",
		Short: "Elsa is a simple Javascript and Typescript runtime written in Go",
	}

	// Run subcommand
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
					Perms:      &options.Perms{fsFlag, netFlag},
					Env:        env,
				}
				elsa.Run(opt)
			}
		},
	}

	// --fs and --net flags
	runCmd.Flags().BoolVar(&fsFlag, "fs", false, "Allow file system access")
	runCmd.Flags().BoolVar(&netFlag, "net", false, "Allow net access")

	// dev subcommand to run script in development mode
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
					Perms:      &options.Perms{fsFlag, netFlag},
					Env:        env,
				}
				og, _ := ioutil.ReadFile(args[0])
				elsa.Dev(string(og), opt)
			}
		},
	}

	// bundle subcommand to bundle a source file
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

	// --minify flag for bundling
	bundleCmd.Flags().BoolVarP(&minifyFlag, "minify", "m", false, "Minify the output bundle")

	// pkg subcommand for trigger the packager
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

	// Add subcommands to root command
	rootCmd.AddCommand(bundleCmd, runCmd, pkgCmd, devCmd)

	// Execute! :)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
