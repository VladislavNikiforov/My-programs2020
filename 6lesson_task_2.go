//2) Найти N-е число Фибоначчи
package main
import "fmt"
func Fibo(n int)int{
	a,b,k:=1,1,0
	i:=2
	
	for i<=n{
		k=a
		a+=b
		b=k
		i++
		
	}
	return b
}
func main(){
var n int
	fmt.Scan(&n)
	fmt.Print(Fibo(n))
}
