package main

import "fmt"

/*
	Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]],
	supplies = ["yeast","flour","meat"]
	Output: ["bread","sandwich"]
 */
/*
    Topological sort
	bread: sandwich

	degree sandwich:1
 */
func main() {
	fmt.Println(findAllRecipes([]string{"bread","sandwich"},
	[][]string{{"yeast","flour"}, {"bread","meat"}}, []string{"yeast","flour","meat"}))
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesHash := make(map[string]bool)
	for _, supply := range supplies {
		suppliesHash[supply] = true
	}

	ingredientToRecipe := make(map[string][]string)
	for i, recipe := range recipes {
		receiptIngredients := ingredients[i]
		for _, ingredient := range receiptIngredients {
			if _, ok := suppliesHash[ingredient]; !ok {
				ingredientToRecipe[ingredient] = append(ingredientToRecipe[ingredient], recipe) // hope it is recipe
			}
		}
	}

	indegree := make(map[string]int)
	for _, recipes := range ingredientToRecipe {
		for _, recipe := range recipes{
			indegree[recipe]++
		}
	}

	q := make([]string, 0)
	for _, recipe := range recipes{ // no need dependent on ingredient
		if _, ok := indegree[recipe]; !ok {
			q = append(q, recipe)
		}
	}

	output := make([]string, 0)
	for len(q) != 0 {
		ingredient := q[0] // recipe becomes ingredient
		q = q[1:]
		output = append(output, ingredient)

		for _, recipe := range ingredientToRecipe[ingredient]{
			indegree[recipe]--
			if indegree[recipe] == 0 {
				q = append(q, recipe)
			}
		}
	}

	return output
}