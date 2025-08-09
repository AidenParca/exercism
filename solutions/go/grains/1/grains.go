package grains

import "math"
import "errors"

func Square(number int) (uint64, error) {
    if number <= 0 || number > 64{
        return 0,errors.New("Error")
    }
    if number > 0{
        number = number -1 
    }
	return uint64(math.Pow(2,float64(number))),nil

}


func Total() uint64 {
	return uint64(18446744073709551615)
}
