package thefarm
import "errors"
import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cownum int) (float64, error) {
    amount, err := fc.FodderAmount(cownum)
    if err != nil {
		return 0, err
	}
    factor, err := fc.FatteningFactor()
    if err != nil {
		return 0, err
	}
    foodPerCow := (amount * factor) / float64(cownum)
    
    return foodPerCow, nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cownum int) (float64, error) {
    if cownum > 0 {
        return DivideFood(fc , cownum)
    } else {
        return  0 , errors.New("invalid number of cows")
    }
}
type MyCustomError struct {
  cownum int
  cm string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.cownum ,e.cm)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(fcownum int) error {
    fcm := ""
    switch  {
        case fcownum < 0 : fcm = "there are no negative cows"
        case fcownum == 0 : fcm = "no cows don't need food"
    }
    if fcm != "" {
        return &MyCustomError{ cownum:fcownum , cm : fcm }
        
    } else {
        return nil
    }
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
