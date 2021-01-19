//1. Есть набор чисел из диапазона от - 2 млрд до 2 млрд.
// Подсчитать количество различных чисел, вывести все числа,
// которые встречаются среди данных (без повторов), подсчитать количество появлений каждого числа.
package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	f,err:=os.Open("input131.txt")
	if err!=nil{
		fmt.Print("filer ead error")
		return
	}
	in:=bufio.NewReader(f)
	m:=make(map[int]int)
	n:=0
	for {
		var c int
		_,err:=fmt.Fscan(in,&c)
		if err!=nil{
			break
		}
		m[c]++
		n++
	}
	fmt.Printf("The amount of numbers in file is %d;\n",n)
	fmt.Println("    Value / Repeats(times)")
	for i,v:=range m{
		fmt.Printf("%10d %d\n",i,v)
	}
}
