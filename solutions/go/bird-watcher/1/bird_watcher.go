package birdwatcher
import "fmt"
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    total :=0
    for i := 0;i < len(birdsPerDay); i++ {
        total = total + birdsPerDay[i]
    }
    return total
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	end := week * 7
	start := end - 7
	total := 0

	for i := start; i < end; i++ {
		if i >= len(birdsPerDay) {
			break
		}
		total += birdsPerDay[i]
	}

	fmt.Println(total)
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0;i < len(birdsPerDay); i+=2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
