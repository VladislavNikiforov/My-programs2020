package main
import "fmt"
import "math"
func main () {
	var a,b,as float64
	fmt.Println("Input two numbers and the programm calculates the precise squares between them")
	fmt.Scan(&a,&b)
	a=math.Ceil(a)
	b=math.Trunc(b)
	if a<b {
		for a<=b {
		as=a*a
		fmt.Println(as)
		a++
		}
	}else{
		for a*a>=b*b {
		as=a*a
		fmt.Println(as)
		a--
		}
	}
}


