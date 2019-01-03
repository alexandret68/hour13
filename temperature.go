package temperature

type Temperature float64

func CtoF(c float64) float64 {
	return(c*(9/5))+32
}

func FtoC(f float64) float64 {
	return(f-32)*(9/5)
}
