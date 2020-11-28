//4. Ввод чисел из файла, а вывод:
// в один файл все введенные четные числа в том порядке, 
//в котором они шли во входном файле, а в другой – нечетные, причём в обратном порядке 
package main
import "fmt"
import "os"
import "bufio"
func main(){
	Even:=make([]int,0)
	Odd:=make([]int,0)
	fin,err:=os.Open("input94.txt")
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
		if c%2==0{
			Even=append(Even,c)
		}else{
			Odd=append(Odd,c)
		}
	}
	fout,err:=os.Create("output94_Even.txt")
	fout2,err2:=os.Create("output94_Odd.txt")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	if err2!=nil{
		fmt.Print("File write error")
		return
	}
	out:=bufio.NewWriter(fout)
	out2:=bufio.NewWriter(fout2)
	defer fout.Close()
	defer out.Flush()
	defer fout2.Close()
	defer out2.Flush()
	
	for i:=range Even{
		fmt.Fprintln(out,Even[i])
	}
	for i:=len(Odd)-1;i>=0;i--{
		fmt.Fprintln(out2,Odd[i])
	}
}
