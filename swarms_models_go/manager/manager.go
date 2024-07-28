package manager


import (
	"fmt"
	"openai-go/llm"
)

type MultiLLMManager struct {
	Models []llm.LLM
}


func NewMultiLLMManager(models []llm.LLN) * MultiLLMManager {
	return &MultiLLMManager{
		Models: models,
	}
}


// GenerateAll generates text from all models in the MultiLLMManager.
// It takes a prompt string, maximum number of tokens, and temperature as input.
// It returns a slice of generated text strings and an error if any.
func (m *MultiLLMManager) GenerateAll(prompt string, maxTokens int, temperature float64) ([]string, error) {
	var results []string
	for _, model := range m.Models {
		result, err := model.Generate(prompt, maxTokens, temperature)
		if err != nil {
			return nil, fmt.Errorf("error generating text from model: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}