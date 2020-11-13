//8. Усредняет слайс (все элементы становятся средним арифметическим)
package main
import "fmt"
func Average(sl[]float64){
	av:=0.0
	for i:=range sl{
		av+=sl[i]
	}
	av/=float64(len(sl))
	for i:=range sl{
		sl[i]=av
	}
}
func main(){
	sl:=make([]float64,5)
	fmt.Print("Input 5 slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	Average(sl)
	fmt.Print(sl)
}
