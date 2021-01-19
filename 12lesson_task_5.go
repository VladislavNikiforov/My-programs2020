//5. Построить функцию, сравнивающую две строки состоящих из русских/латышских букв.
// Сравнивать надо честно, как велит алфавит, а не так, как говорит функция strings.Compare -
// она сравниваем символы по их кодам, и буква "ё" идёт после всех русских букв, за латышские я и не говорю.
// Нет, сравниваем посимвольно, алфавит как-то храним, символы сравниваем по-честному.
package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	Rus:="абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"	
	Lav:="aābcčdeēfgģhiījkķlļmnņoprsštuūvzžAĀBCČDEĒFGĢHIĪJKĶLĻMNŅOPRSŠTUŪVZŽ"
	slR:=[]rune(Rus)
	slL:=[]rune(Lav)
	
	f,err:=os.Open("input125.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(f)
	var slS[]string
	for in.Scan(){
		str:=in.Text()
		slS=append(slS,str)
	}
	s1:=[]rune(slS[0])
	s2:=[]rune(slS[1])
	i:=0
	for {
		
		if s1[i]==s2[i]{
			i++
		}
		if sl[i]>sl[j]{
			fmt.Print("")
			return
		}
	}
}
