package main

import (
	"context"
	"errors"
	"log"
	"os"

	"reumes/resume"

	"github.com/goccy/go-yaml"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "reumes",
		Usage:     "A yaml-based resume builder using the resume.json spec",
		UsageText: "reumes <INPUT_FILE> [OPTIONS]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "inputFile",
				UsageText: "input yaml file",
				Value:     "reumes.yaml",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "the name of the template to use",
				Value:   "reumes",
				Validator: func(t string) error {
					temps := resume.GetTemplates()

					_, exists := temps[t]

					if !exists {
						return errors.New("non-existent template")
					}

					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "print additional debug info",
				Value:   false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			resumePath := cmd.StringArg("inputFile")

			data, err := os.ReadFile(resumePath)
			if err != nil && resumePath == "reumes.yaml" {
				log.Fatal("input not specified and 'reumes.yaml' does not exist")
			} else if err != nil {
				log.Fatalf("failed to read '%s'", resumePath)
			}

			var res resume.Resume
			err = yaml.Unmarshal(data, &res)

			tempArg := cmd.String("template")
			templates := resume.GetTemplates()
			template, exists := templates[tempArg]

			if !exists {
				log.Fatalf("non-existent template")
			}

			outputError := template.Render(res)

			if outputError != nil {
				log.Fatalf("unable to produce output: %v", outputError)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
