package cmd

type Recipe struct {
	name        string
	dateCreated string
	timesCooked int
	ingredients []ingredient
}
type ingredient struct {
	name     string
	quantity string
}