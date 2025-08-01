package raindrops
import "fmt"
func Convert(number int) string {
    three := ""
    five := ""
    seven := ""
    if number % 5 != 0 && number % 3 != 0 && number % 7 != 0{
        return fmt.Sprintf("%d",number)
    } 
    if number % 5 == 0 {
        five = "Plang"
    } 
    if number % 3 == 0{
        three = "Pling"
    }
    if number % 7 == 0{
        seven = "Plong"
    }
    return fmt.Sprintf("%s%s%s",three,five,seven)
	panic("Please implement the Convert function")
}
