# Flow Chat

Chat is a simple yet powerful chat application built with Go. It leverages the power of AI to provide a seamless and interactive chat experience. The application is designed to be flexible, efficient, and easy to use.

## Features

- Interactive chat interface with AI.
- Support for multiple AI models.
- Commands for managing chat history and changing AI models.
- Real-time sentiment analysis (coming soon).
- Support for multiple languages (coming soon).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- An API key from Perplexity AI

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ooyeku/flow
    ```
   
2. Change into the project directory:
3. Run the application:
   ```bash
   go run main.go
   ```
   
4. Install the required dependencies:
   ```bash
   go mod download
   ```
   
5. Set the API key as an environment variable.
   Download the Perplexity AI API key from the [Perplexity AI website](https://perplexity.ai/) and set it as an environment variable:
   ```bash
    export PAI_KEY=your-api-key
    ```

## Usage
Once PAI_KEY is set, you can start the application by running the following command in the project directory:
```bash
go run main.go chat
```

## Commands
- '$history' - View chat history.
- '$model' - Change the AI model.
- '$exit' - Exit the chat application.
- '$clear' - Clear the chat history.