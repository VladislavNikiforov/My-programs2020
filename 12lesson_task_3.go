//3. Отформатировать входной файл следующим образом: 
//весь текст, который в нём содержится выводить строками ровно по 40 символов в каждой.
// Последняя строка может быть неполной.
package main
import "fmt"
import "os"
import "bufio"
func main(){
	var s string
	fin,err:=os.Open("input123.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(fin)
	for in.Scan(){
		str:=in.Text()
		s+=str+" "
	}
	sl:=[]rune(s)
	fout,err:=os.Create("input123.bak")
	if err!=nil{
		fmt.Print("file write error")
		return
	}
	out:=bufio.NewWriter(fout)
	x:=0
	for _=range sl{
		if x+40<len(sl){
			sl2:=sl[x:x+40]
			slStr:=string(sl2)
			fmt.Fprint(out,slStr)
			x+=40
			fmt.Fprint(out,"\n")
		}else{
			sl2:=sl[x:len(sl)]
			slStr:=string(sl2)
			fmt.Fprint(out,slStr)
			break
		}
		
	}
	out.Flush()
	fout.Close()
	fin.Close()
	os.Remove("input123.txt")
	os.Rename("input123.bak","input123.txt")
}
