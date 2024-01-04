package main

import (
	"github.com/fatih/color"
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
			color.Cyan("          :::        :::::::::::       ::::::::        :::        ::::::::::: " +
				"\n       :+: :+:          :+:          :+:    :+:      :+:             :+:      " +
				"\n     +:+   +:+         +:+          +:+    +:+      +:+             +:+       " +
				"\n   +#++:++#++:        +#+          +#+    +:+      +#+             +#+        " +
				"\n  +#+     +#+        +#+          +#+    +#+      +#+             +#+         " +
				"\n #+#     #+#        #+#          #+#    #+#      #+#             #+#          " +
				"\n###     ###    ###########       ########       ##########  ###########       \n")

			color.Cyan("Chat with your own AI Friend **Aioli**!")
			chatWithAioli()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
