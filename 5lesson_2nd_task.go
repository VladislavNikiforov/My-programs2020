//2. Какая степень числа K. `func powerK(n uint, K uint) int. Если n не степень K, то возвращаем –1.`
package main
import "fmt"
func powerK(n int, K int) int{
	var logK int
	for n>1 {
		if n%K!=0 {
			break
		}
		logK++
		n/=K
	}
	if n==1 {
		return logK
	}else{ 
		return -1 
	}
}
func main () {
	var a,l int
	fmt.Println("Input the Number and the potential base of it, programm calculates the logarithm")
	fmt.Scanln(&a,&l)
	fmt.Print(powerK(a,l))

}
