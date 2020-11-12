//5. сравнить два однотипных массива - это bool (a == b не катит)
//Если значения массивов равны, выводит true, иначе false
package main
import "fmt"
const MaxLength =6
type Array[MaxLength]int
func Equality(a,b Array)bool{
	f:=true
	for i:= range a {
		if a[i]!=b[i]{
			f=false
		}
	}
	return f
}
func main(){
	a:=Array{1,2,3,4,5,6}
	b:=Array{1,2,3,4,5,6}
	fmt.Print(Equality(a,b))
}
