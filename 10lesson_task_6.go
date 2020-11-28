//6. удаляет все вхождения данного символа в строку
package main
import "fmt"
func DeleteSymbol(s string, x string)string{
	sl:=[]rune(s)
	slX:=[]rune(x)
	for i:=0;i<len(sl);i++{
		if sl[i]==slX[0]{
			sl=append(sl[:i],sl[i+1:]...)
			i--
		}
	}
	return string(sl)
}
func main(){
	var x string
	s:="heedlessness"
	fmt.Printf("We have word {%s},\nInput the letter you would like to delete:\n(Program will delete it all time it appears)->",s)
	fmt.Scan(&x)
	fmt.Printf("New word is {%s}",DeleteSymbol(s,x))
}
