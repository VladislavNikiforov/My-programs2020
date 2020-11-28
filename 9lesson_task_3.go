//3. Вывести в один файл все введенные числа,
// но сначала положительные, а потом отрицательные.
package main
import "fmt"
import "os"
import "bufio"
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input93.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	in:=bufio.NewReader(fin)
	defer fin.Close()
	for {
		var c int
		_,err:=fmt.Fscan(in,&c)
		if err!=nil{
			break
		}
		sl=append(sl,c)
	}
	for i:=0;i<len(sl);i++{
		for j:=0;j<len(sl);j++{
			if sl[i]>sl[j]{
				sl[i],sl[j]=sl[j],sl[i]
			}
		}
	}
	fout,err:=os.Create("output93.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	out:=bufio.NewWriter(fout)
	defer fout.Close()
	defer out.Flush()
	for i:=range sl{
		fmt.Fprintln(out,sl[i])
	}
}
