package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    x := make(map[string]int)
    x["quarter_of_a_dozen"]	= 3
	x["half_of_a_dozen"]	= 6
	x["dozen"]	= 12
    x["small_gross"]	= 120
    x["gross"]	= 144
    x["great_gross"]	= 1728
    return x
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	q, ok := units[unit]
	if !ok {
		return false
	}
	bill[item] += q
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	amountToRemove, unitExists := units[unit]
	if !unitExists {
		return false
	}
	currentAmount, itemExists := bill[item]
	if !itemExists {
		return false
	}
	newAmount := currentAmount - amountToRemove
	if newAmount < 0 {
		return false
	}
	if newAmount == 0 {
		delete(bill, item)
	} else {
		bill[item] = newAmount
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    _, itemExists := bill[item]
	if !itemExists {
		return 0,false
	} else {
        return bill[item],true

    }
	panic("Please implement the GetItem() function")
}
