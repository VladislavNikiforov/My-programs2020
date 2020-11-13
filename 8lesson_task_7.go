//7. Увеличивает все элементы слайса на n
package main
import "fmt"
func Increase(sl[]int,n int){
	for i:=range sl{
		sl[i]+=n
	}
}
func main(){
	sl:=make([]int,5)
	var n int
	
	fmt.Print("Input 5 slice elements ->")
	for i:=range sl{
		fmt.Scan(&sl[i])
	}
	fmt.Printf("The programm increases all of the slice elements by number you input -> ")
	fmt.Scan(&n)
	Increase(sl,n)
	fmt.Print(sl)
}
