//4. Простое ли число. `func prime(n uint) uint. Функция будет возвращать 0 или 1
package main
import "fmt"
func prime(n uint) uint {
	var i uint =2
	if n<=0 {
		return 0
	}
	for i<n {
		if n%i==0 {
			i=0
			break
		}
			i++
		}
		if i==0 {
			return 0
		}else{
			return 1
		}
}
func main () {
	var a uint
	fmt.Scanln(&a)
	fmt.Print(prime(a))

}
