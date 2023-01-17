package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodle int = 0
	var sauce float64 = 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodle += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	coef := float64(portion) / 2
	var result []float64
	for _, recipe := range quantities {
		result = append(result, recipe * coef)
	}

	return result
}
