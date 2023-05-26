package gpt

import (
	"context"
	"fmt"
	"github.com/DenChenn/blunder/internal/codegen/model"
	"github.com/sashabaranov/go-openai"
	"os"
	"strconv"
	"strings"
)

const (
	fieldSeparator = "/"
	groupSeparator = "#"
)

// CompleteErrorDetail completes the error detail with GPT-3
func CompleteErrorDetail(errorCodes []string) ([]*model.ErrorDescription, error) {
	errorString := strings.Join(errorCodes[:], "#")
	content := formatRequestString(errorString)

	client := openai.NewClient(os.Getenv("OPENAI_API_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("ChatCompletion error: %v\n", err)
	}

	groups := strings.Split(resp.Choices[0].Message.Content, groupSeparator)
	if len(groups) != len(errorCodes) {
		return nil, fmt.Errorf("response with wrong format from GPT-3 in group")
	}
	descriptions := make([]*model.ErrorDescription, len(groups))
	for i, group := range groups {
		fields := strings.Split(group, fieldSeparator)
		if len(fields) != 3 {
			return nil, fmt.Errorf("response with wrong format from GPT-3 in field")
		}

		httpStatusCode, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, fmt.Errorf("http status code is not a number")
		}
		grpcStatusCode, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("grpc status code is not a number")
		}

		descriptions[i] = &model.ErrorDescription{
			Code:           errorCodes[i],
			HttpStatusCode: httpStatusCode,
			GrpcStatusCode: grpcStatusCode,
			Message:        fields[2],
		}
	}

	return descriptions, nil
}

const prompt = `
	What are the http_status_code, grpc_status_code and error_description 
	for the following error codes(separate by #): %s? 
	The resulting output should be in this format: 
	http_status_code/grpc_status_code/error_description#http_status_code/grpc_status_code/error_description. 
	The output should not contain any newline character. For each error code, the output should separate by #.
`

func formatRequestString(errorString string) string {
	return fmt.Sprintf(
		prompt,
		errorString,
	)
}
