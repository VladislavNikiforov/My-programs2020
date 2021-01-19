//2. Небольшое изменение примера 13_1.go: считать не строки целиком, а слова.
// Словом считаем произвольный отрезок строки, не содержащий пробелов, 
//и отделённый пробелами справа и слева (или находящийся на краю строки).
// Уместно применить функцию strings.Split или что-нибудь подобное.
package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main(){
	f,err:=os.Open("input132.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(f)
	m:=make(map[string]int)
	
	for in.Scan(){
		str:=in.Text()
		wordlist:=strings.Split(str, " ")
		for i:=0;i<len(wordlist);i++{
			if wordlist[i]==""{
				wordlist=append(wordlist[:i],wordlist[i+1:]...)
			}else{
				i++
			}
		}
		for j:=range wordlist{
			m[wordlist[j]]++
		}
	}
	fmt.Println("Repeats(times) / Word")
	for i,v:=range m{	
		if v>1{
			fmt.Printf("\t%d\t%s\n",v,i)
		}
	}
}
