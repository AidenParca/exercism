package lasagna
import "fmt"
// TODO: define the 'OvenTime' constant
const (
    OvenTime = 40
)
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven < 40 {
		Remaining := OvenTime - actualMinutesInOven
		fmt.Printf("%v Minutes Remaining", Remaining)
		return Remaining
    } else{
        panic("RemainingOvenTime not implemented")
    }
}
// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    x := numberOfLayers * 2
    return x
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    x := numberOfLayers * 2
    y := actualMinutesInOven + x
    return y
	panic("ElapsedTime not implemented")
}

func main(){
    RemainingOvenTime(30)
}
