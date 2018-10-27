package cmd

import (
	"time"
)

type Recipe struct {
	Name         string       `json:"name"`
	DateCreated  string       `json:"dateCreated"`
	Instructions string       `json:"instructions"`
	TimesCooked  int          `json:"timesCooked"`
	Ingredients  []ingredient `json:"ingredients"`
}
type ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

/*
* NewRecipe returns a pointer to Recipe struct
* name is set to the recipeName given
* dateCreated is set to the current date
* ingredients is initialized as a slice
* timesCooked is defaulted to zero
 */
func NewRecipe(recipeName string) *Recipe {
	recipe := &Recipe{
		Name:         recipeName,
		DateCreated:  time.Now().String(),
		Instructions: "",
		TimesCooked:  0,
		Ingredients:  []ingredient{},
	}
	return recipe
}
func (r *Recipe) addIngredient(name string, quantity string) error {

	ingredient := ingredient{
		Name:     name,
		Quantity: quantity,
	}
	r.Ingredients = append(r.Ingredients, ingredient)
	return nil
}
