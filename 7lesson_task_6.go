/*
6. сравнить два однотипных массива и,
если они отличаются, 
вернуть номер первой позиции, где они отличаются
*/
package main
import "fmt"
const MaxLength=6
type Array[MaxLength]int
func DifferentIndex(a,b Array)int{
	k:=0
	for i:=range a{
		if a[i]!=b[i]{
			k=i
			break
		}
	}
	return k
}
func main(){
	a:=Array{1,2,3,4,5,6}
	b:=Array{7,2,3,1,5,7}
	fmt.Printf("We have Arrays:\n%d\n%d\nThe first different number is in position %d",a,b,DifferentIndex(a,b))
}
