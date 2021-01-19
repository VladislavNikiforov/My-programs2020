//1. Склеить все строки файла в одну строку. Понятно, что хранить все строки в памяти одновременно не надо,
// вполне достаточно считывать строки по одной и выводить их в выходной файл с помощью fmt.Fprint.
package main 
import "fmt"
import "os"
import "bufio"
func main(){
	var s string
	fin,err:=os.Open("input121.txt")
	if err!=nil{
		fmt.Print("File read error")
		return
	}
	in:=bufio.NewScanner(fin)
	for in.Scan(){
		str:=in.Text()
		s+=str+" "
	}
	fout,err:=os.Create("input121.bak")
	if err!=nil{
		fmt.Print("File write error")
		return
	}
	out:=bufio.NewWriter(fout)
	fmt.Fprint(out,s)
	out.Flush()
	fout.Close()
	fin.Close()
	os.Remove("input121.txt")
	os.Rename("input121.bak","input121.txt")
}

