/*
7. подсчитать количество различий в двух однотипных массивах
(различие - на одинаковой позиции стоят различные числа)
*/
package main 
import "fmt"
const MaxLength=6
type Array [MaxLength]int
func DiffCount(a,b Array)int{
	DiffCount:=0
	for i:=range a{
		if a[i]!=b[i]{
			DiffCount++
		}
	}
	return DiffCount
}
func main (){
	a:=Array{1,2,3,4,5,6}
	b:=Array{6,5,4,4,5,6}
	fmt.Printf("We have 2 Arrays:\n%d\n%d\nThe programm counts amount of differences in two arrays:\nAmount is %d",a,b,DiffCount(a,b))
}
