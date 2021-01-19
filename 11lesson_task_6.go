//6. сдвинуть строку на К символов по кругу
package main
import "fmt"
func main(){
	var s,ans string
	var k int
	fmt.Print("Input string ->")
	fmt.Scan(&s)
	fmt.Print("Input value, program will move words right to that value ->")
	fmt.Scan(&k)
	sl:=[]rune(s)
	if k>3{
		k%=len(sl)
	}
	sl2:=sl[:len(sl)-k]
	sl3:=sl[len(sl)-k:]
	ans+=string(sl3)+string(sl2)
	fmt.Print(ans)
}
