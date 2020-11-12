//8. отпечатать массив в строку
//9. отпечатать массив в столбец
package main
import "fmt"
func main() {
	a:=[6]int{1,2,3,4,5}
	fmt.Println(a) //так со скобками (строка)
	for i:=range a{
		fmt.Print(a[i]," ") //так без скобок (строка)
	}
	fmt.Println()
	
	for i:=range a{
		fmt.Println(a[i])//в столбик
	}

}
