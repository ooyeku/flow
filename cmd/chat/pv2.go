package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asdine/storm"
	"github.com/logrusorgru/aurora"
	"github.com/ooyeku/flow/cmd/chat/helpers"
	"github.com/ooyeku/flow/pkg/chat"
	"github.com/theckman/yacspin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// cliSetup initializes the chat app by creating a database, chat service, and chat app instances.
// It also sets up the necessary dependencies for the chat app to function properly.
// The function returns the chat app, chat service, and database instances.
func cliSetup() (*chat.ChatApp, *chat.ChatService, *storm.DB) {
	dbPath := "internal/inmemory/pv2.db"
	db, err := storm.Open(dbPath, storm.BoltOptions(0600, nil))
	if err != nil {
		log.Fatalf("error opening db: %s", err)
	}

	cs := chat.NewStromRepo(db)
	chatservice := chat.NewChatService(cs)

	client := &http.Client{}
	scanner := bufio.NewScanner(os.Stdin)
	apikey := os.Getenv("PAI_KEY")
	models := []string{"pplx-70b-online", "codellama-70b-instruct", "mistral-7b-instruct", "mixtral-8x7b-instruct", "mixtral-8x22b-instruct", "llama-3-70b-instruct"}

	app, err := chat.NewChatApp(client, chatservice, scanner, apikey, models)
	if err != nil {
		log.Fatalf("error creating chat app: %s", err)
	}

	// Get topics from db
	topics, err := chatservice.ListTopics()
	// set current topic to General$h
	app.CurrentTopic = *chat.NewChatTopic("General", "Standard chat topic")
	if err != nil {
		log.Fatalf("error getting topics: %s", err)
	}
	app.Topics = topics

	app.CurrentModel = "pplx-70b-online"

	return app, chatservice, db
}

// run is the main function responsible for running the chat application.
// It takes a pointer to a ChatApp struct as input and returns an error.
func run(app *chat.ChatApp) error {
	au := aurora.NewAurora(true)

	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Suffix:          au.Bold(au.BrightBlue("Pondering...")).String(),
		SuffixAutoColon: true,
		StopCharacter:   "💡",
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		log.Fatal("error creating spinner: ", err)
	}

	introMessage()
	for {
		// display current model
		fmt.Println(au.Bold(au.Gray(12, "Current Model: ")), au.Bold(au.Yellow(app.CurrentModel)))
		// display current topic
		fmt.Println(au.Bold(au.Gray(12, "Current Topic: ")), au.Bold(au.Red(app.CurrentTopic.Name)))
		fmt.Print(au.Green("You: "))
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		userInput := app.Scanner.Text()

		if userInput == "exit" {
			fmt.Println(au.Bold(au.BrightBlue("Goodbye!")))
			return nil
		}

		if strings.HasPrefix(userInput, "$") {
			command := strings.TrimPrefix(userInput, "$")
			switch command {
			case "topics":
				err := topicRun(app)
				if err != nil {
					log.Fatalf("error running topic command: %s", err)
				}
				continue
			case "models":
				err := modelRun(app)
				if err != nil {
					log.Fatalf("error running model command: %s", err)
				}
				continue
			case "history":
				err := historyRun(app)
				if err != nil {
					log.Fatalf("error running history command: %s", err)
				}
				continue
			case "clear":
				err := clearHisRun(app)
				if err != nil {
					log.Fatalf("error running clear command: %s", err)
				}
				continue
			case "exit":
				fmt.Println(au.Bold(au.BrightBlue("Goodbye!")))
				return nil
			default:
				fmt.Println(au.Bold(au.BrightRed("Invalid command")))
			}
			continue
		}

		err := spinner.Start()
		if err != nil {
			return err
		}

		payload := map[string]interface{}{
			"model": app.CurrentModel,
			"messages": []map[string]string{
				{"role": "system", "content": helpers.SysMessage},
				{"role": "user", "content": userInput},
			},
			"temperature":       helpers.Temperature,
			"top_p":             helpers.TopP,
			"frequency_penalty": helpers.FrequencyPenalty,
			"stream":            false,
		}
		payloadBytes, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "https://api.perplexity.ai/chat/completions", bytes.NewReader(payloadBytes))
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "Bearer "+app.ApiKey)

		resp, err := app.Client.Do(req)
		if err != nil {
			log.Fatalf("error sending request: %s", err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var result chat.Response
		_ = json.Unmarshal(body, &result)

		spinner.Stop()

		ai := au.Bold(au.BrightBlue("AI: "))
		aiResponse := fmt.Sprintf("%s%s", ai, result.Choices[0].Message.Content)
		fmt.Println(aiResponse)

		cRes := &chat.ChatResponse{
			ID:        result.ID,
			Model:     result.Model,
			Time:      time.Now().String(),
			UserQuery: chat.NewMessage(userInput, app.CurrentModel),
			Object:    result.Choices[0].Message.Content,
			Topic:     &app.CurrentTopic,
		}
		err = app.ChatService.SaveChatResponse(cRes)
		if err != nil {
			log.Fatalf("error saving chat response: %s", err)
		}
	}
}

