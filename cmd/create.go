// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func shouldEnd(message string) bool {
	return message == "end" || message == "finished"
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a recipe",
	Long:  `Create a recipe that is stored in a json file at ~/.cookitup/storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		text, err := PromptSTDIN(reader, "Enter recipe name")
		if err != nil {
			fmt.Println("An error occured while reading in the user input")
		}
		fmt.Printf("Awesome! I always wanted to Cook%sUp", text)
		fmt.Println("Please enter your ingredients. Enter finished or end to stop accepting ingredients")
		for {
			ingredientName, err := PromptSTDIN(reader, "What is the name of the ingredient?")
			if err != nil {
				fmt.Println("An error occured while reading in the user input")
			}
			if shouldEnd(ingredientName) {
				break
			}
			ingredientQuantity, err := PromptSTDIN(reader, "How many do you need?")
			if err != nil {
				fmt.Println("An error occured while reading in the user input")
			}
			if shouldEnd(ingredientQuantity) {
				break
			}
			fmt.Printf("Need %s of %s\n", ingredientQuantity, ingredientName)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
