// Created By: Lamees Hemdan
// Created On: March 2023
//
// This file contains the main function for the go_app application.

package main

import "fmt"

func main() {
	var length, width, height float64
	fmt.Println("Enter the length of pyramid: ")
	fmt.Scan(&length)
	fmt.Println("Enter the width of pyramid: ")
	fmt.Scan(&width)
	fmt.Println("Enter the height of pyramid: ")
	fmt.Scan(&height)
	volume := (length * width * height) / 3
	fmt.Println("The volume of pyramid is: ", volume, "mm³")

	fmt.Println("\nDone.")
}