// ...
func introMessage() {
	au := aurora.NewAurora(true)
	fmt.Println(au.Bold(au.Cyan(" Welcome to Flow Chat")))
	fmt.Println(au.Bold(au.BgCyan("-----------------------------------------")))
	// options
	fmt.Println(au.Gray(12, "Type 'exit' or $exit to quit the chat"))
	fmt.Println(au.Gray(12, "Type $models to change model"))
	fmt.Println(au.Gray(12, "Type $history to view chat history"))
	fmt.Println(au.Gray(12, "Type $clear to clear chat history"))
	fmt.Println(au.Gray(12, "Type $topics to manage topics"))
	fmt.Println(au.Gray(12, "Type $by-topic after $history to view history by topic"))
}

// topicRun displays the current topic and available topics, allows the user to switch to a new topic or create/delete a topic.
func topicRun(app *chat.ChatApp) error {
	au := aurora.NewAurora(true)

	// display current topic
	fmt.Println(au.Bold(au.BrightMagenta("Current Topic: ")), au.BgBrightCyan(app.CurrentTopic.Name))
	// display available topics
	fmt.Println(au.Bold(au.BrightMagenta("Available Topics: ")))
	for _, topic := range app.Topics {
		fmt.Println(au.BrightCyan(topic.Name))
	}

	fmt.Println(au.Bold(au.BrightMagenta("Enter topic name, $new to create a new topic, or $delete to delete a topic: ")))
	scanned := app.Scanner.Scan()
	if !scanned {
		return app.Scanner.Err()
	}
	topicName := app.Scanner.Text()

	// check if $new was entered
	if topicName == "$new" {
		fmt.Println(au.Bold(au.BrightMagenta("Enter new topic name: ")))
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		newTopicName := app.Scanner.Text()

		// create new topic
		newTopic := chat.NewChatTopic(newTopicName, "")
		err := app.ChatService.CreateTopic(newTopic)
		if err != nil {
			log.Fatalf("error creating topic: %s", err)
		}
		app.Topics = append(app.Topics, newTopic)
		app.CurrentTopic = *newTopic
		fmt.Println(au.Bold(au.Gray(12, "Current Topic: ")), au.Bold(au.Red(app.CurrentTopic.Name)))
		return nil
	}

	if topicName == "$delete" {
		fmt.Println(au.Bold(au.BrightMagenta("Enter topic name to delete: ")))
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		deleteTopicName := app.Scanner.Text()

		// check if topic exists
		for _, topic := range app.Topics {
			// if topic exists, delete it
			if topic.Name == deleteTopicName {
				// if topic is the current topic, set the current topic to nil
				if app.CurrentTopic.Name == deleteTopicName {
					app.CurrentTopic = chat.ChatTopic{}
				}
				err := app.ChatService.DeleteTopic(topic.ID)
				if err != nil {
					log.Fatalf("error deleting topic: %s", err)
				}
				// remove topic from app.Topics
				app.Topics = append(app.Topics[:0], app.Topics[1:]...)
				fmt.Println(au.Bold(au.BrightGreen("Topic deleted")))

				return nil
			}
		}
	}

	// check if topic exists
	for _, topic := range app.Topics {
		// if topic exists, set it as the current topic
		if topic.Name == topicName {
			app.CurrentTopic = *topic
			fmt.Println(au.Bold(au.Gray(12, "Current Topic Changed To: ")), au.Bold(au.Red(app.CurrentTopic.Name)))
			return nil
		}
	}

	// if topic does not exist, notify user
	fmt.Println(au.Bold(au.BrightRed("Topic does not exist")))
	return nil
}

