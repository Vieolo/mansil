package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vieolo/mansil/cmd/internal/generator"
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
			cwd, _ := os.Getwd()
			fmt.Println("Generating ANSI codes from", cwd)
			return generator.Generate(cwd)
		},
	}

	rootCmd.AddCommand(genCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
