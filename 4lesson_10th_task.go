package main
import "fmt"
import "math"
func main () {
var a,b,i float64
	fmt.Println("Input two numbers and the programm calculates the precise squares between them")
	fmt.Scan(&a,&b)
	for a<=b {
		i=math.Round(math.Sqrt(a))
		if i==math.Sqrt(a) {
			fmt.Println(a)
		}
		a++
	}
}
