package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"swarms_models_go/openai"
)

// Agent represents an autonomous agent capable of performing tasks and storing memory.
type Agent struct {
	maxLoops      int
	retryAttempts int
	loopInterval  int
	shortMemory   []string
	openaiClient  *openai.OpenAIClient
	loopCount     int
	systemPrompt  string
	agentName     string

}

// NewAgent creates a new Agent instance.
func NewAgent(apiKey, model string) *Agent {
	client := openai.NewClient(apiKey, model)
	return &Agent{
		maxLoops:      2,
		retryAttempts: 3,
		loopInterval:  2,
		openaiClient:  client,
		loopCount: 2,
		systemPrompt: "The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly.",
		agentName: "Autonomous Agent",
	}
}

// Run executes the agent's main loop, processing tasks and storing responses.
func (a *Agent) Run(task string) {
	a.activateAutonomousAgent()
	a.addTaskToMemory(task)
	a.loopCount = 0

	for a.loopCount < a.maxLoops {
		a.loopCount++
		log.Info().Msgf("Loop %d/%d", a.loopCount, a.maxLoops)

		taskPrompt := a.getShortMemory()

		var response string
		for attempt := 0; attempt < a.retryAttempts; attempt++ {
			resp, err := a.openaiClient.CreateChatCompletion(taskPrompt)
			if err == nil {
				response = resp
				break
			}
			log.Error().Err(err).Msgf("Attempt %d: Error generating response", attempt+1)
		}

		if response == "" {
			log.Error().Msg("Failed to generate a valid response after retry attempts.")
			break
		}

		log.Info().Msgf("Response: %s", response)
		a.addResponseToMemory(response)
		time.Sleep(time.Duration(a.loopInterval) * time.Second)
	}

	finalResponse := a.getShortMemory()
	log.Info().Msgf("Final Response: %s", finalResponse)
}

// activateAutonomousAgent logs the activation of the autonomous agent.
func (a *Agent) activateAutonomousAgent() {
	log.Info().Msg("Autonomous agent activated.")
}

// addTaskToMemory adds a task to the agent's short-term memory.
func (a *Agent) addTaskToMemory(task string) {
	a.shortMemory = append(a.shortMemory, task)
}

// getShortMemory returns a formatted string of the agent's short-term memory.
func (a *Agent) getShortMemory() string {
	return fmt.Sprintf("%v", a.shortMemory)
}

// addResponseToMemory adds a response to the agent's short-term memory.
func (a *Agent) addResponseToMemory(response string) {
	a.shortMemory = append(a.shortMemory, response)
}

func main() {
	// Configure zerolog for console output with time in local time zone.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Log the start of the application.
	log.Info().Msg("Starting the autonomous agent...")

	// Initialize the agent with the OpenAI API key and model.
	agent := NewAgent("sk-proj-nFULo5YbDhjgmOPOysaYT3BlbkFJ2WtN42akJg6gaiX7mNEA", "gpt-3.5-turbo")

	// Run the agent with the initial task.
	if err := runAgent(agent, "What are the best ways to establish a non-profit AI research lab"); err != nil {
		log.Fatal().Err(err).Msg("Agent run failed")
	}
}

// runAgent runs the provided agent with the specified task.
func runAgent(agent *Agent, task string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Recovered from panic: %v", r)
		}
	}()

	agent.Run(task)
	return nil
}
