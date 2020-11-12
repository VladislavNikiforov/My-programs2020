//1. Какая степень двойки. func power2(n uint) int. Если n не степень 2, то возвращаем –1.
package main 
import "fmt"
func power2(n int) int {
	var S,i int
	for S=1; S<n;{
		S*=2
		i++
		}
		if S==n {
			return i
		}else{
			return -1
	}
}
func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Print(power2(a))

}
