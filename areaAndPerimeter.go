// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program calculates the area and perimeter of a rectangle

package main

import "fmt"

func main() {
	// This function calculates the area and perimeter of a rectangle
	var length int
	var width int

	// input
	fmt.Println("This program calculates the area and perimeter of a rectangle")
	fmt.Println()
	fmt.Print("Enter the length (cm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width (cm): ")
	fmt.Scanln(&width)
	fmt.Println()

	// process
	var area = length * width
	var perimeter = 2 * (length + width)

	// output
	fmt.Println("The area is: ", area, " cm²")
	fmt.Println("The perimter is: ", perimeter, " cm")

}
