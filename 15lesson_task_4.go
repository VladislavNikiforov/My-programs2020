/*Задачи на switch:
4. Написать функцию, которая сравнивает лексикографически две строки. 
* */
package main
import "fmt"
func Comp(s1,s2 string)int{
	switch{
		case s1>s2: return 1
		case s2>s1: return -1
		default: return 0
	}
}
func main(){
	var s1,s2 string
	fmt.Print("Input 2 strings -> ")
	fmt.Scan(&s1,&s2)
	ans:=Comp(s1,s2)
	switch {
		case ans>0: fmt.Printf("%s is bigger than %s",s1,s2)
		case ans<0: fmt.Printf("%s is bigger than %s",s2,s1)
		default: fmt.Print("The strings are equal")
	}
}
