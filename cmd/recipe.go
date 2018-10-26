package cmd

type Recipe struct {
	name         string
	dateCreated  string
	instructions string
	timesCooked  int
	ingredients  []ingredient
}

func (r *Recipe) addIngredient(name string, quantity string) error {

	ingredient := ingredient{
		name:     name,
		quantity: quantity,
	}
	r.ingredients = append(r.ingredients, ingredient)
	return nil
}

type ingredient struct {
	name     string
	quantity string
}
