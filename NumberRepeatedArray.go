package main
// thank you so much kio - https://github.com/KioCoan
import ("fmt")

func main() {
	r := 0
	b := [10] int { 1, 1, 4, 10, 5, 7, 7, 4, 10 }

	for _, number := range b {
		r = r ^ number
	}

	fmt.Println("Searching for number repeated in the array")
	fmt.Println(b)
	fmt.Println("")
	fmt.Print("Number that repeats in array: ")
	fmt.Print(r)
}