package concepts

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var result float64 = (float64(productionRate) * successRate/100)/60    
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var costIndividualCar int = 10000
    var cost10Cars int = 95000
    
    if (carsCount > 0 && carsCount < 10 ) {
        return uint(carsCount * costIndividualCar)
    }

    if(carsCount >= 10) {
    	var intValue int = carsCount % 10
    
        fmt.Println("Debug message %i ", intValue)

        if (intValue == 0) {
        	return uint(carsCount/10 * cost10Cars)
        } else {
            return uint((carsCount/10 * cost10Cars) + (intValue * costIndividualCar))
        }
    }

    return 0
}