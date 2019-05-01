//+build !test

package main

import (
	"fmt"
	"os"

	"github.com/michaelkrukov/go-upload-string/internal/myjson"
	"github.com/michaelkrukov/go-upload-string/internal/pkg/payload"
)

var help = `Usage:
- %s save
- %s load <url>

Command accepts string from standart input`

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "save" {
		payload := payload.FromReader(os.Stdin)

		url := myjson.Save(payload)

		fmt.Println(url)
	} else if len(args) == 2 && args[0] == "load" {
		payloadData := myjson.Load(args[1])

		fmt.Printf(
			"%s:\n%s\n",
			payloadData.Created,
			payloadData.Content,
		)
	} else {
		progName := os.Args[0]
		fmt.Printf(help, progName, progName)
		os.Exit(3)
	}
}
