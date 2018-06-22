package main
// thank you so much kio - https://github.com/KioCoan
import ("fmt")

func main() {
	r := 0
	b := [9] int { 1, 1, 2, 5, 5, 7, 9, 7, 9}

	for _, number := range b {
		r = r ^ number
	}
	
	fmt.Println(b)
	fmt.Println("")
	fmt.Print("Number that repeats in array: ")
	fmt.Print(r)
}