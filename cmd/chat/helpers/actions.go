package helpers

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

var au = aurora.NewAurora(true)

func Intro() {
	fmt.Println(au.Bold(au.Cyan(" Welcome to Flow Chat (Perplexity AI) ")))
	fmt.Println(au.Bold(au.BgCyan("-----------------------------------------")))
	// options
	fmt.Println(au.Gray(12, "Type 'exit' or $exit to quit the chat"))
	fmt.Println(au.Gray(12, "Type $models to change model"))
	fmt.Println(au.Gray(12, "Type $history to view chat history"))
	fmt.Println(au.Gray(12, "Type $clear to clear chat history"))
}

func DisplayModels() {
	fmt.Println(au.Bold(au.BgBrown("options")))
	fmt.Println(au.Bold(au.BrightGreen("sonar-small-chat")))
	fmt.Println(au.Bold(au.BrightGreen("sonar-small-online")))
	fmt.Println(au.Bold(au.BrightGreen("sonar-medium-chat")))
	fmt.Println(au.Bold(au.BrightGreen("sonar-medium-online")))
	fmt.Println(au.Bold(au.BrightGreen("pplx-70b-online")))
	fmt.Println(au.Bold(au.BrightGreen("codellama-70b-instruct")))
	fmt.Println(au.Bold(au.BrightGreen("mixtral-8x7b-instruct")))
	fmt.Println("")
}
