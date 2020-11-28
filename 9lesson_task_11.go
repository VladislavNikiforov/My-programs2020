//11. Cимметричен ли числовой массив (заданный в файле)?
package main
import "fmt"
import "os"
import "bufio"
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input911.txt")
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
	f:=true
	
	for i:=0;i<=len(sl)/2;i++{
		if sl[i]!=sl[len(sl)-1-i]{
			f=false
			break
		}
	}
	if f{
		fmt.Print("Array in file IS symmetric")
	}else {
		fmt.Print("Array in file is NOT symmetric")
	}
		
	
}
