package main
import "fmt"
func main () {
	var a,p int
	var k int =1
	fmt.Println("The program calculates the amount of the digits in the inputed natural number")
	fmt.Scan(&a)
	p=a
	for a/10!=0 {
		a=a/10
		k++	
	}
fmt.Printf("The amount of the digits in %d is %d",p,k)
}
