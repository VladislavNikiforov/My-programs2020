//1) Задача Иосифа.
// Храним массив их 0/1 – есть или нет человек на этом месте.
package main //ЗАДАЧА НАПИСАНА САМЫМ НЕОПТИМИЗИРОВАННЫМ ОБРАЗОМ, НО У МЕНЯ УЖЕ НЕТ НЕРВОВ ЧТО-ТО ТУТ МЕНЯТЬ
import "fmt"
const MaxLength=41
type Array[MaxLength]bool
func LastNum(a Array){
	l:=MaxLength
	k:=0
	m:=0
	for i:=range a{
		a[i]=true
	}	
	for l>=1{
		for i:=0;i<=MaxLength-1;i+=2{
			if a[i]{
				k=i+1	
			}
			if l==1{
				l--
				break
			}
			if m==0{
				i++
			}
			if a[i]{
				a[i]=false
				l--
			}
			m++
		}
		fmt.Print(a)
	}
	fmt.Print(k)
}
func main (){
	var a Array
	LastNum(a)
} 
