package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
	prepTime := 2
	if avgPrepTime != 0 {
		prepTime = avgPrepTime
	}
	return prepTime * len(layers)

}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, l := range(layers) {
		if l == "sauce" {
			sauce += 0.2
		} else if l == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friends []string, mine []string) {
	mine[len(mine) - 1] = friends[len(friends) - 1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	mod := float64(scale) / 2.0
	res := []float64{}
	for _, q := range(quantities) {
		res = append(res, q * mod)
	}
	return res
}
