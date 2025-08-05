package darts

func Score(x, y float64) int {
	//(x - center_x)^2 + (y - center_y)^2 < radius^2 --- is in
    point := x*x + y*y
    switch {
        case point <= 1.0 : return 10
        case point <= 25.0 : return 5
        case point <= 100.0 : return 1
    }
    return 0
}
