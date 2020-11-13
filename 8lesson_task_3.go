//3. сколько раз число Х входит в слайс
package main
import "fmt"
func Count(sl[]int,X int)int{
	count:=0
	for _,v:=range sl{
		if v==X{
			count++
		}
	}
	return count
}
func main(){
	var X int
	sl:=make([]int,5)
	
	fmt.Print("Input 5 slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])  
	}
	fmt.Print("Input number -> ")
	fmt.Scan(&X)
	fmt.Printf("Number %d repeats in slice %d times",X,Count(sl,X))
	
}
