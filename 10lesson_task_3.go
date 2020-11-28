//3. заменяет определённый символ (символ с заданным номером) в строке на цепочку символов
package main
import "fmt"
func EditString(s string, word string, n int)string{
	sl:=[]rune(s)
	slW:=[]rune(word)
	var sl2[]rune
	
	sl2=append(sl2,sl[n+1:]...)
	sl=append(sl[:n],slW...)
	sl=append(sl,sl2...)
	return string(sl)
}
func main(){
	s:="apple"
	var n int
	var word string
	fmt.Printf("Input number of symbol in word {%s} you want to change\nand input the word you want to insert ->",s)
	fmt.Scan(&n,&word)
	fmt.Print(EditString(s,word,n))

}
