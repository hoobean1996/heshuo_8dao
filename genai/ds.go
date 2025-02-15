package genai

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const DEEPSEEK_API_KEY = "sk-5cdd440191ac40ff98b8d8a1fb8b3d73"
const GPT_API_KEY = "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"
const BASE_URL = "https://api.deepseek.com"

func GPT(prompt string) string {
	client := openai.NewClient(
		option.WithAPIKey(GPT_API_KEY),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}

	return chatCompletion.Choices[0].Message.Content
}

func Deepseek(prompt string) string {
	client := openai.NewClient(
		option.WithAPIKey(DEEPSEEK_API_KEY),
		option.WithBaseURL(BASE_URL),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F("deepseek-chat"),
	})
	if err != nil {
		panic(err.Error())
	}

	return chatCompletion.Choices[0].Message.Content
}
