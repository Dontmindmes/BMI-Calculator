package main

import (
	"fmt"
	"math"
)

var weight float64
var height float64
var maxsum float64

func calculate() {
	maxsum = (weight / height) / height
	fmt.Printf("Your BMI is %f kg/m2", math.Ceil(maxsum*100)/100) //math.Ceil(x*100)/100
}

func main() {

	fmt.Println("Please enter all data in imperial foramt: kg/meters")

	fmt.Printf("Enter your weight (kilograms): ")
	fmt.Scan(&weight)

	fmt.Printf("Enter your height (meters): ")
	fmt.Scan(&height)

	calculate()

}
