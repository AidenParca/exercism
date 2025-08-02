package purchase
import "fmt"
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    switch {
        case kind == "car" || kind == "truck" : return true
        case kind == "bike" || kind == "stroller"||kind == "e-scooter": return false
        default : 	panic("NeedsLicense not implemented")
    } 

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    ch := ""
    if option1 < option2 {
        ch = option1
    } else {
        ch = option2
    }
    return fmt.Sprintf("%s is clearly the better choice.",ch)
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    x := int(age)
    switch {
           case x < 3 : return originalPrice * 0.80
        	case x >= 3 && x < 10 : return originalPrice * 0.70
			case x > 9 : return originalPrice * 0.50 
        	default : panic("CalculateResellPrice not implemented")
        }
    }
	

