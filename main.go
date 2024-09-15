package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var app = &cli.App{
	Name:  "based16",
	Usage: "a dead simple base16 theme templating tool",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "theme",
			Aliases:  []string{"t"},
			Usage:    "theme file (.yaml file)",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "input",
			Aliases: []string{"i"},
			Usage:   "input template path (use go template syntax)",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "output file path (will be overriden if exists)",
		},
	},
	Action: func(ctx *cli.Context) error {
		templateFile := os.Stdin
		if themePath := ctx.String("input"); themePath != "" {
			file, err := os.Open(themePath)
			if err != nil {
				return err
			}
			templateFile = file
			defer file.Close()
		}

		outputFile := os.Stdout
		if outputPath := ctx.String("output"); outputPath != "" {
			file, err := os.Create(outputPath)
			if err != nil {
				return err
			}
			outputFile = file
			defer file.Close()
		}

		themePath := ctx.String("theme")
		if themePath == "" {
			return errors.New("theme is empty")
		}
		themeFile, err := os.Open(themePath)
		if err != nil {
			return err
		}
		defer themeFile.Close()

		var theme any
		if err := yaml.NewDecoder(themeFile).Decode(&theme); err != nil {
			return err
		}

		templateBody, err := io.ReadAll(templateFile)
		if err != nil {
			return err
		}

		template, err := template.New("").Parse(string(templateBody))
		if err != nil {
			return err
		}
		return template.Execute(outputFile, theme)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}
