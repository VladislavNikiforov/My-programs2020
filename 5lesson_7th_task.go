//7. Первая цифра числа.
package main
import "fmt"
func first(n uint)uint{
	for n/10!=0{
		n/=10
	}
	return n
}
func main (){
	var a uint 
	fmt.Scanln(&a)
	fmt.Print(first(a))

}
