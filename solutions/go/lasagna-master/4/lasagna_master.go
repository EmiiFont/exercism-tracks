package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPerLayer int) int {
	if avgPerLayer == 0 {
		avgPerLayer = 2
	}

	return len(layers) * avgPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, list []string) {
	list[len(list)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quatities []float64, amount int) []float64 {
	quatitiesClone := make([]float64, len(quatities))

	mult := float64(amount) / float64(2)

	for i := 0; i < len(quatities); i++ {
		quatitiesClone[i] = quatities[i] * mult
	}
	return quatitiesClone
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
