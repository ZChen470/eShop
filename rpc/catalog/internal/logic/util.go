package logic

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func GeneratEmbedding(apiKey string, text string, dims int) ([]float32, error) {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model:      openai.SmallEmbedding3,
		Input:      []string{text},
		Dimensions: dims,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}
