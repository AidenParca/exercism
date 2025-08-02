package cards
import "fmt"
// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    return []int{2,6,9}
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if len(slice) >= index && index >= 0 {
    	return slice[index]
    } else {
        return -1
    }
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

    if len(slice) > index && index >= 0 {
    	slice[index] = value
        fmt.Println("value",value,"index",index)
        fmt.Println("slice",slice)
        return slice
    } else {
        slice := append(slice, value)
        fmt.Println("appended slice",slice)
        return slice
    }
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if len(slice) > index && index >= 0 {
    	return append(slice[:(index)],slice[(index+1):]...)
        } else {
        return slice
        }
	panic("Please implement the RemoveItem function")
}