// modelRun is a function in the chat package that allows the user to change the current model used in the chat application.
// Parameters:
// - app: a pointer to a ChatApp struct that contains the chat application state and configuration.
// Errors:
// - Returns an error if there is an error scanning user input or if the specified model does not exist.
// Behavior:
// - Prints the current model and available models.
// - Prompts the user to enter a model name.
// - Checks if the specified model exists.
// - If the model exists, sets it as the current model in the app.
// - If the model does not exist, notifies the user.
// - Returns nil if there are no errors.
func modelRun(app *chat.ChatApp) error {
	au := aurora.NewAurora(true)

	// Display current model
	fmt.Println(au.Bold(au.Gray(12, "Current Model: ")), au.BrightYellow(app.CurrentModel))

	// Display available models
	fmt.Println(au.Bold(au.BrightBlue("Available Models: ")))
	for _, model := range app.Models {
		fmt.Println(au.BrightGreen(model))
	}

	fmt.Println(au.Bold(au.BrightBlue("Enter model name: ")))
	scanned := app.Scanner.Scan()
	if !scanned {
		return app.Scanner.Err()
	}
	modelName := app.Scanner.Text()

	// Check if model exists
	for _, model := range app.Models {
		// If model exists, set it as the current model
		if model == modelName {
			app.CurrentModel = model
			fmt.Println(au.Bold(au.BrightGreen("Current Model: ")), au.Bold(au.BrightGreen(app.CurrentModel)))
			return nil
		}
	}

	// If model does not exist, notify user
	fmt.Println(au.Bold(au.BrightRed("Model does not exist")))
	return nil
}

// historyRun retrieves and displays the chat history based on user input.
func historyRun(app *chat.ChatApp) error {
	au := aurora.NewAurora(true)

	// Retrieve chat history
	chatResponses, err := app.ChatService.ListChatResponses()
	if err != nil {
		return fmt.Errorf("error retrieving chat history: %w", err)
	}

	fmt.Println(au.Bold(au.BrightBlue("Enter $by-topic to view history by topic or press enter to view all history: ")))
	scanned := app.Scanner.Scan()
	if !scanned {
		return app.Scanner.Err()
	}
	command := app.Scanner.Text()

	// Check if $by-topic was entered
	if command == "$by-topic" {
		fmt.Println(au.Bold(au.BrightBlue("Enter topic name: ")))
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		topicName := app.Scanner.Text()

		fmt.Println(au.Bold(au.BrightRed("-------------------------------------------------")))
		fmt.Println(au.Bold(au.BrightRed("Begin Chat History for Topic: ")), au.BrightRed(topicName))
		// Display chat history for the specified topic
		for _, chatResponse := range chatResponses {
			if chatResponse.Topic != nil && chatResponse.Topic.Name == topicName {
				fmt.Println(au.Bold(au.BrightYellow("Time: ")), au.Magenta(chatResponse.Time))
				fmt.Println(au.Bold(au.BrightGreen("You: ")), au.Green(chatResponse.UserQuery.Query))
				fmt.Println(au.Bold(au.BrightBlue("AI: ")), chatResponse.Object)
				fmt.Println(au.Bold(au.BrightYellow("Model: ")), au.Magenta(chatResponse.Model))
				fmt.Println(au.Bold(au.BrightYellow("Topic: ")), au.Magenta(chatResponse.Topic.Name))
				fmt.Println(au.Bold(au.BrightBlue("-------------------------------------------------")))
			}
		}
	} else {
		// Display all chat history
		fmt.Println(au.Bold(au.BrightRed("-------------------------------------------------")))
		fmt.Println(au.Bold(au.BrightRed("Begin Full Chat History")))
		for _, chatResponse := range chatResponses {
			fmt.Println(au.Bold(au.BrightYellow("Time: ")), au.Magenta(chatResponse.Time))
			fmt.Println(au.Bold(au.BrightGreen("You: ")), au.Green(chatResponse.UserQuery.Query))
			fmt.Println(au.Bold(au.BrightBlue("AI: ")), chatResponse.Object)
			fmt.Println(au.Bold(au.BrightYellow("Model: ")), au.Magenta(chatResponse.Model))
			if chatResponse.Topic != nil {
				fmt.Println(au.Bold(au.BrightYellow("Topic: ")), au.Magenta(chatResponse.Topic.Name))
			} else {
				fmt.Println(au.Bold(au.BrightYellow("Topic: ")), au.Magenta("None"))
			}
			fmt.Println(au.Bold(au.BrightBlue("-------------------------------------------------")))
		}
	}

	return nil
}

