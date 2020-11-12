//6. Перевертыш числа.
package main 
import "fmt"
func rev(n uint) uint {
	var dig,i uint
	for n>0 {
		dig=n%10
		n/=10
		i*=10
		i+=dig
	}
	return(i)
}
func main () {
	var a uint
	fmt.Scanln(&a)
	fmt.Print(rev(a))
}
