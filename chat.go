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

	color.Cyan("Chat with your own AI Friend **Aioli**!")

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter message: ")
		scanner.Scan()

		iter := chat.SendMessageStream(ctx, genai.Text(scanner.Text()))

		for {
			resp, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(resp.Candidates[0].Content)

		}

	}
}
