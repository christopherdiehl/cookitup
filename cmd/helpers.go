package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func IndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
func GenerateUniqueRandomNumbers(numbersToGenerate int, upperBound int) []int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var retVals []int
	i := 0
	for {
		randomNumber := r1.Intn(upperBound)
		if IndexOf(randomNumber, retVals) != -1 {
			continue
		}
		retVals = append(retVals, randomNumber)
		i++
		if i == numbersToGenerate {
			break
		}
	}
	return retVals
}
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
