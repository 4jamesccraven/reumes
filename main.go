package main

import (
	"context"
	"log"
	"os"

	"reumes/resume"

	"github.com/goccy/go-yaml"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "reumes",
		Usage: "A resume builder based on the resume.json spec",
		Flags: []cli.Flag{},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() < 1 {
				return cli.Exit("no pos arg", 1)
			}

			resumePath := cmd.Args().Get(0)

			data, err := os.ReadFile(resumePath)
			if err != nil {
				log.Fatalf("failed to read file: %v", err)
			}

			var res resume.Resume
			err = yaml.Unmarshal(data, &res)

			log.Printf("%s", res)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
