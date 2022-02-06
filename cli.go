package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {
	return &cli.App{
		Name:    "ls3",
		Usage:   "A terminal app for accessing files in S3.",
		Version: version,
		Commands: []*cli.Command{
			{
				Name: "init",
				Action: func(c *cli.Context) error {
					return viper.WriteConfig()
				},
			},
		},
		Action: func(c *cli.Context) error {
			// Load the config data
			// err := viper.ReadInConfig()
			// if err != nil {
			// 	fmt.Println("Error reading config file:", err)
			// 	return err
			// }

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
