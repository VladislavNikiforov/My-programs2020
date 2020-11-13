//1. сумму элементов слайса
package main
import "fmt"
func SliceSum(sl[]int)int{
	sum:=0
	for i:=range sl{  
		sum+=sl[i]
	}
	return sum
}
func main (){
	sl:=make([]int,5)
	
	fmt.Print("Input 5 slice elements-> ")
	for i:=0;i<len(sl);i++{
		fmt.Scan(&sl[i])
	}
	fmt.Printf("Sum of slice is %d",SliceSum(sl))
}
