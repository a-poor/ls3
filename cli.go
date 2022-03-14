package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

const configFile = ".ls3.yaml"

func newApp() *cli.App {
	return &cli.App{
		Name:    "ls3",
		Usage:   "A terminal app for accessing files in S3.",
		Version: version,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config-dir",
				Usage:       "Location `DIR` where the ls3 config file is stored.",
				DefaultText: "$HOME",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Create a new ls3 config file.",
				Action: func(c *cli.Context) error {
					return viper.WriteConfig()
				},
			},
			{
				Name:  "local",
				Usage: "Look for a local file to upload to S3.",
				Action: func(c *cli.Context) error {
					return viper.WriteConfig()
				},
			},
			{
				Name:  "remote",
				Usage: "Look for a remote file to download from S3.",
				Action: func(c *cli.Context) error {
					return viper.WriteConfig()
				},
			},
		},
		Action: func(c *cli.Context) error {
			// Create the model
			m, err := newModel()
			if err != nil {
				return err
			}

			p := tea.NewProgram(m, tea.WithAltScreen())
			if err := p.Start(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				return err
			}

			return nil
		},
	}
}