// clearHisRun clears the chat history based on user input.
// If the user enters "$all", it will prompt for confirmation and then clear all history.
// If the user enters "$by-topic", it will display available topics, prompt for a specific topic name,
// check if the topic exists, confirm deletion, and then clear the history of the specified topic.
// If the user enters an invalid command, it will display an error message.
//
// Parameters:
//   - app: A pointer to the ChatApp struct.
//
// Returns:
//   - An error, if any.
func clearHisRun(app *chat.ChatApp) error {
	au := aurora.NewAurora(true)

	fmt.Println(au.Bold(au.BrightBlue("Enter $all to delete all history or $by-topic to delete history of a specific topic: ")))
	scanned := app.Scanner.Scan()
	if !scanned {
		return app.Scanner.Err()
	}
	command := app.Scanner.Text()

	if command == "$all" {
		fmt.Println(au.Bold(au.BrightRed("Are you sure you want to delete all history? (yes/no): ")))
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		confirmation := app.Scanner.Text()
		if confirmation == "yes" {
			err := app.ChatService.ClearEntries()
			if err != nil {
				return fmt.Errorf("error clearing chat history: %w", err)
			}
			fmt.Println(au.Bold(au.BrightGreen("Chat history cleared")))
		}
	} else if command == "$by-topic" {
		fmt.Println(au.Bold(au.BrightBlue("Enter topic name: ")))
		// Display available topics
		for _, topic := range app.Topics {
			fmt.Println(au.BrightCyan(topic.Name))
		}
		scanned := app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}
		topicName := app.Scanner.Text()

		// Check if topic exists
		var topicExists bool
		for _, topic := range app.Topics {
			if topic.Name == topicName {
				topicExists = true
				break
			}
		}
		if !topicExists {
			fmt.Println(au.Bold(au.BrightRed("Topic does not exist")))
			return nil
		}

		// Confirm deletion
		fmt.Println(au.Bold(au.BrightRed("Are you sure you want to delete history of topic " + topicName + "? (yes/no): ")))
		scanned = app.Scanner.Scan()
		if !scanned {
			return app.Scanner.Err()
		}

		confirmation := app.Scanner.Text()
		if confirmation == "yes" {
			err := app.ChatService.ClearEntriesByTopic(topicName)
			if err != nil {
				return fmt.Errorf("error clearing chat history: %w", err)
			}
			fmt.Println(au.Bold(au.BrightGreen("Chat history of topic " + topicName + " cleared")))
		}
	} else {
		fmt.Println(au.Bold(au.BrightRed("Invalid command")))
	}

	return nil
}

func main() {
	app, _, db := cliSetup()
	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	err := run(app)
	if err != nil {
		log.Fatalf("error running app: %s", err)
	}
}
