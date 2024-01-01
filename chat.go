package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

//func initAioli(ctx context.Context) *genai.GenerativeModel {
//
//	return model
//}

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

		//var msg string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter message: ")
		scanner.Scan()
		//msg = scanner.Text()

		if scanner.Text() == "0" {
			break
		}

		resp, err := chat.SendMessage(ctx, genai.Text(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Candidates[0].Content)
	}
}
