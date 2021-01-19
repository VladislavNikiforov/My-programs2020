/*3.Имеется файл, каждая строка файла содержит два слова - две строки символов,
    разделённые одним или несколькими пробелами. Первое слово - это имя покупателя, вторая - товар,
    который он заказал или купил, неважно. Задания:
    3a. Сконструировать мап, в котором ключи - это покупатели, а значения - количество заказанных товаров.
    Написать следующие функции:
    - Определить, сколько всего покупателей
    - Определить, сколько всего сделано заказов
    - Вывести, сколько заказов сделал каждый покупатель
    - Какое максимальное количество заказов сделал один покупатель? Сколько таких покупателей. Вывести их всех.*/
package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
func Amount(m map[string]int)int{
	amount:=len(m)
	return amount
}
func Orders(m map[string]int)int{
	var n int
	for _,v:=range m{
		n+=v
	}
	return n
}
func OrdPerCust(m map[string]int)string{
	list:=""
	for i,v:=range m{
		list+="\t"+strconv.Itoa(v)+"\t/ "+i+"\n"
	}
	return list
}
func MaxOrd(m map[string]int)(int,int,string){
	Vmax:=0
	for _,v:=range m{
		if Vmax==0{
			Vmax=v
		}
		if Vmax<v{
			Vmax=v
		}
	}
	
	var s string
	var cMax int
	
	for i,v:=range m{
		if v==Vmax{
			cMax++
			s+=" "+i+"\n"
		}
	}
	return Vmax, cMax, s
}
func main(){
	f,err:=os.Open("input133.txt")
	if err!=nil{
		fmt.Print("file read error")
		return
	}
	in:=bufio.NewScanner(f)
	m:=make(map[string]int)
	for in.Scan(){
		str:=in.Text()
		wordlist:=strings.Split(str," ")
		m[wordlist[0]]++
	}
	fmt.Printf("Program analyses the file and provides the data\n------------------------------\n")
	fmt.Printf("Amount of customers is %d.\n------------------------------\n",Amount(m))
	fmt.Printf("Amount of orders is %d.\n------------------------------\n",Orders(m))
	fmt.Println("Count of orders / Customer name")
	fmt.Print(OrdPerCust(m),"------------------------------\n")
	max,cMax,s:=MaxOrd(m)
	fmt.Printf("The biggest amount of orders is %d;\n%d person(s) have bought this amount of orders;\nThey are:\n%s",max,cMax,s)
	
}

    
