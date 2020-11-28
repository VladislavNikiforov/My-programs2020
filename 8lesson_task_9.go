//Возводит все элементы в степень n (n - целое)
package main
import "fmt"
func PowerSl(sl[]float64,n float64){
	k:=0.0
	for i:=range sl{
		k=sl[i]
		for j:=n;j>1;j--{
			sl[i]*=k
		}
		if n==0{
			sl[i]=1
		}
		if n<0 {
			for l:=n;l<=0;l++{
				sl[i]/=k
			}
		}
	}
}
func main() {
	sl:=make([]float64,5)
	var n float64
	fmt.Println("Input 5 slice elements")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	fmt.Print("Input degree, in which to power all the slice elements")
	fmt.Scan(&n)
	PowerSl(sl,n)
	fmt.Print(sl)
}
