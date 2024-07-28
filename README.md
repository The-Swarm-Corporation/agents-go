# Agents in Go

The Enterprise-Grade Production-Ready Agent Framework in Golang


## Install

```bash
go install agents
```


## Usage

```go
package main

import (
	"fmt"
	"os"
	"time"
    "agents"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Configure zerolog for console output with time in local time zone.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Log the start of the application.
	log.Info().Msg("Starting the autonomous agent...")

	// Initialize the agent with the OpenAI API key and model.
	agent := NewAgent("your_api_key", "gpt-3.5-turbo")

	// Run the agent with the initial task.
	if err := runAgent(agent, "What are the best ways to establish a non-profit AI research lab"); err != nil {
		log.Fatal().Err(err).Msg("Agent run failed")
	}
}

```