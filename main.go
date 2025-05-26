package main

import (
	"context"
	"errors"
	"fmt"
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
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "the name of the template to use",
				Value:   "pdf",
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
			// Path to resume data
			res, err := GetResumeFileFromArg(cmd.StringArg("inputFile"))
			if err != nil {
				log.Fatalf("reumes error: %v", err)
			}

			// Use the resume data to fill a template
			tempArg := cmd.String("template")
			templates := resume.GetTemplates()
			template, exists := templates[tempArg]

			if !exists {
				log.Fatalf("reumes error: non-existent template")
			}

			outputError := template.Render(res)

			if outputError != nil {
				log.Fatalf("reumes error: unable to produce output:\n%v", outputError)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func GetResumeFileFromArg(resumePath string) (resume.Resume, error) {
	// Read the data file
	data, err := os.ReadFile(resumePath)
	if err != nil {
		if resumePath == "" {
			err = fmt.Errorf("no input data specified")
		} else {
			err = fmt.Errorf("failed to read '%s'", resumePath)
		}
		return resume.Resume{}, err
	}

	var res resume.Resume
	err = yaml.Unmarshal(data, &res)
	if err != nil {
		err = fmt.Errorf(
			"%v is not a data file. See https://jsonresume.org/schema for an example of a valid schema. %v",
			resumePath,
			err,
		)

		return resume.Resume{}, err
	}

	return res, nil
}
