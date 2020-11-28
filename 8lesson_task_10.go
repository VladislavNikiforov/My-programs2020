//1. Есть слайс. Убрать в нём соседние повторы. 
//  [ 1 3 3 2 1 3 5 3 4 4 2 ]  -->  [ 1 3 2 1 3 5 3 4 2 ]
package main 
import "fmt"
func main() {
	a:=[13]int{1,3,3,2,1,3,5,3,4,4,2}
	sl:=a[0:11]
	fmt.Println(sl)
	for i:=0;i+1<len(sl);i++{
		if sl[i]==sl[i+1] && sl[i]>0{
			sl=append(sl[:i],sl[i+1:]...)
			i--
		}
	}
	fmt.Print(sl)
}
