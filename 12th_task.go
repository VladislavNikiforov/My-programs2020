package main
import (
"fmt"
"math"
)
func main () {
var a,b float64 

fmt.Print("Input two numbers and the programm calculates the maximum ->")
fmt.Scan(&a)
fmt.Scan(&b)
max :=(a+b+(math.Abs(a-b)))/2
fmt.Printf("The maximum is %g",max)
}
