package concepts

func PreparationTime2(layers []string, time int) int {
	avgTime := time

    if(time == 0) {
        avgTime = 2
    }

    numLayers := len(layers)
    return numLayers * avgTime
}

// for each noodle layer => 50 gr of noodles
// for each sauce layer => 0.2 liters of sauce
func Quantities(layers []string) (int, float64) {
	var noodleTotal int = 0
    var sauceTotal float64 = 0.0
    
    for _, element := range layers {
        if (element == "noodles") {
            noodleTotal += 50
        }

        if (element == "sauce") {
            sauceTotal += 0.2
        }
    }

    return noodleTotal, sauceTotal      
}

func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
    friendLastElement := friendsIngredients[len(friendsIngredients) - 1]
    myIngredients[len(myIngredients) - 1] = friendLastElement
}


func ScaleRecipe(quantities []float64, portions int) []float64 {
    var result []float64
    var total float64
    
	times := portions / 2
    isOdd := portions % 2 == 1

    for _, amount := range quantities {
        if (isOdd) {
            total = (amount / float64(2)) * float64(portions)
        } else {
        	total = amount * float64(times)    
        }
        
        result = append(result, total)
    }
    return result
}