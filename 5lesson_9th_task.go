//9. Правда ли, что все цифры в числе четные (нечетные, одинаковой четности)
package main
import "fmt"
func parity(n uint)uint{
	var i,j int
	for n>0{
		if n%2==0{
			i++
		}else{
			j++
		}
		n/=10
	}
	if i==0 ||j==0 {
		return 1
	}else{
		return 0
	}
}
func main () {
	var a uint
	fmt.Scanln(&a)
	fmt.Print(parity(a))
}
