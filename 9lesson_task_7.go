//7. Убрать из файла с числами все подряд идущие повторы.
package main

import "fmt"
import "os"
import "bufio"

func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input97.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	defer fin.Close()
	in:=bufio.NewReader(fin)
	
	for {
		var c int
		_,err:=fmt.Fscan(in,&c)
		if err!=nil{
			break
		}
		sl=append(sl,c)
	}
	for i:=1;i<len(sl);i++{
		if sl[i]==sl[i-1]{
			sl=append(sl[:i],sl[i+1:]...)
			i--
		}
	}
	
	fout,err:=os.Create("output97.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	defer fout.Close()
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	for i:=range sl{
		fmt.Fprintln(out,sl[i])
	}
} 
