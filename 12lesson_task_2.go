//2. Форматировать текст в файле - убрать незначащие пробелы. 
//Назначащие пробелы - это все пробелы в начале строки до первого непробела,
//все пробелы в конце стрроки после последнего непробела и все кратные пробелы. 
//Кратные пробелы - это подряд идущие цепочки пробелов. Концевые пробелы надо убирать все,
// а кратные пробелы надо заменять на одиночные. Так, например, строка, состоящая только из пробелов, превращается в пустую строку.
package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main(){
	var s string
	fin,err:=os.Open("input122.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(fin)
	for in.Scan(){
		str:=in.Text()
		strF:=strings.Split(str," ")
		for i:=0;i<len(strF);{
			if strF[i]==""{
				strF=append(strF[:i],strF[i+1:]...)
			}else{
				i++
			}
		}
		check:=""
		for j:=0; j<len(strF);j++{
			s+=strF[j]
			s+=" "
			check+=strF[j]
		}
		if check!=""{
			s+="\n"
		}
		
	}
	fout,err:=os.Create("input122.bak")
	if err!=nil{
		fmt.Print("file write error")
		return
	}
	out:=bufio.NewWriter(fout)
	fmt.Fprint(out,s)
	out.Flush()
	fout.Close()
	fin.Close()
	os.Remove("input122.txt")
	os.Rename("input122.bak","input122.txt")
}



