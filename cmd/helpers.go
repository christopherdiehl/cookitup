package cmd

import (
	"bufio"
	"fmt"
	"strings"
)

func PromptSTDIN(reader *bufio.Reader, message string) (userInput string, err error) {
	fmt.Println(message)
	userInput, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	return
}

func ShouldTerminate(message string) bool {
	return message == "end" || message == "finished"
}
