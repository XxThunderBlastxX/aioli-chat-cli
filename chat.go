package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

func chatWithAioli() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEN_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	aioli := client.GenerativeModel("gemini-pro")

	chat := aioli.StartChat()

	chat.History = []*genai.Content{}

	for {
		scanner := bufio.NewScanner(os.Stdin)

		_, _ = meColor.Print("😃 Me > ")
		scanner.Scan()

		iter := chat.SendMessageStream(ctx, genai.Text(scanner.Text()))

		_, _ = aioliColor.Print("🤖 Aioli > ")
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
