//Write a function that rotates given slice by given number of steps to the right or
//to the left.
package main

import "fmt"

func rotadeSliceRight(s []int, rtd int) {
	l := len(s) // s slice lenght
	/*Algorithm:
	Rotation consists of 2 parts
	First piece: Copy the numbers except the first Rotation number (`rtd`) into the new Slice of `s` Slice
	Second piece: Complete the number of `s` Slice with the Slice from the first piece ignored
	*/
	x := s[(l - rtd):]
	y := s[:(l - rtd)]
	x = append(x, y...)
	fmt.Println("Rotated Right:", x)
}

func rotadeSliceLeft(s []int, rtd int) {
	//same algoritm like rotadeSliceRight
	x := s[rtd:]
	y := s[:rtd]
	x = append(x, y...)
	fmt.Println("Rotated Left:", x)
}

func main() {
	s := make([]int, 0)                         //Create a slice
	s = append(s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9) //assignment of slice
	fmt.Println("Slice:", s)
	rtd := 3 // rotation scale
	rotadeSliceRight(s, rtd)
	rotadeSliceLeft(s, rtd)

}
