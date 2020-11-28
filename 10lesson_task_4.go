//4. заменяет все вхождения некоторого заданного символа 
//(заданный - в данном случае это не номер, а именно символ, некоторая руна)
// на некоорый другой заданный символ или строку
package main
import "fmt"
func SymbolChange(s string, change string, x string)string{
	sl:=[]rune(s)
	slX:=[]rune(x)
	slC:=[]rune(change)
	for i:= 0;i < len(sl);i++{
		if sl[i]==slC[0]{
			sl2 := make([]rune,0)
			sl2=append(sl2,sl[i+1:]...)
			sl=append(sl[:i],slX...)
			sl=append(sl,sl2...)
			i+=len(slX)
		}
	}
	return string(sl)
}
func main(){
	s:="Pacific ocean"
	var x,change string
	fmt.Printf("Input the letter which you would like to change\nin word {%s} and input the one or string you want to insert ->",s)
	fmt.Scan(&change,&x)
	fmt.Print(SymbolChange(s,change,x))
}
