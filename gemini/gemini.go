package gemini

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Gemini struct {
	ctx     context.Context
	client  *genai.Client
	Session *genai.ChatSession
}

func New() *Gemini {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	model := client.GenerativeModel("gemini-1.5-flash")
	session := model.StartChat()
	if err != nil {
		log.Fatal(err)
	}

	return &Gemini{
		ctx:     ctx,
		client:  client,
		Session: session,
	}
}

func (g *Gemini) Close() {
	g.client.Close()
}

func (g *Gemini) SendMessage(message string) (genai.Text, error) {
	resp, err := g.Session.SendMessage(g.ctx, genai.Text(message))
	if err != nil {
		return "", err
	}
	part := resp.Candidates[0].Content.Parts[0].(genai.Text)

	return part, nil
}
