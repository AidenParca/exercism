package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    x := float64(productionRate)
    return x * (successRate*0.01)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    x := CalculateWorkingCarsPerHour(productionRate,successRate)/60
    return int(x)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    return uint((carsCount / 10) * 95000 + (carsCount % 10 ) * 10000)
	panic("CalculateCost not implemented")
}
