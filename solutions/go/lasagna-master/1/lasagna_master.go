package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,time int) int{
    if time == 0 {
        time = 2
    }
    x := len(layers)
    return x * time
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int,float64){
    qSauce := 0.0 
    qNoodle := 0
    for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce"{
        	qSauce += 1.0
    	}
		if layers[i] == "noodles"{
        	qNoodle ++
    	}}
    return qNoodle * 50 , qSauce * 0.2 

    panic("")
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string { 
    secretIngredient := friendsList[len(friendsList)-1] 
    myList = myList[:len(myList)-1] 
    myList = append(myList, secretIngredient)
    return myList
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	sc := float64(portions) / 2.0
	cp := make([]float64, len(quantities))
	for i, amount := range quantities {
		cp[i] = amount * sc
	}
	return cp
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
