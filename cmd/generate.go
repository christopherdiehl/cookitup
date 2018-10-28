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
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func generateRandomNumbers(numbersToGenerate int, upperBound int) []int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var retVals []int
	for i := 1; i <= numbersToGenerate; i++ {
		retVals = append(retVals, r1.Intn(upperBound))
	}
	return retVals
}
func getRecipesFromFiles(numRecipes int) []Recipe {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(home + "/.cookitup/storage/")
	if err != nil {
		log.Fatal(err)
	}
	var recipes []Recipe
	var fileNames []string
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}

	// for _, f := range chosenFiles {
	// 	var recipe Recipe
	// 	fmt.Println(f.Name())
	// 	jsonFile, err := os.Open(home + "/.cookitup/storage/" + f.Name())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer jsonFile.Close()
	// 	byteValue, err := ioutil.ReadAll(jsonFile)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	json.Unmarshal(byteValue, &recipe)
	// 	recipes = append(recipes, recipe)
	// }
	return recipes
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate your mealplan - default is 5 meals. Recieve a text at your specified number",
	Long:  `Generate your mealplan. Default to five meals.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRecipesFromFiles()
		fmt.Println("generate called")

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
