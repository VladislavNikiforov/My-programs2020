package main
import "fmt"
func main () {
	var a,p int
	var k int =0
	fmt.Println("The programm calculates the sum of the digits of the natural number")
	fmt.Println("Input natural number")
	fmt.Scan(&a)
	p=a
	for a!=0 {
		k+=a%10
		a=a/10	
	}
	fmt.Printf("The sum of the digits of %d is %d",p,k)
}

