//2. заменяет определённый символ
// (символ с заданным номером) в строке на другой символ
package main
import "fmt"
func Edit(s string, x string, n int)string{
	sl:=[]rune(s)
	slX:=[]rune(x)
	sl[n]=slX[0]
	return string(sl)
}
func main(){
	s:="programing"
	var x string
	var n int
	fmt.Printf("Input number of symbol you want to change\nin word %s (0-9)\nThe symbol you want to insert->",s)
	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Print(Edit(s,x,n))
	
}
