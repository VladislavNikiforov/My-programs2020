/*
10. отпечатать массив целых чисел в два столбца:
в левом столбце - положительные числа,
в правом - остальные;
числа печатать в том же порядке, как они идут в массиве
*/
package main 
import "fmt"
const MaxLength=10
type Array [MaxLength]int
func Sorting(a Array){
	for i:=range a{
		if a[i]<0{
			fmt.Println("  ",a[i])
		}else{
			if a[i]>0{
				fmt.Println(a[i])
			}else{
				fmt.Println("   ",a[i])
			}
		}
	}
}
func main(){
	a:=Array{1,-11,3,-7,-6,-5,8,0,0,4}
	Sorting(a)
}
