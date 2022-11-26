package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, c := range birdsPerDay {
		count += c
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	zeroIndexWeek := week - 1;
	correctWeek := birdsPerDay[(zeroIndexWeek * 7):(zeroIndexWeek + 1) * 7]
	return TotalBirdCount(correctWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range(birdsPerDay) {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
