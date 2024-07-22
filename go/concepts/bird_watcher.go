package concepts

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int
	for _, number := range birdsPerDay {
    	sum += number
  	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int
    indexStart := 0
	weekCount := 1
    
    for weekCount != week {
        indexStart += 7
        weekCount += 1 
    }
    
	for i := indexStart; i < indexStart + 7; i++ {
    	sum += birdsPerDay[i]
  	}
    
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
    }

    return birdsPerDay
}
