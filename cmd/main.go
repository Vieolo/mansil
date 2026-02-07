package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vieolo/mansil/cmd/internal/generator"
	"github.com/vieolo/mansil/cmd/internal/version"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mansil",
		Short: "Mansil ANSI Code Generator",
	}

	var genCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate ANSI codes for all languages",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Generating the ANSI codes...")
			return generator.Generate()
		},
	}

	var bumpCmd = &cobra.Command{
		Use:   "bump <build|minor|major>",
		Short: "Bump the version in all package files",
		Long: `Bump the semantic version in all package configuration files.
The source version is read from go.yaml and updated in:
  - go.yaml
  - package.json
  - pubspec.yaml
  - pyproject.toml
  - Cargo.toml

Arguments:
  build  - Increment the build/patch version (0.1.1 -> 0.1.2)
  minor  - Increment the minor version (0.1.1 -> 0.2.0)
  major  - Increment the major version (0.1.1 -> 1.0.0)`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			bumpType, err := version.ParseBumpType(args[0])
			if err != nil {
				return err
			}
			return version.Bump(bumpType)
		},
	}

	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(bumpCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
