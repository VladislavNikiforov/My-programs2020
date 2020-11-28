//8. удаляет лидирующие/хвостовые пробелы строки
package main
import "fmt"
func DeleteSpaceFrontBack(s string)string{
	sl:=[]rune(s)
	Front,Back:=-1,-1
	for i:=0; sl[i]==' ';i++{
		Front=i
	}
	
	for i:=len(sl)-1;sl[i]==' ';i--{
		Back=i
	}
	
	if Front > -1{
		sl=append(sl[:Front+1],sl[Front+2:Back]...)
	}
	
	return string(sl)
}
func main(){
	s:="        aaaa       "
	fmt.Printf(DeleteSpaceFrontBack(s))
}
