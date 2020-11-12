package main
import "fmt"
func SumFib(n int)int{
	a,b,k,sum := 1,1,0,1
	for i:=2;i<=n;i++{
		k=a
		a+=b
		b=k
		sum+=b
	
	}
	return sum
}
func main (){
	var n int
	fmt.Scan(&n)
	fmt.Print(SumFib(n))

}
