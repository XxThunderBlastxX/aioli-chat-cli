package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

var meColor = color.New(color.FgYellow)
var aioliColor = color.New(color.FgHiBlue)

func chatWithAioli() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEN_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	aioli := client.GenerativeModel("gemini-pro")

	chat := aioli.StartChat()

	chat.History = []*genai.Content{}

	for {
		scanner := bufio.NewScanner(os.Stdin)

		meColor.Print("😃 Me > ")
		scanner.Scan()

		iter := chat.SendMessageStream(ctx, genai.Text(scanner.Text()))

		aioliColor.Print("🤖 Aioli > ")
		for {
			resp, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}

			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(resp.Candidates[0].Content.Parts[0])

		}
		fmt.Print("\n")
	}
}
