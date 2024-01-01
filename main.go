package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:                 "chat",
		Usage:                "Chat with your own AI Friend **Aioli**",
		EnableBashCompletion: true,
		Action: func(*cli.Context) error {

			chatWithAioli()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
