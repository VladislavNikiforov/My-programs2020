// 12&. Гистограмма (вертикальная) – столбики растут сверху вниз 
// (чисел не больше 80, их величины - от 0 до 25). Вывод в файл.
pcakage main
import "fmt"
import "os"
import "bufio"
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input912.txt")
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
		if len(sl)<=80 && c>=0 && c<=25{
			sl=append(sl,c)
		}
	}
}
/*
package main
import "fmt"
import "os"
import "bufio"
func main(){
	sl:=make([]int,0)
	fin,err:=os.Open("input9122.txt")
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
		if len(sl)<=80 && c>=0 && c<=25{
			sl=append(sl,c)
		}
	}
	
	fout,err:=os.Create("output912.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	defer fout.Close()
	out:=bufio.NewWriter(fout)
	defer out.Flush()
	
	max:=sl[0]
	for maxIndex:=range sl{
		if sl[maxIndex]>max{
			max=sl[maxIndex]
		}
	}
	fmt.Print(max)
	for i:= 1; i <= max; i++{
		for j:=range sl{
			if sl[j]>=i {
				fmt.Fprint(out,"*")
			}else{
				fmt.Fprint(out," ")
			}
		}	
		fmt.Fprintln(out)
	}
}
	*/		
		
		
		
		
		

