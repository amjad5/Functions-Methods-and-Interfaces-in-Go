package main 

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, vo, so float64) func (float64) float64 { 
	fn := func (t float64) float64 { 
		// s = ½ a t2 + vot + so 
		return ( (1/2) * a * (math.Pow(t, 2)) + ( (vo*t) + so))
	}
	return fn
}

func main(){
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement separated by space i.e. 10 2 1")
	var a, vo, so, t float64
	_, err := fmt.Scanf("%f %f %f", &a, &vo, &so)
	if err != nil {
		panic(err)
	}
	GenDisplace := GenDisplaceFn(a, vo, so)	// 10, 2, 1
	fmt.Println("Enter value for time i.e. 3")

	_, timeInputError := fmt.Scanf("%f", &t)
	if timeInputError != nil {
		panic(timeInputError)
	}
	
	fmt.Printf("%f ", GenDisplace(t))
}