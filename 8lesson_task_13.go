//Выкинуть из слайса все четные числа. Эти числа загнать другой слайс.
package main

import "fmt"

func main() {

	sl:=make([]int,10)
	sl2:=make([]int,0)

	fmt.Println("The program cuts even elements and inserts its to second slice")
	fmt.Println("Input 10 slice elements")
	for i:=0;i<len(sl);i++{
		fmt.Scan(&sl[i])
		if sl[i]%2==0{
			sl2=append(sl2,sl[i])
			sl=append(sl[:i],sl[i+1:]...)
			i--
		}
	}
	fmt.Print(sl,sl2)
}
