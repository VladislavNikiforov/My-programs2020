//7. проверяет симметрична ли строка
package main
import "fmt"
func Symmetry(s string)string{
	sl:=[]rune(s)
	var answer string
	f:=true
	for i:=range sl{
		if sl[i]!=sl[len(sl)-1-i]{
			f=false
			answer="NotSymmetric"
			break
		}
	}
	if f{
		answer="Symmetric"
	}
	return answer
}
func main(){
	var s string
	fmt.Printf("Symmetric/NotSymhetric\n------------------------\nInput string ->")
	fmt.Scan(&s)
	fmt.Printf("String is %s",Symmetry(s))
}
