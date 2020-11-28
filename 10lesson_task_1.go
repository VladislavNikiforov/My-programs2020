//1. считает количество вхождений данного символа в данную строку
package main
import "fmt"
func Count(s string, x rune)int{
	var cnt int
	sl:=[]rune(s)
	for _,v:=range sl{
		if v==x{
			cnt++
		}
	}
	return cnt
}
func main(){
	var s string
	var x rune
	//fmt.Print("Input string ->")
	//fmt.Scan(&s)
	//fmt.Print("Input character ->")
	//fmt.Scan(&x) - не работает
	s="hello world"
	x='l'
	fmt.Printf("The amount of %c in string is %d",x,Count(s,x))
}
